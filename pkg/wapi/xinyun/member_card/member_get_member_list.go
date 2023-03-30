package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberGetMemberList
// @id 654
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取会员列表)
func (client *Client) MemberGetMemberList(request *MemberGetMemberListRequest) (response *MemberGetMemberListResponse, err error) {
	rpcResponse := CreateMemberGetMemberListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberGetMemberListRequest struct {
	*api.BaseRequest

	Cursor int64 `json:"cursor,omitempty"`
	Size   int64 `json:"size,omitempty"`
}

type MemberGetMemberListResponse struct {
}

func CreateMemberGetMemberListRequest() (request *MemberGetMemberListRequest) {
	request = &MemberGetMemberListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "member/getMemberList", "api")
	request.Method = api.POST
	return
}

func CreateMemberGetMemberListResponse() (response *api.BaseResponse[MemberGetMemberListResponse]) {
	response = api.CreateResponse[MemberGetMemberListResponse](&MemberGetMemberListResponse{})
	return
}
