package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberCheckVerifyCodeOpen
// @id 1843
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=校验短信验证码)
func (client *Client) MemberCheckVerifyCodeOpen(request *MemberCheckVerifyCodeOpenRequest) (response *MemberCheckVerifyCodeOpenResponse, err error) {
	rpcResponse := CreateMemberCheckVerifyCodeOpenResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberCheckVerifyCodeOpenRequest struct {
	*api.BaseRequest

	Phone      string `json:"phone,omitempty"`
	VerifyCode string `json:"verifyCode,omitempty"`
}

type MemberCheckVerifyCodeOpenResponse struct {
}

func CreateMemberCheckVerifyCodeOpenRequest() (request *MemberCheckVerifyCodeOpenRequest) {
	request = &MemberCheckVerifyCodeOpenRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "member/checkVerifyCodeOpen", "api")
	request.Method = api.POST
	return
}

func CreateMemberCheckVerifyCodeOpenResponse() (response *api.BaseResponse[MemberCheckVerifyCodeOpenResponse]) {
	response = api.CreateResponse[MemberCheckVerifyCodeOpenResponse](&MemberCheckVerifyCodeOpenResponse{})
	return
}
