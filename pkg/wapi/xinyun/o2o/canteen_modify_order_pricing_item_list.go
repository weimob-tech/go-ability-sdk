package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CanteenModifyOrderPricingItemList
// @id 1410
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=点餐订单批价权益更改)
func (client *Client) CanteenModifyOrderPricingItemList(request *CanteenModifyOrderPricingItemListRequest) (response *CanteenModifyOrderPricingItemListResponse, err error) {
	rpcResponse := CreateCanteenModifyOrderPricingItemListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CanteenModifyOrderPricingItemListRequest struct {
	*api.BaseRequest

	PricingItemList []CanteenModifyOrderPricingItemListRequestPricingItemList `json:"pricingItemList,omitempty"`
	ConsumerWid     int64                                                     `json:"consumerWid,omitempty"`
	Operator        string                                                    `json:"operator,omitempty"`
	OrderNo         string                                                    `json:"orderNo,omitempty"`
	PayAmount       float64                                                   `json:"payAmount,omitempty"`
	PayWay          int                                                       `json:"payWay,omitempty"`
	StoreId         int64                                                     `json:"storeId,omitempty"`
	TotalAmount     float64                                                   `json:"totalAmount,omitempty"`
}

type CanteenModifyOrderPricingItemListResponse struct {
}

func CreateCanteenModifyOrderPricingItemListRequest() (request *CanteenModifyOrderPricingItemListRequest) {
	request = &CanteenModifyOrderPricingItemListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "canteen/modifyOrderPricingItemList", "api")
	request.Method = api.POST
	return
}

func CreateCanteenModifyOrderPricingItemListResponse() (response *api.BaseResponse[CanteenModifyOrderPricingItemListResponse]) {
	response = api.CreateResponse[CanteenModifyOrderPricingItemListResponse](&CanteenModifyOrderPricingItemListResponse{})
	return
}

type CanteenModifyOrderPricingItemListRequestPricingItemList struct {
	CardCode              string  `json:"cardCode,omitempty"`
	CardTemplateId        string  `json:"cardTemplateId,omitempty"`
	DiscountAmount        float64 `json:"discountAmount,omitempty"`
	MemberCardCode        string  `json:"memberCardCode,omitempty"`
	Point                 int64   `json:"point,omitempty"`
	PricingType           int     `json:"pricingType,omitempty"`
	PromotionDiscountRule string  `json:"promotionDiscountRule,omitempty"`
}
