package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerAssetAuthorizationGet
// @id 1653
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1653?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取资产核销客户授权配置信息)
func (client *Client) CustomerAssetAuthorizationGet(request *CustomerAssetAuthorizationGetRequest) (response *CustomerAssetAuthorizationGetResponse, err error) {
	rpcResponse := CreateCustomerAssetAuthorizationGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerAssetAuthorizationGetRequest struct {
	*api.BaseRequest

	ConfigTypeList []int64 `json:"configTypeList,omitempty"`
	VidType        int64   `json:"vidType,omitempty"`
	Vid            int64   `json:"vid,omitempty"`
	Wid            int64   `json:"wid,omitempty"`
}

type CustomerAssetAuthorizationGetResponse struct {
	AssetSwitchConfigList []CustomerAssetAuthorizationGetResponseAssetSwitchConfigList `json:"assetSwitchConfigList,omitempty"`
}

func CreateCustomerAssetAuthorizationGetRequest() (request *CustomerAssetAuthorizationGetRequest) {
	request = &CustomerAssetAuthorizationGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "customer/asset/authorization/get", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerAssetAuthorizationGetResponse() (response *api.BaseResponse[CustomerAssetAuthorizationGetResponse]) {
	response = api.CreateResponse[CustomerAssetAuthorizationGetResponse](&CustomerAssetAuthorizationGetResponse{})
	return
}

type CustomerAssetAuthorizationGetResponseAssetSwitchConfigList struct {
	Type   int64 `json:"type,omitempty"`
	IsOpen bool  `json:"isOpen,omitempty"`
}
