package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserDelete
// @id 1640
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除员工)
func (client *Client) UserDelete(request *UserDeleteRequest) (response *UserDeleteResponse, err error) {
	rpcResponse := CreateUserDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserDeleteRequest struct {
	*api.BaseRequest

	WidList         []int `json:"widList,omitempty"`
	TransferUserWid int64 `json:"transferUserWid,omitempty"`
}

type UserDeleteResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateUserDeleteRequest() (request *UserDeleteRequest) {
	request = &UserDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "user/delete", "api")
	request.Method = api.POST
	return
}

func CreateUserDeleteResponse() (response *api.BaseResponse[UserDeleteResponse]) {
	response = api.CreateResponse[UserDeleteResponse](&UserDeleteResponse{})
	return
}
