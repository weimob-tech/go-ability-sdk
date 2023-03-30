package spi

import (
	"context"
	"github.com/weimob-tech/go-project-base/pkg/http"
	"github.com/weimob-tech/go-project-base/pkg/wlog"
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

// XinyunService 提供扩展点基础服务，例如鉴权、加解密、调用链等
type XinyunService struct {
	Service
}

func (s *XinyunService) Register(listener GenericService, iface, action string) {
	s.registry.Register(listener, SpecTypeXinyun.String(), iface, action)
}

// WosService 提供扩展点基础服务，例如鉴权、加解密、调用链等
type WosService struct {
	Service
}

func (s *WosService) Register(listener GenericService, iface, action string) {
	s.registry.Register(listener, SpecTypeWos.String(), iface, action)
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

func NewSpiCallbackConfig() *http.ExtendCallbackConfig {
	return &http.ExtendCallbackConfig{
		Method: "POST",
		Path:   "/weimob/cloud/spi/:action",
		Params: []string{"action"},
		Callback: func(ctx *http.ExtendCallbackContext) (any, error) {
			if getRegistry() == nil {
				wlog.Warnf("[weimob_spi][%s] spi service not found, request: %s", ctx.Params["action"], string(ctx.Body))
				return nil, nil
			}
			return getRegistry().Dispatch(ctx)
		},
	}
}
