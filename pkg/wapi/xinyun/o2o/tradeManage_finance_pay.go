package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TradeManageFinancePay
// @id 2109
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=财务支付)
func (client *Client) TradeManageFinancePay(request *TradeManageFinancePayRequest) (response *TradeManageFinancePayResponse, err error) {
	rpcResponse := CreateTradeManageFinancePayResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TradeManageFinancePayRequest struct {
	*api.BaseRequest

	TradeId        string `json:"tradeId,omitempty"`
	PayAmount      int64  `json:"payAmount,omitempty"`
	ProductNo      int    `json:"productNo,omitempty"`
	ExpireTime     int64  `json:"expireTime,omitempty"`
	OpenId         string `json:"openId,omitempty"`
	AppId          string `json:"appId,omitempty"`
	ClientIp       string `json:"clientIp,omitempty"`
	GoodsDetail    string `json:"goodsDetail,omitempty"`
	AuthCode       string `json:"authCode,omitempty"`
	ContractId     string `json:"contractId,omitempty"`
	ReturnUrl      string `json:"returnUrl,omitempty"`
	UserId         string `json:"userId,omitempty"`
	BackUrl        string `json:"backUrl,omitempty"`
	IdNo           string `json:"idNo,omitempty"`
	IdName         string `json:"idName,omitempty"`
	RiskInfo       string `json:"riskInfo,omitempty"`
	PrepaidCardId  string `json:"prepaidCardId,omitempty"`
	ProfitSharing  string `json:"profitSharing,omitempty"`
	Currency       string `json:"currency,omitempty"`
	SecuredFlag    int    `json:"securedFlag,omitempty"`
	SplitFid       string `json:"splitFid,omitempty"`
	Uid            string `json:"uid,omitempty"`
	TraceNo        string `json:"traceNo,omitempty"`
	CouponDetail   string `json:"couponDetail,omitempty"`
	GoodsTag       string `json:"goodsTag,omitempty"`
	BuyerId        string `json:"buyerId,omitempty"`
	FoodOrderType  string `json:"foodOrderType,omitempty"`
	InstallmentNum int    `json:"installmentNum,omitempty"`
	Phone          string `json:"phone,omitempty"`
	Attach         string `json:"attach,omitempty"`
}

type TradeManageFinancePayResponse struct {
}

func CreateTradeManageFinancePayRequest() (request *TradeManageFinancePayRequest) {
	request = &TradeManageFinancePayRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "tradeManage/financePay", "api")
	request.Method = api.POST
	return
}

func CreateTradeManageFinancePayResponse() (response *api.BaseResponse[TradeManageFinancePayResponse]) {
	response = api.CreateResponse[TradeManageFinancePayResponse](&TradeManageFinancePayResponse{})
	return
}
