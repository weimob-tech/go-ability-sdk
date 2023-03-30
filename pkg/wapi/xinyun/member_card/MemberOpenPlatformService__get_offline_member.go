package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberOpenPlatformServiceGetOfflineMember
// @id 202
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取线下会员列表)
func (client *Client) MemberOpenPlatformServiceGetOfflineMember(request *MemberOpenPlatformServiceGetOfflineMemberRequest) (response *MemberOpenPlatformServiceGetOfflineMemberResponse, err error) {
	rpcResponse := CreateMemberOpenPlatformServiceGetOfflineMemberResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberOpenPlatformServiceGetOfflineMemberRequest struct {
	*api.BaseRequest

	Searchby  int    `json:"searchby,omitempty"`
	Keyword   string `json:"keyword,omitempty"`
	Pageindex int    `json:"pageindex,omitempty"`
	Pagesize  int    `json:"pagesize,omitempty"`
}

type MemberOpenPlatformServiceGetOfflineMemberResponse struct {
}

func CreateMemberOpenPlatformServiceGetOfflineMemberRequest() (request *MemberOpenPlatformServiceGetOfflineMemberRequest) {
	request = &MemberOpenPlatformServiceGetOfflineMemberRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "MemberOpenPlatformService/GetOfflineMember", "api")
	request.Method = api.POST
	return
}

func CreateMemberOpenPlatformServiceGetOfflineMemberResponse() (response *api.BaseResponse[MemberOpenPlatformServiceGetOfflineMemberResponse]) {
	response = api.CreateResponse[MemberOpenPlatformServiceGetOfflineMemberResponse](&MemberOpenPlatformServiceGetOfflineMemberResponse{})
	return
}
