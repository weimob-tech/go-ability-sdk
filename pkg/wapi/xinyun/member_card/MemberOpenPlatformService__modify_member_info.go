package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberOpenPlatformServiceModifyMemberInfo
// @id 165
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改会员资料)
func (client *Client) MemberOpenPlatformServiceModifyMemberInfo(request *MemberOpenPlatformServiceModifyMemberInfoRequest) (response *MemberOpenPlatformServiceModifyMemberInfoResponse, err error) {
	rpcResponse := CreateMemberOpenPlatformServiceModifyMemberInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberOpenPlatformServiceModifyMemberInfoRequest struct {
	*api.BaseRequest

	MemberProperty MemberOpenPlatformServiceModifyMemberInfoRequestMember_property `json:"member_property,omitempty"`
	CustomerNo     string                                                          `json:"customer_no,omitempty"`
}

type MemberOpenPlatformServiceModifyMemberInfoResponse struct {
}

func CreateMemberOpenPlatformServiceModifyMemberInfoRequest() (request *MemberOpenPlatformServiceModifyMemberInfoRequest) {
	request = &MemberOpenPlatformServiceModifyMemberInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "MemberOpenPlatformService/ModifyMemberInfo", "api")
	request.Method = api.POST
	return
}

func CreateMemberOpenPlatformServiceModifyMemberInfoResponse() (response *api.BaseResponse[MemberOpenPlatformServiceModifyMemberInfoResponse]) {
	response = api.CreateResponse[MemberOpenPlatformServiceModifyMemberInfoResponse](&MemberOpenPlatformServiceModifyMemberInfoResponse{})
	return
}

type MemberOpenPlatformServiceModifyMemberInfoRequestMember_property struct {
}
