package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreUpdateMerchantActivitySwitchStatus
// @id 2094
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新商户业务开关状态)
func (client *Client) StoreUpdateMerchantActivitySwitchStatus(request *StoreUpdateMerchantActivitySwitchStatusRequest) (response *StoreUpdateMerchantActivitySwitchStatusResponse, err error) {
	rpcResponse := CreateStoreUpdateMerchantActivitySwitchStatusResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreUpdateMerchantActivitySwitchStatusRequest struct {
	*api.BaseRequest

	BusinessType int `json:"businessType,omitempty"`
	SwitchStatus int `json:"switchStatus,omitempty"`
	ActivityType int `json:"activityType,omitempty"`
}

type StoreUpdateMerchantActivitySwitchStatusResponse struct {
}

func CreateStoreUpdateMerchantActivitySwitchStatusRequest() (request *StoreUpdateMerchantActivitySwitchStatusRequest) {
	request = &StoreUpdateMerchantActivitySwitchStatusRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "store/updateMerchantActivitySwitchStatus", "api")
	request.Method = api.POST
	return
}

func CreateStoreUpdateMerchantActivitySwitchStatusResponse() (response *api.BaseResponse[StoreUpdateMerchantActivitySwitchStatusResponse]) {
	response = api.CreateResponse[StoreUpdateMerchantActivitySwitchStatusResponse](&StoreUpdateMerchantActivitySwitchStatusResponse{})
	return
}
