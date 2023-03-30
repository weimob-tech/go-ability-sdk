package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TradeManageFinanceCreateTrade
// @id 2107
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=财务创建交易)
func (client *Client) TradeManageFinanceCreateTrade(request *TradeManageFinanceCreateTradeRequest) (response *TradeManageFinanceCreateTradeResponse, err error) {
	rpcResponse := CreateTradeManageFinanceCreateTradeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TradeManageFinanceCreateTradeRequest struct {
	*api.BaseRequest

	TraceNo         string `json:"traceNo,omitempty"`
	OrderNo         string `json:"orderNo,omitempty"`
	RequestTime     int    `json:"requestTime,omitempty"`
	TradeDesc       string `json:"tradeDesc,omitempty"`
	TradeDetail     string `json:"tradeDetail,omitempty"`
	TradeAmount     int64  `json:"tradeAmount,omitempty"`
	Currency        string `json:"currency,omitempty"`
	TradeType       int    `json:"tradeType,omitempty"`
	NotifyUrl       string `json:"notifyUrl,omitempty"`
	ReturnUrl       string `json:"returnUrl,omitempty"`
	CouponAmount    int64  `json:"couponAmount,omitempty"`
	CouponType      string `json:"couponType,omitempty"`
	SecuredFlag     int    `json:"securedFlag,omitempty"`
	SecondSplitFlag int    `json:"secondSplitFlag,omitempty"`
	RiskInfo        string `json:"riskInfo,omitempty"`
	ExtParams       string `json:"extParams,omitempty"`
	PaymentMode     string `json:"paymentMode,omitempty"`
	ProductNo       int    `json:"productNo,omitempty"`
	BusinessCode    string `json:"businessCode,omitempty"`
	Fid             string `json:"fid,omitempty"`
	ChannelType     int    `json:"channelType,omitempty"`
	OriginalId      string `json:"originalId,omitempty"`
}

type TradeManageFinanceCreateTradeResponse struct {
}

func CreateTradeManageFinanceCreateTradeRequest() (request *TradeManageFinanceCreateTradeRequest) {
	request = &TradeManageFinanceCreateTradeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "tradeManage/financeCreateTrade", "api")
	request.Method = api.POST
	return
}

func CreateTradeManageFinanceCreateTradeResponse() (response *api.BaseResponse[TradeManageFinanceCreateTradeResponse]) {
	response = api.CreateResponse[TradeManageFinanceCreateTradeResponse](&TradeManageFinanceCreateTradeResponse{})
	return
}
