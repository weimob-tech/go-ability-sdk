package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantUpdateStoreStatus
// @id 1810
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=启用/停用门店)
func (client *Client) MerchantUpdateStoreStatus(request *MerchantUpdateStoreStatusRequest) (response *MerchantUpdateStoreStatusResponse, err error) {
	rpcResponse := CreateMerchantUpdateStoreStatusResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantUpdateStoreStatusRequest struct {
	*api.BaseRequest

	StoreId     int64 `json:"storeId,omitempty"`
	StoreStatus int   `json:"storeStatus,omitempty"`
}

type MerchantUpdateStoreStatusResponse struct {
}

func CreateMerchantUpdateStoreStatusRequest() (request *MerchantUpdateStoreStatusRequest) {
	request = &MerchantUpdateStoreStatusRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchant/updateStoreStatus", "api")
	request.Method = api.POST
	return
}

func CreateMerchantUpdateStoreStatusResponse() (response *api.BaseResponse[MerchantUpdateStoreStatusResponse]) {
	response = api.CreateResponse[MerchantUpdateStoreStatusResponse](&MerchantUpdateStoreStatusResponse{})
	return
}
