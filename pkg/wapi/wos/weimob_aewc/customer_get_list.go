package weimob_aewc

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerGetList
// @id 2223
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2223?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=筛选客户列表)
func (client *Client) CustomerGetList(request *CustomerGetListRequest) (response *CustomerGetListResponse, err error) {
	rpcResponse := CreateCustomerGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerGetListRequest struct {
	*api.BaseRequest

	NickName           string  `json:"nickName,omitempty"`
	Gender             int64   `json:"gender,omitempty"`
	TagIds             []int64 `json:"tagIds,omitempty"`
	AttIds             []int64 `json:"attIds,omitempty"`
	AddMultiUsers      bool    `json:"addMultiUsers,omitempty"`
	AddFriendStartTime string  `json:"addFriendStartTime,omitempty"`
	AddFriendEndTime   string  `json:"addFriendEndTime,omitempty"`
	FollowStatus       bool    `json:"followStatus,omitempty"`
	OrgUserIdList      []int64 `json:"orgUserIdList,omitempty"`
	Phone              string  `json:"phone,omitempty"`
	AddWay             int64   `json:"addWay,omitempty"`
	SubAddWay          int64   `json:"subAddWay,omitempty"`
	PageNum            int64   `json:"pageNum,omitempty"`
	PageSize           int64   `json:"pageSize,omitempty"`
}

type CustomerGetListResponse struct {
	Results        []CustomerGetListResponseResults `json:"results,omitempty"`
	LoseTotalCount int64                            `json:"loseTotalCount,omitempty"`
	TotalCount     int64                            `json:"totalCount,omitempty"`
}

func CreateCustomerGetListRequest() (request *CustomerGetListRequest) {
	request = &CustomerGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_aewc", "v2.0", "customer/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerGetListResponse() (response *api.BaseResponse[CustomerGetListResponse]) {
	response = api.CreateResponse[CustomerGetListResponse](&CustomerGetListResponse{})
	return
}

type CustomerGetListResponseResults struct {
	TagBOList        []CustomerGetListResponseTagBOList `json:"tagBOList,omitempty"`
	UserList         []CustomerGetListResponseUserList  `json:"userList,omitempty"`
	AddWay           int64                              `json:"addWay,omitempty"`
	AddWayDesc       string                             `json:"addWayDesc,omitempty"`
	Avatar           string                             `json:"avatar,omitempty"`
	CorpId           string                             `json:"corpId,omitempty"`
	CustomerId       string                             `json:"customerId,omitempty"`
	ExternalUserId   string                             `json:"externalUserId,omitempty"`
	ExternalUserName string                             `json:"externalUserName,omitempty"`
	ExternalUserType int64                              `json:"externalUserType,omitempty"`
	FlowStatus       bool                               `json:"flowStatus,omitempty"`
	Gender           int64                              `json:"gender,omitempty"`
	Phone            string                             `json:"phone,omitempty"`
	RecentlyAddTime  int64                              `json:"recentlyAddTime,omitempty"`
	Status           int64                              `json:"status,omitempty"`
	Unionid          string                             `json:"unionid,omitempty"`
}

type CustomerGetListResponseTagBOList struct {
	TagName   string `json:"tagName,omitempty"`
	TagSource int64  `json:"tagSource,omitempty"`
}

type CustomerGetListResponseUserList struct {
	OrgUserId string `json:"orgUserId,omitempty"`
	UserName  string `json:"userName,omitempty"`
}
