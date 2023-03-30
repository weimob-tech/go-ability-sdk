package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreQueryStoreBizConfig
// @id 1417
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询门店业务基础配置)
func (client *Client) StoreQueryStoreBizConfig(request *StoreQueryStoreBizConfigRequest) (response *StoreQueryStoreBizConfigResponse, err error) {
	rpcResponse := CreateStoreQueryStoreBizConfigResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreQueryStoreBizConfigRequest struct {
	*api.BaseRequest

	BizTypes []int   `json:"bizTypes,omitempty"`
	StoreIds []int64 `json:"storeIds,omitempty"`
}

type StoreQueryStoreBizConfigResponse struct {
}

func CreateStoreQueryStoreBizConfigRequest() (request *StoreQueryStoreBizConfigRequest) {
	request = &StoreQueryStoreBizConfigRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "store/queryStoreBizConfig", "api")
	request.Method = api.POST
	return
}

func CreateStoreQueryStoreBizConfigResponse() (response *api.BaseResponse[StoreQueryStoreBizConfigResponse]) {
	response = api.CreateResponse[StoreQueryStoreBizConfigResponse](&StoreQueryStoreBizConfigResponse{})
	return
}
