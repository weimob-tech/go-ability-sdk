package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderImport
// @id 2096
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2096?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=导入CRM订单)
func (client *Client) OrderImport(request *OrderImportRequest) (response *OrderImportResponse, err error) {
	rpcResponse := CreateOrderImportResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderImportRequest struct {
	*api.BaseRequest

	OrderInfo    OrderImportRequestOrderInfo `json:"orderInfo,omitempty"`
	OperatorWid  int64                       `json:"operatorWid,omitempty"`
	OperatorName string                      `json:"operatorName,omitempty"`
}

type OrderImportResponse struct {
	OutputInfo OrderImportResponseOutputInfo `json:"outputInfo,omitempty"`
	Success    bool                          `json:"success,omitempty"`
}

func CreateOrderImportRequest() (request *OrderImportRequest) {
	request = &OrderImportRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "order/import", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderImportResponse() (response *api.BaseResponse[OrderImportResponse]) {
	response = api.CreateResponse[OrderImportResponse](&OrderImportResponse{})
	return
}

type OrderImportRequestOrderInfo struct {
	OrderBaseInfo     OrderImportRequestOrderBaseInfo       `json:"orderBaseInfo,omitempty"`
	BuyerInfo         OrderImportRequestBuyerInfo           `json:"buyerInfo,omitempty"`
	SellerInfo        OrderImportRequestSellerInfo          `json:"sellerInfo,omitempty"`
	PaymentInfo       OrderImportRequestPaymentInfo         `json:"paymentInfo,omitempty"`
	DiscountInfoList  []OrderImportRequestDiscountInfoList  `json:"discountInfoList,omitempty"`
	OrderItemInfoList []OrderImportRequestOrderItemInfoList `json:"orderItemInfoList,omitempty"`
	PosInfo           OrderImportRequestPosInfo             `json:"posInfo,omitempty"`
	GuiderInfo        OrderImportRequestGuiderInfo          `json:"guiderInfo,omitempty"`
	DeliveryInfo      OrderImportRequestDeliveryInfo        `json:"deliveryInfo,omitempty"`
	AssociateBizList  []OrderImportRequestAssociateBizList  `json:"associateBizList,omitempty"`
}

type OrderImportRequestOrderBaseInfo struct {
	OrderType          int64  `json:"orderType,omitempty"`
	Status             int64  `json:"status,omitempty"`
	PaymentType        int64  `json:"paymentType,omitempty"`
	ChannelType        int64  `json:"channelType,omitempty"`
	OrderTime          int64  `json:"orderTime,omitempty"`
	FinishPaymentTime  int64  `json:"finishPaymentTime,omitempty"`
	OutOrderNo         string `json:"outOrderNo,omitempty"`
	OutOriginalOrderNo string `json:"outOriginalOrderNo,omitempty"`
	SaleChannelType    int64  `json:"saleChannelType,omitempty"`
	BizSourceType      int64  `json:"bizSourceType,omitempty"`
	PayStatus          int64  `json:"payStatus,omitempty"`
	FinishTime         int64  `json:"finishTime,omitempty"`
	Remark             string `json:"remark,omitempty"`
	PlatformType       int64  `json:"platformType,omitempty"`
}

type OrderImportRequestBuyerInfo struct {
	Wid         int64  `json:"wid,omitempty"`
	BuyerName   string `json:"buyerName,omitempty"`
	BuyerRemark string `json:"buyerRemark,omitempty"`
	BuyerPhone  string `json:"buyerPhone,omitempty"`
	OutCardNo   string `json:"outCardNo,omitempty"`
}

type OrderImportRequestSellerInfo struct {
	BosId        int64  `json:"bosId,omitempty"`
	MerchantId   int64  `json:"merchantId,omitempty"`
	MerchantName string `json:"merchantName,omitempty"`
	Vid          int64  `json:"vid,omitempty"`
	VidName      string `json:"vidName,omitempty"`
	ProcessVid   int64  `json:"processVid,omitempty"`
}

type OrderImportRequestPaymentInfo struct {
	FeeInfoList         []OrderImportRequestFeeInfoList `json:"feeInfoList,omitempty"`
	PayItems            []OrderImportRequestPayItems    `json:"payItems,omitempty"`
	PaymentAmount       string                          `json:"paymentAmount,omitempty"`
	TotalAmount         string                          `json:"totalAmount,omitempty"`
	ShouldPaymentAmount string                          `json:"shouldPaymentAmount,omitempty"`
}

type OrderImportRequestFeeInfoList struct {
	Type        int64  `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	Amount      string `json:"amount,omitempty"`
}

type OrderImportRequestPayItems struct {
	PayTime      int64   `json:"payTime,omitempty"`
	PayTradeId   int64   `json:"payTradeId,omitempty"`
	Phase        int64   `json:"phase,omitempty"`
	TradeId      int64   `json:"tradeId,omitempty"`
	PayId        int64   `json:"payId,omitempty"`
	PayMethodIds []int64 `json:"payMethodIds,omitempty"`
	ChannelTrxNo string  `json:"channelTrxNo,omitempty"`
}

type OrderImportRequestDiscountInfoList struct {
	DiscountId     string `json:"discountId,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	DiscountLevel  int64  `json:"discountLevel,omitempty"`
	SubType        int64  `json:"subType,omitempty"`
	Name           string `json:"name,omitempty"`
	DiscountAmount string `json:"discountAmount,omitempty"`
	CostAmount     string `json:"costAmount,omitempty"`
}

type OrderImportRequestOrderItemInfoList struct {
	GoodsInfoDto     OrderImportRequestGoodsInfoDto        `json:"goodsInfoDto,omitempty"`
	PaymentInfo      OrderImportRequestPaymentInfo2        `json:"paymentInfo,omitempty"`
	DiscountInfoList []OrderImportRequestDiscountInfoList2 `json:"discountInfoList,omitempty"`
	ItemId           int64                                 `json:"itemId,omitempty"`
	Num              int64                                 `json:"num,omitempty"`
	TagInfo          string                                `json:"tagInfo,omitempty"`
}

type OrderImportRequestGoodsInfoDto struct {
	GoodsType    int64  `json:"goodsType,omitempty"`
	SubGoodsType int64  `json:"subGoodsType,omitempty"`
	GoodsId      int64  `json:"goodsId,omitempty"`
	GoodsTitle   string `json:"goodsTitle,omitempty"`
	SkuId        int64  `json:"skuId,omitempty"`
	SkuAttrInfo  string `json:"skuAttrInfo,omitempty"`
	ImageUrl     string `json:"imageUrl,omitempty"`
	UnitType     int64  `json:"unitType,omitempty"`
	SalePrice    string `json:"salePrice,omitempty"`
	OutGoodsId   string `json:"outGoodsId,omitempty"`
	OutSkuId     string `json:"outSkuId,omitempty"`
	MarketPrice  string `json:"marketPrice,omitempty"`
	GuiderWid    int64  `json:"guiderWid,omitempty"`
	GuiderName   string `json:"guiderName,omitempty"`
	GuiderPhone  string `json:"guiderPhone,omitempty"`
	OutGuiderNo  string `json:"outGuiderNo,omitempty"`
	GuiderNo     string `json:"guiderNo,omitempty"`
}

type OrderImportRequestPaymentInfo2 struct {
	FeeInfoList         []OrderImportRequestFeeInfoList2 `json:"feeInfoList,omitempty"`
	PaymentAmount       string                           `json:"paymentAmount,omitempty"`
	TotalAmount         string                           `json:"totalAmount,omitempty"`
	ShouldPaymentAmount string                           `json:"shouldPaymentAmount,omitempty"`
}

type OrderImportRequestFeeInfoList2 struct {
	Type        int64  `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	Amount      string `json:"amount,omitempty"`
}

type OrderImportRequestDiscountInfoList2 struct {
	DiscountId     string `json:"discountId,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	DiscountLevel  int64  `json:"discountLevel,omitempty"`
	SubType        int64  `json:"subType,omitempty"`
	Name           string `json:"name,omitempty"`
	DiscountAmount string `json:"discountAmount,omitempty"`
	CostAmount     string `json:"costAmount,omitempty"`
}

type OrderImportRequestPosInfo struct {
	Pos          string `json:"pos,omitempty"`
	CashierName  string `json:"cashierName,omitempty"`
	CashierPhone string `json:"cashierPhone,omitempty"`
}

type OrderImportRequestGuiderInfo struct {
	GuiderName         string `json:"guiderName,omitempty"`
	GuiderPhone        string `json:"guiderPhone,omitempty"`
	GuiderWid          int64  `json:"guiderWid,omitempty"`
	PrivateGuiderName  string `json:"privateGuiderName,omitempty"`
	PrivateGuiderPhone string `json:"privateGuiderPhone,omitempty"`
	PrivateGuiderWid   int64  `json:"privateGuiderWid,omitempty"`
	OutGuiderNo        string `json:"outGuiderNo,omitempty"`
	GuiderNo           string `json:"guiderNo,omitempty"`
	PrivateGuiderNo    string `json:"privateGuiderNo,omitempty"`
	PrivateGuiderVid   int64  `json:"privateGuiderVid,omitempty"`
	GuiderId           string `json:"guiderId,omitempty"`
	PrivateGuiderId    string `json:"privateGuiderId,omitempty"`
}

type OrderImportRequestDeliveryInfo struct {
	DeliveryType   int64  `json:"deliveryType,omitempty"`
	ReceiverName   string `json:"receiverName,omitempty"`
	ReceiverMobile string `json:"receiverMobile,omitempty"`
	ProvinceCode   string `json:"provinceCode,omitempty"`
	Province       string `json:"province,omitempty"`
	CityCode       string `json:"cityCode,omitempty"`
	City           string `json:"city,omitempty"`
	CountyCode     string `json:"countyCode,omitempty"`
	County         string `json:"county,omitempty"`
	AreaCode       string `json:"areaCode,omitempty"`
	Area           string `json:"area,omitempty"`
	Address        string `json:"address,omitempty"`
	Zip            string `json:"zip,omitempty"`
	PickUpName     string `json:"pickUpName,omitempty"`
	PickUpVid      string `json:"pickUpVid,omitempty"`
}

type OrderImportRequestAssociateBizList struct {
	BizValue int64 `json:"bizValue,omitempty"`
}

type OrderImportResponseOutputInfo struct {
	OrderNo int64 `json:"orderNo,omitempty"`
}
