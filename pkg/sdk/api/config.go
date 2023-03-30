package api

import (
	"github.com/weimob-tech/go-project-base/pkg/auth"
	"github.com/weimob-tech/go-project-base/pkg/config"
)

type Config struct {
	clientInfo    *auth.ClientInfo
	defaultSchema string
	defaultDomain string
}

func NewConfig() (conf *Config) {
	conf = &Config{
		defaultSchema: config.GetString("client.schema"),
		defaultDomain: config.GetString("client.domain"),
	}
	return
}
