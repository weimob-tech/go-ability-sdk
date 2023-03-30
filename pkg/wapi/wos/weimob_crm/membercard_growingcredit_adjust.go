package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembercardGrowingcreditAdjust
// @id 1512
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1512?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=调整会员卡成长值)
func (client *Client) MembercardGrowingcreditAdjust(request *MembercardGrowingcreditAdjustRequest) (response *MembercardGrowingcreditAdjustResponse, err error) {
	rpcResponse := CreateMembercardGrowingcreditAdjustResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembercardGrowingcreditAdjustRequest struct {
	*api.BaseRequest

	MembershipPlanId int64  `json:"membershipPlanId,omitempty"`
	EventSource      int64  `json:"eventSource,omitempty"`
	EventType        string `json:"eventType,omitempty"`
	EventId          string `json:"eventId,omitempty"`
	ChangeGrowth     int64  `json:"changeGrowth,omitempty"`
	Reason           string `json:"reason,omitempty"`
	OperatorChannel  int64  `json:"operatorChannel,omitempty"`
	ChangeRuleType   int64  `json:"changeRuleType,omitempty"`
	OperatorWid      int64  `json:"operatorWid,omitempty"`
	Wid              int64  `json:"wid,omitempty"`
	Vid              int64  `json:"vid,omitempty"`
	ChannelType      int64  `json:"channelType,omitempty"`
	ChannelId        int64  `json:"channelId,omitempty"`
}

type MembercardGrowingcreditAdjustResponse struct {
	GrowthBillId int64 `json:"growthBillId,omitempty"`
}

func CreateMembercardGrowingcreditAdjustRequest() (request *MembercardGrowingcreditAdjustRequest) {
	request = &MembercardGrowingcreditAdjustRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "membercard/growingcredit/adjust", "apigw")
	request.Method = api.POST
	return
}

func CreateMembercardGrowingcreditAdjustResponse() (response *api.BaseResponse[MembercardGrowingcreditAdjustResponse]) {
	response = api.CreateResponse[MembercardGrowingcreditAdjustResponse](&MembercardGrowingcreditAdjustResponse{})
	return
}
