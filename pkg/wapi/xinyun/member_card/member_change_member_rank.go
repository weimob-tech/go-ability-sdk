package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberChangeMemberRank
// @id 677
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改会员等级)
func (client *Client) MemberChangeMemberRank(request *MemberChangeMemberRankRequest) (response *MemberChangeMemberRankResponse, err error) {
	rpcResponse := CreateMemberChangeMemberRankResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberChangeMemberRankRequest struct {
	*api.BaseRequest

	Mid         int64  `json:"mid,omitempty"`
	RankId      int    `json:"rankId,omitempty"`
	Reason      string `json:"reason,omitempty"`
	ChannelType int    `json:"channelType,omitempty"`
	AttachId    string `json:"attachId,omitempty"`
	RequestId   string `json:"requestId,omitempty"`
	Wid         int64  `json:"wid,omitempty"`
}

type MemberChangeMemberRankResponse struct {
}

func CreateMemberChangeMemberRankRequest() (request *MemberChangeMemberRankRequest) {
	request = &MemberChangeMemberRankRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "member/changeMemberRank", "api")
	request.Method = api.POST
	return
}

func CreateMemberChangeMemberRankResponse() (response *api.BaseResponse[MemberChangeMemberRankResponse]) {
	response = api.CreateResponse[MemberChangeMemberRankResponse](&MemberChangeMemberRankResponse{})
	return
}
