package handlers

import "errors"

var (
	roleDoesNotExist = errors.New("role does not exist").Error()
	permissionDenied = errors.New("permission denied").Error()
	errCannotVoteYet = errors.New("users can vote once per hour")
	errSelfVote      = errors.New("cannot vote for yourself")
)
