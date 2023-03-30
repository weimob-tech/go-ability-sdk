package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberChangeMemberGrowth
// @id 674
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改会员成长值)
func (client *Client) MemberChangeMemberGrowth(request *MemberChangeMemberGrowthRequest) (response *MemberChangeMemberGrowthResponse, err error) {
	rpcResponse := CreateMemberChangeMemberGrowthResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberChangeMemberGrowthRequest struct {
	*api.BaseRequest

	Mid         int64  `json:"mid,omitempty"`
	Growth      int    `json:"growth,omitempty"`
	Reason      string `json:"reason,omitempty"`
	ChannelType int    `json:"channelType,omitempty"`
	AttachId    string `json:"attachId,omitempty"`
	RequestId   string `json:"requestId,omitempty"`
	Wid         int64  `json:"wid,omitempty"`
}

type MemberChangeMemberGrowthResponse struct {
}

func CreateMemberChangeMemberGrowthRequest() (request *MemberChangeMemberGrowthRequest) {
	request = &MemberChangeMemberGrowthRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "member/changeMemberGrowth", "api")
	request.Method = api.POST
	return
}

func CreateMemberChangeMemberGrowthResponse() (response *api.BaseResponse[MemberChangeMemberGrowthResponse]) {
	response = api.CreateResponse[MemberChangeMemberGrowthResponse](&MemberChangeMemberGrowthResponse{})
	return
}
