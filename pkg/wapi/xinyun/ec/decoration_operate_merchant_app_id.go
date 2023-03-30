package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DecorationOperateMerchantAppId
// @id 1046
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=跳转外部小程序的appId配置)
func (client *Client) DecorationOperateMerchantAppId(request *DecorationOperateMerchantAppIdRequest) (response *DecorationOperateMerchantAppIdResponse, err error) {
	rpcResponse := CreateDecorationOperateMerchantAppIdResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DecorationOperateMerchantAppIdRequest struct {
	*api.BaseRequest

	OperateType   int    `json:"operateType,omitempty"`
	OriginalAppId string `json:"originalAppId,omitempty"`
	CurrentAppId  string `json:"currentAppId,omitempty"`
}

type DecorationOperateMerchantAppIdResponse struct {
}

func CreateDecorationOperateMerchantAppIdRequest() (request *DecorationOperateMerchantAppIdRequest) {
	request = &DecorationOperateMerchantAppIdRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "decoration/operateMerchantAppId", "api")
	request.Method = api.POST
	return
}

func CreateDecorationOperateMerchantAppIdResponse() (response *api.BaseResponse[DecorationOperateMerchantAppIdResponse]) {
	response = api.CreateResponse[DecorationOperateMerchantAppIdResponse](&DecorationOperateMerchantAppIdResponse{})
	return
}
