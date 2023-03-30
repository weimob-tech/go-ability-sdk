package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberOpenPlatformServiceDeleteOfflineMember
// @id 201
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除线下会员)
func (client *Client) MemberOpenPlatformServiceDeleteOfflineMember(request *MemberOpenPlatformServiceDeleteOfflineMemberRequest) (response *MemberOpenPlatformServiceDeleteOfflineMemberResponse, err error) {
	rpcResponse := CreateMemberOpenPlatformServiceDeleteOfflineMemberResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberOpenPlatformServiceDeleteOfflineMemberRequest struct {
	*api.BaseRequest

	Phones []string `json:"phones,omitempty"`
}

type MemberOpenPlatformServiceDeleteOfflineMemberResponse struct {
}

func CreateMemberOpenPlatformServiceDeleteOfflineMemberRequest() (request *MemberOpenPlatformServiceDeleteOfflineMemberRequest) {
	request = &MemberOpenPlatformServiceDeleteOfflineMemberRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "MemberOpenPlatformService/DeleteOfflineMember", "api")
	request.Method = api.POST
	return
}

func CreateMemberOpenPlatformServiceDeleteOfflineMemberResponse() (response *api.BaseResponse[MemberOpenPlatformServiceDeleteOfflineMemberResponse]) {
	response = api.CreateResponse[MemberOpenPlatformServiceDeleteOfflineMemberResponse](&MemberOpenPlatformServiceDeleteOfflineMemberResponse{})
	return
}
