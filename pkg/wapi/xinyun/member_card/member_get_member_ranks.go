package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberGetMemberRanks
// @id 978
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询会员等级)
func (client *Client) MemberGetMemberRanks(request *MemberGetMemberRanksRequest) (response *MemberGetMemberRanksResponse, err error) {
	rpcResponse := CreateMemberGetMemberRanksResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberGetMemberRanksRequest struct {
	*api.BaseRequest

	RankIds []int `json:"rankIds,omitempty"`
	Type    int   `json:"type,omitempty"`
}

type MemberGetMemberRanksResponse struct {
}

func CreateMemberGetMemberRanksRequest() (request *MemberGetMemberRanksRequest) {
	request = &MemberGetMemberRanksRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "member/getMemberRanks", "api")
	request.Method = api.POST
	return
}

func CreateMemberGetMemberRanksResponse() (response *api.BaseResponse[MemberGetMemberRanksResponse]) {
	response = api.CreateResponse[MemberGetMemberRanksResponse](&MemberGetMemberRanksResponse{})
	return
}
