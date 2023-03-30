package weimob_aewc

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StaffGetList
// @id 2215
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2215?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询员工/部门信息)
func (client *Client) StaffGetList(request *StaffGetListRequest) (response *StaffGetListResponse, err error) {
	rpcResponse := CreateStaffGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StaffGetListRequest struct {
	*api.BaseRequest

	Keyword           string `json:"keyword,omitempty"`
	KeyType           string `json:"keyType,omitempty"`
	FetchAllStatus    bool   `json:"fetchAllStatus,omitempty"`
	ExpectInvalidSeat bool   `json:"expectInvalidSeat,omitempty"`
	Pid               int64  `json:"pid,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	BosId             int64  `json:"bosId,omitempty"`
	Vid               int64  `json:"vid,omitempty"`
	SuiteId           string `json:"suiteId,omitempty"`
	CorpId            string `json:"corpId,omitempty"`
	AgentId           string `json:"agentId,omitempty"`
	PageNumber        int64  `json:"pageNumber,omitempty"`
	PageSize          int64  `json:"pageSize,omitempty"`
}

type StaffGetListResponse struct {
	OrgDepartmentBOS    []StaffGetListResponseOrgDepartmentBOS `json:"orgDepartmentBOS,omitempty"`
	OrgUserSimpleBOS    []StaffGetListResponseOrgUserSimpleBOS `json:"orgUserSimpleBOS,omitempty"`
	UserExistMark       bool                                   `json:"userExistMark,omitempty"`
	DepartmentExistMark bool                                   `json:"departmentExistMark,omitempty"`
}

func CreateStaffGetListRequest() (request *StaffGetListRequest) {
	request = &StaffGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_aewc", "v2.0", "staff/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateStaffGetListResponse() (response *api.BaseResponse[StaffGetListResponse]) {
	response = api.CreateResponse[StaffGetListResponse](&StaffGetListResponse{})
	return
}

type StaffGetListResponseOrgDepartmentBOS struct {
	Pid               int64  `json:"pid,omitempty"`
	BosId             int64  `json:"bosId,omitempty"`
	Vid               int64  `json:"vid,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	DepartmentId      int64  `json:"departmentId,omitempty"`
	CorpDepartmentId  int64  `json:"corpDepartmentId,omitempty"`
	CorpId            string `json:"corpId,omitempty"`
	Name              string `json:"name,omitempty"`
	ParentId          int64  `json:"parentId,omitempty"`
	Level             int64  `json:"level,omitempty"`
	Sort              int64  `json:"sort,omitempty"`
	Deleted           bool   `json:"deleted,omitempty"`
}

type StaffGetListResponseOrgUserSimpleBOS struct {
	OrgDepartmentListBOS    []StaffGetListResponseOrgDepartmentListBOS `json:"orgDepartmentListBOS,omitempty"`
	OrgUserId               string                                     `json:"orgUserId,omitempty"`
	CorpId                  string                                     `json:"corpId,omitempty"`
	UserId                  string                                     `json:"userId,omitempty"`
	Name                    string                                     `json:"name,omitempty"`
	Mobile                  string                                     `json:"mobile,omitempty"`
	Gender                  int64                                      `json:"gender,omitempty"`
	Position                string                                     `json:"position,omitempty"`
	Email                   string                                     `json:"email,omitempty"`
	IsLeaderInDept          bool                                       `json:"isLeaderInDept,omitempty"`
	Avatar                  string                                     `json:"avatar,omitempty"`
	ThumbAvatar             string                                     `json:"thumbAvatar,omitempty"`
	Telephone               string                                     `json:"telephone,omitempty"`
	Alias                   string                                     `json:"alias,omitempty"`
	Status                  int64                                      `json:"status,omitempty"`
	Extattr                 string                                     `json:"extattr,omitempty"`
	QrCode                  string                                     `json:"qrCode,omitempty"`
	ExternalProfile         string                                     `json:"externalProfile,omitempty"`
	ExternalPosition        string                                     `json:"externalPosition,omitempty"`
	Address                 string                                     `json:"address,omitempty"`
	OpenUserid              string                                     `json:"openUserid,omitempty"`
	LeaderInDept            []int64                                    `json:"leaderInDept,omitempty"`
	Deleted                 bool                                       `json:"deleted,omitempty"`
	CreateTime              string                                     `json:"createTime,omitempty"`
	HaveSessionArchiving    bool                                       `json:"haveSessionArchiving,omitempty"`
	ActiveStatus            int64                                      `json:"activeStatus,omitempty"`
	HaveBindGuide           bool                                       `json:"haveBindGuide,omitempty"`
	GuiderProductInstanceId int64                                      `json:"guiderProductInstanceId,omitempty"`
	GuiderWid               int64                                      `json:"guiderWid,omitempty"`
	BelongToStoreName       string                                     `json:"belongToStoreName,omitempty"`
	BelongToStoreId         int64                                      `json:"belongToStoreId,omitempty"`
	GuideName               string                                     `json:"guideName,omitempty"`
	GuideDetailUrl          string                                     `json:"guideDetailUrl,omitempty"`
	BcActiveStatus          int64                                      `json:"bcActiveStatus,omitempty"`
}

type StaffGetListResponseOrgDepartmentListBOS struct {
	Pid               int64  `json:"pid,omitempty"`
	BosId             int64  `json:"bosId,omitempty"`
	Vid               int64  `json:"vid,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	DepartmentId      int64  `json:"departmentId,omitempty"`
	CorpDepartmentId  int64  `json:"corpDepartmentId,omitempty"`
	CorpId            string `json:"corpId,omitempty"`
	Name              string `json:"name,omitempty"`
	ParentId          int64  `json:"parentId,omitempty"`
	Level             int64  `json:"level,omitempty"`
	Sort              int64  `json:"sort,omitempty"`
	Deleted           bool   `json:"deleted,omitempty"`
}
