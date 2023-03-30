package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AuthRoleGet
// @id 2122
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2122?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询角色的基本信息)
func (client *Client) AuthRoleGet(request *AuthRoleGetRequest) (response *AuthRoleGetResponse, err error) {
	rpcResponse := CreateAuthRoleGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AuthRoleGetRequest struct {
	*api.BaseRequest

	RoleId      int64 `json:"roleId,omitempty"`
	NeedVidType int64 `json:"needVidType,omitempty"`
}

type AuthRoleGetResponse struct {
	RoleId   AuthRoleGetResponseRoleId `json:"roleId,omitempty"`
	RoleDesc string                    `json:"roleDesc,omitempty"`
	RoleName string                    `json:"roleName,omitempty"`
	RoleType int64                     `json:"roleType,omitempty"`
	VidTypes []int64                   `json:"vidTypes,omitempty"`
}

func CreateAuthRoleGetRequest() (request *AuthRoleGetRequest) {
	request = &AuthRoleGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "auth/role/get", "apigw")
	request.Method = api.POST
	return
}

func CreateAuthRoleGetResponse() (response *api.BaseResponse[AuthRoleGetResponse]) {
	response = api.CreateResponse[AuthRoleGetResponse](&AuthRoleGetResponse{})
	return
}

type AuthRoleGetResponseRoleId struct {
}
