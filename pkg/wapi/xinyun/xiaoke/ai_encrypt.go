package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AiEncrypt
// @id 3933
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=加密)
func (client *Client) AiEncrypt(request *AiEncryptRequest) (response *AiEncryptResponse, err error) {
	rpcResponse := CreateAiEncryptResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AiEncryptRequest struct {
	*api.BaseRequest

	SourceList []map[string]any `json:"sourceList,omitempty"`
}

type AiEncryptResponse struct {
	EncryptionList []AiEncryptResponseEncryptionList `json:"encryptionList,omitempty"`
}

func CreateAiEncryptRequest() (request *AiEncryptRequest) {
	request = &AiEncryptRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "ai/encrypt", "api")
	request.Method = api.POST
	return
}

func CreateAiEncryptResponse() (response *api.BaseResponse[AiEncryptResponse]) {
	response = api.CreateResponse[AiEncryptResponse](&AiEncryptResponse{})
	return
}

type AiEncryptResponseEncryptionList struct {
	Source     string `json:"source,omitempty"`
	Result     bool   `json:"result,omitempty"`
	Encryption string `json:"encryption,omitempty"`
	Msg        string `json:"msg,omitempty"`
}
