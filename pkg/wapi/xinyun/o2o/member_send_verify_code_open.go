package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberSendVerifyCodeOpen
// @id 1844
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=发送短信验证码)
func (client *Client) MemberSendVerifyCodeOpen(request *MemberSendVerifyCodeOpenRequest) (response *MemberSendVerifyCodeOpenResponse, err error) {
	rpcResponse := CreateMemberSendVerifyCodeOpenResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberSendVerifyCodeOpenRequest struct {
	*api.BaseRequest

	Phone string `json:"phone,omitempty"`
}

type MemberSendVerifyCodeOpenResponse struct {
}

func CreateMemberSendVerifyCodeOpenRequest() (request *MemberSendVerifyCodeOpenRequest) {
	request = &MemberSendVerifyCodeOpenRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "member/sendVerifyCodeOpen", "api")
	request.Method = api.POST
	return
}

func CreateMemberSendVerifyCodeOpenResponse() (response *api.BaseResponse[MemberSendVerifyCodeOpenResponse]) {
	response = api.CreateResponse[MemberSendVerifyCodeOpenResponse](&MemberSendVerifyCodeOpenResponse{})
	return
}
