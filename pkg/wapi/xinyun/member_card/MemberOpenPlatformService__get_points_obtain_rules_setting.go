package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberOpenPlatformServiceGetPointsObtainRulesSetting
// @id 180
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=积分获取配置)
func (client *Client) MemberOpenPlatformServiceGetPointsObtainRulesSetting(request *MemberOpenPlatformServiceGetPointsObtainRulesSettingRequest) (response *MemberOpenPlatformServiceGetPointsObtainRulesSettingResponse, err error) {
	rpcResponse := CreateMemberOpenPlatformServiceGetPointsObtainRulesSettingResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberOpenPlatformServiceGetPointsObtainRulesSettingRequest struct {
	*api.BaseRequest
}

type MemberOpenPlatformServiceGetPointsObtainRulesSettingResponse struct {
}

func CreateMemberOpenPlatformServiceGetPointsObtainRulesSettingRequest() (request *MemberOpenPlatformServiceGetPointsObtainRulesSettingRequest) {
	request = &MemberOpenPlatformServiceGetPointsObtainRulesSettingRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "MemberOpenPlatformService/GetPointsObtainRulesSetting", "api")
	request.Method = api.POST
	return
}

func CreateMemberOpenPlatformServiceGetPointsObtainRulesSettingResponse() (response *api.BaseResponse[MemberOpenPlatformServiceGetPointsObtainRulesSettingResponse]) {
	response = api.CreateResponse[MemberOpenPlatformServiceGetPointsObtainRulesSettingResponse](&MemberOpenPlatformServiceGetPointsObtainRulesSettingResponse{})
	return
}
