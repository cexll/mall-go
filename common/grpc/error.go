package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mall-go/app/user/cmd/pb"
)

func GrpcError(rpcError error, errCode int32) error {
	errStatus, err := status.New(codes.Aborted, rpcError.Error()).WithDetails(&pb.Error{
		Code:    errCode,
		Message: rpcError.Error(),
	})
	if err != nil {
		return status.New(codes.Aborted, rpcError.Error()).Err()
	}
	return errStatus.Err()
}
