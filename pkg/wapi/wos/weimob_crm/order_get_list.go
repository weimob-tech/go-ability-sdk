package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderGetList
// @id 2097
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2097?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询订单列表)
func (client *Client) OrderGetList(request *OrderGetListRequest) (response *OrderGetListResponse, err error) {
	rpcResponse := CreateOrderGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderGetListRequest struct {
	*api.BaseRequest

	PageNum        int64   `json:"pageNum,omitempty"`
	PageSize       int64   `json:"pageSize,omitempty"`
	OrderTypes     []int64 `json:"orderTypes,omitempty"`
	Statuses       []int64 `json:"statuses,omitempty"`
	Wid            int64   `json:"wid,omitempty"`
	MerchantId     int64   `json:"merchantId,omitempty"`
	OrderNos       []int64 `json:"orderNos,omitempty"`
	OrderTimeBegin string  `json:"orderTimeBegin,omitempty"`
	OrderTimeEnd   string  `json:"orderTimeEnd,omitempty"`
}

type OrderGetListResponse struct {
	OrderInfoList []OrderGetListResponseOrderInfoList `json:"orderInfoList,omitempty"`
	PageSize      int64                               `json:"pageSize,omitempty"`
	TotalCount    int64                               `json:"totalCount,omitempty"`
	PageNum       int64                               `json:"pageNum,omitempty"`
	TotalPage     int64                               `json:"totalPage,omitempty"`
}

func CreateOrderGetListRequest() (request *OrderGetListRequest) {
	request = &OrderGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "order/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderGetListResponse() (response *api.BaseResponse[OrderGetListResponse]) {
	response = api.CreateResponse[OrderGetListResponse](&OrderGetListResponse{})
	return
}

type OrderGetListResponseOrderInfoList struct {
	OrderInfo   OrderGetListResponseOrderInfo     `json:"orderInfo,omitempty"`
	RightsInfos []OrderGetListResponseRightsInfos `json:"rightsInfos,omitempty"`
}

type OrderGetListResponseOrderInfo struct {
	OrderBaseInfo     OrderGetListResponseOrderBaseInfo       `json:"orderBaseInfo,omitempty"`
	BuyerInfo         OrderGetListResponseBuyerInfo           `json:"buyerInfo,omitempty"`
	CancelInfo        OrderGetListResponseCancelInfo          `json:"cancelInfo,omitempty"`
	DiscountInfoList  []OrderGetListResponseDiscountInfoList  `json:"discountInfoList,omitempty"`
	OrderExtInfo      OrderGetListResponseOrderExtInfo        `json:"orderExtInfo,omitempty"`
	SellerInfo        OrderGetListResponseSellerInfo          `json:"sellerInfo,omitempty"`
	PaymentInfo       OrderGetListResponsePaymentInfo         `json:"paymentInfo,omitempty"`
	OrderItemInfoList []OrderGetListResponseOrderItemInfoList `json:"orderItemInfoList,omitempty"`
}

type OrderGetListResponseOrderBaseInfo struct {
	OrderType         int64  `json:"orderType,omitempty"`
	OrderSource       int64  `json:"orderSource,omitempty"`
	OrderNo           int64  `json:"orderNo,omitempty"`
	Status            int64  `json:"status,omitempty"`
	ChannelType       int64  `json:"channelType,omitempty"`
	UpdateTime        string `json:"updateTime,omitempty"`
	PaymentType       int64  `json:"paymentType,omitempty"`
	Delete            int64  `json:"delete,omitempty"`
	OrderTime         string `json:"orderTime,omitempty"`
	FinishPaymentTime string `json:"finishPaymentTime,omitempty"`
}

type OrderGetListResponseBuyerInfo struct {
	BuyerRemark string `json:"buyerRemark,omitempty"`
	Wid         int64  `json:"wid,omitempty"`
	BuyerName   string `json:"buyerName,omitempty"`
}

type OrderGetListResponseCancelInfo struct {
	CancelTime   string `json:"cancelTime,omitempty"`
	CancelType   int64  `json:"cancelType,omitempty"`
	CancelReason string `json:"cancelReason,omitempty"`
}

type OrderGetListResponseDiscountInfoList struct {
	DiscountExtInfo OrderGetListResponseDiscountExtInfo `json:"discountExtInfo,omitempty"`
	DiscountAmount  int64                               `json:"discountAmount,omitempty"`
	DiscountType    int64                               `json:"discountType,omitempty"`
	DiscountId      int64                               `json:"discountId,omitempty"`
	DiscountLevel   int64                               `json:"discountLevel,omitempty"`
	SubType         int64                               `json:"subType,omitempty"`
	Name            string                              `json:"name,omitempty"`
	CostAmount      int64                               `json:"costAmount,omitempty"`
}

type OrderGetListResponseDiscountExtInfo struct {
}

type OrderGetListResponseOrderExtInfo struct {
	OrderExtInfo OrderGetListResponseOrderExtInfo2 `json:"orderExtInfo,omitempty"`
}

type OrderGetListResponseOrderExtInfo2 struct {
	CrmGrantAmount string `json:"crmGrantAmount,omitempty"`
	BdOutOrder     string `json:"bdOutOrder,omitempty"`
}

type OrderGetListResponseSellerInfo struct {
	Vid          int64  `json:"vid,omitempty"`
	MerchantId   int64  `json:"merchantId,omitempty"`
	BosId        int64  `json:"bosId,omitempty"`
	MerchantName string `json:"merchantName,omitempty"`
	VidName      string `json:"vidName,omitempty"`
}

type OrderGetListResponsePaymentInfo struct {
	FeeInfoList         []OrderGetListResponseFeeInfoList `json:"feeInfoList,omitempty"`
	PayItems            []OrderGetListResponsePayItems    `json:"payItems,omitempty"`
	TotalAmount         int64                             `json:"totalAmount,omitempty"`
	PaymentAmount       int64                             `json:"paymentAmount,omitempty"`
	ShouldPaymentAmount int64                             `json:"shouldPaymentAmount,omitempty"`
	CurrencyType        int64                             `json:"currencyType,omitempty"`
}

type OrderGetListResponseFeeInfoList struct {
	AmountExtInfo OrderGetListResponseAmountExtInfo `json:"amountExtInfo,omitempty"`
	Amount        int64                             `json:"amount,omitempty"`
	Description   string                            `json:"description,omitempty"`
	Type          int64                             `json:"type,omitempty"`
	CurrencyType  int64                             `json:"currencyType,omitempty"`
}

type OrderGetListResponseAmountExtInfo struct {
}

type OrderGetListResponsePayItems struct {
	PayTime      int64   `json:"payTime,omitempty"`
	PayType      int64   `json:"payType,omitempty"`
	PayMethodIds []int64 `json:"payMethodIds,omitempty"`
	PayId        int64   `json:"payId,omitempty"`
	ChannelTrxNo string  `json:"channelTrxNo,omitempty"`
	PayTradeId   int64   `json:"payTradeId,omitempty"`
	Phase        int64   `json:"phase,omitempty"`
	TradeId      string  `json:"tradeId,omitempty"`
}

type OrderGetListResponseOrderItemInfoList struct {
	GoodsInfoDto     OrderGetListResponseGoodsInfoDto        `json:"goodsInfoDto,omitempty"`
	PaymentInfo      OrderGetListResponsePaymentInfo2        `json:"paymentInfo,omitempty"`
	DiscountInfoList []OrderGetListResponseDiscountInfoList2 `json:"discountInfoList,omitempty"`
	BizExtMap        OrderGetListResponseBizExtMap           `json:"bizExtMap,omitempty"`
	OrderItemExtInfo OrderGetListResponseOrderItemExtInfo    `json:"orderItemExtInfo,omitempty"`
	ItemId           int64                                   `json:"itemId,omitempty"`
	Num              int64                                   `json:"num,omitempty"`
	ActivityTypeList []int64                                 `json:"activityTypeList,omitempty"`
	TagInfo          string                                  `json:"tagInfo,omitempty"`
}

type OrderGetListResponseGoodsInfoDto struct {
	PriceInfoList []OrderGetListResponsePriceInfoList `json:"priceInfoList,omitempty"`
	GoodsType     int64                               `json:"goodsType,omitempty"`
	SubGoodsType  int64                               `json:"subGoodsType,omitempty"`
	GoodsId       int64                               `json:"goodsId,omitempty"`
	GoodsTitle    string                              `json:"goodsTitle,omitempty"`
	SkuId         int64                               `json:"skuId,omitempty"`
	SkuAttrInfo   string                              `json:"skuAttrInfo,omitempty"`
	ImageUrl      string                              `json:"imageUrl,omitempty"`
	UnitType      int64                               `json:"unitType,omitempty"`
	SalePrice     int64                               `json:"salePrice,omitempty"`
	CurrencyType  int64                               `json:"currencyType,omitempty"`
}

type OrderGetListResponsePriceInfoList struct {
	AmountExtInfo OrderGetListResponseAmountExtInfo2 `json:"amountExtInfo,omitempty"`
	Type          int64                              `json:"type,omitempty"`
	Description   string                             `json:"description,omitempty"`
	CurrencyType  int64                              `json:"currencyType,omitempty"`
	Amount        int64                              `json:"amount,omitempty"`
}

type OrderGetListResponseAmountExtInfo2 struct {
}

type OrderGetListResponsePaymentInfo2 struct {
	FeeInfoList         []OrderGetListResponseFeeInfoList2 `json:"feeInfoList,omitempty"`
	PayItems            []OrderGetListResponsePayItems2    `json:"payItems,omitempty"`
	PaymentAmount       int64                              `json:"paymentAmount,omitempty"`
	TotalAmount         int64                              `json:"totalAmount,omitempty"`
	ShouldPaymentAmount int64                              `json:"shouldPaymentAmount,omitempty"`
	CurrencyType        int64                              `json:"currencyType,omitempty"`
}

type OrderGetListResponseFeeInfoList2 struct {
	AmountExtInfo OrderGetListResponseAmountExtInfo3 `json:"amountExtInfo,omitempty"`
	Type          int64                              `json:"type,omitempty"`
	Description   string                             `json:"description,omitempty"`
	CurrencyType  int64                              `json:"currencyType,omitempty"`
	Amount        int64                              `json:"amount,omitempty"`
}

type OrderGetListResponseAmountExtInfo3 struct {
}

type OrderGetListResponsePayItems2 struct {
	PayTime      int64   `json:"payTime,omitempty"`
	PayTradeId   int64   `json:"payTradeId,omitempty"`
	Phase        int64   `json:"phase,omitempty"`
	TradeId      string  `json:"tradeId,omitempty"`
	PayId        int64   `json:"payId,omitempty"`
	PayMethodIds []int64 `json:"payMethodIds,omitempty"`
	ChannelTrxNo string  `json:"channelTrxNo,omitempty"`
	PayType      int64   `json:"payType,omitempty"`
}

type OrderGetListResponseDiscountInfoList2 struct {
}

type OrderGetListResponseBizExtMap struct {
}

type OrderGetListResponseOrderItemExtInfo struct {
}

type OrderGetListResponseRightsInfos struct {
	OrderNo          int64  `json:"orderNo,omitempty"`
	RightsStatus     int64  `json:"rightsStatus,omitempty"`
	RightsId         int64  `json:"rightsId,omitempty"`
	RightsStatusName string `json:"rightsStatusName,omitempty"`
	RightsItemId     int64  `json:"rightsItemId,omitempty"`
}
