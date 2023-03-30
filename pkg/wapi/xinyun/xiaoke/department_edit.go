package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DepartmentEdit
// @id 1633
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=编辑部门)
func (client *Client) DepartmentEdit(request *DepartmentEditRequest) (response *DepartmentEditResponse, err error) {
	rpcResponse := CreateDepartmentEditResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DepartmentEditRequest struct {
	*api.BaseRequest

	DepartmentId int64  `json:"departmentId,omitempty"`
	Name         string `json:"name,omitempty"`
}

type DepartmentEditResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateDepartmentEditRequest() (request *DepartmentEditRequest) {
	request = &DepartmentEditRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "department/edit", "api")
	request.Method = api.POST
	return
}

func CreateDepartmentEditResponse() (response *api.BaseResponse[DepartmentEditResponse]) {
	response = api.CreateResponse[DepartmentEditResponse](&DepartmentEditResponse{})
	return
}
