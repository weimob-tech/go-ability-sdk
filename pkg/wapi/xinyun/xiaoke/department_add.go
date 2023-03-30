package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DepartmentAdd
// @id 1632
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=创建部门)
func (client *Client) DepartmentAdd(request *DepartmentAddRequest) (response *DepartmentAddResponse, err error) {
	rpcResponse := CreateDepartmentAddResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DepartmentAddRequest struct {
	*api.BaseRequest

	ParentDepartmentId int64  `json:"parentDepartmentId,omitempty"`
	Name               string `json:"name,omitempty"`
}

type DepartmentAddResponse struct {
	DepartmentId int64 `json:"departmentId,omitempty"`
}

func CreateDepartmentAddRequest() (request *DepartmentAddRequest) {
	request = &DepartmentAddRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "department/add", "api")
	request.Method = api.POST
	return
}

func CreateDepartmentAddResponse() (response *api.BaseResponse[DepartmentAddResponse]) {
	response = api.CreateResponse[DepartmentAddResponse](&DepartmentAddResponse{})
	return
}
