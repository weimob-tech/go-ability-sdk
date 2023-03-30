package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreQueryStoreList
// @id 47
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询门店列表)
func (client *Client) StoreQueryStoreList(request *StoreQueryStoreListRequest) (response *StoreQueryStoreListResponse, err error) {
	rpcResponse := CreateStoreQueryStoreListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreQueryStoreListRequest struct {
	*api.BaseRequest

	StoreStatus int `json:"storeStatus,omitempty"`
	Page        int `json:"page,omitempty"`
	PageSize    int `json:"pageSize,omitempty"`
	PageNum     int `json:"pageNum,omitempty"`
}

type StoreQueryStoreListResponse struct {
}

func CreateStoreQueryStoreListRequest() (request *StoreQueryStoreListRequest) {
	request = &StoreQueryStoreListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "store/queryStoreList", "api")
	request.Method = api.POST
	return
}

func CreateStoreQueryStoreListResponse() (response *api.BaseResponse[StoreQueryStoreListResponse]) {
	response = api.CreateResponse[StoreQueryStoreListResponse](&StoreQueryStoreListResponse{})
	return
}
