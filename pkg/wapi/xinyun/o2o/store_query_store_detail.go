package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreQueryStoreDetail
// @id 48
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询门店详情)
func (client *Client) StoreQueryStoreDetail(request *StoreQueryStoreDetailRequest) (response *StoreQueryStoreDetailResponse, err error) {
	rpcResponse := CreateStoreQueryStoreDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreQueryStoreDetailRequest struct {
	*api.BaseRequest

	StoreIds string `json:"storeIds,omitempty"`
}

type StoreQueryStoreDetailResponse struct {
}

func CreateStoreQueryStoreDetailRequest() (request *StoreQueryStoreDetailRequest) {
	request = &StoreQueryStoreDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "store/queryStoreDetail", "api")
	request.Method = api.POST
	return
}

func CreateStoreQueryStoreDetailResponse() (response *api.BaseResponse[StoreQueryStoreDetailResponse]) {
	response = api.CreateResponse[StoreQueryStoreDetailResponse](&StoreQueryStoreDetailResponse{})
	return
}
