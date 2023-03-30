package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DepartmentTree
// @id 1634
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询部门及下级部门)
func (client *Client) DepartmentTree(request *DepartmentTreeRequest) (response *DepartmentTreeResponse, err error) {
	rpcResponse := CreateDepartmentTreeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DepartmentTreeRequest struct {
	*api.BaseRequest

	DepartmentId int64 `json:"departmentId,omitempty"`
}

type DepartmentTreeResponse struct {
	ChildList      []map[string]any `json:"childList,omitempty"`
	DepartmentName string           `json:"departmentName,omitempty"`
	AreaName       string           `json:"areaName,omitempty"`
	DepartmentId   int64            `json:"departmentId,omitempty"`
	ParentId       int64            `json:"parentId,omitempty"`
	Id             int64            `json:"id,omitempty"`
}

func CreateDepartmentTreeRequest() (request *DepartmentTreeRequest) {
	request = &DepartmentTreeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "department/tree", "api")
	request.Method = api.POST
	return
}

func CreateDepartmentTreeResponse() (response *api.BaseResponse[DepartmentTreeResponse]) {
	response = api.CreateResponse[DepartmentTreeResponse](&DepartmentTreeResponse{})
	return
}
