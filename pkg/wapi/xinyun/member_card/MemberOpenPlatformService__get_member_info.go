package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberOpenPlatformServiceGetMemberInfo
// @id 185
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取客户详情)
func (client *Client) MemberOpenPlatformServiceGetMemberInfo(request *MemberOpenPlatformServiceGetMemberInfoRequest) (response *MemberOpenPlatformServiceGetMemberInfoResponse, err error) {
	rpcResponse := CreateMemberOpenPlatformServiceGetMemberInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberOpenPlatformServiceGetMemberInfoRequest struct {
	*api.BaseRequest

	CustomerNo     string `json:"customer_no,omitempty"`
	MenbercardNo   string `json:"menbercard_no,omitempty"`
	WeimobOpenid   string `json:"weimob_openid,omitempty"`
	Openid         string `json:"openid,omitempty"`
	Phone          string `json:"phone,omitempty"`
	IsNeedTagsInfo bool   `json:"is_need_tags_info,omitempty"`
	IsNeedUserinfo bool   `json:"is_need_userinfo,omitempty"`
}

type MemberOpenPlatformServiceGetMemberInfoResponse struct {
}

func CreateMemberOpenPlatformServiceGetMemberInfoRequest() (request *MemberOpenPlatformServiceGetMemberInfoRequest) {
	request = &MemberOpenPlatformServiceGetMemberInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "MemberOpenPlatformService/GetMemberInfo", "api")
	request.Method = api.POST
	return
}

func CreateMemberOpenPlatformServiceGetMemberInfoResponse() (response *api.BaseResponse[MemberOpenPlatformServiceGetMemberInfoResponse]) {
	response = api.CreateResponse[MemberOpenPlatformServiceGetMemberInfoResponse](&MemberOpenPlatformServiceGetMemberInfoResponse{})
	return
}
