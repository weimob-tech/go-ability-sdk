package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NicheCancels
// @id 1736
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除商机)
func (client *Client) NicheCancels(request *NicheCancelsRequest) (response *NicheCancelsResponse, err error) {
	rpcResponse := CreateNicheCancelsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NicheCancelsRequest struct {
	*api.BaseRequest

	Keys []map[string]any `json:"keys,omitempty"`
	Wid  string           `json:"wid,omitempty"`
}

type NicheCancelsResponse struct {
	Message string `json:"message,omitempty"`
}

func CreateNicheCancelsRequest() (request *NicheCancelsRequest) {
	request = &NicheCancelsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "niche/cancels", "api")
	request.Method = api.POST
	return
}

func CreateNicheCancelsResponse() (response *api.BaseResponse[NicheCancelsResponse]) {
	response = api.CreateResponse[NicheCancelsResponse](&NicheCancelsResponse{})
	return
}
