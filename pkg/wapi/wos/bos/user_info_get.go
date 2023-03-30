package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserInfoGet
// @id 1251
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1251?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商家查询用户信息)
func (client *Client) UserInfoGet(request *UserInfoGetRequest) (response *UserInfoGetResponse, err error) {
	rpcResponse := CreateUserInfoGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserInfoGetRequest struct {
	*api.BaseRequest

	BosId int64  `json:"bosId,omitempty"`
	Wid   int64  `json:"wid,omitempty"`
	Phone string `json:"phone,omitempty"`
	Zone  string `json:"zone,omitempty"`
}

type UserInfoGetResponse struct {
	BosId    int64  `json:"bosId,omitempty"`
	Wid      int64  `json:"wid,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Zone     string `json:"zone,omitempty"`
	NickName string `json:"nickName,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
}

func CreateUserInfoGetRequest() (request *UserInfoGetRequest) {
	request = &UserInfoGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "user/info/get", "apigw")
	request.Method = api.POST
	return
}

func CreateUserInfoGetResponse() (response *api.BaseResponse[UserInfoGetResponse]) {
	response = api.CreateResponse[UserInfoGetResponse](&UserInfoGetResponse{})
	return
}
