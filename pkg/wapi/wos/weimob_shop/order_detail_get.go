package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderDetailGet
// @id 1620
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1620?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询订单详情)
func (client *Client) OrderDetailGet(request *OrderDetailGetRequest) (response *OrderDetailGetResponse, err error) {
	rpcResponse := CreateOrderDetailGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderDetailGetRequest struct {
	*api.BaseRequest

	OrderDomains []int64 `json:"orderDomains,omitempty"`
	OrderNo      int64   `json:"orderNo,omitempty"`
}

type OrderDetailGetResponse struct {
	OrderInfo       OrderDetailGetResponseOrderInfo         `json:"orderInfo,omitempty"`
	RightsInfos     []OrderDetailGetResponseRightsInfos     `json:"rightsInfos,omitempty"`
	FulfillInfoList []OrderDetailGetResponseFulfillInfoList `json:"fulfillInfoList,omitempty"`
}

func CreateOrderDetailGetRequest() (request *OrderDetailGetRequest) {
	request = &OrderDetailGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "order/detail/get", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderDetailGetResponse() (response *api.BaseResponse[OrderDetailGetResponse]) {
	response = api.CreateResponse[OrderDetailGetResponse](&OrderDetailGetResponse{})
	return
}

type OrderDetailGetResponseOrderInfo struct {
	BuyerInfo      OrderDetailGetResponseBuyerInfo        `json:"buyerInfo,omitempty"`
	CancelInfo     OrderDetailGetResponseCancelInfo       `json:"cancelInfo,omitempty"`
	DiscountInfos  []OrderDetailGetResponseDiscountInfos  `json:"discountInfos,omitempty"`
	FlagInfo       OrderDetailGetResponseFlagInfo         `json:"flagInfo,omitempty"`
	Items          []OrderDetailGetResponseItems          `json:"items,omitempty"`
	MerchantInfo   OrderDetailGetResponseMerchantInfo     `json:"merchantInfo,omitempty"`
	OrderBaseInfo  OrderDetailGetResponseOrderBaseInfo    `json:"orderBaseInfo,omitempty"`
	OrderBizExt    OrderDetailGetResponseOrderBizExt      `json:"orderBizExt,omitempty"`
	OrderFulfill   OrderDetailGetResponseOrderFulfill     `json:"orderFulfill,omitempty"`
	PayInfo        OrderDetailGetResponsePayInfo2         `json:"payInfo,omitempty"`
	TotalDiscounts []OrderDetailGetResponseTotalDiscounts `json:"totalDiscounts,omitempty"`
	GuideInfo      OrderDetailGetResponseGuideInfo        `json:"guideInfo,omitempty"`
	InvoiceInfo    OrderDetailGetResponseInvoiceInfo      `json:"invoiceInfo,omitempty"`
	PreSellInfo    OrderDetailGetResponsePreSellInfo      `json:"preSellInfo,omitempty"`
}

type OrderDetailGetResponseBuyerInfo struct {
	MemberBenefits []OrderDetailGetResponseMemberBenefits `json:"memberBenefits,omitempty"`
	UserNickName   string                                 `json:"userNickName,omitempty"`
	Wid            int64                                  `json:"wid,omitempty"`
	BuyerRemark    string                                 `json:"buyerRemark,omitempty"`
}

type OrderDetailGetResponseMemberBenefits struct {
	BenefitType int64 `json:"benefitType,omitempty"`
}

type OrderDetailGetResponseCancelInfo struct {
	CancelType    int64  `json:"cancelType,omitempty"`
	Id            int64  `json:"id,omitempty"`
	Reason        string `json:"reason,omitempty"`
	SpecialReason string `json:"specialReason,omitempty"`
}

type OrderDetailGetResponseDiscountInfos struct {
	DiscountExt    OrderDetailGetResponseDiscountExt `json:"discountExt,omitempty"`
	CostAmount     string                            `json:"costAmount,omitempty"`
	DiscountAmount string                            `json:"discountAmount,omitempty"`
	DiscountId     string                            `json:"discountId,omitempty"`
	DiscountLevel  int64                             `json:"discountLevel,omitempty"`
	DiscountType   int64                             `json:"discountType,omitempty"`
	Name           string                            `json:"name,omitempty"`
	SubType        int64                             `json:"subType,omitempty"`
}

type OrderDetailGetResponseDiscountExt struct {
	AttributionType int64 `json:"attributionType,omitempty"`
}

type OrderDetailGetResponseFlagInfo struct {
	FlagContent string `json:"flagContent,omitempty"`
	FlagRank    int64  `json:"flagRank,omitempty"`
}

type OrderDetailGetResponseItems struct {
	DiscountInfos    []OrderDetailGetResponseDiscountInfos2 `json:"discountInfos,omitempty"`
	GoodsExt         OrderDetailGetResponseGoodsExt         `json:"goodsExt,omitempty"`
	ItemBizExt       OrderDetailGetResponseItemBizExt       `json:"itemBizExt,omitempty"`
	PayInfo          OrderDetailGetResponsePayInfo          `json:"payInfo,omitempty"`
	PriceInfos       []OrderDetailGetResponsePriceInfos     `json:"priceInfos,omitempty"`
	ProductInfos     []OrderDetailGetResponseProductInfos   `json:"productInfos,omitempty"`
	ActivityTypeList []int64                                `json:"activityTypeList,omitempty"`
	CategoryId       int64                                  `json:"categoryId,omitempty"`
	CategoryTitle    []int64                                `json:"categoryTitle,omitempty"`
	GoodsCode        string                                 `json:"goodsCode,omitempty"`
	GoodsId          int64                                  `json:"goodsId,omitempty"`
	GoodsTitle       string                                 `json:"goodsTitle,omitempty"`
	GoodsType        int64                                  `json:"goodsType,omitempty"`
	ImageUrl         string                                 `json:"imageUrl,omitempty"`
	ItemId           int64                                  `json:"itemId,omitempty"`
	SalePrice        string                                 `json:"salePrice,omitempty"`
	SkuAttrInfo      string                                 `json:"skuAttrInfo,omitempty"`
	SkuBarCode       string                                 `json:"skuBarCode,omitempty"`
	SkuCode          string                                 `json:"skuCode,omitempty"`
	SkuId            int64                                  `json:"skuId,omitempty"`
	SkuNum           string                                 `json:"skuNum,omitempty"`
	SubGoodsType     int64                                  `json:"subGoodsType,omitempty"`
	UnitType         int64                                  `json:"unitType,omitempty"`
	UnitName         string                                 `json:"unitName,omitempty"`
}

type OrderDetailGetResponseDiscountInfos2 struct {
	DiscountExt    OrderDetailGetResponseDiscountExt2 `json:"discountExt,omitempty"`
	CostAmount     string                             `json:"costAmount,omitempty"`
	DiscountAmount string                             `json:"discountAmount,omitempty"`
	DiscountId     string                             `json:"discountId,omitempty"`
	DiscountLevel  int64                              `json:"discountLevel,omitempty"`
	DiscountType   int64                              `json:"discountType,omitempty"`
	Name           string                             `json:"name,omitempty"`
	SubType        int64                              `json:"subType,omitempty"`
}

type OrderDetailGetResponseDiscountExt2 struct {
	AttributionType int64 `json:"attributionType,omitempty"`
}

type OrderDetailGetResponseGoodsExt struct {
	GroupInfos      []OrderDetailGetResponseGroupInfos `json:"groupInfos,omitempty"`
	OriginSkuName   string                             `json:"originSkuName,omitempty"`
	OriginSkuNum    int64                              `json:"originSkuNum,omitempty"`
	ProductCategory string                             `json:"productCategory,omitempty"`
}

type OrderDetailGetResponseGroupInfos struct {
	Grade   int64 `json:"grade,omitempty"`
	GroupId int64 `json:"groupId,omitempty"`
}

type OrderDetailGetResponseItemBizExt struct {
	BizInfos              []OrderDetailGetResponseBizInfos          `json:"bizInfos,omitempty"`
	GoodsCustom           OrderDetailGetResponseGoodsCustom         `json:"goodsCustom,omitempty"`
	LabelInfos            []OrderDetailGetResponseLabelInfos        `json:"labelInfos,omitempty"`
	AbilityCode           []OrderDetailGetResponseAbilityCode       `json:"abilityCode,omitempty"`
	GoodsMessageInfos     []OrderDetailGetResponseGoodsMessageInfos `json:"goodsMessageInfos,omitempty"`
	ActivityStockType     int64                                     `json:"activityStockType,omitempty"`
	ExpandField           string                                    `json:"expandField,omitempty"`
	GoodsGuideType        int64                                     `json:"goodsGuideType,omitempty"`
	GoodsLimitSwitch      int64                                     `json:"goodsLimitSwitch,omitempty"`
	GoodsPromotionOrderId string                                    `json:"goodsPromotionOrderId,omitempty"`
	GoodsPromotionType    int64                                     `json:"goodsPromotionType,omitempty"`
	GoodsSellMode         int64                                     `json:"goodsSellMode,omitempty"`
	GoodsSourceType       int64                                     `json:"goodsSourceType,omitempty"`
	OuterGoodsId          string                                    `json:"outerGoodsId,omitempty"`
	OuterMerchantId       string                                    `json:"outerMerchantId,omitempty"`
	OuterSkuId            string                                    `json:"outerSkuId,omitempty"`
	RightsServiceType     int64                                     `json:"rightsServiceType,omitempty"`
}

type OrderDetailGetResponseBizInfos struct {
	BizId      int64  `json:"bizId,omitempty"`
	BizOrderId string `json:"bizOrderId,omitempty"`
	BizType    int64  `json:"bizType,omitempty"`
	SubBizType int64  `json:"subBizType,omitempty"`
}

type OrderDetailGetResponseGoodsCustom struct {
	CustomFields []OrderDetailGetResponseCustomFields `json:"customFields,omitempty"`
	CloudCustom  string                               `json:"cloudCustom,omitempty"`
}

type OrderDetailGetResponseCustomFields struct {
	Key   string `json:"key,omitempty"`
	Name  string `json:"name,omitempty"`
	Sort  int64  `json:"sort,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type OrderDetailGetResponseLabelInfos struct {
	AttachId   string `json:"attachId,omitempty"`
	Attachment string `json:"attachment,omitempty"`
	LabelType  string `json:"labelType,omitempty"`
}

type OrderDetailGetResponseAbilityCode struct {
}

type OrderDetailGetResponseGoodsMessageInfos struct {
}

type OrderDetailGetResponsePayInfo struct {
	AmountInfos         []OrderDetailGetResponseAmountInfos `json:"amountInfos,omitempty"`
	PayAmount           string                              `json:"payAmount,omitempty"`
	ShouldPayAmount     string                              `json:"shouldPayAmount,omitempty"`
	TotalAmount         string                              `json:"totalAmount,omitempty"`
	TotalDiscountAmount string                              `json:"totalDiscountAmount,omitempty"`
}

type OrderDetailGetResponseAmountInfos struct {
	Amount          string `json:"amount,omitempty"`
	Description     string `json:"description,omitempty"`
	PayAmount       string `json:"payAmount,omitempty"`
	ShouldPayAmount string `json:"shouldPayAmount,omitempty"`
	Type            int64  `json:"type,omitempty"`
}

type OrderDetailGetResponsePriceInfos struct {
	Amount      string `json:"amount,omitempty"`
	Description string `json:"description,omitempty"`
	Type        int64  `json:"type,omitempty"`
}

type OrderDetailGetResponseProductInfos struct {
	ProductExt      OrderDetailGetResponseProductExt       `json:"productExt,omitempty"`
	WarehouseInfos  []OrderDetailGetResponseWarehouseInfos `json:"warehouseInfos,omitempty"`
	CombSkuId       int64                                  `json:"combSkuId,omitempty"`
	CombTitle       string                                 `json:"combTitle,omitempty"`
	ItemSkuQuantity string                                 `json:"itemSkuQuantity,omitempty"`
	Price           string                                 `json:"price,omitempty"`
	ProductType     int64                                  `json:"productType,omitempty"`
	Title           string                                 `json:"title,omitempty"`
}

type OrderDetailGetResponseProductExt struct {
	ImageUrl         string `json:"imageUrl,omitempty"`
	ProductAttribute string `json:"productAttribute,omitempty"`
	ProductCode      string `json:"productCode,omitempty"`
	Unit             string `json:"unit,omitempty"`
}

type OrderDetailGetResponseWarehouseInfos struct {
	Quantity      int64  `json:"quantity,omitempty"`
	WarehouseId   int64  `json:"warehouseId,omitempty"`
	WarehouseName string `json:"warehouseName,omitempty"`
	WarehouseType string `json:"warehouseType,omitempty"`
}

type OrderDetailGetResponseMerchantInfo struct {
	MerchantExtInfo   OrderDetailGetResponseMerchantExtInfo `json:"merchantExtInfo,omitempty"`
	BosId             int64                                 `json:"bosId,omitempty"`
	BosName           string                                `json:"bosName,omitempty"`
	MerchantId        int64                                 `json:"merchantId,omitempty"`
	ProcessVid        int64                                 `json:"processVid,omitempty"`
	ProcessVidName    string                                `json:"processVidName,omitempty"`
	ProcessVidType    int64                                 `json:"processVidType,omitempty"`
	ProductId         int64                                 `json:"productId,omitempty"`
	ProductInstanceId int64                                 `json:"productInstanceId,omitempty"`
	ProductName       string                                `json:"productName,omitempty"`
	Vid               int64                                 `json:"vid,omitempty"`
	VidName           string                                `json:"vidName,omitempty"`
	VidType           int64                                 `json:"vidType,omitempty"`
}

type OrderDetailGetResponseMerchantExtInfo struct {
	DeliveryVid      int64  `json:"deliveryVid,omitempty"`
	ProcessVidNumber string `json:"processVidNumber,omitempty"`
	VidNumber        string `json:"vidNumber,omitempty"`
}

type OrderDetailGetResponseOrderBaseInfo struct {
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

type OrderDetailGetResponseOrderBizExt struct {
	OperatorInfo    OrderDetailGetResponseOperatorInfo  `json:"operatorInfo,omitempty"`
	FinishInfo      OrderDetailGetResponseFinishInfo    `json:"finishInfo,omitempty"`
	LabelInfos      []OrderDetailGetResponseLabelInfos2 `json:"labelInfos,omitempty"`
	FeatureType     int64                               `json:"featureType,omitempty"`
	SaleChannelType int64                               `json:"saleChannelType,omitempty"`
}

type OrderDetailGetResponseOperatorInfo struct {
	OperatorId    string `json:"operatorId,omitempty"`
	OperatorName  string `json:"operatorName,omitempty"`
	OperatorPhone string `json:"operatorPhone,omitempty"`
}

type OrderDetailGetResponseFinishInfo struct {
	FinishOrderType int64 `json:"finishOrderType,omitempty"`
}

type OrderDetailGetResponseLabelInfos2 struct {
	AttachId   string `json:"attachId,omitempty"`
	Attachment string `json:"attachment,omitempty"`
	LabelType  string `json:"labelType,omitempty"`
}

type OrderDetailGetResponseOrderFulfill struct {
	ReceiverInfo    OrderDetailGetResponseReceiverInfo      `json:"receiverInfo,omitempty"`
	CustomFieldInfo []OrderDetailGetResponseCustomFieldInfo `json:"customFieldInfo,omitempty"`
	DeliveryCoName  string                                  `json:"deliveryCoName,omitempty"`
	DeliveryCode    string                                  `json:"deliveryCode,omitempty"`
	DeliveryTime    int64                                   `json:"deliveryTime,omitempty"`
	DeliveryType    int64                                   `json:"deliveryType,omitempty"`
	ExpRcvDate      int64                                   `json:"expRcvDate,omitempty"`
	ExpRcvEndTime   int64                                   `json:"expRcvEndTime,omitempty"`
	ExpRcvStartTime int64                                   `json:"expRcvStartTime,omitempty"`
	ExpRcvType      int64                                   `json:"expRcvType,omitempty"`
	FulfillStatus   int64                                   `json:"fulfillStatus,omitempty"`
	DeliverySubCode string                                  `json:"deliverySubCode,omitempty"`
}

type OrderDetailGetResponseReceiverInfo struct {
	CertificateInfo OrderDetailGetResponseCertificateInfo `json:"certificateInfo,omitempty"`
	Address         string                                `json:"address,omitempty"`
	Area            string                                `json:"area,omitempty"`
	AreaCode        string                                `json:"areaCode,omitempty"`
	City            string                                `json:"city,omitempty"`
	CityCode        string                                `json:"cityCode,omitempty"`
	County          string                                `json:"county,omitempty"`
	CountyCode      string                                `json:"countyCode,omitempty"`
	Latitude        string                                `json:"latitude,omitempty"`
	Longitude       string                                `json:"longitude,omitempty"`
	Name            string                                `json:"name,omitempty"`
	Phone           string                                `json:"phone,omitempty"`
	Province        string                                `json:"province,omitempty"`
	ProvinceCode    string                                `json:"provinceCode,omitempty"`
	Zip             string                                `json:"zip,omitempty"`
}

type OrderDetailGetResponseCertificateInfo struct {
	BehindImg       string `json:"behindImg,omitempty"`
	CertificateNo   string `json:"certificateNo,omitempty"`
	CertificateType string `json:"certificateType,omitempty"`
	FrontImg        string `json:"frontImg,omitempty"`
	UserName        string `json:"userName,omitempty"`
}

type OrderDetailGetResponseCustomFieldInfo struct {
	ShowScene []OrderDetailGetResponseShowScene `json:"showScene,omitempty"`
	Key       string                            `json:"key,omitempty"`
	Name      string                            `json:"name,omitempty"`
	Value     string                            `json:"value,omitempty"`
	Sort      int64                             `json:"sort,omitempty"`
	Type      string                            `json:"type,omitempty"`
}

type OrderDetailGetResponseShowScene struct {
}

type OrderDetailGetResponsePayInfo2 struct {
	AmountInfos         []OrderDetailGetResponseAmountInfos2 `json:"amountInfos,omitempty"`
	PayItems            []OrderDetailGetResponsePayItems     `json:"payItems,omitempty"`
	PayAmount           string                               `json:"payAmount,omitempty"`
	ShouldPayAmount     string                               `json:"shouldPayAmount,omitempty"`
	TotalAmount         string                               `json:"totalAmount,omitempty"`
	TotalDiscountAmount string                               `json:"totalDiscountAmount,omitempty"`
}

type OrderDetailGetResponseAmountInfos2 struct {
	Amount          string `json:"amount,omitempty"`
	Description     string `json:"description,omitempty"`
	PayAmount       string `json:"payAmount,omitempty"`
	ShouldPayAmount string `json:"shouldPayAmount,omitempty"`
	Type            int64  `json:"type,omitempty"`
}

type OrderDetailGetResponsePayItems struct {
	PayItemExtInfo OrderDetailGetResponsePayItemExtInfo `json:"payItemExtInfo,omitempty"`
	PayTime        int64                                `json:"payTime,omitempty"`
	PayTradeId     int64                                `json:"payTradeId,omitempty"`
	Phase          int64                                `json:"phase,omitempty"`
	TradeId        string                               `json:"tradeId,omitempty"`
	PayId          int64                                `json:"payId,omitempty"`
	PayMethodIds   []int64                              `json:"payMethodIds,omitempty"`
	ChannelTrxNo   string                               `json:"channelTrxNo,omitempty"`
	PayType        int64                                `json:"payType,omitempty"`
}

type OrderDetailGetResponsePayItemExtInfo struct {
	Amount     string `json:"amount,omitempty"`
	InteractId string `json:"interactId,omitempty"`
	ChannelMid string `json:"channelMid,omitempty"`
}

type OrderDetailGetResponseTotalDiscounts struct {
	AttributionType int64  `json:"attributionType,omitempty"`
	DiscountAmount  string `json:"discountAmount,omitempty"`
	DiscountType    int64  `json:"discountType,omitempty"`
}

type OrderDetailGetResponseGuideInfo struct {
	BuyerExpandInfo   OrderDetailGetResponseBuyerExpandInfo `json:"buyerExpandInfo,omitempty"`
	GuiderWid         int64                                 `json:"guiderWid,omitempty"`
	GuiderName        string                                `json:"guiderName,omitempty"`
	PrivateGuiderWid  int64                                 `json:"privateGuiderWid,omitempty"`
	PrivateGuiderName string                                `json:"privateGuiderName,omitempty"`
	PersonalGuiderNo  string                                `json:"personalGuiderNo,omitempty"`
	GuiderNo          string                                `json:"guiderNo,omitempty"`
}

type OrderDetailGetResponseBuyerExpandInfo struct {
	AttributionStoreId    int64  `json:"attributionStoreId,omitempty"`
	AttributionStoreName  string `json:"attributionStoreName,omitempty"`
	PersonalGuiderStoreId int64  `json:"personalGuiderStoreId,omitempty"`
	PersonalGuiderWid     int64  `json:"personalGuiderWid,omitempty"`
	PersonalGuiderName    string `json:"personalGuiderName,omitempty"`
}

type OrderDetailGetResponseInvoiceInfo struct {
	TitleInfo     OrderDetailGetResponseTitleInfo     `json:"titleInfo,omitempty"`
	ReceiverInfo  OrderDetailGetResponseReceiverInfo2 `json:"receiverInfo,omitempty"`
	InvoiceStatus int64                               `json:"invoiceStatus,omitempty"`
	InvoiceType   int64                               `json:"invoiceType,omitempty"`
	TitleType     int64                               `json:"titleType,omitempty"`
}

type OrderDetailGetResponseTitleInfo struct {
	TitleName      string `json:"titleName,omitempty"`
	TaxNumber      string `json:"taxNumber,omitempty"`
	CompanyAddress string `json:"companyAddress,omitempty"`
	CompanyTel     string `json:"companyTel,omitempty"`
	BankAddress    string `json:"bankAddress,omitempty"`
	BankAccount    string `json:"bankAccount,omitempty"`
}

type OrderDetailGetResponseReceiverInfo2 struct {
	ReceiverTel   string `json:"receiverTel,omitempty"`
	ReceiverEmail string `json:"receiverEmail,omitempty"`
}

type OrderDetailGetResponsePreSellInfo struct {
	DepositPayType      int64 `json:"depositPayType,omitempty"`
	PreSellDeliveryDate int64 `json:"preSellDeliveryDate,omitempty"`
}

type OrderDetailGetResponseRightsInfos struct {
	OrderNo          int64  `json:"orderNo,omitempty"`
	RightsId         int64  `json:"rightsId,omitempty"`
	RightsItemId     int64  `json:"rightsItemId,omitempty"`
	RightsStatus     int64  `json:"rightsStatus,omitempty"`
	RightsStatusName string `json:"rightsStatusName,omitempty"`
}

type OrderDetailGetResponseFulfillInfoList struct {
	BuyerInfo         OrderDetailGetResponseBuyerInfo2        `json:"buyerInfo,omitempty"`
	ConsignOrder      OrderDetailGetResponseConsignOrder      `json:"consignOrder,omitempty"`
	DeliveryInfo      OrderDetailGetResponseDeliveryInfo      `json:"deliveryInfo,omitempty"`
	FulfillItemList   []OrderDetailGetResponseFulfillItemList `json:"fulfillItemList,omitempty"`
	ReceiveInfo       OrderDetailGetResponseReceiveInfo       `json:"receiveInfo,omitempty"`
	SendInfo          OrderDetailGetResponseSendInfo          `json:"sendInfo,omitempty"`
	AutoDeliveryTime  int64                                   `json:"autoDeliveryTime,omitempty"`
	AutoReceivingTime int64                                   `json:"autoReceivingTime,omitempty"`
	CancelTime        int64                                   `json:"cancelTime,omitempty"`
	DeliveryTime      int64                                   `json:"deliveryTime,omitempty"`
	DeliveryVid       int64                                   `json:"deliveryVid,omitempty"`
	Exception         string                                  `json:"exception,omitempty"`
	ExpectFulfillTime int64                                   `json:"expectFulfillTime,omitempty"`
	FulfillMethod     int64                                   `json:"fulfillMethod,omitempty"`
	FulfillNo         int64                                   `json:"fulfillNo,omitempty"`
	FulfillStatus     int64                                   `json:"fulfillStatus,omitempty"`
	FulfillType       int64                                   `json:"fulfillType,omitempty"`
	IsSplitPackage    int64                                   `json:"isSplitPackage,omitempty"`
	OrderNo           int64                                   `json:"orderNo,omitempty"`
	ReceivingTime     int64                                   `json:"receivingTime,omitempty"`
	Remark            string                                  `json:"remark,omitempty"`
	UpdateTime        string                                  `json:"updateTime,omitempty"`
}

type OrderDetailGetResponseBuyerInfo2 struct {
	BuyerRemark string `json:"buyerRemark,omitempty"`
	Wid         int64  `json:"wid,omitempty"`
}

type OrderDetailGetResponseConsignOrder struct {
	ConfirmInfo    OrderDetailGetResponseConfirmInfo    `json:"confirmInfo,omitempty"`
	ConsignItems   []OrderDetailGetResponseConsignItems `json:"consignItems,omitempty"`
	LogisticsInfo  OrderDetailGetResponseLogisticsInfo  `json:"logisticsInfo,omitempty"`
	CancelTime     int64                                `json:"cancelTime,omitempty"`
	CancelType     int64                                `json:"cancelType,omitempty"`
	ConfirmEndTime int64                                `json:"confirmEndTime,omitempty"`
	ConfirmTime    int64                                `json:"confirmTime,omitempty"`
	ConfirmType    int64                                `json:"confirmType,omitempty"`
	ConsignTime    int64                                `json:"consignTime,omitempty"`
	ConsignVidName string                               `json:"consignVidName,omitempty"`
	FulfillNo      int64                                `json:"fulfillNo,omitempty"`
	PickupCode     string                               `json:"pickupCode,omitempty"`
}

type OrderDetailGetResponseConfirmInfo struct {
	Name    string `json:"name,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Vid     int64  `json:"vid,omitempty"`
	VidName string `json:"vidName,omitempty"`
	VidType int64  `json:"vidType,omitempty"`
	Wid     int64  `json:"wid,omitempty"`
}

type OrderDetailGetResponseConsignItems struct {
	ConsignExts []OrderDetailGetResponseConsignExts `json:"consignExts,omitempty"`
}

type OrderDetailGetResponseConsignExts struct {
	TemplateId int64  `json:"templateId,omitempty"`
	Name       string `json:"name,omitempty"`
	Code       string `json:"code,omitempty"`
	Type       int64  `json:"type,omitempty"`
}

type OrderDetailGetResponseLogisticsInfo struct {
	Appointment int64  `json:"appointment,omitempty"`
	CompanyCode string `json:"companyCode,omitempty"`
	CompanyName string `json:"companyName,omitempty"`
	DeliveryNo  string `json:"deliveryNo,omitempty"`
	OutOrderNo  string `json:"outOrderNo,omitempty"`
	Remark      string `json:"remark,omitempty"`
}

type OrderDetailGetResponseDeliveryInfo struct {
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

type OrderDetailGetResponseFulfillItemList struct {
	Product     OrderDetailGetResponseProduct `json:"product,omitempty"`
	DeliveryNum string                        `json:"deliveryNum,omitempty"`
	GoodsId     int64                         `json:"goodsId,omitempty"`
	OrderItemId int64                         `json:"orderItemId,omitempty"`
}

type OrderDetailGetResponseProduct struct {
	ItemProducts []OrderDetailGetResponseItemProducts `json:"itemProducts,omitempty"`
}

type OrderDetailGetResponseItemProducts struct {
	Warehouses      []OrderDetailGetResponseWarehouses `json:"warehouses,omitempty"`
	CombSkuId       int64                              `json:"combSkuId,omitempty"`
	ItemSkuId       int64                              `json:"itemSkuId,omitempty"`
	ItemSkuQuantity string                             `json:"itemSkuQuantity,omitempty"`
	ProductType     int64                              `json:"productType,omitempty"`
}

type OrderDetailGetResponseWarehouses struct {
	Quantity      string `json:"quantity,omitempty"`
	WarehouseId   int64  `json:"warehouseId,omitempty"`
	WarehouseName string `json:"warehouseName,omitempty"`
	WarehouseType string `json:"warehouseType,omitempty"`
}

type OrderDetailGetResponseReceiveInfo struct {
	AddressInfo     OrderDetailGetResponseAddressInfo `json:"addressInfo,omitempty"`
	Receiver        OrderDetailGetResponseReceiver    `json:"receiver,omitempty"`
	PickUpName      string                            `json:"pickUpName,omitempty"`
	PickUpVid       int64                             `json:"pickUpVid,omitempty"`
	ReceiverAddress string                            `json:"receiverAddress,omitempty"`
}

type OrderDetailGetResponseAddressInfo struct {
	AddressExt OrderDetailGetResponseAddressExt `json:"addressExt,omitempty"`
	Address    string                           `json:"address,omitempty"`
	Area       string                           `json:"area,omitempty"`
	City       string                           `json:"city,omitempty"`
	County     string                           `json:"county,omitempty"`
	Latitude   string                           `json:"latitude,omitempty"`
	Longitude  string                           `json:"longitude,omitempty"`
	Province   string                           `json:"province,omitempty"`
	Zip        string                           `json:"zip,omitempty"`
}

type OrderDetailGetResponseAddressExt struct {
	AreaCode     string `json:"areaCode,omitempty"`
	CityCode     string `json:"cityCode,omitempty"`
	CountyCode   string `json:"countyCode,omitempty"`
	ProvinceCode string `json:"provinceCode,omitempty"`
}

type OrderDetailGetResponseReceiver struct {
	IdCardExt      OrderDetailGetResponseIdCardExt `json:"idCardExt,omitempty"`
	IdCardNo       string                          `json:"idCardNo,omitempty"`
	ReceiverMobile string                          `json:"receiverMobile,omitempty"`
	ReceiverName   string                          `json:"receiverName,omitempty"`
}

type OrderDetailGetResponseIdCardExt struct {
	BehindImg  string `json:"behindImg,omitempty"`
	FrontImg   string `json:"frontImg,omitempty"`
	IsVerified bool   `json:"isVerified,omitempty"`
	UserName   string `json:"userName,omitempty"`
}

type OrderDetailGetResponseSendInfo struct {
	AddressInfo   OrderDetailGetResponseAddressInfo2 `json:"addressInfo,omitempty"`
	Sender        OrderDetailGetResponseSender       `json:"sender,omitempty"`
	SenderAddress string                             `json:"senderAddress,omitempty"`
}

type OrderDetailGetResponseAddressInfo2 struct {
	AddressExt OrderDetailGetResponseAddressExt2 `json:"addressExt,omitempty"`
	Address    string                            `json:"address,omitempty"`
	Area       string                            `json:"area,omitempty"`
	City       string                            `json:"city,omitempty"`
	County     string                            `json:"county,omitempty"`
	Latitude   string                            `json:"latitude,omitempty"`
	Longitude  string                            `json:"longitude,omitempty"`
	Province   string                            `json:"province,omitempty"`
	Zip        string                            `json:"zip,omitempty"`
}

type OrderDetailGetResponseAddressExt2 struct {
	AreaCode     string `json:"areaCode,omitempty"`
	CityCode     string `json:"cityCode,omitempty"`
	CountyCode   string `json:"countyCode,omitempty"`
	ProvinceCode string `json:"provinceCode,omitempty"`
}

type OrderDetailGetResponseSender struct {
	SenderMobile string `json:"senderMobile,omitempty"`
	SenderName   string `json:"senderName,omitempty"`
}
