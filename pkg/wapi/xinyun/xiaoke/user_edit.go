package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserEdit
// @id 1639
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=编辑员工)
func (client *Client) UserEdit(request *UserEditRequest) (response *UserEditResponse, err error) {
	rpcResponse := CreateUserEditResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserEditRequest struct {
	*api.BaseRequest

	DepartmentIds []int  `json:"departmentIds,omitempty"`
	Wid           int64  `json:"wid,omitempty"`
	UserName      string `json:"userName,omitempty"`
	Email         string `json:"email,omitempty"`
}

type UserEditResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateUserEditRequest() (request *UserEditRequest) {
	request = &UserEditRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "user/edit", "api")
	request.Method = api.POST
	return
}

func CreateUserEditResponse() (response *api.BaseResponse[UserEditResponse]) {
	response = api.CreateResponse[UserEditResponse](&UserEditResponse{})
	return
}
