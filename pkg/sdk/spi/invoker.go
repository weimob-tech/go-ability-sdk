package spi

import (
	"context"
)

// XinyunInvocationHandler 为处理SPI扩展点调用的闭包，类似Java的FunctionalInterface
type XinyunInvocationHandler[T, R any] func(ctx context.Context, request *XinyunInvocationRequest[T]) (response InvocationResponse[R], err error)

// WosInvocationHandler 为处理SPI扩展点调用的闭包，类似Java的FunctionalInterface
type WosInvocationHandler[T, R any] func(ctx context.Context, request *WosInvocationRequest[T]) (response InvocationResponse[R], err error)

// InvocationService 是SPI扩展点的接口
type InvocationService[PARAM, RESULT any] interface {
	// NewRequest 每次有新的SPI回调时返回一个新的入参实例
	NewRequest() InvocationRequest[PARAM]

	// Invoke 处理SPI扩展点回调
	Invoke(ctx context.Context, message InvocationRequest[PARAM]) (InvocationResponse[RESULT], error)
}

// GenericService 是SPI扩展点的代理方法
type GenericService interface {
	// NewRequest 新回调参数实例代理方法
	NewRequest() any
	// Invoke 处理SPI回调请求代理方法
	Invoke(context.Context, any) (any, error)
}

type delegateWrapper struct {
	delegateNewRequest func() any
	delegateInvoke     func(context.Context, any) (any, error)
}

func (d delegateWrapper) NewRequest() any {
	return d.delegateNewRequest()
}

func (d delegateWrapper) Invoke(c context.Context, request any) (any, error) {
	return d.delegateInvoke(c, request)
}

func WrapInvoker[PARAM, RESULT any](listener InvocationService[PARAM, RESULT]) GenericService {
	var delegate = delegateWrapper{}
	delegate.delegateNewRequest = func() any {
		return listener.NewRequest()
	}
	delegate.delegateInvoke = func(c context.Context, a any) (any, error) {
		request := a.(InvocationRequest[PARAM])
		return listener.Invoke(c, request)
	}
	return &delegate
}
