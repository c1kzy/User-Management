package handler

import (
	"errors"
	"fmt"
)

var (
	userUpdated = "user updated"
	userDeleted = "user deleted"

	postCreated = "post created"
	postUpdated = "post updated"
	postDeleted = "post deleted"

	userVoted   = "user voted"
	voteUpdated = "user vote updated"

	/// Errors
	//lint:file-ignore U1000 Ignore report
	//lint:file-ignore ST1012 Ignore report
	selfVoteErr    = errors.New("cannot vote for yourself")
	cannotVoteErr  = errors.New("cannot vote yet")
	invalidIDErr   = fmt.Errorf("invalid id")
	invalidPageErr = fmt.Errorf("invalid page")
)
