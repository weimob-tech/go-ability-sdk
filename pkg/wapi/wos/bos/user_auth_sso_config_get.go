package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserAuthSsoConfigGet
// @id 1935
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1935?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=APP开店-查询商家设置)
func (client *Client) UserAuthSsoConfigGet(request *UserAuthSsoConfigGetRequest) (response *UserAuthSsoConfigGetResponse, err error) {
	rpcResponse := CreateUserAuthSsoConfigGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserAuthSsoConfigGetRequest struct {
	*api.BaseRequest
}

type UserAuthSsoConfigGetResponse struct {
	LoginCallbackUrl string `json:"loginCallbackUrl,omitempty"`
	NumberScale      int64  `json:"numberScale,omitempty"`
	Expires          int64  `json:"expires,omitempty"`
}

func CreateUserAuthSsoConfigGetRequest() (request *UserAuthSsoConfigGetRequest) {
	request = &UserAuthSsoConfigGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "user/auth/sso/config/get", "apigw")
	request.Method = api.POST
	return
}

func CreateUserAuthSsoConfigGetResponse() (response *api.BaseResponse[UserAuthSsoConfigGetResponse]) {
	response = api.CreateResponse[UserAuthSsoConfigGetResponse](&UserAuthSsoConfigGetResponse{})
	return
}
