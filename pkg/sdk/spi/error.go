package spi

import (
	"fmt"
	"strings"
)

type SpiError struct {
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

func (r SpiError) GetErrcode() int64 {
	return r.Errcode
}

func (r SpiError) GetErrmsg() string {
	return r.Errmsg
}

func (r SpiError) Error() string {
	return fmt.Sprintf("SpiError{errcode: %d, errmsg: %s}", r.Errcode, r.Errmsg)
}

func NewInvokerNotFoundErr(path ...string) error {
	// todo refine this
	return SpiError{
		Errcode: 90404,
		Errmsg:  "wspi invoker not found for event " + strings.Join(path, "/"),
	}
}

func NewBadRequestErr(err error) error {
	return SpiError{
		Errcode: 90400,
		Errmsg:  "bad request data " + err.Error(),
	}
}

func NewNoActionSpecifiedErr() error {
	return SpiError{
		Errcode: 90400,
		Errmsg:  "no action specified in request url",
	}
}
