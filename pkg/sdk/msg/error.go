package msg

import (
	"fmt"
	"strings"
)

type MsgError struct {
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

func (r MsgError) GetErrcode() int64 {
	return r.Errcode
}

func (r MsgError) GetErrmsg() string {
	return r.Errmsg
}

func (r MsgError) Error() string {
	return fmt.Sprintf("RpcError{errcode: %d, errmsg: %s}", r.Errcode, r.Errmsg)
}

func NewListenerNotFoundErr(path ...string) error {
	// todo refine this
	return MsgError{
		Errcode: 90404,
		Errmsg:  "wmsg listener not found for event " + strings.Join(path, "/"),
	}
}

func NewBadRequestErr(err error) error {
	return MsgError{
		Errcode: 90400,
		Errmsg:  "bad request data " + err.Error(),
	}
}
