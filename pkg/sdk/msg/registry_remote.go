package msg

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/auth"
	ba "github.com/weimob-tech/go-project-base/pkg/auth"
	"github.com/weimob-tech/go-project-base/pkg/codec"
	"github.com/weimob-tech/go-project-base/pkg/config"
	"github.com/weimob-tech/go-project-base/pkg/hook"
	"github.com/weimob-tech/go-project-base/pkg/http"
	"github.com/weimob-tech/go-project-base/pkg/wlog"
	"strconv"
	"strings"
	"time"
)

type SpecType int

const (
	SpecTypeXinyun SpecType = iota + 1
	SpecTypeWos
)

func (s SpecType) String() string {
	switch s {
	case 1:
		return "xinyun"
	case 2:
		return "wos"
	default:
		return "unknown"
	}
}

const (
	SdkVersionV1 = "v1"
	SdkVersionV2 = "v2"
)

func init() {
	config.AddStoreSetHook(func() {
		config.SetDefault("msg.sdk.version", SdkVersionV2)
		_ = config.BindEnv("msg.register.url", "postUrl")
	})
	hook.AddPostStartHook(func() {
		err := RegistryRemote()
		if err != nil {
			wlog.Errorf("[weimob_msg] 信息扩展点注册失败: %v", err)
		}
	})
}

type RegisterRequest struct {
	Sign      string `json:"sign"`
	Timestamp string `json:"timestamp"`
	ClientId  string `json:"clientId"`
	AppId     string `json:"appId"`
	Env       string `json:"env"`
	Params    string `json:"params"`
}

type RegisterInfo struct {
	SdkVersion       string          `json:"sdkVersion"`
	ClientId         string          `json:"clientId"`
	RegisterMsgInfo  RegisterMsgInfo `json:"messageExtensionDTO"`
	InterfacePathVos []any           `json:"interfacePathVos"`
}

type RegisterMsgInfo struct {
	ClientId    string   `json:"clientId"`
	HostAddress string   `json:"hostAddress"`
	Path        string   `json:"path"`
	SpecsType   SpecType `json:"specsType"`
}

func RegistryRemote() (err error) {
	if len(msgRegistryInfo) == 0 {
		// 跳过注册
		return
	}

	var clientInfoMap = auth.CreateClientInfo()
	for _, clientInfo := range clientInfoMap {
		// 为多个 client 注册扩展点实现
		for _, msgInfo := range msgRegistryInfo {
			register := RegisterInfo{RegisterMsgInfo: msgInfo, InterfacePathVos: []any{}}
			err = RegistryRemoteForClient(register, *clientInfo)
			if err != nil {
				return
			}
		}
	}
	return
}

func RegistryRemoteForClient(register RegisterInfo, clientInfo ba.ClientInfo) (err error) {
	client := http.NewHttpClient()

	register.ClientId = clientInfo.ClientId
	register.RegisterMsgInfo.ClientId = clientInfo.ClientId
	register.SdkVersion = config.GetString("msg.register.version")

	appId := config.GetString("weimob_cloud_appId")
	env := config.GetString("weimob_cloud_env")
	param := codec.ToJson[RegisterInfo](&register)
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)

	var sb strings.Builder
	sb.WriteString(clientInfo.ClientId)
	sb.WriteString(clientInfo.ClientSecret)
	sb.WriteString(timestamp)
	sb.WriteString(codec.Md5String(param))

	var sign = codec.Md5String(sb.String())
	var request = &RegisterRequest{
		Sign:      sign,
		Timestamp: timestamp,
		ClientId:  clientInfo.ClientId,
		AppId:     appId,
		Env:       env,
		Params:    param,
	}
	return doRegister(client, request)
}

func doRegister(client http.Client, request *RegisterRequest) (err error) {
	wlog.Infof("[weimob_msg] 注册信息: %s", codec.ToJson[RegisterRequest](request))
	httpRes := client.NewResponse()
	httpReq := client.NewRequest()
	httpReq.GetHeader().SetMethod("POST")
	httpReq.GetHeader().SetContentTypeBytes(http.ContentTypeJsonByte)
	httpReq.SetBody(codec.ToJsonByte[RegisterRequest](request))
	httpReq.SetRequestURI(config.GetString("msg.register.url"))

	err = client.Do(context.Background(), httpReq, httpRes)
	if err != nil {
		wlog.Errorf("[weimob_msg] 注册信息失败: %s，%s", codec.ToJson[RegisterRequest](request), err)
		return
	}
	wlog.Infof("[weimob_msg] 注册返回信息: %s", string(httpRes.Body()))
	return
}
