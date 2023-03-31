package api

import (
	"github.com/weimob-tech/go-project-base/pkg/codec"
	"github.com/weimob-tech/go-project-base/pkg/http"
)

type RpcResponse interface {
	IsSuccess() bool
	GetHttpStatus() int
	GetHttpContentBytes() []byte
	Unmarshal(http.Response) error
}

type Code struct {
	Errcode string `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

type baseResponse struct {
	GlobalTicket string `json:"globalTicket,omitempty"`
	Code         Code   `json:"code,omitempty"`
}

type BaseResponse[T any] struct {
	*baseResponse
	httpStatus       int
	httpContentBytes []byte
	Data             *T
}

func CreateResponse[T any](data *T) (response *BaseResponse[T]) {
	response = &BaseResponse[T]{
		baseResponse: &baseResponse{},
		Data:         data,
	}
	return
}

func (response *BaseResponse[T]) Unmarshal(res http.Response) (err error) {
	response.httpStatus = res.StatusCode()
	response.httpContentBytes = res.Body()
	err = codec.Json.Unmarshal(response.httpContentBytes, response)
	return
}

func (response *BaseResponse[T]) GetHttpStatus() int {
	return response.httpStatus
}

func (response *BaseResponse[T]) GetHttpContentBytes() []byte {
	return response.httpContentBytes
}

func (response *BaseResponse[T]) IsSuccess() bool {
	if response.GetHttpStatus() >= 200 && response.GetHttpStatus() < 300 {
		return true
	}
	return false
}

func (response *BaseResponse[T]) GetData() (data *T, err error) {
	if response.Code.Errcode == "0" {
		data = response.Data
		return
	}
	return nil, NewRpcError(response.Code.Errcode, response.Code.Errmsg)
}
