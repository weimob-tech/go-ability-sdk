package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// EmployeeRoleUpdate
// @id 2126
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2126?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新员工组织角色关系)
func (client *Client) EmployeeRoleUpdate(request *EmployeeRoleUpdateRequest) (response *EmployeeRoleUpdateResponse, err error) {
	rpcResponse := CreateEmployeeRoleUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type EmployeeRoleUpdateRequest struct {
	*api.BaseRequest

	Wid         int64   `json:"wid,omitempty"`
	VidList     []int64 `json:"vidList,omitempty"`
	RoleIdList  []int64 `json:"roleIdList,omitempty"`
	OperateType int64   `json:"operateType,omitempty"`
	VidType     int64   `json:"vidType,omitempty"`
}

type EmployeeRoleUpdateResponse struct {
}

func CreateEmployeeRoleUpdateRequest() (request *EmployeeRoleUpdateRequest) {
	request = &EmployeeRoleUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "employee/role/update", "apigw")
	request.Method = api.POST
	return
}

func CreateEmployeeRoleUpdateResponse() (response *api.BaseResponse[EmployeeRoleUpdateResponse]) {
	response = api.CreateResponse[EmployeeRoleUpdateResponse](&EmployeeRoleUpdateResponse{})
	return
}
