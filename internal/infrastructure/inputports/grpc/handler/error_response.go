package handler

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ResponseWithError(statusCode codes.Code, err error) error {
	return status.Error(statusCode, err.Error())
}
