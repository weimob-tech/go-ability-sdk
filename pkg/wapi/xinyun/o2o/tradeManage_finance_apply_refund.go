package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TradeManageFinanceApplyRefund
// @id 2110
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=财务申请退款)
func (client *Client) TradeManageFinanceApplyRefund(request *TradeManageFinanceApplyRefundRequest) (response *TradeManageFinanceApplyRefundResponse, err error) {
	rpcResponse := CreateTradeManageFinanceApplyRefundResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TradeManageFinanceApplyRefundRequest struct {
	*api.BaseRequest

	DetailVOList      []TradeManageFinanceApplyRefundRequestDetailVOList `json:"detailVOList,omitempty"`
	TradeId           string                                             `json:"tradeId,omitempty"`
	OrderNo           string                                             `json:"orderNo,omitempty"`
	RefundNo          string                                             `json:"refundNo,omitempty"`
	RefundAmount      int64                                              `json:"refundAmount,omitempty"`
	RefundType        int                                                `json:"refundType,omitempty"`
	Source            int                                                `json:"source,omitempty"`
	RefundDesc        string                                             `json:"refundDesc,omitempty"`
	BizName           string                                             `json:"bizName,omitempty"`
	NotifyUrl         string                                             `json:"notifyUrl,omitempty"`
	Fid               string                                             `json:"fid,omitempty"`
	GoodsRefundNo     string                                             `json:"goodsRefundNo,omitempty"`
	CouponDetail      string                                             `json:"couponDetail,omitempty"`
	OrderRefundAmount int64                                              `json:"orderRefundAmount,omitempty"`
}

type TradeManageFinanceApplyRefundResponse struct {
}

func CreateTradeManageFinanceApplyRefundRequest() (request *TradeManageFinanceApplyRefundRequest) {
	request = &TradeManageFinanceApplyRefundRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "tradeManage/financeApplyRefund", "api")
	request.Method = api.POST
	return
}

func CreateTradeManageFinanceApplyRefundResponse() (response *api.BaseResponse[TradeManageFinanceApplyRefundResponse]) {
	response = api.CreateResponse[TradeManageFinanceApplyRefundResponse](&TradeManageFinanceApplyRefundResponse{})
	return
}

type TradeManageFinanceApplyRefundRequestDetailVOList struct {
	DetailSharingNo string `json:"detailSharingNo,omitempty"`
	RefundAmount    int64  `json:"refundAmount,omitempty"`
}
