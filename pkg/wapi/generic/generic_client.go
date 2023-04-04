package generic

import (
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"
	"github.com/weimob-tech/go-project-base/pkg/codec"
)

// DoRequest 允许调用者自定义出入参
func (client *Client) DoRequest(request api.RpcRequest, response api.RpcResponse) (err error) {
	return client.DoAction(request, response)
}

type RawContentRequest struct {
	*api.BaseRequest

	content []byte
}

func (request *RawContentRequest) SetBody(payload any) error {
	// raw slice
	if data, ok := payload.([]byte); ok {
		request.content = data
		return nil
	}
	// other type
	var data, err = codec.Json.Marshal(payload)
	if err != nil {
		return err
	}
	request.content = data
	return nil
}

func (request *RawContentRequest) GetContent() []byte {
	return request.content
}

func NewPostRequest(path string) (request *RawContentRequest) {
	request = &RawContentRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithPath(path)
	request.Method = api.POST
	return
}
