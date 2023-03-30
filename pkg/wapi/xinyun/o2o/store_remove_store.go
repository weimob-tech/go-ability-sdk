package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreRemoveStore
// @id 51
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除门店)
func (client *Client) StoreRemoveStore(request *StoreRemoveStoreRequest) (response *StoreRemoveStoreResponse, err error) {
	rpcResponse := CreateStoreRemoveStoreResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreRemoveStoreRequest struct {
	*api.BaseRequest

	StoreId int64 `json:"storeId,omitempty"`
	PoiId   int64 `json:"poiId,omitempty"`
}

type StoreRemoveStoreResponse struct {
}

func CreateStoreRemoveStoreRequest() (request *StoreRemoveStoreRequest) {
	request = &StoreRemoveStoreRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "store/removeStore", "api")
	request.Method = api.POST
	return
}

func CreateStoreRemoveStoreResponse() (response *api.BaseResponse[StoreRemoveStoreResponse]) {
	response = api.CreateResponse[StoreRemoveStoreResponse](&StoreRemoveStoreResponse{})
	return
}
