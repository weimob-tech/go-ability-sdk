package msg

// Config 配置项
// ListenerNotFound 扩展点未找到时的处理方法
// OnRegistryHook 注册扩展点时的回调方法
type Config struct {
	ListenerNotFound func(path ...string) (GenericListener, error)
	OnRegistryHook   func(listener GenericListener, path ...string)
}

func NewConfig() *Config {
	return &Config{
		ListenerNotFound: defaultListenerNotfound,
		OnRegistryHook:   defaultOnRegistryHook,
	}
}
