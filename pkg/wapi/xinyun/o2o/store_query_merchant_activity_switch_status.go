package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreQueryMerchantActivitySwitchStatus
// @id 2093
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询商户业务开关状态)
func (client *Client) StoreQueryMerchantActivitySwitchStatus(request *StoreQueryMerchantActivitySwitchStatusRequest) (response *StoreQueryMerchantActivitySwitchStatusResponse, err error) {
	rpcResponse := CreateStoreQueryMerchantActivitySwitchStatusResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreQueryMerchantActivitySwitchStatusRequest struct {
	*api.BaseRequest

	BusinessType int `json:"businessType,omitempty"`
	ActivityType int `json:"activityType,omitempty"`
}

type StoreQueryMerchantActivitySwitchStatusResponse struct {
}

func CreateStoreQueryMerchantActivitySwitchStatusRequest() (request *StoreQueryMerchantActivitySwitchStatusRequest) {
	request = &StoreQueryMerchantActivitySwitchStatusRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "store/queryMerchantActivitySwitchStatus", "api")
	request.Method = api.POST
	return
}

func CreateStoreQueryMerchantActivitySwitchStatusResponse() (response *api.BaseResponse[StoreQueryMerchantActivitySwitchStatusResponse]) {
	response = api.CreateResponse[StoreQueryMerchantActivitySwitchStatusResponse](&StoreQueryMerchantActivitySwitchStatusResponse{})
	return
}
