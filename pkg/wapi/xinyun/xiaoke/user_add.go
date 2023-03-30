package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserAdd
// @id 1641
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增员工)
func (client *Client) UserAdd(request *UserAddRequest) (response *UserAddResponse, err error) {
	rpcResponse := CreateUserAddResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserAddRequest struct {
	*api.BaseRequest

	DepartmentIds []int  `json:"departmentIds,omitempty"`
	UserName      string `json:"userName,omitempty"`
	Email         string `json:"email,omitempty"`
	Phone         string `json:"phone,omitempty"`
}

type UserAddResponse struct {
	Wid int64 `json:"wid,omitempty"`
}

func CreateUserAddRequest() (request *UserAddRequest) {
	request = &UserAddRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "user/add", "api")
	request.Method = api.POST
	return
}

func CreateUserAddResponse() (response *api.BaseResponse[UserAddResponse]) {
	response = api.CreateResponse[UserAddResponse](&UserAddResponse{})
	return
}
