package msg

import (
	"context"
	"github.com/weimob-tech/go-project-base/pkg/wlog"

	"github.com/weimob-tech/go-project-base/pkg/http"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type Dispatcher func(ctx context.Context, payload []byte) (x.Result, error)

// Service 提供扩展点基础服务，例如鉴权、加解密、调用链等
type Service struct {
	// 提供扩展点生命周期管理
	*registry
	// 配置项
	config *Config
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

// XinyunService 提供扩展点基础服务，例如鉴权、加解密、调用链等
type XinyunService struct {
	Service
}

func (service *XinyunService) Register(listener GenericListener, path ...string) {
	service.registry.Register(listener, SpecTypeXinyun, path...)
}

// WosService 提供扩展点基础服务，例如鉴权、加解密、调用链等
type WosService struct {
	Service
}

func (service *WosService) Register(listener GenericListener, path ...string) {
	service.registry.Register(listener, SpecTypeWos, path...)
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
