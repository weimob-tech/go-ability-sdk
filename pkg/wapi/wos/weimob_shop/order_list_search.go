package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderListSearch
// @id 1621
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1621?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询订单列表)
func (client *Client) OrderListSearch(request *OrderListSearchRequest) (response *OrderListSearchResponse, err error) {
	rpcResponse := CreateOrderListSearchResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderListSearchRequest struct {
	*api.BaseRequest

	QueryParameter OrderListSearchRequestQueryParameter `json:"queryParameter,omitempty"`
	PageNum        int64                                `json:"pageNum,omitempty"`
	PageSize       int64                                `json:"pageSize,omitempty"`
}

type OrderListSearchResponse struct {
	PageList   []OrderListSearchResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                             `json:"pageNum,omitempty"`
	PageSize   int64                             `json:"pageSize,omitempty"`
	TotalCount int64                             `json:"totalCount,omitempty"`
}

func CreateOrderListSearchRequest() (request *OrderListSearchRequest) {
	request = &OrderListSearchRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "order/list/search", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderListSearchResponse() (response *api.BaseResponse[OrderListSearchResponse]) {
	response = api.CreateResponse[OrderListSearchResponse](&OrderListSearchResponse{})
	return
}

type OrderListSearchRequestQueryParameter struct {
	QueryTime      OrderListSearchRequestQueryTime `json:"queryTime,omitempty"`
	Activities     []int64                         `json:"activities,omitempty"`
	ChannelTypes   []int64                         `json:"channelTypes,omitempty"`
	DeliveryTypes  []int64                         `json:"deliveryTypes,omitempty"`
	GoodsSellModes []int64                         `json:"goodsSellModes,omitempty"`
	GoodsTypes     []int64                         `json:"goodsTypes,omitempty"`
	OrderDomains   []int64                         `json:"orderDomains,omitempty"`
	OrderSources   []int64                         `json:"orderSources,omitempty"`
	OrderTypes     []int64                         `json:"orderTypes,omitempty"`
	PaymentTypes   []int64                         `json:"paymentTypes,omitempty"`
	ProcessVids    []int64                         `json:"processVids,omitempty"`
	Statuses       []int64                         `json:"statuses,omitempty"`
	SubGoodsTypes  []int64                         `json:"subGoodsTypes,omitempty"`
	UserWid        int64                           `json:"userWid,omitempty"`
	ParentOrderNo  int64                           `json:"parentOrderNo,omitempty"`
	SearchType     int64                           `json:"searchType,omitempty"`
	Keyword        string                          `json:"keyword,omitempty"`
}

type OrderListSearchRequestQueryTime struct {
	EndTime   int64 `json:"endTime,omitempty"`
	StartTime int64 `json:"startTime,omitempty"`
	Type      int64 `json:"type,omitempty"`
}

type OrderListSearchResponsePageList struct {
	OrderInfo        OrderListSearchResponseOrderInfo          `json:"orderInfo,omitempty"`
	RightsInfos      []OrderListSearchResponseRightsInfos      `json:"rightsInfos,omitempty"`
	FulfillOrderList []OrderListSearchResponseFulfillOrderList `json:"fulfillOrderList,omitempty"`
}

type OrderListSearchResponseOrderInfo struct {
	BuyerInfo      OrderListSearchResponseBuyerInfo        `json:"buyerInfo,omitempty"`
	CancelInfo     OrderListSearchResponseCancelInfo       `json:"cancelInfo,omitempty"`
	DiscountInfos  []OrderListSearchResponseDiscountInfos  `json:"discountInfos,omitempty"`
	FlagInfo       OrderListSearchResponseFlagInfo         `json:"flagInfo,omitempty"`
	Items          []OrderListSearchResponseItems          `json:"items,omitempty"`
	MerchantInfo   OrderListSearchResponseMerchantInfo     `json:"merchantInfo,omitempty"`
	OrderBaseInfo  OrderListSearchResponseOrderBaseInfo    `json:"orderBaseInfo,omitempty"`
	OrderBizExt    OrderListSearchResponseOrderBizExt      `json:"orderBizExt,omitempty"`
	OrderFulfill   OrderListSearchResponseOrderFulfill     `json:"orderFulfill,omitempty"`
	PayInfo        OrderListSearchResponsePayInfo2         `json:"payInfo,omitempty"`
	TotalDiscounts []OrderListSearchResponseTotalDiscounts `json:"totalDiscounts,omitempty"`
	GuideInfo      OrderListSearchResponseGuideInfo        `json:"guideInfo,omitempty"`
	BookOrderInfo  OrderListSearchResponseBookOrderInfo    `json:"bookOrderInfo,omitempty"`
}

type OrderListSearchResponseBuyerInfo struct {
	MemberBenefits []OrderListSearchResponseMemberBenefits `json:"memberBenefits,omitempty"`
	UserNickName   string                                  `json:"userNickName,omitempty"`
	Wid            int64                                   `json:"wid,omitempty"`
	BuyerRemark    string                                  `json:"buyerRemark,omitempty"`
}

type OrderListSearchResponseMemberBenefits struct {
	BenefitType int64 `json:"benefitType,omitempty"`
}

type OrderListSearchResponseCancelInfo struct {
	CancelType    int64  `json:"cancelType,omitempty"`
	Id            int64  `json:"id,omitempty"`
	Reason        string `json:"reason,omitempty"`
	SpecialReason string `json:"specialReason,omitempty"`
}

type OrderListSearchResponseDiscountInfos struct {
	DiscountExt    OrderListSearchResponseDiscountExt `json:"discountExt,omitempty"`
	CostAmount     string                             `json:"costAmount,omitempty"`
	DiscountAmount string                             `json:"discountAmount,omitempty"`
	DiscountId     string                             `json:"discountId,omitempty"`
	DiscountLevel  int64                              `json:"discountLevel,omitempty"`
	DiscountType   int64                              `json:"discountType,omitempty"`
	Name           string                             `json:"name,omitempty"`
	SubType        int64                              `json:"subType,omitempty"`
}

type OrderListSearchResponseDiscountExt struct {
	AttributionType int64 `json:"attributionType,omitempty"`
}

type OrderListSearchResponseFlagInfo struct {
	FlagContent string `json:"flagContent,omitempty"`
	FlagRank    int64  `json:"flagRank,omitempty"`
}

type OrderListSearchResponseItems struct {
	DiscountInfos    []OrderListSearchResponseDiscountInfos2 `json:"discountInfos,omitempty"`
	GoodsExt         OrderListSearchResponseGoodsExt         `json:"goodsExt,omitempty"`
	ItemBizExt       OrderListSearchResponseItemBizExt       `json:"itemBizExt,omitempty"`
	PayInfo          OrderListSearchResponsePayInfo          `json:"payInfo,omitempty"`
	PriceInfos       []OrderListSearchResponsePriceInfos     `json:"priceInfos,omitempty"`
	ProductInfos     []OrderListSearchResponseProductInfos   `json:"productInfos,omitempty"`
	ActivityTypeList []int64                                 `json:"activityTypeList,omitempty"`
	CategoryId       int64                                   `json:"categoryId,omitempty"`
	CategoryTitle    []int64                                 `json:"categoryTitle,omitempty"`
	GoodsCode        string                                  `json:"goodsCode,omitempty"`
	GoodsId          int64                                   `json:"goodsId,omitempty"`
	GoodsTitle       string                                  `json:"goodsTitle,omitempty"`
	GoodsType        int64                                   `json:"goodsType,omitempty"`
	ImageUrl         string                                  `json:"imageUrl,omitempty"`
	ItemId           int64                                   `json:"itemId,omitempty"`
	SalePrice        string                                  `json:"salePrice,omitempty"`
	SkuAttrInfo      string                                  `json:"skuAttrInfo,omitempty"`
	SkuBarCode       string                                  `json:"skuBarCode,omitempty"`
	SkuCode          string                                  `json:"skuCode,omitempty"`
	SkuId            int64                                   `json:"skuId,omitempty"`
	SkuNum           string                                  `json:"skuNum,omitempty"`
	SubGoodsType     int64                                   `json:"subGoodsType,omitempty"`
	UnitType         int64                                   `json:"unitType,omitempty"`
}

type OrderListSearchResponseDiscountInfos2 struct {
	DiscountExt    OrderListSearchResponseDiscountExt2 `json:"discountExt,omitempty"`
	CostAmount     string                              `json:"costAmount,omitempty"`
	DiscountAmount string                              `json:"discountAmount,omitempty"`
	DiscountId     string                              `json:"discountId,omitempty"`
	DiscountLevel  int64                               `json:"discountLevel,omitempty"`
	DiscountType   int64                               `json:"discountType,omitempty"`
	Name           string                              `json:"name,omitempty"`
	SubType        int64                               `json:"subType,omitempty"`
}

type OrderListSearchResponseDiscountExt2 struct {
	AttributionType int64 `json:"attributionType,omitempty"`
}

type OrderListSearchResponseGoodsExt struct {
	GroupInfos      []OrderListSearchResponseGroupInfos `json:"groupInfos,omitempty"`
	OriginSkuName   string                              `json:"originSkuName,omitempty"`
	OriginSkuNum    int64                               `json:"originSkuNum,omitempty"`
	ProductCategory string                              `json:"productCategory,omitempty"`
}

type OrderListSearchResponseGroupInfos struct {
	Grade   int64 `json:"grade,omitempty"`
	GroupId int64 `json:"groupId,omitempty"`
}

type OrderListSearchResponseItemBizExt struct {
	BizInfos              []OrderListSearchResponseBizInfos          `json:"bizInfos,omitempty"`
	GoodsCustom           OrderListSearchResponseGoodsCustom         `json:"goodsCustom,omitempty"`
	LabelInfos            []OrderListSearchResponseLabelInfos        `json:"labelInfos,omitempty"`
	GoodsMessageInfos     []OrderListSearchResponseGoodsMessageInfos `json:"goodsMessageInfos,omitempty"`
	ActivityStockType     int64                                      `json:"activityStockType,omitempty"`
	ExpandField           string                                     `json:"expandField,omitempty"`
	GoodsGuideType        int64                                      `json:"goodsGuideType,omitempty"`
	GoodsLimitSwitch      int64                                      `json:"goodsLimitSwitch,omitempty"`
	GoodsPromotionOrderId string                                     `json:"goodsPromotionOrderId,omitempty"`
	GoodsPromotionType    int64                                      `json:"goodsPromotionType,omitempty"`
	GoodsSellMode         int64                                      `json:"goodsSellMode,omitempty"`
	GoodsSourceType       int64                                      `json:"goodsSourceType,omitempty"`
	OuterGoodsId          string                                     `json:"outerGoodsId,omitempty"`
	OuterMerchantId       string                                     `json:"outerMerchantId,omitempty"`
	OuterSkuId            string                                     `json:"outerSkuId,omitempty"`
	RightsServiceType     int64                                      `json:"rightsServiceType,omitempty"`
}

type OrderListSearchResponseBizInfos struct {
	BizId      int64  `json:"bizId,omitempty"`
	BizOrderId string `json:"bizOrderId,omitempty"`
	BizType    int64  `json:"bizType,omitempty"`
	SubBizType int64  `json:"subBizType,omitempty"`
}

type OrderListSearchResponseGoodsCustom struct {
	CustomFields []OrderListSearchResponseCustomFields `json:"customFields,omitempty"`
	CloudCustom  string                                `json:"cloudCustom,omitempty"`
}

type OrderListSearchResponseCustomFields struct {
	Key   string `json:"key,omitempty"`
	Name  string `json:"name,omitempty"`
	Sort  int64  `json:"sort,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type OrderListSearchResponseLabelInfos struct {
	AttachId   string `json:"attachId,omitempty"`
	Attachment string `json:"attachment,omitempty"`
	LabelType  string `json:"labelType,omitempty"`
}

type OrderListSearchResponseGoodsMessageInfos struct {
}

type OrderListSearchResponsePayInfo struct {
	AmountInfos         []OrderListSearchResponseAmountInfos `json:"amountInfos,omitempty"`
	PayAmount           string                               `json:"payAmount,omitempty"`
	ShouldPayAmount     string                               `json:"shouldPayAmount,omitempty"`
	TotalAmount         string                               `json:"totalAmount,omitempty"`
	TotalDiscountAmount string                               `json:"totalDiscountAmount,omitempty"`
}

type OrderListSearchResponseAmountInfos struct {
	Amount          string `json:"amount,omitempty"`
	Description     string `json:"description,omitempty"`
	PayAmount       string `json:"payAmount,omitempty"`
	ShouldPayAmount string `json:"shouldPayAmount,omitempty"`
	Type            int64  `json:"type,omitempty"`
}

type OrderListSearchResponsePriceInfos struct {
	Amount      string `json:"amount,omitempty"`
	Description string `json:"description,omitempty"`
	Type        int64  `json:"type,omitempty"`
}

type OrderListSearchResponseProductInfos struct {
	ProductExt      OrderListSearchResponseProductExt       `json:"productExt,omitempty"`
	WarehouseInfos  []OrderListSearchResponseWarehouseInfos `json:"warehouseInfos,omitempty"`
	CombSkuId       int64                                   `json:"combSkuId,omitempty"`
	CombTitle       string                                  `json:"combTitle,omitempty"`
	ItemSkuQuantity string                                  `json:"itemSkuQuantity,omitempty"`
	Price           string                                  `json:"price,omitempty"`
	ProductType     int64                                   `json:"productType,omitempty"`
	Title           string                                  `json:"title,omitempty"`
}

type OrderListSearchResponseProductExt struct {
	ImageUrl         string `json:"imageUrl,omitempty"`
	ProductAttribute string `json:"productAttribute,omitempty"`
	ProductCode      string `json:"productCode,omitempty"`
	Unit             string `json:"unit,omitempty"`
}

type OrderListSearchResponseWarehouseInfos struct {
	Quantity      string `json:"quantity,omitempty"`
	WarehouseId   int64  `json:"warehouseId,omitempty"`
	WarehouseName string `json:"warehouseName,omitempty"`
	WarehouseType string `json:"warehouseType,omitempty"`
}

type OrderListSearchResponseMerchantInfo struct {
	MerchantExtInfo   OrderListSearchResponseMerchantExtInfo `json:"merchantExtInfo,omitempty"`
	BosId             int64                                  `json:"bosId,omitempty"`
	BosName           string                                 `json:"bosName,omitempty"`
	MerchantId        int64                                  `json:"merchantId,omitempty"`
	ProcessVid        int64                                  `json:"processVid,omitempty"`
	ProcessVidName    string                                 `json:"processVidName,omitempty"`
	ProcessVidType    int64                                  `json:"processVidType,omitempty"`
	ProductId         int64                                  `json:"productId,omitempty"`
	ProductInstanceId int64                                  `json:"productInstanceId,omitempty"`
	ProductName       string                                 `json:"productName,omitempty"`
	Vid               int64                                  `json:"vid,omitempty"`
	VidName           string                                 `json:"vidName,omitempty"`
	VidType           int64                                  `json:"vidType,omitempty"`
}

type OrderListSearchResponseMerchantExtInfo struct {
	DeliveryVid      int64  `json:"deliveryVid,omitempty"`
	ProcessVidNumber string `json:"processVidNumber,omitempty"`
	VidNumber        string `json:"vidNumber,omitempty"`
}

type OrderListSearchResponseOrderBaseInfo struct {
	AutoCancelTime     int64  `json:"autoCancelTime,omitempty"`
	BizSourceType      int64  `json:"bizSourceType,omitempty"`
	ChannelType        int64  `json:"channelType,omitempty"`
	ConfirmTime        int64  `json:"confirmTime,omitempty"`
	CreateTime         int64  `json:"createTime,omitempty"`
	DeliveryType       int64  `json:"deliveryType,omitempty"`
	FinishDeliveryTime int64  `json:"finishDeliveryTime,omitempty"`
	FinishTime         int64  `json:"finishTime,omitempty"`
	IsDeleted          int64  `json:"isDeleted,omitempty"`
	OrderNo            int64  `json:"orderNo,omitempty"`
	OrderSource        int64  `json:"orderSource,omitempty"`
	OrderStatus        int64  `json:"orderStatus,omitempty"`
	OrderType          int64  `json:"orderType,omitempty"`
	ParentOrderNo      int64  `json:"parentOrderNo,omitempty"`
	PayStatus          int64  `json:"payStatus,omitempty"`
	PayTime            int64  `json:"payTime,omitempty"`
	PayType            int64  `json:"payType,omitempty"`
	ThirdOrderNo       string `json:"thirdOrderNo,omitempty"`
	UpdateTime         int64  `json:"updateTime,omitempty"`
}

type OrderListSearchResponseOrderBizExt struct {
	OperatorInfo    OrderListSearchResponseOperatorInfo  `json:"operatorInfo,omitempty"`
	LabelInfos      []OrderListSearchResponseLabelInfos2 `json:"labelInfos,omitempty"`
	FeatureType     int64                                `json:"featureType,omitempty"`
	SaleChannelType int64                                `json:"saleChannelType,omitempty"`
}

type OrderListSearchResponseOperatorInfo struct {
	OperatorId    string `json:"operatorId,omitempty"`
	OperatorName  string `json:"operatorName,omitempty"`
	OperatorPhone string `json:"operatorPhone,omitempty"`
}

type OrderListSearchResponseLabelInfos2 struct {
	AttachId   string `json:"attachId,omitempty"`
	Attachment string `json:"attachment,omitempty"`
	LabelType  string `json:"labelType,omitempty"`
}

type OrderListSearchResponseOrderFulfill struct {
	ReceiverInfo    OrderListSearchResponseReceiverInfo `json:"receiverInfo,omitempty"`
	DeliveryCoName  string                              `json:"deliveryCoName,omitempty"`
	DeliveryCode    string                              `json:"deliveryCode,omitempty"`
	DeliveryTime    int64                               `json:"deliveryTime,omitempty"`
	DeliveryType    int64                               `json:"deliveryType,omitempty"`
	ExpRcvDate      int64                               `json:"expRcvDate,omitempty"`
	ExpRcvEndTime   int64                               `json:"expRcvEndTime,omitempty"`
	ExpRcvStartTime int64                               `json:"expRcvStartTime,omitempty"`
	ExpRcvType      int64                               `json:"expRcvType,omitempty"`
	FulfillStatus   int64                               `json:"fulfillStatus,omitempty"`
}

type OrderListSearchResponseReceiverInfo struct {
	CertificateInfo OrderListSearchResponseCertificateInfo `json:"certificateInfo,omitempty"`
	Address         string                                 `json:"address,omitempty"`
	Area            string                                 `json:"area,omitempty"`
	AreaCode        string                                 `json:"areaCode,omitempty"`
	City            string                                 `json:"city,omitempty"`
	CityCode        string                                 `json:"cityCode,omitempty"`
	County          string                                 `json:"county,omitempty"`
	CountyCode      string                                 `json:"countyCode,omitempty"`
	Latitude        string                                 `json:"latitude,omitempty"`
	Longitude       string                                 `json:"longitude,omitempty"`
	Name            string                                 `json:"name,omitempty"`
	Phone           string                                 `json:"phone,omitempty"`
	Province        string                                 `json:"province,omitempty"`
	ProvinceCode    string                                 `json:"provinceCode,omitempty"`
	Zip             string                                 `json:"zip,omitempty"`
}

type OrderListSearchResponseCertificateInfo struct {
	BehindImg       string `json:"behindImg,omitempty"`
	CertificateNo   string `json:"certificateNo,omitempty"`
	CertificateType string `json:"certificateType,omitempty"`
	FrontImg        string `json:"frontImg,omitempty"`
	UserName        string `json:"userName,omitempty"`
}

type OrderListSearchResponsePayInfo2 struct {
	AmountInfos         []OrderListSearchResponseAmountInfos2 `json:"amountInfos,omitempty"`
	PayItems            []OrderListSearchResponsePayItems     `json:"payItems,omitempty"`
	PayAmount           string                                `json:"payAmount,omitempty"`
	ShouldPayAmount     string                                `json:"shouldPayAmount,omitempty"`
	TotalAmount         string                                `json:"totalAmount,omitempty"`
	TotalDiscountAmount string                                `json:"totalDiscountAmount,omitempty"`
}

type OrderListSearchResponseAmountInfos2 struct {
	Amount          string `json:"amount,omitempty"`
	Description     string `json:"description,omitempty"`
	PayAmount       string `json:"payAmount,omitempty"`
	ShouldPayAmount string `json:"shouldPayAmount,omitempty"`
	Type            int64  `json:"type,omitempty"`
}

type OrderListSearchResponsePayItems struct {
	PayItemExtInfo OrderListSearchResponsePayItemExtInfo `json:"payItemExtInfo,omitempty"`
	PayTime        int64                                 `json:"payTime,omitempty"`
	PayTradeId     int64                                 `json:"payTradeId,omitempty"`
	Phase          int64                                 `json:"phase,omitempty"`
	TradeId        string                                `json:"tradeId,omitempty"`
	PayId          int64                                 `json:"payId,omitempty"`
	PayMethodIds   []int64                               `json:"payMethodIds,omitempty"`
	ChannelTrxNo   string                                `json:"channelTrxNo,omitempty"`
	PayType        int64                                 `json:"payType,omitempty"`
}

type OrderListSearchResponsePayItemExtInfo struct {
	Amount     string `json:"amount,omitempty"`
	InteractId string `json:"interactId,omitempty"`
}

type OrderListSearchResponseTotalDiscounts struct {
	AttributionType int64  `json:"attributionType,omitempty"`
	DiscountAmount  string `json:"discountAmount,omitempty"`
	DiscountType    int64  `json:"discountType,omitempty"`
}

type OrderListSearchResponseGuideInfo struct {
	BuyerExpandInfo   OrderListSearchResponseBuyerExpandInfo `json:"buyerExpandInfo,omitempty"`
	GuiderWid         int64                                  `json:"guiderWid,omitempty"`
	GuiderName        string                                 `json:"guiderName,omitempty"`
	PrivateGuiderWid  int64                                  `json:"privateGuiderWid,omitempty"`
	PrivateGuiderName string                                 `json:"privateGuiderName,omitempty"`
	PersonalGuiderNo  string                                 `json:"personalGuiderNo,omitempty"`
	GuiderNo          string                                 `json:"guiderNo,omitempty"`
}

type OrderListSearchResponseBuyerExpandInfo struct {
	AttributionStoreId    int64  `json:"attributionStoreId,omitempty"`
	AttributionStoreName  string `json:"attributionStoreName,omitempty"`
	PersonalGuiderStoreId int64  `json:"personalGuiderStoreId,omitempty"`
	PersonalGuiderWid     int64  `json:"personalGuiderWid,omitempty"`
	PersonalGuiderName    string `json:"personalGuiderName,omitempty"`
}

type OrderListSearchResponseBookOrderInfo struct {
	BookOrder               bool  `json:"bookOrder,omitempty"`
	ExpectReceivedStartTime int64 `json:"expectReceivedStartTime,omitempty"`
	ExpectReceivedEndTime   int64 `json:"expectReceivedEndTime,omitempty"`
}

type OrderListSearchResponseRightsInfos struct {
	OrderNo          int64  `json:"orderNo,omitempty"`
	RightsId         int64  `json:"rightsId,omitempty"`
	RightsItemId     int64  `json:"rightsItemId,omitempty"`
	RightsStatus     int64  `json:"rightsStatus,omitempty"`
	RightsStatusName string `json:"rightsStatusName,omitempty"`
}

type OrderListSearchResponseFulfillOrderList struct {
	BuyerInfo         OrderListSearchResponseBuyerInfo2        `json:"buyerInfo,omitempty"`
	ConsignOrder      OrderListSearchResponseConsignOrder      `json:"consignOrder,omitempty"`
	DeliveryInfo      OrderListSearchResponseDeliveryInfo      `json:"deliveryInfo,omitempty"`
	FulfillItemList   []OrderListSearchResponseFulfillItemList `json:"fulfillItemList,omitempty"`
	ReceiveInfo       OrderListSearchResponseReceiveInfo       `json:"receiveInfo,omitempty"`
	SendInfo          OrderListSearchResponseSendInfo          `json:"sendInfo,omitempty"`
	AutoDeliveryTime  int64                                    `json:"autoDeliveryTime,omitempty"`
	AutoReceivingTime int64                                    `json:"autoReceivingTime,omitempty"`
	CancelTime        int64                                    `json:"cancelTime,omitempty"`
	DeliveryTime      int64                                    `json:"deliveryTime,omitempty"`
	DeliveryVid       int64                                    `json:"deliveryVid,omitempty"`
	Exception         string                                   `json:"exception,omitempty"`
	ExpectFulfillTime int64                                    `json:"expectFulfillTime,omitempty"`
	FulfillMethod     int64                                    `json:"fulfillMethod,omitempty"`
	FulfillNo         int64                                    `json:"fulfillNo,omitempty"`
	FulfillStatus     int64                                    `json:"fulfillStatus,omitempty"`
	FulfillType       int64                                    `json:"fulfillType,omitempty"`
	IsSplitPackage    int64                                    `json:"isSplitPackage,omitempty"`
	OrderNo           int64                                    `json:"orderNo,omitempty"`
	ReceivingTime     int64                                    `json:"receivingTime,omitempty"`
	Remark            string                                   `json:"remark,omitempty"`
	UpdateTime        int64                                    `json:"updateTime,omitempty"`
}

type OrderListSearchResponseBuyerInfo2 struct {
	BuyerRemark string `json:"buyerRemark,omitempty"`
	Wid         int64  `json:"wid,omitempty"`
}

type OrderListSearchResponseConsignOrder struct {
	ConfirmInfo    OrderListSearchResponseConfirmInfo    `json:"confirmInfo,omitempty"`
	ConsignItems   []OrderListSearchResponseConsignItems `json:"consignItems,omitempty"`
	LogisticsInfo  OrderListSearchResponseLogisticsInfo  `json:"logisticsInfo,omitempty"`
	CancelTime     int64                                 `json:"cancelTime,omitempty"`
	CancelType     int64                                 `json:"cancelType,omitempty"`
	ConfirmEndTime int64                                 `json:"confirmEndTime,omitempty"`
	ConfirmTime    int64                                 `json:"confirmTime,omitempty"`
	ConfirmType    int64                                 `json:"confirmType,omitempty"`
	ConsignTime    int64                                 `json:"consignTime,omitempty"`
	FulfillNo      int64                                 `json:"fulfillNo,omitempty"`
	ConsignVidName string                                `json:"consignVidName,omitempty"`
	PickupCode     string                                `json:"pickupCode,omitempty"`
}

type OrderListSearchResponseConfirmInfo struct {
	Name    string `json:"name,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Vid     int64  `json:"vid,omitempty"`
	VidName string `json:"vidName,omitempty"`
	VidType int64  `json:"vidType,omitempty"`
	Wid     int64  `json:"wid,omitempty"`
}

type OrderListSearchResponseConsignItems struct {
	ConsignExts []OrderListSearchResponseConsignExts `json:"consignExts,omitempty"`
}

type OrderListSearchResponseConsignExts struct {
	TemplateId int64  `json:"templateId,omitempty"`
	Name       string `json:"name,omitempty"`
	Code       string `json:"code,omitempty"`
	Type       int64  `json:"type,omitempty"`
}

type OrderListSearchResponseLogisticsInfo struct {
	Appointment int64  `json:"appointment,omitempty"`
	CompanyCode string `json:"companyCode,omitempty"`
	CompanyName string `json:"companyName,omitempty"`
	DeliveryNo  string `json:"deliveryNo,omitempty"`
	OutOrderNo  string `json:"outOrderNo,omitempty"`
	Remark      string `json:"remark,omitempty"`
}

type OrderListSearchResponseDeliveryInfo struct {
	CompanyCode             string `json:"companyCode,omitempty"`
	CompanyName             string `json:"companyName,omitempty"`
	ExpectReceivedDate      int64  `json:"expectReceivedDate,omitempty"`
	ExpectReceivedEndTime   int64  `json:"expectReceivedEndTime,omitempty"`
	ExpectReceivedStartTime int64  `json:"expectReceivedStartTime,omitempty"`
	ExpectReceivedType      int64  `json:"expectReceivedType,omitempty"`
	ExpectReceivedTypeName  string `json:"expectReceivedTypeName,omitempty"`
	Number                  string `json:"number,omitempty"`
	Status                  int64  `json:"status,omitempty"`
	StatusName              string `json:"statusName,omitempty"`
	WriteOffId              int64  `json:"writeOffId,omitempty"`
	WriteOffName            string `json:"writeOffName,omitempty"`
}

type OrderListSearchResponseFulfillItemList struct {
	Product     OrderListSearchResponseProduct `json:"product,omitempty"`
	DeliveryNum string                         `json:"deliveryNum,omitempty"`
	GoodsId     int64                          `json:"goodsId,omitempty"`
	OrderItemId int64                          `json:"orderItemId,omitempty"`
}

type OrderListSearchResponseProduct struct {
	ItemProducts []OrderListSearchResponseItemProducts `json:"itemProducts,omitempty"`
}

type OrderListSearchResponseItemProducts struct {
	Warehouses      []OrderListSearchResponseWarehouses `json:"warehouses,omitempty"`
	CombSkuId       int64                               `json:"combSkuId,omitempty"`
	ItemSkuId       int64                               `json:"itemSkuId,omitempty"`
	ItemSkuQuantity string                              `json:"itemSkuQuantity,omitempty"`
	ProductType     int64                               `json:"productType,omitempty"`
}

type OrderListSearchResponseWarehouses struct {
	Quantity      string `json:"quantity,omitempty"`
	WarehouseId   int64  `json:"warehouseId,omitempty"`
	WarehouseName string `json:"warehouseName,omitempty"`
	WarehouseType int64  `json:"warehouseType,omitempty"`
}

type OrderListSearchResponseReceiveInfo struct {
	AddressInfo     OrderListSearchResponseAddressInfo `json:"addressInfo,omitempty"`
	Receiver        OrderListSearchResponseReceiver    `json:"receiver,omitempty"`
	PickUpName      string                             `json:"pickUpName,omitempty"`
	PickUpVid       int64                              `json:"pickUpVid,omitempty"`
	ReceiverAddress string                             `json:"receiverAddress,omitempty"`
}

type OrderListSearchResponseAddressInfo struct {
	AddressExt OrderListSearchResponseAddressExt `json:"addressExt,omitempty"`
	Address    string                            `json:"address,omitempty"`
	Area       string                            `json:"area,omitempty"`
	City       string                            `json:"city,omitempty"`
	County     string                            `json:"county,omitempty"`
	Latitude   string                            `json:"latitude,omitempty"`
	Longitude  string                            `json:"longitude,omitempty"`
	Province   string                            `json:"province,omitempty"`
	Zip        string                            `json:"zip,omitempty"`
}

type OrderListSearchResponseAddressExt struct {
	AreaCode     string `json:"areaCode,omitempty"`
	CityCode     string `json:"cityCode,omitempty"`
	CountyCode   string `json:"countyCode,omitempty"`
	ProvinceCode string `json:"provinceCode,omitempty"`
}

type OrderListSearchResponseReceiver struct {
	IdCardExt OrderListSearchResponseIdCardExt `json:"idCardExt,omitempty"`
	IdCardNo  string                           `json:"idCardNo,omitempty"`
}

type OrderListSearchResponseIdCardExt struct {
	BehindImg  string `json:"behindImg,omitempty"`
	FrontImg   string `json:"frontImg,omitempty"`
	IsVerified bool   `json:"isVerified,omitempty"`
	UserName   string `json:"userName,omitempty"`
}

type OrderListSearchResponseSendInfo struct {
	AddressInfo   OrderListSearchResponseAddressInfo2 `json:"addressInfo,omitempty"`
	Sender        OrderListSearchResponseSender       `json:"sender,omitempty"`
	SenderAddress string                              `json:"senderAddress,omitempty"`
}

type OrderListSearchResponseAddressInfo2 struct {
	AddressExt OrderListSearchResponseAddressExt2 `json:"addressExt,omitempty"`
	Address    string                             `json:"address,omitempty"`
	Area       string                             `json:"area,omitempty"`
	City       string                             `json:"city,omitempty"`
	County     string                             `json:"county,omitempty"`
	Latitude   string                             `json:"latitude,omitempty"`
	Longitude  string                             `json:"longitude,omitempty"`
	Province   string                             `json:"province,omitempty"`
	Zip        string                             `json:"zip,omitempty"`
}

type OrderListSearchResponseAddressExt2 struct {
	AreaCode     string `json:"areaCode,omitempty"`
	CityCode     string `json:"cityCode,omitempty"`
	CountyCode   string `json:"countyCode,omitempty"`
	ProvinceCode string `json:"provinceCode,omitempty"`
}

type OrderListSearchResponseSender struct {
	SenderMobile string `json:"senderMobile,omitempty"`
	SenderName   string `json:"senderName,omitempty"`
}
