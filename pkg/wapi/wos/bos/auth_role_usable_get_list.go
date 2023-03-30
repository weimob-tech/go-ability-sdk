package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AuthRoleUsableGetList
// @id 2248
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2248?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询指定组织可分配的员工角色列表)
func (client *Client) AuthRoleUsableGetList(request *AuthRoleUsableGetListRequest) (response *AuthRoleUsableGetListResponse, err error) {
	rpcResponse := CreateAuthRoleUsableGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AuthRoleUsableGetListRequest struct {
	*api.BaseRequest

	Vid      int64 `json:"vid,omitempty"`
	PageSize int64 `json:"pageSize,omitempty"`
	LastId   int64 `json:"lastId,omitempty"`
}

type AuthRoleUsableGetListResponse struct {
	RoleOpenApiVOList []AuthRoleUsableGetListResponseRoleOpenApiVOList `json:"roleOpenApiVOList,omitempty"`
	LastId            int64                                            `json:"lastId,omitempty"`
	More              bool                                             `json:"more,omitempty"`
}

func CreateAuthRoleUsableGetListRequest() (request *AuthRoleUsableGetListRequest) {
	request = &AuthRoleUsableGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "auth/role/usable/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateAuthRoleUsableGetListResponse() (response *api.BaseResponse[AuthRoleUsableGetListResponse]) {
	response = api.CreateResponse[AuthRoleUsableGetListResponse](&AuthRoleUsableGetListResponse{})
	return
}

type AuthRoleUsableGetListResponseRoleOpenApiVOList struct {
	RoleId   int64  `json:"roleId,omitempty"`
	RoleName string `json:"roleName,omitempty"`
	RoleDesc string `json:"roleDesc,omitempty"`
	RoleType int64  `json:"roleType,omitempty"`
}
