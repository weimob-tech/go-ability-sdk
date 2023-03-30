package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreGetStoreInfo
// @id 510
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=酒店信息)
func (client *Client) StoreGetStoreInfo(request *StoreGetStoreInfoRequest) (response *StoreGetStoreInfoResponse, err error) {
	rpcResponse := CreateStoreGetStoreInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreGetStoreInfoRequest struct {
	*api.BaseRequest

	StoreId int64 `json:"storeId,omitempty"`
}

type StoreGetStoreInfoResponse struct {
}

func CreateStoreGetStoreInfoRequest() (request *StoreGetStoreInfoRequest) {
	request = &StoreGetStoreInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "store/getStoreInfo", "api")
	request.Method = api.POST
	return
}

func CreateStoreGetStoreInfoResponse() (response *api.BaseResponse[StoreGetStoreInfoResponse]) {
	response = api.CreateResponse[StoreGetStoreInfoResponse](&StoreGetStoreInfoResponse{})
	return
}
