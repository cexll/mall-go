package response

import (
	"google.golang.org/grpc/status"
	"mall-go/common/mrpc/pb"
	"net/http"
)

type Response struct {
	Code    uint64 `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (t *Response) Success(data any) Response {
	return Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	}
}

func (t *Response) Fail(message string) Response {
	return Response{
		Code:    http.StatusInternalServerError,
		Message: message,
		Data:    nil,
	}
}

func (t *Response) Error(err error) Response {
	if rpcErr, ok := status.FromError(err); ok {
		ditails := rpcErr.Details()
		pbError := ditails[0].(*pb.Error)
		return Response{
			Code:    uint64(pbError.Code),
			Message: pbError.Message,
			Data:    nil,
		}
	}
	return Response{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
		Data:    nil,
	}
}
