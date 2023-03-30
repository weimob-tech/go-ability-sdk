package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberChangeMemberGrowth
// @id 1673
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改会员成长值V2)
func (client *Client) MemberChangeMemberGrowthV2_0(request *MemberChangeMemberGrowthRequestV2_0) (response *MemberChangeMemberGrowthResponseV2_0, err error) {
	rpcResponse := CreateMemberChangeMemberGrowthResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberChangeMemberGrowthRequestV2_0 struct {
	*api.BaseRequest

	Wid          int64  `json:"wid,omitempty"`
	McTemplateId int64  `json:"mcTemplateId,omitempty"`
	Growth       int    `json:"growth,omitempty"`
	Reason       string `json:"reason,omitempty"`
	ChannelType  int    `json:"channelType,omitempty"`
	AttachId     string `json:"attachId,omitempty"`
	RequestId    string `json:"requestId,omitempty"`
}

type MemberChangeMemberGrowthResponseV2_0 struct {
}

func CreateMemberChangeMemberGrowthRequestV2_0() (request *MemberChangeMemberGrowthRequestV2_0) {
	request = &MemberChangeMemberGrowthRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "2_0", "member/changeMemberGrowth", "api")
	request.Method = api.POST
	return
}

func CreateMemberChangeMemberGrowthResponseV2_0() (response *api.BaseResponse[MemberChangeMemberGrowthResponseV2_0]) {
	response = api.CreateResponse[MemberChangeMemberGrowthResponseV2_0](&MemberChangeMemberGrowthResponseV2_0{})
	return
}
