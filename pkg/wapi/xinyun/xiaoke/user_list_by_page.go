package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserListByPage
// @id 1635
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=分页查询员工列表)
func (client *Client) UserListByPage(request *UserListByPageRequest) (response *UserListByPageResponse, err error) {
	rpcResponse := CreateUserListByPageResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserListByPageRequest struct {
	*api.BaseRequest

	WidList                []int `json:"widList,omitempty"`
	IsNeedDepartmentDetail int   `json:"isNeedDepartmentDetail,omitempty"`
	IsNeedPermissionDetail int   `json:"isNeedPermissionDetail,omitempty"`
	PageNum                int   `json:"pageNum,omitempty"`
	PageSize               int   `json:"pageSize,omitempty"`
}

type UserListByPageResponse struct {
	DepartmentList       []UserListByPageResponseDepartmentList       `json:"departmentList,omitempty"`
	PermissionDetailList []UserListByPageResponsePermissionDetailList `json:"permissionDetailList,omitempty"`
	Wid                  int64                                        `json:"wid,omitempty"`
	Zone                 string                                       `json:"zone,omitempty"`
	Email                string                                       `json:"email,omitempty"`
	Qq                   string                                       `json:"qq,omitempty"`
	HeadImage            string                                       `json:"headImage,omitempty"`
	IsSeat               int64                                        `json:"isSeat,omitempty"`
	Telephone            string                                       `json:"telephone,omitempty"`
	UserName             string                                       `json:"userName,omitempty"`
	Weixin               string                                       `json:"weixin,omitempty"`
	Position             string                                       `json:"position,omitempty"`
	UserType             int64                                        `json:"userType,omitempty"`
	PageNum              int64                                        `json:"pageNum,omitempty"`
	TotalCount           int64                                        `json:"totalCount,omitempty"`
	PageSize             int64                                        `json:"pageSize,omitempty"`
	TotalPage            int64                                        `json:"totalPage,omitempty"`
}

func CreateUserListByPageRequest() (request *UserListByPageRequest) {
	request = &UserListByPageRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "user/listByPage", "api")
	request.Method = api.POST
	return
}

func CreateUserListByPageResponse() (response *api.BaseResponse[UserListByPageResponse]) {
	response = api.CreateResponse[UserListByPageResponse](&UserListByPageResponse{})
	return
}

type UserListByPageResponseDepartmentList struct {
	DepartmentName string `json:"departmentName,omitempty"`
	Level          int    `json:"level,omitempty"`
	DepartmentId   int64  `json:"departmentId,omitempty"`
	ParentId       int64  `json:"parentId,omitempty"`
}

type UserListByPageResponsePermissionDetailList struct {
	PermissionType int64  `json:"permissionType,omitempty"`
	PermissionCode string `json:"permissionCode,omitempty"`
	PermissionName string `json:"permissionName,omitempty"`
}
