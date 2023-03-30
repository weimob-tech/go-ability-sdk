package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberGetDynamicCodeInfo
// @id 1593
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询会员动态码信息)
func (client *Client) MemberGetDynamicCodeInfo(request *MemberGetDynamicCodeInfoRequest) (response *MemberGetDynamicCodeInfoResponse, err error) {
	rpcResponse := CreateMemberGetDynamicCodeInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberGetDynamicCodeInfoRequest struct {
	*api.BaseRequest

	DynamicCode string `json:"dynamicCode,omitempty"`
}

type MemberGetDynamicCodeInfoResponse struct {
}

func CreateMemberGetDynamicCodeInfoRequest() (request *MemberGetDynamicCodeInfoRequest) {
	request = &MemberGetDynamicCodeInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "member/getDynamicCodeInfo", "api")
	request.Method = api.POST
	return
}

func CreateMemberGetDynamicCodeInfoResponse() (response *api.BaseResponse[MemberGetDynamicCodeInfoResponse]) {
	response = api.CreateResponse[MemberGetDynamicCodeInfoResponse](&MemberGetDynamicCodeInfoResponse{})
	return
}
