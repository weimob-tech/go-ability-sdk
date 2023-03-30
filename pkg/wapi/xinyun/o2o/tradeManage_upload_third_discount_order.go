package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TradeManageUploadThirdDiscountOrder
// @id 1412
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=第三方优惠买单订单上传)
func (client *Client) TradeManageUploadThirdDiscountOrder(request *TradeManageUploadThirdDiscountOrderRequest) (response *TradeManageUploadThirdDiscountOrderResponse, err error) {
	rpcResponse := CreateTradeManageUploadThirdDiscountOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TradeManageUploadThirdDiscountOrderRequest struct {
	*api.BaseRequest

	PricingItemList []TradeManageUploadThirdDiscountOrderRequestPricingItemList `json:"pricingItemList,omitempty"`
	ConsumerWid     int64                                                       `json:"consumerWid,omitempty"`
	Operator        string                                                      `json:"operator,omitempty"`
	OrderStatus     int                                                         `json:"orderStatus,omitempty"`
	PayAmount       float64                                                     `json:"payAmount,omitempty"`
	PayWay          int                                                         `json:"payWay,omitempty"`
	StoreId         int64                                                       `json:"storeId,omitempty"`
	ThirdOrderNo    string                                                      `json:"thirdOrderNo,omitempty"`
	TotalAmount     float64                                                     `json:"totalAmount,omitempty"`
}

type TradeManageUploadThirdDiscountOrderResponse struct {
}

func CreateTradeManageUploadThirdDiscountOrderRequest() (request *TradeManageUploadThirdDiscountOrderRequest) {
	request = &TradeManageUploadThirdDiscountOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "tradeManage/uploadThirdDiscountOrder", "api")
	request.Method = api.POST
	return
}

func CreateTradeManageUploadThirdDiscountOrderResponse() (response *api.BaseResponse[TradeManageUploadThirdDiscountOrderResponse]) {
	response = api.CreateResponse[TradeManageUploadThirdDiscountOrderResponse](&TradeManageUploadThirdDiscountOrderResponse{})
	return
}

type TradeManageUploadThirdDiscountOrderRequestPricingItemList struct {
	CardCode              string  `json:"cardCode,omitempty"`
	Point                 int     `json:"point,omitempty"`
	PricingAmount         float64 `json:"pricingAmount,omitempty"`
	PricingType           int     `json:"pricingType,omitempty"`
	PromotionDiscountRule string  `json:"promotionDiscountRule,omitempty"`
	MemberLevel           string  `json:"memberLevel,omitempty"`
	MemberDiscount        int     `json:"memberDiscount,omitempty"`
}
