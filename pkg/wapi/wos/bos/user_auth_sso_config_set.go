package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserAuthSsoConfigSet
// @id 1936
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1936?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=APP开店-设置商家设置)
func (client *Client) UserAuthSsoConfigSet(request *UserAuthSsoConfigSetRequest) (response *UserAuthSsoConfigSetResponse, err error) {
	rpcResponse := CreateUserAuthSsoConfigSetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserAuthSsoConfigSetRequest struct {
	*api.BaseRequest

	LoginCallbackUrl string `json:"loginCallbackUrl,omitempty"`
	NumberScale      int64  `json:"numberScale,omitempty"`
	Expires          int64  `json:"expires,omitempty"`
}

type UserAuthSsoConfigSetResponse struct {
}

func CreateUserAuthSsoConfigSetRequest() (request *UserAuthSsoConfigSetRequest) {
	request = &UserAuthSsoConfigSetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "user/auth/sso/config/set", "apigw")
	request.Method = api.POST
	return
}

func CreateUserAuthSsoConfigSetResponse() (response *api.BaseResponse[UserAuthSsoConfigSetResponse]) {
	response = api.CreateResponse[UserAuthSsoConfigSetResponse](&UserAuthSsoConfigSetResponse{})
	return
}
