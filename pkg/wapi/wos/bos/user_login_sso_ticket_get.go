package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserLoginSsoTicketGet
// @id 2120
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2120?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取单点登陆临时凭证)
func (client *Client) UserLoginSsoTicketGet(request *UserLoginSsoTicketGetRequest) (response *UserLoginSsoTicketGetResponse, err error) {
	rpcResponse := CreateUserLoginSsoTicketGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserLoginSsoTicketGetRequest struct {
	*api.BaseRequest

	Wid     int64  `json:"wid,omitempty"`
	Channel string `json:"channel,omitempty"`
}

type UserLoginSsoTicketGetResponse struct {
	Ticket string `json:"ticket,omitempty"`
}

func CreateUserLoginSsoTicketGetRequest() (request *UserLoginSsoTicketGetRequest) {
	request = &UserLoginSsoTicketGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "user/login/sso/ticket/get", "apigw")
	request.Method = api.POST
	return
}

func CreateUserLoginSsoTicketGetResponse() (response *api.BaseResponse[UserLoginSsoTicketGetResponse]) {
	response = api.CreateResponse[UserLoginSsoTicketGetResponse](&UserLoginSsoTicketGetResponse{})
	return
}
