package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrGeneralFailure     = status.Error(codes.Internal, "ErrGeneralFailure")
	ErrProfileNotFound    = status.Error(codes.NotFound, "ErrProfileNotFound")
	ErrKnapsackNotFound   = status.Error(codes.NotFound, "ErrKnapsackNotFound")
	ErrClientParamFailure = status.Error(codes.InvalidArgument, "ErrClientParamFailure")
)
