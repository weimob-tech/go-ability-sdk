package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserAuthSsoLogout
// @id 1938
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1938?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=三方用户登出接口)
func (client *Client) UserAuthSsoLogout(request *UserAuthSsoLogoutRequest) (response *UserAuthSsoLogoutResponse, err error) {
	rpcResponse := CreateUserAuthSsoLogoutResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserAuthSsoLogoutRequest struct {
	*api.BaseRequest

	OpenUserId string `json:"openUserId,omitempty"`
}

type UserAuthSsoLogoutResponse struct {
}

func CreateUserAuthSsoLogoutRequest() (request *UserAuthSsoLogoutRequest) {
	request = &UserAuthSsoLogoutRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "user/auth/sso/logout", "apigw")
	request.Method = api.POST
	return
}

func CreateUserAuthSsoLogoutResponse() (response *api.BaseResponse[UserAuthSsoLogoutResponse]) {
	response = api.CreateResponse[UserAuthSsoLogoutResponse](&UserAuthSsoLogoutResponse{})
	return
}
