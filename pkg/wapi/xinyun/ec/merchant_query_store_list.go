package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantQueryStoreList
// @id 539
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取商户的门店分页列表)
func (client *Client) MerchantQueryStoreList(request *MerchantQueryStoreListRequest) (response *MerchantQueryStoreListResponse, err error) {
	rpcResponse := CreateMerchantQueryStoreListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantQueryStoreListRequest struct {
	*api.BaseRequest

	PageNum          int    `json:"pageNum,omitempty"`
	PageSize         int    `json:"pageSize,omitempty"`
	ContainHeadStore int    `json:"containHeadStore,omitempty"`
	NeedDeleteStore  int    `json:"needDeleteStore,omitempty"`
	StoreName        string `json:"storeName,omitempty"`
}

type MerchantQueryStoreListResponse struct {
	PageList      []int64 `json:"pageList,omitempty"`
	HeadStoreInfo []int64 `json:"headStoreInfo,omitempty"`
	PageNum       int64   `json:"pageNum,omitempty"`
	TotalCount    int64   `json:"totalCount,omitempty"`
	PageSize      int64   `json:"pageSize,omitempty"`
	StoreNumber   string  `json:"storeNumber,omitempty"`
}

func CreateMerchantQueryStoreListRequest() (request *MerchantQueryStoreListRequest) {
	request = &MerchantQueryStoreListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchant/queryStoreList", "api")
	request.Method = api.POST
	return
}

func CreateMerchantQueryStoreListResponse() (response *api.BaseResponse[MerchantQueryStoreListResponse]) {
	response = api.CreateResponse[MerchantQueryStoreListResponse](&MerchantQueryStoreListResponse{})
	return
}
