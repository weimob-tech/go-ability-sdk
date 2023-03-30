package msg

import (
	"context"
	"github.com/weimob-tech/go-project-base/pkg/wlog"

	"github.com/weimob-tech/go-project-base/pkg/codec"
	"github.com/weimob-tech/go-project-base/pkg/http"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

const (
	_topic = "topic"
	_event = "event"
)

type Dispatcher func(ctx context.Context, payload []byte) (x.Result, error)

// Service 提供扩展点基础服务，例如鉴权、加解密、调用链等
type Service struct {
	// 提供扩展点生命周期管理
	*registry
	// 配置项
	config *Config
}

// XinyunService 提供扩展点基础服务，例如鉴权、加解密、调用链等
type XinyunService struct {
	Service
}

// WosService 提供扩展点基础服务，例如鉴权、加解密、调用链等
type WosService struct {
	Service
}

func (service *Service) Setup() (err error) {
	config := service.InitConfig()

	return service.InitWithOptions(config)
}

func (service *Service) InitConfig() (config *Config) {
	if service.config != nil {
		return service.config
	} else {
		return NewConfig()
	}
}

func (service *Service) InitWithOptions(config *Config) error {
	if service.registry == nil {
		service.registry = initRegistry(config)
	}
	return nil
}

func messageMetaFrom(data []byte) (topic string, event string, err error) {
	// get topic
	topic, err = codec.Json.GetString(data, _topic)
	if err != nil {
		return
	}
	// get event
	event, err = codec.Json.GetString(data, _event)
	return
}

func NewWosCallbackConfig() *http.ExtendCallbackConfig {
	return &http.ExtendCallbackConfig{
		NullToEmpty: true,
		Method:      "POST",
		Path:        "/weimob/cloud/wos/message/receive",
		Callback: func(ctx *http.ExtendCallbackContext) (any, error) {
			if getRegistry() == nil {
				wlog.Warnf("[weimob_msg] listener not found, message: %s", string(ctx.Body))
				return nil, nil
			}
			return getRegistry().Dispatch(ctx)
		},
	}
}

func NewXinyunCallbackConfig() *http.ExtendCallbackConfig {
	return &http.ExtendCallbackConfig{
		NullToEmpty: true,
		Method:      "POST",
		Path:        "/weimob/cloud/xinyun/message/receive",
		Callback: func(ctx *http.ExtendCallbackContext) (any, error) {
			if getRegistry() == nil {
				wlog.Warnf("[weimob_msg] listener not found, message: %s", string(ctx.Body))
				return nil, nil
			}
			// todo: xinyun use independent registry?
			return getRegistry().Dispatch(ctx)
		},
	}
}
