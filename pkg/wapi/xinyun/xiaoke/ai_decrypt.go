package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AiDecrypt
// @id 3935
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=解密)
func (client *Client) AiDecrypt(request *AiDecryptRequest) (response *AiDecryptResponse, err error) {
	rpcResponse := CreateAiDecryptResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AiDecryptRequest struct {
	*api.BaseRequest

	EncryptionList []map[string]any `json:"encryptionList,omitempty"`
}

type AiDecryptResponse struct {
	SourceList []AiDecryptResponseSourceList `json:"sourceList,omitempty"`
}

func CreateAiDecryptRequest() (request *AiDecryptRequest) {
	request = &AiDecryptRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "ai/decrypt", "api")
	request.Method = api.POST
	return
}

func CreateAiDecryptResponse() (response *api.BaseResponse[AiDecryptResponse]) {
	response = api.CreateResponse[AiDecryptResponse](&AiDecryptResponse{})
	return
}

type AiDecryptResponseSourceList struct {
	Source     string `json:"source,omitempty"`
	Result     bool   `json:"result,omitempty"`
	Encryption string `json:"encryption,omitempty"`
	Msg        string `json:"msg,omitempty"`
}
