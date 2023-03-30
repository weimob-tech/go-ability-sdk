package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ActivityTriggerConsume
// @id 1416
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=消费有礼触发)
func (client *Client) ActivityTriggerConsume(request *ActivityTriggerConsumeRequest) (response *ActivityTriggerConsumeResponse, err error) {
	rpcResponse := CreateActivityTriggerConsumeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ActivityTriggerConsumeRequest struct {
	*api.BaseRequest

	TriggerAmountInfo ActivityTriggerConsumeRequestTriggerAmountInfo `json:"triggerAmountInfo,omitempty"`
	NotSendCollection []int                                          `json:"notSendCollection,omitempty"`
	OrderNo           string                                         `json:"orderNo,omitempty"`
	TriggerAmount     float64                                        `json:"triggerAmount,omitempty"`
}

type ActivityTriggerConsumeResponse struct {
}

func CreateActivityTriggerConsumeRequest() (request *ActivityTriggerConsumeRequest) {
	request = &ActivityTriggerConsumeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "activity/triggerConsume", "api")
	request.Method = api.POST
	return
}

func CreateActivityTriggerConsumeResponse() (response *api.BaseResponse[ActivityTriggerConsumeResponse]) {
	response = api.CreateResponse[ActivityTriggerConsumeResponse](&ActivityTriggerConsumeResponse{})
	return
}

type ActivityTriggerConsumeRequestTriggerAmountInfo struct {
	EntityCard        float64 `json:"entityCard,omitempty"`
	Actual            float64 `json:"actual,omitempty"`
	Balance           float64 `json:"balance,omitempty"`
	StoreRechargeCard float64 `json:"storeRechargeCard,omitempty"`
	GiftCard          float64 `json:"giftCard,omitempty"`
}
