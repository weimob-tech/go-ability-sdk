package auth

import (
	"github.com/weimob-tech/go-project-base/pkg/auth"
	"github.com/weimob-tech/go-project-base/pkg/config"
	"strings"
)

const (
	ClientPrefix       = "weimob.cloud"
	ClientIdSuffix     = ".clientId"
	ClientSecretSuffix = ".clientSecret"
)

func init() {
	config.AddStoreSetHook(func() {
		_ = config.BindEnv("weimob_cloud_appId", "weimob_cloud_appId")
		_ = config.BindEnv("weimob_cloud_env", "weimob_cloud_env")
	})
}

// CreateClientInfo 返回一个
// weimob.cloud.foo.clientId = clientId 组装成一个 map
// weimob.cloud.foo.clientSecret = clientSecret 组装成一个 map
func CreateClientInfo() (clients map[string]*auth.ClientInfo) {
	var clientInfoMap = config.GetStringMap(ClientPrefix)
	if len(clientInfoMap) == 0 {
		return
	}

	clients = make(map[string]*auth.ClientInfo)
	// extract client id
	for k, v := range clientInfoMap {
		// key.clientId 格式
		if val, ok := v.(string); ok {
			if strings.HasSuffix(k, ClientIdSuffix) {
				key := strings.TrimSuffix(k, ClientIdSuffix)
				if client, ok := clients[key]; ok {
					client.ClientId = val
				} else {
					clients[key] = &auth.ClientInfo{ClientId: val}
				}
			}
			if strings.HasSuffix(k, ClientSecretSuffix) {
				key := strings.TrimSuffix(k, ClientSecretSuffix)
				if client, ok := clients[key]; ok {
					client.ClientSecret = val
				} else {
					clients[key] = &auth.ClientInfo{ClientSecret: val}
				}
			}
		}
		// key -> {clientId} 格式
		if val, ok := v.(map[string]interface{}); ok {
			c := &auth.ClientInfo{}
			for kk, vv := range val {
				if strings.ToLower(kk) == "clientid" {
					c.ClientId = vv.(string)
				}
				if strings.ToLower(kk) == "clientsecret" {
					c.ClientSecret = vv.(string)
				}
			}
			clients[k] = c
		}
	}
	return
}
