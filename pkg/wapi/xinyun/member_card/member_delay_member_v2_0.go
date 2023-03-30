package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberDelayMember
// @id 1795
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=会员卡续期)
func (client *Client) MemberDelayMemberV2_0(request *MemberDelayMemberRequestV2_0) (response *MemberDelayMemberResponseV2_0, err error) {
	rpcResponse := CreateMemberDelayMemberResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberDelayMemberRequestV2_0 struct {
	*api.BaseRequest

	Wid                  int64  `json:"wid,omitempty"`
	MemberCardTemplateId int64  `json:"memberCardTemplateId,omitempty"`
	DelayType            int    `json:"delayType,omitempty"`
	ExpireDate           string `json:"expireDate,omitempty"`
	OperatorType         int    `json:"operatorType,omitempty"`
	OperatorId           int64  `json:"operatorId,omitempty"`
	OperatorReason       string `json:"operatorReason,omitempty"`
}

type MemberDelayMemberResponseV2_0 struct {
	Status bool `json:"status,omitempty"`
}

func CreateMemberDelayMemberRequestV2_0() (request *MemberDelayMemberRequestV2_0) {
	request = &MemberDelayMemberRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "2_0", "member/delayMember", "api")
	request.Method = api.POST
	return
}

func CreateMemberDelayMemberResponseV2_0() (response *api.BaseResponse[MemberDelayMemberResponseV2_0]) {
	response = api.CreateResponse[MemberDelayMemberResponseV2_0](&MemberDelayMemberResponseV2_0{})
	return
}
