package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserList
// @id 1637
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询员工列表)
func (client *Client) UserList(request *UserListRequest) (response *UserListResponse, err error) {
	rpcResponse := CreateUserListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserListRequest struct {
	*api.BaseRequest

	WidList                []int `json:"widList,omitempty"`
	IsNeedDepartmentDetail int64 `json:"isNeedDepartmentDetail,omitempty"`
	IsNeedPermissionDetail int64 `json:"isNeedPermissionDetail,omitempty"`
}

type UserListResponse struct {
	DepartmentList       []UserListResponseDepartmentList       `json:"departmentList,omitempty"`
	PermissionDetailList []UserListResponsePermissionDetailList `json:"permissionDetailList,omitempty"`
	Wid                  int64                                  `json:"wid,omitempty"`
	Zone                 string                                 `json:"zone,omitempty"`
	Email                string                                 `json:"email,omitempty"`
	Qq                   string                                 `json:"qq,omitempty"`
	HeadImage            string                                 `json:"headImage,omitempty"`
	IsSeat               int64                                  `json:"isSeat,omitempty"`
	Telephone            string                                 `json:"telephone,omitempty"`
	UserName             string                                 `json:"userName,omitempty"`
	Weixin               string                                 `json:"weixin,omitempty"`
	Position             string                                 `json:"position,omitempty"`
	UserType             int64                                  `json:"userType,omitempty"`
}

func CreateUserListRequest() (request *UserListRequest) {
	request = &UserListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "user/list", "api")
	request.Method = api.POST
	return
}

func CreateUserListResponse() (response *api.BaseResponse[UserListResponse]) {
	response = api.CreateResponse[UserListResponse](&UserListResponse{})
	return
}

type UserListResponseDepartmentList struct {
	DepartmentName string `json:"departmentName,omitempty"`
	Level          int64  `json:"level,omitempty"`
	DepartmentId   int64  `json:"departmentId,omitempty"`
	ParentId       int64  `json:"parentId,omitempty"`
}

type UserListResponsePermissionDetailList struct {
	PermissionType int64  `json:"permissionType,omitempty"`
	PermissionCode string `json:"permissionCode,omitempty"`
	PermissionName string `json:"permissionName,omitempty"`
}
