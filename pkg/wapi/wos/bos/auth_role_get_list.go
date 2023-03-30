package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AuthRoleGetList
// @id 2123
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2123?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询店铺下组织的角色列表)
func (client *Client) AuthRoleGetList(request *AuthRoleGetListRequest) (response *AuthRoleGetListResponse, err error) {
	rpcResponse := CreateAuthRoleGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AuthRoleGetListRequest struct {
	*api.BaseRequest

	Vid          int64   `json:"vid,omitempty"`
	RoleTypeList []int64 `json:"roleTypeList,omitempty"`
	RoleName     string  `json:"roleName,omitempty"`
	PageSize     int64   `json:"pageSize,omitempty"`
	PageNum      int64   `json:"pageNum,omitempty"`
}

type AuthRoleGetListResponse struct {
	Data      []AuthRoleGetListResponseData `json:"data,omitempty"`
	TotalSize int64                         `json:"totalSize,omitempty"`
	PageSize  int64                         `json:"pageSize,omitempty"`
	PageNum   int64                         `json:"pageNum,omitempty"`
}

func CreateAuthRoleGetListRequest() (request *AuthRoleGetListRequest) {
	request = &AuthRoleGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "auth/role/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateAuthRoleGetListResponse() (response *api.BaseResponse[AuthRoleGetListResponse]) {
	response = api.CreateResponse[AuthRoleGetListResponse](&AuthRoleGetListResponse{})
	return
}

type AuthRoleGetListResponseData struct {
	RoleDesc string `json:"roleDesc,omitempty"`
	RoleId   int64  `json:"roleId,omitempty"`
	RoleName string `json:"roleName,omitempty"`
	RoleType int64  `json:"roleType,omitempty"`
}
