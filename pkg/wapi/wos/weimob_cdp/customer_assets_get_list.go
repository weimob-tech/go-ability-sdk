package weimob_cdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerAssetsGetList
// @id 1541
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1541?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询客户资产列表)
func (client *Client) CustomerAssetsGetList(request *CustomerAssetsGetListRequest) (response *CustomerAssetsGetListResponse, err error) {
	rpcResponse := CreateCustomerAssetsGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerAssetsGetListRequest struct {
	*api.BaseRequest

	TargetProductInstanceId int64   `json:"targetProductInstanceId,omitempty"`
	UkeyList                []int64 `json:"ukeyList,omitempty"`
	StartDay                string  `json:"startDay,omitempty"`
	EndDay                  string  `json:"endDay,omitempty"`
	AssetsType              int64   `json:"assetsType,omitempty"`
	PageNum                 int64   `json:"pageNum,omitempty"`
	PageSize                int64   `json:"pageSize,omitempty"`
	BosId                   int64   `json:"bosId,omitempty"`
}

type CustomerAssetsGetListResponse struct {
	CustomerGrowthVal []CustomerAssetsGetListResponseCustomerGrowthVal `json:"CustomerGrowthVal,omitempty"`
	CustomerBalance   []CustomerAssetsGetListResponseCustomerBalance   `json:"CustomerBalance,omitempty"`
	CustomerPoint     CustomerAssetsGetListResponseCustomerPoint       `json:"CustomerPoint,omitempty"`
}

func CreateCustomerAssetsGetListRequest() (request *CustomerAssetsGetListRequest) {
	request = &CustomerAssetsGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_cdp", "v2.0", "customer/assets/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerAssetsGetListResponse() (response *api.BaseResponse[CustomerAssetsGetListResponse]) {
	response = api.CreateResponse[CustomerAssetsGetListResponse](&CustomerAssetsGetListResponse{})
	return
}

type CustomerAssetsGetListResponseCustomerGrowthVal struct {
	UkeyType         string `json:"ukeyType,omitempty"`
	Ukey             string `json:"ukey,omitempty"`
	ChangeValue      int64  `json:"changeValue,omitempty"`
	BosId            int64  `json:"bosId,omitempty"`
	MembershipPlanId string `json:"membershipPlanId,omitempty"`
	ChannelId        string `json:"channelId,omitempty"`
}

type CustomerAssetsGetListResponseCustomerBalance struct {
	List     []CustomerAssetsGetListResponselist `json:"list,omitempty"`
	PageSize int64                               `json:"pageSize,omitempty"`
	Total    int64                               `json:"total,omitempty"`
	PageNum  int64                               `json:"pageNum,omitempty"`
}

type CustomerAssetsGetListResponselist struct {
	ChannelType      string `json:"channelType,omitempty"`
	AcctId           int64  `json:"acctId,omitempty"`
	FrozenAmount     int64  `json:"frozenAmount,omitempty"`
	BosId            int64  `json:"bosId,omitempty"`
	Vid              int64  `json:"vid,omitempty"`
	PlanId           int64  `json:"planId,omitempty"`
	Ukey             string `json:"ukey,omitempty"`
	SchannelType     int64  `json:"schannelType,omitempty"`
	CurrentAmount    int64  `json:"currentAmount,omitempty"`
	SchannelId       string `json:"schannelId,omitempty"`
	ChannelId        string `json:"channelId,omitempty"`
	UpdateTime       int64  `json:"updateTime,omitempty"`
	IsDeleted        int64  `json:"isDeleted,omitempty"`
	TotalAmount      int64  `json:"totalAmount,omitempty"`
	UkeyType         string `json:"ukeyType,omitempty"`
	TotalIssueAmount int64  `json:"totalIssueAmount,omitempty"`
	CreateTime       int64  `json:"createTime,omitempty"`
}

type CustomerAssetsGetListResponseCustomerPoint struct {
	List     []CustomerAssetsGetListResponselist2 `json:"list,omitempty"`
	Total    int64                                `json:"total,omitempty"`
	PageNum  int64                                `json:"pageNum,omitempty"`
	PageSize int64                                `json:"pageSize,omitempty"`
}

type CustomerAssetsGetListResponselist2 struct {
	PlanId          int64  `json:"planId,omitempty"`
	CreateTime      int64  `json:"createTime,omitempty"`
	ChannelType     string `json:"channelType,omitempty"`
	IsDeleted       int64  `json:"isDeleted,omitempty"`
	TotalPoint      int64  `json:"totalPoint,omitempty"`
	UpdateTime      int64  `json:"updateTime,omitempty"`
	Ukey            string `json:"ukey,omitempty"`
	CurrentPoint    int64  `json:"currentPoint,omitempty"`
	Vid             int64  `json:"vid,omitempty"`
	FrozenPoint     int64  `json:"frozenPoint,omitempty"`
	SchannelType    string `json:"schannelType,omitempty"`
	BosId           int64  `json:"bosId,omitempty"`
	TotalIssuePoint int64  `json:"totalIssuePoint,omitempty"`
	SchannelId      string `json:"schannelId,omitempty"`
	AcctId          int64  `json:"acctId,omitempty"`
	UkeyType        string `json:"ukeyType,omitempty"`
	ChannelId       int64  `json:"channelId,omitempty"`
}
