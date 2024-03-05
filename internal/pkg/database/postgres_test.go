package database

import (
	"fmt"
	"restapi/internal/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDB_Exec(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := mocks.NewDBService(ctrl)

	tests := []struct {
		name        string
		mocks       func(db *mocks.DBService)
		sp          string
		args        []any
		expectedErr error
	}{
		{
			name: "ok",
			mocks: func(db *mocks.DBService) {
				db.EXPECT().Exec("insertUser", []interface{}{"test1", "test2"}).Return([]interface{}{"test1", "test2"}, nil)
			},
			sp:          "insertUser",
			args:        []interface{}{"test1", "test2"},
			expectedErr: nil,
		},
		{
			name: "empty sp",
			mocks: func(db *mocks.DBService) {
				db.EXPECT().Exec("", []interface{}{"test1", "test2"}).Return([]interface{}{}, fmt.Errorf("empty input: sp:%v, args:%v", "", []interface{}{"test1", "test2"}))
			},
			sp:          "",
			args:        []interface{}{"test1", "test2"},
			expectedErr: fmt.Errorf("empty input: sp:%v, args:%v", "", []interface{}{"test1", "test2"}),
		},
		{
			name: "empty args",
			mocks: func(db *mocks.DBService) {
				db.EXPECT().Exec("insertUser", []interface{}{}).Return([]interface{}{}, fmt.Errorf("empty input: sp:%v, args:%v", "insertUser", []interface{}{}))
			},
			sp:          "insertUser",
			args:        []interface{}{},
			expectedErr: fmt.Errorf("empty input: sp:%v, args:%v", "insertUser", []interface{}{}),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mocks(db)
			_, err := db.Exec(tc.sp, tc.args...)
			assert.Equal(t, err, tc.expectedErr)
		})
	}
}
