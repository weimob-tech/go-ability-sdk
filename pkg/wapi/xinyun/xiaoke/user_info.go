package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserInfo
// @id 1638
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询员工详情)
func (client *Client) UserInfo(request *UserInfoRequest) (response *UserInfoResponse, err error) {
	rpcResponse := CreateUserInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserInfoRequest struct {
	*api.BaseRequest

	Wid                    int64 `json:"wid,omitempty"`
	IsNeedDepartmentDetail int64 `json:"isNeedDepartmentDetail,omitempty"`
	IsNeedPermissionDetail int64 `json:"isNeedPermissionDetail,omitempty"`
}

type UserInfoResponse struct {
	DepartmentList       []UserInfoResponseDepartmentList       `json:"departmentList,omitempty"`
	PermissionDetailList []UserInfoResponsePermissionDetailList `json:"permissionDetailList,omitempty"`
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

func CreateUserInfoRequest() (request *UserInfoRequest) {
	request = &UserInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "user/info", "api")
	request.Method = api.POST
	return
}

func CreateUserInfoResponse() (response *api.BaseResponse[UserInfoResponse]) {
	response = api.CreateResponse[UserInfoResponse](&UserInfoResponse{})
	return
}

type UserInfoResponseDepartmentList struct {
	DepartmentName string `json:"departmentName,omitempty"`
	Level          int64  `json:"level,omitempty"`
	DepartmentId   int64  `json:"departmentId,omitempty"`
	ParentId       int64  `json:"parentId,omitempty"`
}

type UserInfoResponsePermissionDetailList struct {
	PermissionType int64  `json:"permissionType,omitempty"`
	PermissionCode string `json:"permissionCode,omitempty"`
	PermissionName string `json:"permissionName,omitempty"`
}
