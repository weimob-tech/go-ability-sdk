package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AiSaveEncryptConfig
// @id 3932
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=保存密钥)
func (client *Client) AiSaveEncryptConfig(request *AiSaveEncryptConfigRequest) (response *AiSaveEncryptConfigResponse, err error) {
	rpcResponse := CreateAiSaveEncryptConfigResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AiSaveEncryptConfigRequest struct {
	*api.BaseRequest

	Key string `json:"key,omitempty"`
}

type AiSaveEncryptConfigResponse struct {
}

func CreateAiSaveEncryptConfigRequest() (request *AiSaveEncryptConfigRequest) {
	request = &AiSaveEncryptConfigRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "ai/saveEncryptConfig", "api")
	request.Method = api.POST
	return
}

func CreateAiSaveEncryptConfigResponse() (response *api.BaseResponse[AiSaveEncryptConfigResponse]) {
	response = api.CreateResponse[AiSaveEncryptConfigResponse](&AiSaveEncryptConfigResponse{})
	return
}
