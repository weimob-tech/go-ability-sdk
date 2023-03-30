package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantUpdateStoreConfig
// @id 2681
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新门店设置)
func (client *Client) MerchantUpdateStoreConfig(request *MerchantUpdateStoreConfigRequest) (response *MerchantUpdateStoreConfigResponse, err error) {
	rpcResponse := CreateMerchantUpdateStoreConfigResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantUpdateStoreConfigRequest struct {
	*api.BaseRequest
}

type MerchantUpdateStoreConfigResponse struct {
}

func CreateMerchantUpdateStoreConfigRequest() (request *MerchantUpdateStoreConfigRequest) {
	request = &MerchantUpdateStoreConfigRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchant/updateStoreConfig", "api")
	request.Method = api.POST
	return
}

func CreateMerchantUpdateStoreConfigResponse() (response *api.BaseResponse[MerchantUpdateStoreConfigResponse]) {
	response = api.CreateResponse[MerchantUpdateStoreConfigResponse](&MerchantUpdateStoreConfigResponse{})
	return
}
