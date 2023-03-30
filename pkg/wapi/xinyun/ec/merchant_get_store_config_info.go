package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantGetStoreConfigInfo
// @id 1495
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询门店配置信息)
func (client *Client) MerchantGetStoreConfigInfo(request *MerchantGetStoreConfigInfoRequest) (response *MerchantGetStoreConfigInfoResponse, err error) {
	rpcResponse := CreateMerchantGetStoreConfigInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantGetStoreConfigInfoRequest struct {
	*api.BaseRequest

	StoreId int64 `json:"storeId,omitempty"`
}

type MerchantGetStoreConfigInfoResponse struct {
}

func CreateMerchantGetStoreConfigInfoRequest() (request *MerchantGetStoreConfigInfoRequest) {
	request = &MerchantGetStoreConfigInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchant/getStoreConfigInfo", "api")
	request.Method = api.POST
	return
}

func CreateMerchantGetStoreConfigInfoResponse() (response *api.BaseResponse[MerchantGetStoreConfigInfoResponse]) {
	response = api.CreateResponse[MerchantGetStoreConfigInfoResponse](&MerchantGetStoreConfigInfoResponse{})
	return
}
