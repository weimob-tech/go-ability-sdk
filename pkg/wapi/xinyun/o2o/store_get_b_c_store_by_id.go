package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreGetBCStoreById
// @id 2092
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询微盟门店信息)
func (client *Client) StoreGetBCStoreById(request *StoreGetBCStoreByIdRequest) (response *StoreGetBCStoreByIdResponse, err error) {
	rpcResponse := CreateStoreGetBCStoreByIdResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreGetBCStoreByIdRequest struct {
	*api.BaseRequest

	Id int64 `json:"id,omitempty"`
}

type StoreGetBCStoreByIdResponse struct {
}

func CreateStoreGetBCStoreByIdRequest() (request *StoreGetBCStoreByIdRequest) {
	request = &StoreGetBCStoreByIdRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "store/getBCStoreById", "api")
	request.Method = api.POST
	return
}

func CreateStoreGetBCStoreByIdResponse() (response *api.BaseResponse[StoreGetBCStoreByIdResponse]) {
	response = api.CreateResponse[StoreGetBCStoreByIdResponse](&StoreGetBCStoreByIdResponse{})
	return
}
