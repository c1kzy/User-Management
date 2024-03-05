package commands

import (
	"restapi/internal/domain"
	"restapi/internal/pkg/database"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestVoteService_postUpdater(t *testing.T) {
	votes := 0

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	db := database.DB{Db: sqlxDB}

	v := NewVoteService(&db)

	mock.ExpectExec("call voteSumUpdater(?, ?)").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(0, 1))

	tests := []struct {
		name    string
		values  domain.VoteSumUpdate
		wantErr error
	}{
		{
			name: "1",
			values: domain.VoteSumUpdate{
				Vote:   1,
				PostID: 1,
			},
			wantErr: nil,
		},
		{
			name: "2",
			values: domain.VoteSumUpdate{
				Vote:   1,
				PostID: 1,
			},
			wantErr: nil,
		},

		{
			name: "3",
			values: domain.VoteSumUpdate{
				Vote:   -1,
				PostID: 1,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			votes++
			tt := tt // go vet fix
			t.Parallel()
			v.UpdatePostVotes(tt.values)

			assert.Equal(t, 3, votes)
		})
	}
}
