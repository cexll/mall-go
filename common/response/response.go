package response

import "net/http"

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
