package spi

import (
	"fmt"
	"sync"

	"github.com/weimob-tech/go-project-base/pkg/codec"
	"github.com/weimob-tech/go-project-base/pkg/http"
	"github.com/weimob-tech/go-project-base/pkg/wlog"
)

var (
	registryInitialize sync.Once
	globalRegistry     *registry
	spiRegistryInfo    []RegisterSpiInfo

	defaultInvokerNotfound = func(action string) (GenericService, error) {
		// 返回空，不抛异常，开放平台会不断重试并将消息放入失败队列
		return nil, nil
	}

	defaultOnRegistryHook = func(listener GenericService, spec, iface, action string) {
		// 注册到扩展点实现注册中心
		spiRegistryInfo = append(spiRegistryInfo, RegisterSpiInfo{
			InterfaceName: iface,
			ImplFullName:  fmt.Sprintf("go.impl.%s.%s", spec, iface),
			BeanName:      fmt.Sprintf("%s.%s", spec, iface),
			MethodName:    "excute",
			SpiBelongType: SpecTypeWos,
		})
	}
)

// 提供扩展点服务注册查找等生命周期管理
type registry struct {
	store            sync.Map
	onRegistryHook   func(listener GenericService, spec, iface, action string)
	listenerNotFound func(action string) (GenericService, error)
}

func initRegistry(config *Config) *registry {
	registryInitialize.Do(func() {
		globalRegistry = &registry{
			listenerNotFound: config.InvokerNotFound,
			onRegistryHook:   config.OnRegistryHook,
		}
	})
	return globalRegistry
}

func getRegistry() *registry {
	return globalRegistry
}

// Lookup 查找扩展点，如果找不到则调用 registry.listenerNotFound 方法处理
func (reg *registry) Lookup(action string) (listener GenericService, err error) {
	if res, ok := reg.store.Load(action); ok && res != nil {
		return res.(GenericService), nil
	} else {
		return reg.listenerNotFound(action)
	}
}

// Register 将扩展点注册到注册中心
func (reg *registry) Register(listener GenericService, spec, iface, action string) {
	// 注册
	reg.store.Store(spec+"."+iface, listener)
	// 调用钩子
	if reg.onRegistryHook != nil {
		reg.onRegistryHook(listener, spec, iface, action)
	}
}

// Dispatch 处理 wos 开放消息扩展点
func (reg *registry) Dispatch(c *http.ExtendCallbackContext) (result any, err error) {
	var (
		ctx     = c.Context
		payload = c.Body
	)
	if len(payload) == 0 {
		return nil, nil
	}
	// 定位扩展点
	var service GenericService
	if action, ok := c.Params["action"]; ok {
		// wos.InterfaceName
		service, err = reg.Lookup(action)
		if err != nil {
			wlog.Warnf("[weimob_spi][%s] spi service not found, err: %s", action, err)
			return nil, err
		}
		if service == nil {
			wlog.Warnf("[weimob_spi][%s] spi service not found", action)
			return
		}
	} else {
		wlog.Errorf("[weimob_spi][%s]: no action specified in url: %s", action, c.Path)
		return nil, NewNoActionSpecifiedErr()
	}
	// 反序列化
	request := service.NewRequest()
	err = codec.Json.Unmarshal(payload, request)
	if err != nil {
		wlog.Errorf("[weimob_spi]: unmarshal request failed, err: %s, payload: %s", err, string(payload))
		err = NewBadRequestErr(err)
		return
	}
	err = request.(InvocationUnmarshaler).Unmarshal()
	if err != nil {
		err = NewBadRequestErr(err)
		return
	}

	return service.Invoke(ctx, request)
}
