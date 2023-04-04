package api

import (
	"context"
	"fmt"
	"io"
	"reflect"

	"github.com/weimob-tech/go-project-base/pkg/auth"
	"github.com/weimob-tech/go-project-base/pkg/codec"
	"github.com/weimob-tech/go-project-base/pkg/http"
)

type Client struct {
	Domain     string
	config     *Config
	httpClient http.Client
}

type XinyunClient struct {
	Client
}

type WosClient struct {
	Client
}

type GenericUploadRequest struct {
	super uintptr
	Name  string
	File  io.Reader
}

func (client *Client) InitWithAccessKey(clientId, clientSecret string) (err error) {
	config := client.InitClientConfig()
	credential := &auth.ClientInfo{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}
	return client.InitWithOptions(config, credential)
}

func (client *Client) InitClientConfig() (config *Config) {
	if client.config != nil {
		return client.config
	}
	return NewConfig()
}

func (client *Client) InitWithOptions(config *Config, clientInfo *auth.ClientInfo) (err error) {
	config.clientInfo = clientInfo
	client.config = config
	client.httpClient = http.NewHttpClient()

	// todo setup timeout async
	// if config.Timeout > 0 {
	// 	client.httpClient.Timeout = config.Timeout
	// }
	// todo setup async
	// if config.EnableAsync {
	// 	client.EnableAsync(config.GoRoutinePoolSize, config.MaxTaskQueueSize)
	// }
	// todo setup signer
	// client.signer, err = auth.NewSignerWithCredential(credential, client.SignCallback)
	return
}

func (client *Client) DoAction(request RpcRequest, response RpcResponse) (err error) {
	switch request.GetRPCType() {
	case RPCTypeJson:
		return client.doJsonAction(request, response)
	case RPCTypeMultiPart:
		return client.doMultipartAction(request, response)
	default:
		return NewRpcError("90500", "unknown rpc type")
	}
}

func (client *Client) doJsonAction(request RpcRequest, response RpcResponse) (err error) {
	req := client.buildRequest(request)
	res := client.httpClient.NewResponse()

	if len(request.GetContent()) == 0 {
		// 没有指定 content，用 request 整个作为入参
		payload, err := codec.Json.Marshal(request)
		if err != nil {
			return NewRpcError("90500", fmt.Sprintf("json marshal error, %s", err.Error()))
		}
		req.SetBody(payload)
	} else {
		//  指定了 content，用 content 作为入参
		req.SetBody(request.GetContent())
	}

	err = client.httpClient.Do(context.Background(), req, res)
	if err != nil {
		return
	}

	err = response.Unmarshal(res)
	return
}

func (client *Client) doMultipartAction(request RpcRequest, response RpcResponse) (err error) {
	req := client.buildRequest(request)
	res := client.httpClient.NewResponse()

	fileInfo := (*GenericUploadRequest)(reflect.ValueOf(request).UnsafePointer())
	req.SetFile("file", fileInfo.Name, fileInfo.File)

	err = client.httpClient.Do(context.Background(), req, res)
	if err != nil {
		return
	}

	err = response.Unmarshal(res)
	return
}

func (client *Client) buildRequest(request RpcRequest) (req http.Request) {
	req = client.httpClient.NewRequest()
	req.GetHeader().SetMethod(request.GetMethod())
	req.GetHeader().SetContentTypeBytes(request.GetContentTypeByte())
	req.SetRequestURI(request.BuildUrl(client.config))

	query := request.GetQueryParams()
	if _, ok := query["accesstoken"]; !ok {
		// token 不存在，自动注入
		query["accesstoken"] = auth.GetClientCCToken(
			context.Background(),
			client.config.clientInfo.ClientId,
			client.config.clientInfo.ClientSecret,
			request.GetShopId(),
			request.GetShopType())
	}

	req.SetQueryString(request.BuildQueries())
	return
}
