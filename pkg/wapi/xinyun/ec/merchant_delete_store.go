package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantDeleteStore
// @id 1426
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除门店)
func (client *Client) MerchantDeleteStore(request *MerchantDeleteStoreRequest) (response *MerchantDeleteStoreResponse, err error) {
	rpcResponse := CreateMerchantDeleteStoreResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantDeleteStoreRequest struct {
	*api.BaseRequest

	EcBizStoreId int64 `json:"ecBizStoreId,omitempty"`
}

type MerchantDeleteStoreResponse struct {
}

func CreateMerchantDeleteStoreRequest() (request *MerchantDeleteStoreRequest) {
	request = &MerchantDeleteStoreRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchant/deleteStore", "api")
	request.Method = api.POST
	return
}

func CreateMerchantDeleteStoreResponse() (response *api.BaseResponse[MerchantDeleteStoreResponse]) {
	response = api.CreateResponse[MerchantDeleteStoreResponse](&MerchantDeleteStoreResponse{})
	return
}
