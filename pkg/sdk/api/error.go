package api

import "fmt"

type RpcError struct {
	Errcode string `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

func (r RpcError) Error() string {
	return fmt.Sprintf("RpcError{errcode: %d, errmsg: %s}", r.Errcode, r.Errmsg)
}

func NewRpcError(code, msg string) error {
	return &RpcError{code, msg}
}
