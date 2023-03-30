package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// EmployeeRoleGet
// @id 2081
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2081?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询员工组织下的角色)
func (client *Client) EmployeeRoleGet(request *EmployeeRoleGetRequest) (response *EmployeeRoleGetResponse, err error) {
	rpcResponse := CreateEmployeeRoleGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type EmployeeRoleGetRequest struct {
	*api.BaseRequest

	Vid      int64 `json:"vid,omitempty"`
	Wid      int64 `json:"wid,omitempty"`
	PageNum  int64 `json:"pageNum,omitempty"`
	PageSize int64 `json:"pageSize,omitempty"`
}

type EmployeeRoleGetResponse struct {
	Data      []EmployeeRoleGetResponseData `json:"data,omitempty"`
	PageSize  int64                         `json:"pageSize,omitempty"`
	PageNum   int64                         `json:"pageNum,omitempty"`
	TotalSize int64                         `json:"totalSize,omitempty"`
}

func CreateEmployeeRoleGetRequest() (request *EmployeeRoleGetRequest) {
	request = &EmployeeRoleGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "employee/role/get", "apigw")
	request.Method = api.POST
	return
}

func CreateEmployeeRoleGetResponse() (response *api.BaseResponse[EmployeeRoleGetResponse]) {
	response = api.CreateResponse[EmployeeRoleGetResponse](&EmployeeRoleGetResponse{})
	return
}

type EmployeeRoleGetResponseData struct {
	RoleName string `json:"roleName,omitempty"`
	RoleId   int64  `json:"roleId,omitempty"`
}
