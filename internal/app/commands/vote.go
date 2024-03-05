package commands

import (
	"fmt"
	"restapi/internal/domain"
	"restapi/internal/pkg/database"
	"sync"

	"github.com/phuslu/log"
)

var (
	lock                = sync.Mutex{}
	voteServiceInstance *VoteService
)

type VoteService struct {
	votesChan chan domain.VoteSumUpdate
	errorChan chan error
	db        *database.DB
}

func NewVoteService(database *database.DB) *VoteService {
	if voteServiceInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if voteServiceInstance == nil {
			voteServiceInstance = &VoteService{
				votesChan: make(chan domain.VoteSumUpdate, 100),
				errorChan: make(chan error),
				db:        database,
			}

			go voteServiceInstance.postUpdater()

			log.Info().Msg("Vote and error chans created")
		}
	}

	return voteServiceInstance
}

func (v *VoteService) postUpdater() {

	for {
		data, ok := <-v.votesChan
		if !ok {
			v.errorChan <- fmt.Errorf("channel closed")
			log.Error().Err(fmt.Errorf("channel closed"))

			break
		}

		_, err := v.db.Exec("call voteSumUpdater($1, $2)", data.PostID, data.Vote)
		if err != nil {
			v.errorChan <- err
			log.Error().Err(fmt.Errorf("vote error occured. See error:%w", err))
			continue
		}

		v.errorChan <- nil

	}

	v.restartPostUpdater()

}

func (v *VoteService) UpdatePostVotes(vote domain.VoteSumUpdate) error {
	v.votesChan <- vote

	return <-v.errorChan
}

func (v *VoteService) restartPostUpdater() {
	if v.isClosed() {
		log.Error().Err(fmt.Errorf("channel is closed"))
		return
	}

	close(v.votesChan)

	newVoteChan := make(chan domain.VoteSumUpdate)
	v.votesChan = newVoteChan

	go func() {
		v.postUpdater()
	}()
}

func (v *VoteService) isClosed() bool {
	_, ok := <-v.votesChan

	return ok
}
