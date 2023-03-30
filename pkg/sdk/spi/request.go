package spi

import (
	"github.com/weimob-tech/go-project-base/pkg/codec"
	"github.com/weimob-tech/go-project-base/pkg/wlog"
)

type InvocationUnmarshaler interface {
	Unmarshal() error
}

type InvocationRequest[T any] interface {
	GetParams() *T
}

type baseRequest struct {
	ParamsRaw string `json:"params,omitempty"`
}

type WosInvocationRequest[T any] struct {
	baseRequest
	BosId      int64  `json:"bosId,omitempty"`
	ActionKey  string `json:"actionKey,omitempty"`
	FunctionId int64  `json:"functionId,omitempty"`
	Vid        int64  `json:"vid,omitempty"`
	VType      int64  `json:"vType,omitempty"`
	Params     *T     `json:"-"`
}

func (r *WosInvocationRequest[T]) GetParams() *T {
	return r.Params
}

func (r *WosInvocationRequest[T]) Unmarshal() error {
	if len(r.ParamsRaw) == 0 {
		return nil
	}
	r.Params = new(T)
	err := codec.Json.UnmarshalString(r.ParamsRaw, r.Params)
	if err != nil {
		wlog.Errorf("[weimob_spi]: unmarshal request failed, params: %s, err: %v", r.ParamsRaw, err)
	}
	return err
}

type XinyunInvocationRequest[T any] struct {
	baseRequest
	Sign                    string `json:"sign,omitempty"`
	Timestamp               string `json:"timestamp,omitempty"`
	Pid                     int64  `json:"pid,omitempty"`
	BosId                   int64  `json:"bosId,omitempty"`
	ActionKey               string `json:"actionKey,omitempty"`
	Vid                     int64  `json:"vid,omitempty"`
	VType                   int    `json:"vType,omitempty"`
	OriginProductId         int64  `json:"originProductId,omitempty"`
	TargetProductId         int64  `json:"targetProductId,omitempty"`
	OriginProductInstanceId int64  `json:"originProductInstanceId,omitempty"`
	TargetProductInstanceId int64  `json:"targetProductInstanceId,omitempty"`
	FunctionId              int64  `json:"functionId,omitempty"`
	Params                  *T     `json:"-"`
}

func (r *XinyunInvocationRequest[T]) GetParams() *T {
	return r.Params
}

func (r *XinyunInvocationRequest[T]) Unmarshal() error {
	if len(r.ParamsRaw) == 0 {
		return nil
	}
	r.Params = new(T)
	err := codec.Json.UnmarshalString(r.ParamsRaw, r.Params)
	if err != nil {
		wlog.Errorf("[weimob_spi]: unmarshal request failed, params: %s, err: %v", r.ParamsRaw, err)
	}
	return err
}
