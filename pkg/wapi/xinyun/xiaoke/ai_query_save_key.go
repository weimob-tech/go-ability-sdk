package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AiQuerySaveKey
// @id 3934
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=是否保存密钥)
func (client *Client) AiQuerySaveKey(request *AiQuerySaveKeyRequest) (response *AiQuerySaveKeyResponse, err error) {
	rpcResponse := CreateAiQuerySaveKeyResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AiQuerySaveKeyRequest struct {
	*api.BaseRequest
}

type AiQuerySaveKeyResponse struct {
}

func CreateAiQuerySaveKeyRequest() (request *AiQuerySaveKeyRequest) {
	request = &AiQuerySaveKeyRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "ai/querySaveKey", "api")
	request.Method = api.POST
	return
}

func CreateAiQuerySaveKeyResponse() (response *api.BaseResponse[AiQuerySaveKeyResponse]) {
	response = api.CreateResponse[AiQuerySaveKeyResponse](&AiQuerySaveKeyResponse{})
	return
}
