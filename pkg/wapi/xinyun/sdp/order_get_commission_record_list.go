package sdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderGetCommissionRecordList
// @id 747
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取商户分账记录列表)
func (client *Client) OrderGetCommissionRecordList(request *OrderGetCommissionRecordListRequest) (response *OrderGetCommissionRecordListResponse, err error) {
	rpcResponse := CreateOrderGetCommissionRecordListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderGetCommissionRecordListRequest struct {
	*api.BaseRequest

	QueryParameter OrderGetCommissionRecordListRequestQueryParameter `json:"queryParameter,omitempty"`
	BizType        int                                               `json:"bizType,omitempty"`
	BizId          int64                                             `json:"bizId,omitempty"`
	PageNum        int                                               `json:"pageNum,omitempty"`
	PageSize       int                                               `json:"pageSize,omitempty"`
}

type OrderGetCommissionRecordListResponse struct {
	PageList   []OrderGetCommissionRecordListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                          `json:"pageNum,omitempty"`
	PageSize   int64                                          `json:"pageSize,omitempty"`
	TotalCount int64                                          `json:"totalCount,omitempty"`
}

func CreateOrderGetCommissionRecordListRequest() (request *OrderGetCommissionRecordListRequest) {
	request = &OrderGetCommissionRecordListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("sdp", "1_0", "order/getCommissionRecordList", "api")
	request.Method = api.POST
	return
}

func CreateOrderGetCommissionRecordListResponse() (response *api.BaseResponse[OrderGetCommissionRecordListResponse]) {
	response = api.CreateResponse[OrderGetCommissionRecordListResponse](&OrderGetCommissionRecordListResponse{})
	return
}

type OrderGetCommissionRecordListRequestQueryParameter struct {
	Wid       int64 `json:"wid,omitempty"`
	StartTime int   `json:"startTime,omitempty"`
	EndTime   int   `json:"endTime,omitempty"`
}

type OrderGetCommissionRecordListResponsePageList struct {
	WeikeAward            OrderGetCommissionRecordListResponseWeikeAward   `json:"weikeAward,omitempty"`
	InviterAward          OrderGetCommissionRecordListResponseInviterAward `json:"inviterAward,omitempty"`
	TradeTime             int64                                            `json:"tradeTime,omitempty"`
	Wid                   int64                                            `json:"wid,omitempty"`
	TradeType             int64                                            `json:"tradeType,omitempty"`
	TradeId               int64                                            `json:"tradeId,omitempty"`
	SplitAmount           int64                                            `json:"splitAmount,omitempty"`
	BalanceDiscountAmount int64                                            `json:"balanceDiscountAmount,omitempty"`
	TradeAmount           int64                                            `json:"tradeAmount,omitempty"`
	RightsRefundBalance   int64                                            `json:"rightsRefundBalance,omitempty"`
	RightsRefundAmount    int64                                            `json:"rightsRefundAmount,omitempty"`
	RewardType            string                                           `json:"rewardType,omitempty"`
	State                 string                                           `json:"state,omitempty"`
	UpdateTime            int64                                            `json:"updateTime,omitempty"`
	ScrollId              int64                                            `json:"scrollId,omitempty"`
}

type OrderGetCommissionRecordListResponseWeikeAward struct {
	Points  int64 `json:"points,omitempty"`
	Amount  int64 `json:"amount,omitempty"`
	Balance int64 `json:"balance,omitempty"`
}

type OrderGetCommissionRecordListResponseInviterAward struct {
	Points  int64 `json:"points,omitempty"`
	Amount  int64 `json:"amount,omitempty"`
	Balance int64 `json:"balance,omitempty"`
}
