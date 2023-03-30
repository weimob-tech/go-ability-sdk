package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderGet
// @id 2066
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2066?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询订单详情)
func (client *Client) OrderGet(request *OrderGetRequest) (response *OrderGetResponse, err error) {
	rpcResponse := CreateOrderGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderGetRequest struct {
	*api.BaseRequest

	OrderNo int64 `json:"orderNo,omitempty"`
}

type OrderGetResponse struct {
	OrderInfo   OrderGetResponseOrderInfo     `json:"orderInfo,omitempty"`
	RightsInfos []OrderGetResponseRightsInfos `json:"rightsInfos,omitempty"`
}

func CreateOrderGetRequest() (request *OrderGetRequest) {
	request = &OrderGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "order/get", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderGetResponse() (response *api.BaseResponse[OrderGetResponse]) {
	response = api.CreateResponse[OrderGetResponse](&OrderGetResponse{})
	return
}

type OrderGetResponseOrderInfo struct {
	OrderExtInfo      OrderGetResponseOrderExtInfo        `json:"orderExtInfo,omitempty"`
	OrderBaseInfo     OrderGetResponseOrderBaseInfo       `json:"orderBaseInfo,omitempty"`
	BuyerInfo         OrderGetResponseBuyerInfo           `json:"buyerInfo,omitempty"`
	PaymentInfo       OrderGetResponsePaymentInfo         `json:"paymentInfo,omitempty"`
	DiscountInfoList  []OrderGetResponseDiscountInfoList  `json:"discountInfoList,omitempty"`
	SellerInfo        OrderGetResponseSellerInfo          `json:"sellerInfo,omitempty"`
	CancelInfo        OrderGetResponseCancelInfo          `json:"cancelInfo,omitempty"`
	OrderItemInfoList []OrderGetResponseOrderItemInfoList `json:"orderItemInfoList,omitempty"`
}

type OrderGetResponseOrderExtInfo struct {
	OrderExtInfo OrderGetResponseOrderExtInfo2 `json:"orderExtInfo,omitempty"`
}

type OrderGetResponseOrderExtInfo2 struct {
}

type OrderGetResponseOrderBaseInfo struct {
	OrderNo           int64  `json:"orderNo,omitempty"`
	OrderType         int64  `json:"orderType,omitempty"`
	Status            int64  `json:"status,omitempty"`
	PaymentType       int64  `json:"paymentType,omitempty"`
	ChannelType       int64  `json:"channelType,omitempty"`
	OrderTime         string `json:"orderTime,omitempty"`
	Delete            int64  `json:"delete,omitempty"`
	OrderSource       int64  `json:"orderSource,omitempty"`
	UpdateTime        string `json:"updateTime,omitempty"`
	FinishPaymentTime string `json:"finishPaymentTime,omitempty"`
}

type OrderGetResponseBuyerInfo struct {
	Wid         int64  `json:"wid,omitempty"`
	BuyerRemark string `json:"buyerRemark,omitempty"`
	BuyerName   string `json:"buyerName,omitempty"`
}

type OrderGetResponsePaymentInfo struct {
	PayItems            []OrderGetResponsePayItems    `json:"payItems,omitempty"`
	FeeInfoList         []OrderGetResponseFeeInfoList `json:"feeInfoList,omitempty"`
	PaymentAmount       int64                         `json:"paymentAmount,omitempty"`
	TotalAmount         int64                         `json:"totalAmount,omitempty"`
	ShouldPaymentAmount int64                         `json:"shouldPaymentAmount,omitempty"`
	CurrencyType        int64                         `json:"currencyType,omitempty"`
}

type OrderGetResponsePayItems struct {
	PayType      int64   `json:"payType,omitempty"`
	TradeId      string  `json:"tradeId,omitempty"`
	PayId        int64   `json:"payId,omitempty"`
	PayTradeId   int64   `json:"payTradeId,omitempty"`
	ChannelTrxNo string  `json:"channelTrxNo,omitempty"`
	Phase        int64   `json:"phase,omitempty"`
	PayTime      int64   `json:"payTime,omitempty"`
	PayMethodIds []int64 `json:"payMethodIds,omitempty"`
}

type OrderGetResponseFeeInfoList struct {
	AmountExtInfo OrderGetResponseAmountExtInfo `json:"amountExtInfo,omitempty"`
	Type          int64                         `json:"type,omitempty"`
	CurrencyType  int64                         `json:"currencyType,omitempty"`
	Description   string                        `json:"description,omitempty"`
	Amount        int64                         `json:"amount,omitempty"`
}

type OrderGetResponseAmountExtInfo struct {
}

type OrderGetResponseDiscountInfoList struct {
	DiscountExtInfo OrderGetResponseDiscountExtInfo `json:"discountExtInfo,omitempty"`
	DiscountType    int64                           `json:"discountType,omitempty"`
	SubType         int64                           `json:"subType,omitempty"`
	CostAmount      int64                           `json:"costAmount,omitempty"`
	Name            string                          `json:"name,omitempty"`
	DiscountLevel   int64                           `json:"discountLevel,omitempty"`
	DiscountAmount  int64                           `json:"discountAmount,omitempty"`
	DiscountId      string                          `json:"discountId,omitempty"`
}

type OrderGetResponseDiscountExtInfo struct {
}

type OrderGetResponseSellerInfo struct {
	BosId        int64  `json:"bosId,omitempty"`
	MerchantId   int64  `json:"merchantId,omitempty"`
	Vid          int64  `json:"vid,omitempty"`
	MerchantName string `json:"merchantName,omitempty"`
	VidName      string `json:"vidName,omitempty"`
}

type OrderGetResponseCancelInfo struct {
	CancelReason string `json:"cancelReason,omitempty"`
	CancelTime   string `json:"cancelTime,omitempty"`
	CancelType   int64  `json:"cancelType,omitempty"`
}

type OrderGetResponseOrderItemInfoList struct {
	GoodsInfoDto     OrderGetResponseGoodsInfoDto        `json:"goodsInfoDto,omitempty"`
	PaymentInfo      OrderGetResponsePaymentInfo2        `json:"paymentInfo,omitempty"`
	DiscountInfoList []OrderGetResponseDiscountInfoList2 `json:"discountInfoList,omitempty"`
	BizExtMap        OrderGetResponseBizExtMap           `json:"bizExtMap,omitempty"`
	OrderItemExtInfo OrderGetResponseOrderItemExtInfo    `json:"orderItemExtInfo,omitempty"`
	ItemId           int64                               `json:"itemId,omitempty"`
	Num              int64                               `json:"num,omitempty"`
	ActivityTypeList []int64                             `json:"activityTypeList,omitempty"`
	TagInfo          string                              `json:"tagInfo,omitempty"`
}

type OrderGetResponseGoodsInfoDto struct {
	PriceInfoList []OrderGetResponsePriceInfoList `json:"priceInfoList,omitempty"`
	GoodsType     int64                           `json:"goodsType,omitempty"`
	GoodsId       int64                           `json:"goodsId,omitempty"`
	SkuId         int64                           `json:"skuId,omitempty"`
	SalePrice     int64                           `json:"salePrice,omitempty"`
	ImageUrl      string                          `json:"imageUrl,omitempty"`
	SkuAttrInfo   string                          `json:"skuAttrInfo,omitempty"`
	CurrencyType  int64                           `json:"currencyType,omitempty"`
	GoodsTitle    string                          `json:"goodsTitle,omitempty"`
	UnitType      int64                           `json:"unitType,omitempty"`
	SubGoodsType  int64                           `json:"subGoodsType,omitempty"`
}

type OrderGetResponsePriceInfoList struct {
	AmountExtInfo OrderGetResponseAmountExtInfo2 `json:"amountExtInfo,omitempty"`
	Amount        int64                          `json:"amount,omitempty"`
	Type          int64                          `json:"type,omitempty"`
	CurrencyType  int64                          `json:"currencyType,omitempty"`
	Description   string                         `json:"description,omitempty"`
}

type OrderGetResponseAmountExtInfo2 struct {
}

type OrderGetResponsePaymentInfo2 struct {
	PayItems            []OrderGetResponsePayItems2    `json:"payItems,omitempty"`
	FeeInfoList         []OrderGetResponseFeeInfoList2 `json:"feeInfoList,omitempty"`
	PaymentAmount       int64                          `json:"paymentAmount,omitempty"`
	TotalAmount         int64                          `json:"totalAmount,omitempty"`
	ShouldPaymentAmount int64                          `json:"shouldPaymentAmount,omitempty"`
	CurrencyType        int64                          `json:"currencyType,omitempty"`
}

type OrderGetResponsePayItems2 struct {
	PayTradeId   int64   `json:"payTradeId,omitempty"`
	PayTime      int64   `json:"payTime,omitempty"`
	TradeId      string  `json:"tradeId,omitempty"`
	PayId        int64   `json:"payId,omitempty"`
	Phase        int64   `json:"phase,omitempty"`
	PayType      int64   `json:"payType,omitempty"`
	ChannelTrxNo string  `json:"channelTrxNo,omitempty"`
	PayMethodIds []int64 `json:"payMethodIds,omitempty"`
}

type OrderGetResponseFeeInfoList2 struct {
	AmountExtInfo OrderGetResponseAmountExtInfo3 `json:"amountExtInfo,omitempty"`
	Type          int64                          `json:"type,omitempty"`
	Description   string                         `json:"description,omitempty"`
	Amount        int64                          `json:"amount,omitempty"`
	CurrencyType  int64                          `json:"currencyType,omitempty"`
}

type OrderGetResponseAmountExtInfo3 struct {
}

type OrderGetResponseDiscountInfoList2 struct {
	DiscountExtInfo OrderGetResponseDiscountExtInfo2 `json:"discountExtInfo,omitempty"`
	DiscountType    int64                            `json:"discountType,omitempty"`
	Name            string                           `json:"name,omitempty"`
	DiscountAmount  int64                            `json:"discountAmount,omitempty"`
	CostAmount      int64                            `json:"costAmount,omitempty"`
	DiscountLevel   int64                            `json:"discountLevel,omitempty"`
	DiscountId      string                           `json:"discountId,omitempty"`
	SubType         int64                            `json:"subType,omitempty"`
}

type OrderGetResponseDiscountExtInfo2 struct {
}

type OrderGetResponseBizExtMap struct {
}

type OrderGetResponseOrderItemExtInfo struct {
}

type OrderGetResponseRightsInfos struct {
	RightsId         int64  `json:"rightsId,omitempty"`
	RightsStatusName string `json:"rightsStatusName,omitempty"`
	OrderNo          int64  `json:"orderNo,omitempty"`
	RightsItemId     int64  `json:"rightsItemId,omitempty"`
	RightsStatus     int64  `json:"rightsStatus,omitempty"`
}
