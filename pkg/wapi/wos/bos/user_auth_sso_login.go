package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserAuthSsoLogin
// @id 1937
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1937?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=三方用户登入接口)
func (client *Client) UserAuthSsoLogin(request *UserAuthSsoLoginRequest) (response *UserAuthSsoLoginResponse, err error) {
	rpcResponse := CreateUserAuthSsoLoginResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserAuthSsoLoginRequest struct {
	*api.BaseRequest

	OpenUserId string `json:"openUserId,omitempty"`
	Nickname   string `json:"nickname,omitempty"`
	HeadUrl    string `json:"headUrl,omitempty"`
	Phone      string `json:"phone,omitempty"`
}

type UserAuthSsoLoginResponse struct {
	AppTicket string `json:"appTicket,omitempty"`
	Expires   int64  `json:"expires,omitempty"`
}

func CreateUserAuthSsoLoginRequest() (request *UserAuthSsoLoginRequest) {
	request = &UserAuthSsoLoginRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "user/auth/sso/login", "apigw")
	request.Method = api.POST
	return
}

func CreateUserAuthSsoLoginResponse() (response *api.BaseResponse[UserAuthSsoLoginResponse]) {
	response = api.CreateResponse[UserAuthSsoLoginResponse](&UserAuthSsoLoginResponse{})
	return
}
