package uc

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserBindSourceAndWid
// @id 1407
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=绑定渠道和wid)
func (client *Client) UserBindSourceAndWid(request *UserBindSourceAndWidRequest) (response *UserBindSourceAndWidResponse, err error) {
	rpcResponse := CreateUserBindSourceAndWidResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserBindSourceAndWidRequest struct {
	*api.BaseRequest

	Pid          int64  `json:"pid,omitempty"`
	Wid          int64  `json:"wid,omitempty"`
	SourceOpenid string `json:"sourceOpenid,omitempty"`
	Source       int64  `json:"source,omitempty"`
	SourceAppid  string `json:"sourceAppid,omitempty"`
}

type UserBindSourceAndWidResponse struct {
}

func CreateUserBindSourceAndWidRequest() (request *UserBindSourceAndWidRequest) {
	request = &UserBindSourceAndWidRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("uc", "1_0", "user/bindSourceAndWid", "api")
	request.Method = api.POST
	return
}

func CreateUserBindSourceAndWidResponse() (response *api.BaseResponse[UserBindSourceAndWidResponse]) {
	response = api.CreateResponse[UserBindSourceAndWidResponse](&UserBindSourceAndWidResponse{})
	return
}
