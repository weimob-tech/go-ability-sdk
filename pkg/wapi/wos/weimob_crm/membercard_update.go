package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembercardUpdate
// @id 1742
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1742?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新会员卡信息)
func (client *Client) MembercardUpdate(request *MembercardUpdateRequest) (response *MembercardUpdateResponse, err error) {
	rpcResponse := CreateMembercardUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembercardUpdateRequest struct {
	*api.BaseRequest

	Wid              int64  `json:"wid,omitempty"`
	DelayType        int64  `json:"delayType,omitempty"`
	NewCustomCardNo  string `json:"newCustomCardNo,omitempty"`
	Reason           string `json:"reason,omitempty"`
	OperatorName     string `json:"operatorName,omitempty"`
	OldCustomCardNo  string `json:"oldCustomCardNo,omitempty"`
	UpdateType       int64  `json:"updateType,omitempty"`
	LevelId          int64  `json:"levelId,omitempty"`
	OperatorWId      int64  `json:"operatorWId,omitempty"`
	ExpireTime       int64  `json:"expireTime,omitempty"`
	MembershipPlanId int64  `json:"membershipPlanId,omitempty"`
	ChannelType      int64  `json:"channelType,omitempty"`
	ChannelId        int64  `json:"channelId,omitempty"`
}

type MembercardUpdateResponse struct {
	Status bool `json:"status,omitempty"`
}

func CreateMembercardUpdateRequest() (request *MembercardUpdateRequest) {
	request = &MembercardUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "membercard/update", "apigw")
	request.Method = api.POST
	return
}

func CreateMembercardUpdateResponse() (response *api.BaseResponse[MembercardUpdateResponse]) {
	response = api.CreateResponse[MembercardUpdateResponse](&MembercardUpdateResponse{})
	return
}
