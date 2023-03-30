package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberOpenPlatformServiceModifyMemberStatus
// @id 167
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改会员状态)
func (client *Client) MemberOpenPlatformServiceModifyMemberStatus(request *MemberOpenPlatformServiceModifyMemberStatusRequest) (response *MemberOpenPlatformServiceModifyMemberStatusResponse, err error) {
	rpcResponse := CreateMemberOpenPlatformServiceModifyMemberStatusResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberOpenPlatformServiceModifyMemberStatusRequest struct {
	*api.BaseRequest

	CustomerNos []string `json:"customer_nos,omitempty"`
	Isfrozen    bool     `json:"isfrozen,omitempty"`
}

type MemberOpenPlatformServiceModifyMemberStatusResponse struct {
}

func CreateMemberOpenPlatformServiceModifyMemberStatusRequest() (request *MemberOpenPlatformServiceModifyMemberStatusRequest) {
	request = &MemberOpenPlatformServiceModifyMemberStatusRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "MemberOpenPlatformService/ModifyMemberStatus", "api")
	request.Method = api.POST
	return
}

func CreateMemberOpenPlatformServiceModifyMemberStatusResponse() (response *api.BaseResponse[MemberOpenPlatformServiceModifyMemberStatusResponse]) {
	response = api.CreateResponse[MemberOpenPlatformServiceModifyMemberStatusResponse](&MemberOpenPlatformServiceModifyMemberStatusResponse{})
	return
}
