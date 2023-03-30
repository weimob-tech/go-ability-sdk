package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberOpenPlatformServiceGetPointsConsumeSetting
// @id 179
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=积分消耗配置)
func (client *Client) MemberOpenPlatformServiceGetPointsConsumeSetting(request *MemberOpenPlatformServiceGetPointsConsumeSettingRequest) (response *MemberOpenPlatformServiceGetPointsConsumeSettingResponse, err error) {
	rpcResponse := CreateMemberOpenPlatformServiceGetPointsConsumeSettingResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberOpenPlatformServiceGetPointsConsumeSettingRequest struct {
	*api.BaseRequest
}

type MemberOpenPlatformServiceGetPointsConsumeSettingResponse struct {
}

func CreateMemberOpenPlatformServiceGetPointsConsumeSettingRequest() (request *MemberOpenPlatformServiceGetPointsConsumeSettingRequest) {
	request = &MemberOpenPlatformServiceGetPointsConsumeSettingRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "MemberOpenPlatformService/GetPointsConsumeSetting", "api")
	request.Method = api.POST
	return
}

func CreateMemberOpenPlatformServiceGetPointsConsumeSettingResponse() (response *api.BaseResponse[MemberOpenPlatformServiceGetPointsConsumeSettingResponse]) {
	response = api.CreateResponse[MemberOpenPlatformServiceGetPointsConsumeSettingResponse](&MemberOpenPlatformServiceGetPointsConsumeSettingResponse{})
	return
}
