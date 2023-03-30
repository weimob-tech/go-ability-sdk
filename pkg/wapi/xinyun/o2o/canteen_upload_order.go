package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CanteenUploadOrder
// @id 1409
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=上传点餐订单)
func (client *Client) CanteenUploadOrder(request *CanteenUploadOrderRequest) (response *CanteenUploadOrderResponse, err error) {
	rpcResponse := CreateCanteenUploadOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CanteenUploadOrderRequest struct {
	*api.BaseRequest

	OrderDetailList []CanteenUploadOrderRequestOrderDetailList `json:"orderDetailList,omitempty"`
	PricingItemList []CanteenUploadOrderRequestPricingItemList `json:"pricingItemList,omitempty"`
	ConsumerWid     int64                                      `json:"consumerWid,omitempty"`
	DeskName        string                                     `json:"deskName,omitempty"`
	Operator        string                                     `json:"operator,omitempty"`
	OrderSource     string                                     `json:"orderSource,omitempty"`
	PayAmount       float64                                    `json:"payAmount,omitempty"`
	PayStatus       int                                        `json:"payStatus,omitempty"`
	PayWay          int                                        `json:"payWay,omitempty"`
	StoreId         int64                                      `json:"storeId,omitempty"`
	ThirdOrderNo    string                                     `json:"thirdOrderNo,omitempty"`
	TotalAmount     float64                                    `json:"totalAmount,omitempty"`
}

type CanteenUploadOrderResponse struct {
}

func CreateCanteenUploadOrderRequest() (request *CanteenUploadOrderRequest) {
	request = &CanteenUploadOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "canteen/uploadOrder", "api")
	request.Method = api.POST
	return
}

func CreateCanteenUploadOrderResponse() (response *api.BaseResponse[CanteenUploadOrderResponse]) {
	response = api.CreateResponse[CanteenUploadOrderResponse](&CanteenUploadOrderResponse{})
	return
}

type CanteenUploadOrderRequestOrderDetailList struct {
	AdditionList   []CanteenUploadOrderRequestAdditionList   `json:"additionList,omitempty"`
	GoodsComboList []CanteenUploadOrderRequestGoodsComboList `json:"goodsComboList,omitempty"`
	SkuInfo        CanteenUploadOrderRequestSkuInfo          `json:"skuInfo,omitempty"`
	TasteList      []CanteenUploadOrderRequestTasteList      `json:"tasteList,omitempty"`
	GoodsId        int64                                     `json:"goodsId,omitempty"`
	GoodsName      string                                    `json:"goodsName,omitempty"`
	GoodsNum       int                                       `json:"goodsNum,omitempty"`
	GoodsPrice     float64                                   `json:"goodsPrice,omitempty"`
	GoodsType      int                                       `json:"goodsType,omitempty"`
	ThirdGoodsId   string                                    `json:"thirdGoodsId,omitempty"`
}

type CanteenUploadOrderRequestAdditionList struct {
	AdditionId      int64   `json:"additionId,omitempty"`
	AdditionName    string  `json:"additionName,omitempty"`
	AdditionNum     int     `json:"additionNum,omitempty"`
	AdditionPrice   float64 `json:"additionPrice,omitempty"`
	AdditionThirdId string  `json:"additionThirdId,omitempty"`
}

type CanteenUploadOrderRequestGoodsComboList struct {
	GoodsComboId    int64   `json:"goodsComboId,omitempty"`
	GoodsComboName  string  `json:"goodsComboName,omitempty"`
	GoodsComboPrice float64 `json:"goodsComboPrice,omitempty"`
	ThirdGoodsId    string  `json:"thirdGoodsId,omitempty"`
}

type CanteenUploadOrderRequestSkuInfo struct {
	SkuId      int64   `json:"skuId,omitempty"`
	SkuName    string  `json:"skuName,omitempty"`
	SkuPrice   float64 `json:"skuPrice,omitempty"`
	ThirdSkuId string  `json:"thirdSkuId,omitempty"`
}

type CanteenUploadOrderRequestTasteList struct {
	TasteId   int64  `json:"tasteId,omitempty"`
	TasteName string `json:"tasteName,omitempty"`
}

type CanteenUploadOrderRequestPricingItemList struct {
	CardCode               string  `json:"cardCode,omitempty"`
	CardTemplateId         string  `json:"cardTemplateId,omitempty"`
	DiscountAmount         float64 `json:"discountAmount,omitempty"`
	MemberCardCode         string  `json:"memberCardCode,omitempty"`
	Point                  int64   `json:"point,omitempty"`
	PricingType            int     `json:"pricingType,omitempty"`
	PromotionDiscountRule  string  `json:"promotionDiscountRule,omitempty"`
	GiftCardTicket         string  `json:"giftCardTicket,omitempty"`
	GiftCardNo             string  `json:"giftCardNo,omitempty"`
	GiftCardType           int     `json:"giftCardType,omitempty"`
	GiftCardDeductingMoney float64 `json:"giftCardDeductingMoney,omitempty"`
}
