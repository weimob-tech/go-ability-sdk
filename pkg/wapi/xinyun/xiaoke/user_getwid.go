package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserGetwid
// @id 1893
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=根据员工手机号,pid获取wid)
func (client *Client) UserGetwid(request *UserGetwidRequest) (response *UserGetwidResponse, err error) {
	rpcResponse := CreateUserGetwidResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserGetwidRequest struct {
	*api.BaseRequest

	Phone string `json:"phone,omitempty"`
}

type UserGetwidResponse struct {
}

func CreateUserGetwidRequest() (request *UserGetwidRequest) {
	request = &UserGetwidRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "user/getwid", "api")
	request.Method = api.POST
	return
}

func CreateUserGetwidResponse() (response *api.BaseResponse[UserGetwidResponse]) {
	response = api.CreateResponse[UserGetwidResponse](&UserGetwidResponse{})
	return
}
