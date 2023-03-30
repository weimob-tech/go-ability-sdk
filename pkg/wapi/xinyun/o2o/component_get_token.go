package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ComponentGetToken
// @id 2097
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取Token)
func (client *Client) ComponentGetToken(request *ComponentGetTokenRequest) (response *ComponentGetTokenResponse, err error) {
	rpcResponse := CreateComponentGetTokenResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ComponentGetTokenRequest struct {
	*api.BaseRequest

	ClientId     string `json:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
	Key          string `json:"key,omitempty"`
}

type ComponentGetTokenResponse struct {
}

func CreateComponentGetTokenRequest() (request *ComponentGetTokenRequest) {
	request = &ComponentGetTokenRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "component/getToken", "api")
	request.Method = api.POST
	return
}

func CreateComponentGetTokenResponse() (response *api.BaseResponse[ComponentGetTokenResponse]) {
	response = api.CreateResponse[ComponentGetTokenResponse](&ComponentGetTokenResponse{})
	return
}
