package msg

import (
	"context"
	"github.com/weimob-tech/go-project-base/pkg/codec"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

// XinyunEventHandler 为处理扩展点调用的闭包，类似Java的FunctionalInterface
type XinyunEventHandler[T any] func(ctx context.Context, message *XinyunOpenMessage[T]) (response x.Result, err error)

// WosEventHandler 为处理扩展点调用的闭包，类似Java的FunctionalInterface
type WosEventHandler[T any] func(ctx context.Context, messgae *WosOpenMessage[T]) (response x.Result, err error)

type EventListener[EVENT any] interface {
	// NewMessage 每次有新的消息回调时返回一个新的消息实例
	NewMessage() MsgRequest[EVENT]

	// OnMessage 处理消息回调
	OnMessage(ctx context.Context, message MsgRequest[EVENT]) (x.Result, error)
}

// GenericListener 是开放消息代理方法
type GenericListener interface {
	// NewMessage 新消息实例代理方法
	NewMessage() any
	// OnMessage 消息处理回调代理方法
	OnMessage(context.Context, any) (x.Result, error)
}

type delegateWrapper struct {
	delegateNewMessage func() any
	delegateOnMessage  func(context.Context, any) (x.Result, error)
}

func (d delegateWrapper) NewMessage() any {
	return d.delegateNewMessage()
}

func (d delegateWrapper) OnMessage(c context.Context, msg any) (x.Result, error) {
	return d.delegateOnMessage(c, msg)
}

func WrapListener[EVENT any](listener EventListener[EVENT]) GenericListener {
	var delegate = delegateWrapper{}
	delegate.delegateNewMessage = func() any {
		return listener.NewMessage()
	}
	delegate.delegateOnMessage = func(c context.Context, a any) (x.Result, error) {
		msg := a.(MsgRequest[EVENT])
		if len(msg.GetRawMsgBody()) > 0 {
			err := codec.Json.UnmarshalString(msg.GetRawMsgBody(), msg.GetMsgBody())
			if err != nil {
				return nil, NewBadRequestErr(err)
			}
		}
		return listener.OnMessage(c, msg)
	}
	return &delegate
}
