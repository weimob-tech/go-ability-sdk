package weimob_guide

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RewardPerformanceGetList
// @id 1753
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1753?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询导购业绩明细)
func (client *Client) RewardPerformanceGetList(request *RewardPerformanceGetListRequest) (response *RewardPerformanceGetListResponse, err error) {
	rpcResponse := CreateRewardPerformanceGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RewardPerformanceGetListRequest struct {
	*api.BaseRequest

	PageNum     int64  `json:"pageNum,omitempty"`
	PageSize    int64  `json:"pageSize,omitempty"`
	CreateTime  string `json:"createTime,omitempty"`
	EndTime     string `json:"endTime,omitempty"`
	JobNumber   string `json:"jobNumber,omitempty"`
	GuiderPhone string `json:"guiderPhone,omitempty"`
	GuiderWid   int64  `json:"guiderWid,omitempty"`
	Vid         int64  `json:"vid,omitempty"`
	VidType     int64  `json:"vidType,omitempty"`
}

type RewardPerformanceGetListResponse struct {
	PageList   []RewardPerformanceGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                      `json:"pageNum,omitempty"`
	PageSize   int64                                      `json:"pageSize,omitempty"`
	TotalCount int64                                      `json:"totalCount,omitempty"`
}

func CreateRewardPerformanceGetListRequest() (request *RewardPerformanceGetListRequest) {
	request = &RewardPerformanceGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_guide", "v2.0", "reward/performance/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateRewardPerformanceGetListResponse() (response *api.BaseResponse[RewardPerformanceGetListResponse]) {
	response = api.CreateResponse[RewardPerformanceGetListResponse](&RewardPerformanceGetListResponse{})
	return
}

type RewardPerformanceGetListResponsePageList struct {
	StoreId                     int64  `json:"storeId,omitempty"`
	TradeId                     int64  `json:"tradeId,omitempty"`
	GuiderWid                   int64  `json:"guiderWid,omitempty"`
	GuiderName                  string `json:"guiderName,omitempty"`
	GuiderPhone                 string `json:"guiderPhone,omitempty"`
	TradeType                   int64  `json:"tradeType,omitempty"`
	OrderType                   int64  `json:"orderType,omitempty"`
	PaymentAmount               int64  `json:"paymentAmount,omitempty"`
	PointsAmount                int64  `json:"pointsAmount,omitempty"`
	BalanceDiscountAmount       int64  `json:"balanceDiscountAmount,omitempty"`
	SharedPayAmount             int64  `json:"sharedPayAmount,omitempty"`
	SharedBalanceDiscountAmount int64  `json:"sharedBalanceDiscountAmount,omitempty"`
	SharedPointsAmount          int64  `json:"sharedPointsAmount,omitempty"`
	Wid                         int64  `json:"wid,omitempty"`
	UserNickname                string `json:"userNickname,omitempty"`
	TradeTime                   string `json:"tradeTime,omitempty"`
	TradeUpdateTime             string `json:"tradeUpdateTime,omitempty"`
	CommissionAmount            int64  `json:"commissionAmount,omitempty"`
	GuiderStoreName             string `json:"guiderStoreName,omitempty"`
	JobNumber                   string `json:"jobNumber,omitempty"`
}
