package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberOpenPlatformServiceModifyMemberLevel
// @id 168
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改会员等级)
func (client *Client) MemberOpenPlatformServiceModifyMemberLevel(request *MemberOpenPlatformServiceModifyMemberLevelRequest) (response *MemberOpenPlatformServiceModifyMemberLevelResponse, err error) {
	rpcResponse := CreateMemberOpenPlatformServiceModifyMemberLevelResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberOpenPlatformServiceModifyMemberLevelRequest struct {
	*api.BaseRequest

	CustomerNos []string `json:"customer_nos,omitempty"`
	LevelCode   string   `json:"level_code,omitempty"`
}

type MemberOpenPlatformServiceModifyMemberLevelResponse struct {
}

func CreateMemberOpenPlatformServiceModifyMemberLevelRequest() (request *MemberOpenPlatformServiceModifyMemberLevelRequest) {
	request = &MemberOpenPlatformServiceModifyMemberLevelRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "MemberOpenPlatformService/ModifyMemberLevel", "api")
	request.Method = api.POST
	return
}

func CreateMemberOpenPlatformServiceModifyMemberLevelResponse() (response *api.BaseResponse[MemberOpenPlatformServiceModifyMemberLevelResponse]) {
	response = api.CreateResponse[MemberOpenPlatformServiceModifyMemberLevelResponse](&MemberOpenPlatformServiceModifyMemberLevelResponse{})
	return
}
