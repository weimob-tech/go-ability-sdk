package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantQueryStoreInfoByStoreNumberList
// @id 1179
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=通过门店外部编码批量查询门店详情)
func (client *Client) MerchantQueryStoreInfoByStoreNumberList(request *MerchantQueryStoreInfoByStoreNumberListRequest) (response *MerchantQueryStoreInfoByStoreNumberListResponse, err error) {
	rpcResponse := CreateMerchantQueryStoreInfoByStoreNumberListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantQueryStoreInfoByStoreNumberListRequest struct {
	*api.BaseRequest

	StoreNumberList []string `json:"storeNumberList,omitempty"`
}

type MerchantQueryStoreInfoByStoreNumberListResponse struct {
}

func CreateMerchantQueryStoreInfoByStoreNumberListRequest() (request *MerchantQueryStoreInfoByStoreNumberListRequest) {
	request = &MerchantQueryStoreInfoByStoreNumberListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchant/queryStoreInfoByStoreNumberList", "api")
	request.Method = api.POST
	return
}

func CreateMerchantQueryStoreInfoByStoreNumberListResponse() (response *api.BaseResponse[MerchantQueryStoreInfoByStoreNumberListResponse]) {
	response = api.CreateResponse[MerchantQueryStoreInfoByStoreNumberListResponse](&MerchantQueryStoreInfoByStoreNumberListResponse{})
	return
}
