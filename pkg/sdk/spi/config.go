package spi

// Config 配置项
// InvokerNotFound 扩展点未找到时的处理方法
// OnRegistryHook 注册扩展点时的回调方法
type Config struct {
	InvokerNotFound func(action string) (GenericService, error)
	OnRegistryHook  func(listener GenericService, spec, iface, action string)
}

func NewConfig() *Config {
	return &Config{
		InvokerNotFound: defaultInvokerNotfound,
		OnRegistryHook:  defaultOnRegistryHook,
	}
}
