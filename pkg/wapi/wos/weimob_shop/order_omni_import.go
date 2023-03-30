package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderOmniImport
// @id 1977
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1977?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=外部订单导入)
func (client *Client) OrderOmniImport(request *OrderOmniImportRequest) (response *OrderOmniImportResponse, err error) {
	rpcResponse := CreateOrderOmniImportResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderOmniImportRequest struct {
	*api.BaseRequest

	OrderInfo OrderOmniImportRequestOrderInfo `json:"orderInfo,omitempty"`
}

type OrderOmniImportResponse struct {
	OutputInfo OrderOmniImportResponseOutputInfo `json:"outputInfo,omitempty"`
	Success    bool                              `json:"success,omitempty"`
}

func CreateOrderOmniImportRequest() (request *OrderOmniImportRequest) {
	request = &OrderOmniImportRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "order/omni/import", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderOmniImportResponse() (response *api.BaseResponse[OrderOmniImportResponse]) {
	response = api.CreateResponse[OrderOmniImportResponse](&OrderOmniImportResponse{})
	return
}

type OrderOmniImportRequestOrderInfo struct {
	BuyerInfo        OrderOmniImportRequestBuyerInfo          `json:"buyerInfo,omitempty"`
	CancelInfo       OrderOmniImportRequestCancelInfo         `json:"cancelInfo,omitempty"`
	DeliveryInfo     OrderOmniImportRequestDeliveryInfo       `json:"deliveryInfo,omitempty"`
	DiscountInfoList []OrderOmniImportRequestDiscountInfoList `json:"discountInfoList,omitempty"`
	GuideInfo        OrderOmniImportRequestGuideInfo          `json:"guideInfo,omitempty"`
	ItemInfoList     []OrderOmniImportRequestItemInfoList     `json:"itemInfoList,omitempty"`
	MerchantInfo     OrderOmniImportRequestMerchantInfo       `json:"merchantInfo,omitempty"`
	OrderBaseInfo    OrderOmniImportRequestOrderBaseInfo      `json:"orderBaseInfo,omitempty"`
	PayInfo          OrderOmniImportRequestPayInfo2           `json:"payInfo,omitempty"`
	PaymentDetails   []OrderOmniImportRequestPaymentDetails   `json:"paymentDetails,omitempty"`
	AssociateBizList []OrderOmniImportRequestAssociateBizList `json:"associateBizList,omitempty"`
}

type OrderOmniImportRequestBuyerInfo struct {
	Phone        string `json:"phone,omitempty"`
	UserNickName string `json:"userNickName,omitempty"`
	BuyerRemark  string `json:"buyerRemark,omitempty"`
	Wid          int64  `json:"wid,omitempty"`
}

type OrderOmniImportRequestCancelInfo struct {
	CancelType    int64  `json:"cancelType,omitempty"`
	Reason        string `json:"reason,omitempty"`
	SpecialReason string `json:"specialReason,omitempty"`
}

type OrderOmniImportRequestDeliveryInfo struct {
	ReceiveInfo             OrderOmniImportRequestReceiveInfo   `json:"receiveInfo,omitempty"`
	SendInfo                OrderOmniImportRequestSendInfo      `json:"sendInfo,omitempty"`
	PackageList             []OrderOmniImportRequestPackageList `json:"packageList,omitempty"`
	DeliveryType            int64                               `json:"deliveryType,omitempty"`
	ExpectReceivedType      int64                               `json:"expectReceivedType,omitempty"`
	ExpectReceivedDate      int64                               `json:"expectReceivedDate,omitempty"`
	ExpectReceivedStartTime int64                               `json:"expectReceivedStartTime,omitempty"`
	ExpectReceivedEndTime   int64                               `json:"expectReceivedEndTime,omitempty"`
}

type OrderOmniImportRequestReceiveInfo struct {
	AddressInfo  OrderOmniImportRequestAddressInfo  `json:"addressInfo,omitempty"`
	ReceiverInfo OrderOmniImportRequestReceiverInfo `json:"receiverInfo,omitempty"`
	PickUpName   string                             `json:"pickUpName,omitempty"`
	PickUpVid    int64                              `json:"pickUpVid,omitempty"`
}

type OrderOmniImportRequestAddressInfo struct {
	AddressExt OrderOmniImportRequestAddressExt `json:"addressExt,omitempty"`
	Address    string                           `json:"address,omitempty"`
	Area       string                           `json:"area,omitempty"`
	City       string                           `json:"city,omitempty"`
	County     string                           `json:"county,omitempty"`
	Latitude   string                           `json:"latitude,omitempty"`
	Longitude  string                           `json:"longitude,omitempty"`
	Province   string                           `json:"province,omitempty"`
	Zip        string                           `json:"zip,omitempty"`
}

type OrderOmniImportRequestAddressExt struct {
	AreaCode     string `json:"areaCode,omitempty"`
	CityCode     string `json:"cityCode,omitempty"`
	CountyCode   string `json:"countyCode,omitempty"`
	ProvinceCode string `json:"provinceCode,omitempty"`
}

type OrderOmniImportRequestReceiverInfo struct {
	ReceiverMobile string `json:"receiverMobile,omitempty"`
	ReceiverName   string `json:"receiverName,omitempty"`
}

type OrderOmniImportRequestSendInfo struct {
	AddressInfo    OrderOmniImportRequestAddressInfo2 `json:"addressInfo,omitempty"`
	SenderInfo     OrderOmniImportRequestSenderInfo   `json:"senderInfo,omitempty"`
	ProcessVid     int64                              `json:"processVid,omitempty"`
	ProcessVidName string                             `json:"processVidName,omitempty"`
}

type OrderOmniImportRequestAddressInfo2 struct {
	Address   string `json:"address,omitempty"`
	Area      string `json:"area,omitempty"`
	City      string `json:"city,omitempty"`
	County    string `json:"county,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
	Province  string `json:"province,omitempty"`
	Zip       string `json:"zip,omitempty"`
}

type OrderOmniImportRequestSenderInfo struct {
	SenderMobile string `json:"senderMobile,omitempty"`
	SenderName   string `json:"senderName,omitempty"`
}

type OrderOmniImportRequestPackageList struct {
	DeliveryImportInfo OrderOmniImportRequestDeliveryImportInfo `json:"deliveryImportInfo,omitempty"`
	ReceiveInfo        OrderOmniImportRequestReceiveInfo2       `json:"receiveInfo,omitempty"`
	SendInfo           OrderOmniImportRequestSendInfo2          `json:"sendInfo,omitempty"`
	PackageItems       []OrderOmniImportRequestPackageItems     `json:"packageItems,omitempty"`
	DeliveryMethod     int64                                    `json:"deliveryMethod,omitempty"`
	DeliveryTime       int64                                    `json:"deliveryTime,omitempty"`
	PackageName        string                                   `json:"packageName,omitempty"`
	ConfirmType        int64                                    `json:"confirmType,omitempty"`
}

type OrderOmniImportRequestDeliveryImportInfo struct {
	CompanyCode            string `json:"companyCode,omitempty"`
	CompanyName            string `json:"companyName,omitempty"`
	ExpectReceivedType     int64  `json:"expectReceivedType,omitempty"`
	ExpectReceivedTypeName string `json:"expectReceivedTypeName,omitempty"`
	Number                 string `json:"number,omitempty"`
	WriteOffName           string `json:"writeOffName,omitempty"`
}

type OrderOmniImportRequestReceiveInfo2 struct {
	AddressInfo  OrderOmniImportRequestAddressInfo3  `json:"addressInfo,omitempty"`
	ReceiverInfo OrderOmniImportRequestReceiverInfo2 `json:"receiverInfo,omitempty"`
	PickUpName   string                              `json:"pickUpName,omitempty"`
	PickUpVid    int64                               `json:"pickUpVid,omitempty"`
}

type OrderOmniImportRequestAddressInfo3 struct {
	AddressExt OrderOmniImportRequestAddressExt2 `json:"addressExt,omitempty"`
	Address    string                            `json:"address,omitempty"`
	Area       string                            `json:"area,omitempty"`
	City       string                            `json:"city,omitempty"`
	County     string                            `json:"county,omitempty"`
	Latitude   string                            `json:"latitude,omitempty"`
	Longitude  string                            `json:"longitude,omitempty"`
	Province   string                            `json:"province,omitempty"`
	Zip        string                            `json:"zip,omitempty"`
}

type OrderOmniImportRequestAddressExt2 struct {
	AreaCode     string `json:"areaCode,omitempty"`
	CityCode     string `json:"cityCode,omitempty"`
	CountyCode   string `json:"countyCode,omitempty"`
	ProvinceCode string `json:"provinceCode,omitempty"`
}

type OrderOmniImportRequestReceiverInfo2 struct {
	ReceiverMobile string `json:"receiverMobile,omitempty"`
	ReceiverName   string `json:"receiverName,omitempty"`
}

type OrderOmniImportRequestSendInfo2 struct {
	AddressInfo    OrderOmniImportRequestAddressInfo4 `json:"addressInfo,omitempty"`
	SenderInfo     OrderOmniImportRequestSenderInfo2  `json:"senderInfo,omitempty"`
	ProcessVid     int64                              `json:"processVid,omitempty"`
	ProcessVidName string                             `json:"processVidName,omitempty"`
}

type OrderOmniImportRequestAddressInfo4 struct {
	Address   string `json:"address,omitempty"`
	Area      string `json:"area,omitempty"`
	City      string `json:"city,omitempty"`
	County    string `json:"county,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
	Province  string `json:"province,omitempty"`
	Zip       string `json:"zip,omitempty"`
}

type OrderOmniImportRequestSenderInfo2 struct {
	SenderMobile string `json:"senderMobile,omitempty"`
	SenderName   string `json:"senderName,omitempty"`
}

type OrderOmniImportRequestPackageItems struct {
	OutItemId int64  `json:"outItemId,omitempty"`
	SkuNum    string `json:"skuNum,omitempty"`
}

type OrderOmniImportRequestDiscountInfoList struct {
	DiscountExt    OrderOmniImportRequestDiscountExt `json:"discountExt,omitempty"`
	DiscountId     string                            `json:"discountId,omitempty"`
	DiscountType   int64                             `json:"discountType,omitempty"`
	DiscountLevel  int64                             `json:"discountLevel,omitempty"`
	SubType        int64                             `json:"subType,omitempty"`
	Name           string                            `json:"name,omitempty"`
	DiscountAmount string                            `json:"discountAmount,omitempty"`
	CostAmount     string                            `json:"costAmount,omitempty"`
}

type OrderOmniImportRequestDiscountExt struct {
	DiscountExtInfo OrderOmniImportRequestDiscountExtInfo `json:"discountExtInfo,omitempty"`
	AttributionType int64                                 `json:"attributionType,omitempty"`
	Participate     string                                `json:"participate,omitempty"`
}

type OrderOmniImportRequestDiscountExtInfo struct {
}

type OrderOmniImportRequestGuideInfo struct {
	GuiderName          string `json:"guiderName,omitempty"`
	GuiderPhone         string `json:"guiderPhone,omitempty"`
	GuiderWid           int64  `json:"guiderWid,omitempty"`
	PersonalGuiderPhone string `json:"personalGuiderPhone,omitempty"`
	PersonalGuiderNo    string `json:"personalGuiderNo,omitempty"`
	GuiderNo            string `json:"guiderNo,omitempty"`
}

type OrderOmniImportRequestItemInfoList struct {
	DiscountInfoList []OrderOmniImportRequestDiscountInfoList2 `json:"discountInfoList,omitempty"`
	GoodsInfo        OrderOmniImportRequestGoodsInfo           `json:"goodsInfo,omitempty"`
	PayInfo          OrderOmniImportRequestPayInfo             `json:"payInfo,omitempty"`
	OutItemId        string                                    `json:"outItemId,omitempty"`
}

type OrderOmniImportRequestDiscountInfoList2 struct {
	DiscountExt    OrderOmniImportRequestDiscountExt2 `json:"discountExt,omitempty"`
	DiscountId     string                             `json:"discountId,omitempty"`
	DiscountType   int64                              `json:"discountType,omitempty"`
	DiscountLevel  int64                              `json:"discountLevel,omitempty"`
	SubType        int64                              `json:"subType,omitempty"`
	Name           string                             `json:"name,omitempty"`
	DiscountAmount string                             `json:"discountAmount,omitempty"`
	CostAmount     string                             `json:"costAmount,omitempty"`
}

type OrderOmniImportRequestDiscountExt2 struct {
	DiscountExtInfo OrderOmniImportRequestDiscountExtInfo2 `json:"discountExtInfo,omitempty"`
	AttributionType int64                                  `json:"attributionType,omitempty"`
	Participate     string                                 `json:"participate,omitempty"`
}

type OrderOmniImportRequestDiscountExtInfo2 struct {
}

type OrderOmniImportRequestGoodsInfo struct {
	CategoryTitle []OrderOmniImportRequestCategoryTitle `json:"categoryTitle,omitempty"`
	SalePrice     string                                `json:"salePrice,omitempty"`
	SkuId         int64                                 `json:"skuId,omitempty"`
	SkuNum        string                                `json:"skuNum,omitempty"`
	GoodsSellMode int64                                 `json:"goodsSellMode,omitempty"`
	SkuCode       string                                `json:"skuCode,omitempty"`
	CategoryId    int64                                 `json:"categoryId,omitempty"`
}

type OrderOmniImportRequestCategoryTitle struct {
}

type OrderOmniImportRequestPayInfo struct {
	AmountInfos           []OrderOmniImportRequestAmountInfos `json:"amountInfos,omitempty"`
	PayAmount             string                              `json:"payAmount,omitempty"`
	TotalAmount           string                              `json:"totalAmount,omitempty"`
	TotalDiscountAmount   string                              `json:"totalDiscountAmount,omitempty"`
	ShouldPayAmount       string                              `json:"shouldPayAmount,omitempty"`
	SettlementFee         string                              `json:"settlementFee,omitempty"`
	PlatformSubsidyAmount string                              `json:"platformSubsidyAmount,omitempty"`
}

type OrderOmniImportRequestAmountInfos struct {
	Type            int64  `json:"type,omitempty"`
	Description     string `json:"description,omitempty"`
	Amount          string `json:"amount,omitempty"`
	PayAmount       string `json:"payAmount,omitempty"`
	ShouldPayAmount string `json:"shouldPayAmount,omitempty"`
}

type OrderOmniImportRequestMerchantInfo struct {
	BosId          int64  `json:"bosId,omitempty"`
	BosName        string `json:"bosName,omitempty"`
	MerchantId     int64  `json:"merchantId,omitempty"`
	ProcessVid     int64  `json:"processVid,omitempty"`
	ProcessVidName string `json:"processVidName,omitempty"`
	ProcessVidType int64  `json:"processVidType,omitempty"`
	Vid            int64  `json:"vid,omitempty"`
	VidName        string `json:"vidName,omitempty"`
	VidType        int64  `json:"vidType,omitempty"`
}

type OrderOmniImportRequestOrderBaseInfo struct {
	TimeList        []OrderOmniImportRequestTimeList `json:"timeList,omitempty"`
	ChannelType     int64                            `json:"channelType,omitempty"`
	PayStatus       int64                            `json:"payStatus,omitempty"`
	PayTime         int64                            `json:"payTime,omitempty"`
	PayType         int64                            `json:"payType,omitempty"`
	OuterOrderNo    string                           `json:"outerOrderNo,omitempty"`
	SaleChannelType int64                            `json:"saleChannelType,omitempty"`
	BizSourceType   int64                            `json:"bizSourceType,omitempty"`
	OrderStatus     int64                            `json:"orderStatus,omitempty"`
	InsuranceType   int64                            `json:"insuranceType,omitempty"`
	DepositPayType  int64                            `json:"depositPayType,omitempty"`
}

type OrderOmniImportRequestTimeList struct {
	Type  int64 `json:"type,omitempty"`
	Value int64 `json:"value,omitempty"`
}

type OrderOmniImportRequestPayInfo2 struct {
	AmountInfos           []OrderOmniImportRequestAmountInfos2 `json:"amountInfos,omitempty"`
	PayAmount             string                               `json:"payAmount,omitempty"`
	TotalAmount           string                               `json:"totalAmount,omitempty"`
	TotalDiscountAmount   string                               `json:"totalDiscountAmount,omitempty"`
	ShouldPayAmount       string                               `json:"shouldPayAmount,omitempty"`
	SettlementFee         string                               `json:"settlementFee,omitempty"`
	PlatformSubsidyAmount string                               `json:"platformSubsidyAmount,omitempty"`
}

type OrderOmniImportRequestAmountInfos2 struct {
	Type            int64  `json:"type,omitempty"`
	Description     string `json:"description,omitempty"`
	Amount          string `json:"amount,omitempty"`
	PayAmount       string `json:"payAmount,omitempty"`
	ShouldPayAmount string `json:"shouldPayAmount,omitempty"`
}

type OrderOmniImportRequestPaymentDetails struct {
	InstallmentInfo OrderOmniImportRequestInstallmentInfo `json:"installmentInfo,omitempty"`
	PayTime         int64                                 `json:"payTime,omitempty"`
	Phase           int64                                 `json:"phase,omitempty"`
	ChannelTrxNo    string                                `json:"channelTrxNo,omitempty"`
	PayMethodIds    []int64                               `json:"payMethodIds,omitempty"`
	InterestAmount  string                                `json:"interestAmount,omitempty"`
}

type OrderOmniImportRequestInstallmentInfo struct {
	BuyerInterestAmount    string `json:"buyerInterestAmount,omitempty"`
	SellerInterestAmount   string `json:"sellerInterestAmount,omitempty"`
	PlatformInterestAmount string `json:"platformInterestAmount,omitempty"`
}

type OrderOmniImportRequestAssociateBizList struct {
	BizValue int64 `json:"bizValue,omitempty"`
}

type OrderOmniImportResponseOutputInfo struct {
	OrderNo int64 `json:"orderNo,omitempty"`
}
