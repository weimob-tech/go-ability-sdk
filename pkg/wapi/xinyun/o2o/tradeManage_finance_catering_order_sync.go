package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TradeManageFinanceCateringOrderSync
// @id 2108
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=财务餐饮订单同步)
func (client *Client) TradeManageFinanceCateringOrderSync(request *TradeManageFinanceCateringOrderSyncRequest) (response *TradeManageFinanceCateringOrderSyncResponse, err error) {
	rpcResponse := CreateTradeManageFinanceCateringOrderSyncResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TradeManageFinanceCateringOrderSyncRequest struct {
	*api.BaseRequest

	DishList       []map[string]any `json:"dishList,omitempty"`
	TradeId        string           `json:"tradeId,omitempty"`
	TraceNo        string           `json:"traceNo,omitempty"`
	OutShopNo      string           `json:"outShopNo,omitempty"`
	OpenId         string           `json:"openId,omitempty"`
	LoginToken     string           `json:"loginToken,omitempty"`
	OrderEntry     string           `json:"orderEntry,omitempty"`
	TotalAmount    string           `json:"totalAmount,omitempty"`
	DiscountAmount string           `json:"discountAmount,omitempty"`
	UserAmount     string           `json:"userAmount,omitempty"`
	OutTableNo     string           `json:"outTableNo,omitempty"`
	PeopleCount    int              `json:"peopleCount,omitempty"`
}

type TradeManageFinanceCateringOrderSyncResponse struct {
}

func CreateTradeManageFinanceCateringOrderSyncRequest() (request *TradeManageFinanceCateringOrderSyncRequest) {
	request = &TradeManageFinanceCateringOrderSyncRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "tradeManage/financeCateringOrderSync", "api")
	request.Method = api.POST
	return
}

func CreateTradeManageFinanceCateringOrderSyncResponse() (response *api.BaseResponse[TradeManageFinanceCateringOrderSyncResponse]) {
	response = api.CreateResponse[TradeManageFinanceCateringOrderSyncResponse](&TradeManageFinanceCateringOrderSyncResponse{})
	return
}
