package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TradeInitConfirmOrder
// @id 1825
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=初始化结算)
func (client *Client) TradeInitConfirmOrder(request *TradeInitConfirmOrderRequest) (response *TradeInitConfirmOrderResponse, err error) {
	rpcResponse := CreateTradeInitConfirmOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TradeInitConfirmOrderRequest struct {
	*api.BaseRequest

	GoodsList   []TradeInitConfirmOrderRequestGoodsList `json:"goodsList,omitempty"`
	TradeWid    int64                                   `json:"tradeWid,omitempty"`
	BizLineType int64                                   `json:"bizLineType,omitempty"`
	Source      int64                                   `json:"source,omitempty"`
	TemplateId  int64                                   `json:"templateId,omitempty"`
}

type TradeInitConfirmOrderResponse struct {
	ValidBizResult  TradeInitConfirmOrderResponseValidBizResult `json:"validBizResult,omitempty"`
	BizLineType     int64                                       `json:"bizLineType,omitempty"`
	ConfirmOrderKey string                                      `json:"confirmOrderKey,omitempty"`
	NextPage        string                                      `json:"nextPage,omitempty"`
	TradeTrackId    string                                      `json:"tradeTrackId,omitempty"`
}

func CreateTradeInitConfirmOrderRequest() (request *TradeInitConfirmOrderRequest) {
	request = &TradeInitConfirmOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "trade/initConfirmOrder", "api")
	request.Method = api.POST
	return
}

func CreateTradeInitConfirmOrderResponse() (response *api.BaseResponse[TradeInitConfirmOrderResponse]) {
	response = api.CreateResponse[TradeInitConfirmOrderResponse](&TradeInitConfirmOrderResponse{})
	return
}

type TradeInitConfirmOrderRequestGoodsList struct {
	InputActivities []TradeInitConfirmOrderRequestInputActivities `json:"inputActivities,omitempty"`
	StoreId         int64                                         `json:"storeId,omitempty"`
	GoodsId         int64                                         `json:"goodsId,omitempty"`
	SkuId           int64                                         `json:"skuId,omitempty"`
	BuyNum          int64                                         `json:"buyNum,omitempty"`
	TradeGoodsType  int64                                         `json:"tradeGoodsType,omitempty"`
}

type TradeInitConfirmOrderRequestInputActivities struct {
	ActivityId   int64 `json:"activityId,omitempty"`
	ActivityType int64 `json:"activityType,omitempty"`
}

type TradeInitConfirmOrderResponseValidBizResult struct {
	ValidBizInfo int64 `json:"validBizInfo,omitempty"`
	ValidBizType int64 `json:"validBizType,omitempty"`
	ValidSuccess bool  `json:"validSuccess,omitempty"`
}
