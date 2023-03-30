package uc

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserGetSuperWidBySource
// @id 1152
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=根据属性获取主wid)
func (client *Client) UserGetSuperWidBySource(request *UserGetSuperWidBySourceRequest) (response *UserGetSuperWidBySourceResponse, err error) {
	rpcResponse := CreateUserGetSuperWidBySourceResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserGetSuperWidBySourceRequest struct {
	*api.BaseRequest

	Source       int    `json:"source,omitempty"`
	SourceAppid  string `json:"sourceAppid,omitempty"`
	SourceOpenid string `json:"sourceOpenid,omitempty"`
}

type UserGetSuperWidBySourceResponse struct {
}

func CreateUserGetSuperWidBySourceRequest() (request *UserGetSuperWidBySourceRequest) {
	request = &UserGetSuperWidBySourceRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("uc", "1_0", "user/getSuperWidBySource", "api")
	request.Method = api.POST
	return
}

func CreateUserGetSuperWidBySourceResponse() (response *api.BaseResponse[UserGetSuperWidBySourceResponse]) {
	response = api.CreateResponse[UserGetSuperWidBySourceResponse](&UserGetSuperWidBySourceResponse{})
	return
}
