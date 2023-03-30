package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantQueryStoreListByIdList
// @id 1682
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量查询门店基础信息)
func (client *Client) MerchantQueryStoreListByIdList(request *MerchantQueryStoreListByIdListRequest) (response *MerchantQueryStoreListByIdListResponse, err error) {
	rpcResponse := CreateMerchantQueryStoreListByIdListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantQueryStoreListByIdListRequest struct {
	*api.BaseRequest

	StoreIdList []int64 `json:"storeIdList,omitempty"`
}

type MerchantQueryStoreListByIdListResponse struct {
	StoreInfoVos []int64 `json:"StoreInfoVos,omitempty"`
}

func CreateMerchantQueryStoreListByIdListRequest() (request *MerchantQueryStoreListByIdListRequest) {
	request = &MerchantQueryStoreListByIdListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchant/queryStoreListByIdList", "api")
	request.Method = api.POST
	return
}

func CreateMerchantQueryStoreListByIdListResponse() (response *api.BaseResponse[MerchantQueryStoreListByIdListResponse]) {
	response = api.CreateResponse[MerchantQueryStoreListByIdListResponse](&MerchantQueryStoreListByIdListResponse{})
	return
}
