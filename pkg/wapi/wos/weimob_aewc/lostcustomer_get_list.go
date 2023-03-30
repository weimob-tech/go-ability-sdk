package weimob_aewc

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// LostcustomerGetList
// @id 2221
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2221?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=筛选流失客户列表)
func (client *Client) LostcustomerGetList(request *LostcustomerGetListRequest) (response *LostcustomerGetListResponse, err error) {
	rpcResponse := CreateLostcustomerGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type LostcustomerGetListRequest struct {
	*api.BaseRequest

	FriendStartTime    string  `json:"friendStartTime,omitempty"`
	FriendEndTime      string  `json:"friendEndTime,omitempty"`
	LoseStartTime      string  `json:"loseStartTime,omitempty"`
	LoseEndTime        string  `json:"loseEndTime,omitempty"`
	StartRetentionDays int64   `json:"startRetentionDays,omitempty"`
	EndRetentionDays   int64   `json:"endRetentionDays,omitempty"`
	LoseType           int64   `json:"loseType,omitempty"`
	OrgUserIdList      []int64 `json:"orgUserIdList,omitempty"`
	DeleteStartTime    int64   `json:"deleteStartTime,omitempty"`
	DeleteEndTime      int64   `json:"deleteEndTime,omitempty"`
	PageNum            int64   `json:"pageNum,omitempty"`
	PageSize           int64   `json:"pageSize,omitempty"`
}

type LostcustomerGetListResponse struct {
	Results    []LostcustomerGetListResponseResults `json:"results,omitempty"`
	PageNum    int64                                `json:"pageNum,omitempty"`
	PageSize   int64                                `json:"pageSize,omitempty"`
	TotalCount int64                                `json:"totalCount,omitempty"`
	TotalPage  int64                                `json:"totalPage,omitempty"`
}

func CreateLostcustomerGetListRequest() (request *LostcustomerGetListRequest) {
	request = &LostcustomerGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_aewc", "v2.0", "lostcustomer/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateLostcustomerGetListResponse() (response *api.BaseResponse[LostcustomerGetListResponse]) {
	response = api.CreateResponse[LostcustomerGetListResponse](&LostcustomerGetListResponse{})
	return
}

type LostcustomerGetListResponseResults struct {
	Tags          []LostcustomerGetListResponseTags `json:"tags,omitempty"`
	AddTime       int64                             `json:"addTime,omitempty"`
	CreateTime    int64                             `json:"createTime,omitempty"`
	CustomerId    string                            `json:"customerId,omitempty"`
	CustomerName  string                            `json:"customerName,omitempty"`
	Id            int64                             `json:"id,omitempty"`
	LoseTime      int64                             `json:"loseTime,omitempty"`
	LoseType      int64                             `json:"loseType,omitempty"`
	OrgUserId     string                            `json:"orgUserId,omitempty"`
	RetentionDays int64                             `json:"retentionDays,omitempty"`
	UpdateTime    int64                             `json:"updateTime,omitempty"`
	UserName      string                            `json:"userName,omitempty"`
}

type LostcustomerGetListResponseTags struct {
	TagName   string `json:"tagName,omitempty"`
	TagSource int64  `json:"tagSource,omitempty"`
}
