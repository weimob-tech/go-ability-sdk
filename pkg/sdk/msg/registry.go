package msg

import (
	"fmt"
	"strings"
	"sync"

	"github.com/weimob-tech/go-project-base/pkg/codec"
	"github.com/weimob-tech/go-project-base/pkg/http"
	"github.com/weimob-tech/go-project-base/pkg/wlog"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

var (
	registryInitialize sync.Once
	globalRegistry     *registry
	msgRegistryInfo    = map[SpecType]RegisterMsgInfo{}

	defaultListenerNotfound = func(path ...string) (GenericListener, error) {
		// 返回空，不抛异常，开放平台会不断重试并将消息放入失败队列
		return nil, nil
	}

	defaultOnRegistryHook = func(listener GenericListener, specType SpecType) {
		// 注册到扩展点实现注册中心
		if _, ok := msgRegistryInfo[specType]; !ok {
			msgRegistryInfo[specType] = RegisterMsgInfo{
				Path:      fmt.Sprintf("weimob/cloud/%s/message/receive", specType.String()),
				SpecsType: specType,
			}
		}
	}
)

// 提供扩展点服务注册查找等生命周期管理
type registry struct {
	store            sync.Map
	onRegistryHook   func(listener GenericListener, specType SpecType)
	listenerNotFound func(path ...string) (GenericListener, error)
}

func initRegistry(config *Config) *registry {
	registryInitialize.Do(func() {
		globalRegistry = &registry{
			listenerNotFound: config.ListenerNotFound,
			onRegistryHook:   config.OnRegistryHook,
		}
	})
	return globalRegistry
}

func getRegistry() *registry {
	return globalRegistry
}

// Lookup 查找扩展点，如果找不到则调用 registry.listenerNotFound 方法处理
func (reg *registry) Lookup(path ...string) (listener GenericListener, err error) {
	if res, ok := reg.store.Load(strings.Join(path, "/")); ok && res != nil {
		return res.(GenericListener), nil
	} else {
		return reg.listenerNotFound(path...)
	}
}

// Register 将扩展点注册到注册中心
func (reg *registry) Register(listener GenericListener, spec SpecType, path ...string) {
	// 注册
	reg.store.Store(strings.Join(path, "/"), listener)
	// 调用钩子
	if reg.onRegistryHook != nil {
		reg.onRegistryHook(listener, spec)
	}
}

// Dispatch 处理 wos 开放消息扩展点
func (reg *registry) Dispatch(c *http.ExtendCallbackContext) (result x.Result, err error) {
	var (
		ctx     = c.Context
		payload = c.Body
	)
	if len(payload) == 0 {
		return nil, nil
	}
	// 定位扩展点
	topic, event, err := messageMetaFrom(payload)
	if err != nil {
		wlog.Errorf("[weimob_msg][%s/%s]: read message meta failed, payload: %s, err: %s",
			topic, event, string(payload), err)
		err = NewBadRequestErr(err)
		return
	}
	listener, err := reg.Lookup(topic, event)
	if err != nil {
		wlog.Warnf("[weimob_msg][%s/%s] lookup listener failed, err: %s", topic, event, err)
		return
	}
	if listener == nil {
		wlog.Warnf("[weimob_msg][%s/%s] listener not found", topic, event)
		return
	}
	// 反序列化
	message := listener.NewMessage()
	err = codec.Json.Unmarshal(payload, message)
	if err != nil {
		wlog.Errorf("[weimob_msg][%s/%s]: unmarshal message failed, payload: %s, err: %s",
			topic, event, string(payload), err)
		err = NewBadRequestErr(err)
		return
	}

	return listener.OnMessage(ctx, message)
}

type msgMeta struct {
	Topic string `json:"topic"`
	Event string `json:"event"`
}

func messageMetaFrom(data []byte) (topic string, event string, err error) {
	// get topic
	var meta msgMeta
	err = codec.Json.Unmarshal(data, &meta)
	if err != nil {
		wlog.Errorf("[weimob_msg]: unmarshal message meta failed, payload: %s, err: %s", string(data), err)
		return "", "", err
	}
	if meta.Topic == "" || meta.Event == "" {
		err = x.Error("90400", fmt.Sprintf("消息Topic与Event不能为空，Topic=%s, Event=%s", meta.Topic, meta.Event))
		return
	}
	return meta.Topic, meta.Event, nil
}
