package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FinanceBillCustomTransactionGet
// @id 1738
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1738?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=海关支付单反查接口)
func (client *Client) FinanceBillCustomTransactionGet(request *FinanceBillCustomTransactionGetRequest) (response *FinanceBillCustomTransactionGetResponse, err error) {
	rpcResponse := CreateFinanceBillCustomTransactionGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FinanceBillCustomTransactionGetRequest struct {
	*api.BaseRequest

	LastId int64  `json:"lastId,omitempty"`
	BosId  string `json:"bosId,omitempty"`
}

type FinanceBillCustomTransactionGetResponse struct {
	List         []FinanceBillCustomTransactionGetResponselist `json:"list,omitempty"`
	LastId       int64                                         `json:"lastId,omitempty"`
	HaveNextPage bool                                          `json:"haveNextPage,omitempty"`
}

func CreateFinanceBillCustomTransactionGetRequest() (request *FinanceBillCustomTransactionGetRequest) {
	request = &FinanceBillCustomTransactionGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "finance/bill/custom/transaction/get", "apigw")
	request.Method = api.POST
	return
}

func CreateFinanceBillCustomTransactionGetResponse() (response *api.BaseResponse[FinanceBillCustomTransactionGetResponse]) {
	response = api.CreateResponse[FinanceBillCustomTransactionGetResponse](&FinanceBillCustomTransactionGetResponse{})
	return
}

type FinanceBillCustomTransactionGetResponselist struct {
	TradingTime      FinanceBillCustomTransactionGetResponseTradingTime `json:"tradingTime,omitempty"`
	Guid             string                                             `json:"guid,omitempty"`
	TradeId          string                                             `json:"tradeId,omitempty"`
	Fid              string                                             `json:"fid,omitempty"`
	InitalRequest    string                                             `json:"initalRequest,omitempty"`
	InitalResponse   string                                             `json:"initalResponse,omitempty"`
	EbpCode          string                                             `json:"ebpCode,omitempty"`
	PayCode          string                                             `json:"payCode,omitempty"`
	PayTransactionId string                                             `json:"payTransactionId,omitempty"`
	TotalAmount      string                                             `json:"totalAmount,omitempty"`
	Currency         string                                             `json:"currency,omitempty"`
	VerDept          string                                             `json:"verDept,omitempty"`
	PayType          string                                             `json:"payType,omitempty"`
	OrderNo          string                                             `json:"orderNo,omitempty"`
	GoodsInfo        string                                             `json:"goodsInfo,omitempty"`
	RecpAccount      string                                             `json:"recpAccount,omitempty"`
	RecpCode         string                                             `json:"recpCode,omitempty"`
	RecpName         string                                             `json:"recpName,omitempty"`
	SessionId        string                                             `json:"sessionId,omitempty"`
}

type FinanceBillCustomTransactionGetResponseTradingTime struct {
}
