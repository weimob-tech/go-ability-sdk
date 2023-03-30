package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreGetStoreList
// @id 507
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=门店列表/酒店列表)
func (client *Client) StoreGetStoreList(request *StoreGetStoreListRequest) (response *StoreGetStoreListResponse, err error) {
	rpcResponse := CreateStoreGetStoreListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreGetStoreListRequest struct {
	*api.BaseRequest

	QueryName string `json:"queryName,omitempty"`
	PageIndex int    `json:"pageIndex,omitempty"`
	PageSize  int    `json:"pageSize,omitempty"`
}

type StoreGetStoreListResponse struct {
}

func CreateStoreGetStoreListRequest() (request *StoreGetStoreListRequest) {
	request = &StoreGetStoreListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "store/getStoreList", "api")
	request.Method = api.POST
	return
}

func CreateStoreGetStoreListResponse() (response *api.BaseResponse[StoreGetStoreListResponse]) {
	response = api.CreateResponse[StoreGetStoreListResponse](&StoreGetStoreListResponse{})
	return
}
