package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberChangeMemberRank
// @id 1674
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改会员等级V2)
func (client *Client) MemberChangeMemberRankV2_0(request *MemberChangeMemberRankRequestV2_0) (response *MemberChangeMemberRankResponseV2_0, err error) {
	rpcResponse := CreateMemberChangeMemberRankResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberChangeMemberRankRequestV2_0 struct {
	*api.BaseRequest

	Wid          int64  `json:"wid,omitempty"`
	McTemplateId int64  `json:"mcTemplateId,omitempty"`
	RankId       int    `json:"rankId,omitempty"`
	Reason       string `json:"reason,omitempty"`
	ChannelType  int    `json:"channelType,omitempty"`
	AttachId     string `json:"attachId,omitempty"`
	RequestId    string `json:"requestId,omitempty"`
}

type MemberChangeMemberRankResponseV2_0 struct {
}

func CreateMemberChangeMemberRankRequestV2_0() (request *MemberChangeMemberRankRequestV2_0) {
	request = &MemberChangeMemberRankRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "2_0", "member/changeMemberRank", "api")
	request.Method = api.POST
	return
}

func CreateMemberChangeMemberRankResponseV2_0() (response *api.BaseResponse[MemberChangeMemberRankResponseV2_0]) {
	response = api.CreateResponse[MemberChangeMemberRankResponseV2_0](&MemberChangeMemberRankResponseV2_0{})
	return
}
