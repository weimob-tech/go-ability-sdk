package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideOrderGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideOrderGetListRequest, PaasWeimobGuideOrderGetListResponse]
}

func (s PaasWeimobGuideOrderGetListService) NewRequest() spi.InvocationRequest[PaasWeimobGuideOrderGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideOrderGetListRequest]{
		Params: &PaasWeimobGuideOrderGetListRequest{},
	}
}

func (s PaasWeimobGuideOrderGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideOrderGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideOrderGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideOrderGetListRequest]))
}

type PaasWeimobGuideOrderGetListRequest struct {
}

type PaasWeimobGuideOrderGetListResponse struct {
	BuyerExpandInfoVo   PaasWeimobGuideOrderGetListResponseBuyerExpandInfoVo `json:"buyerExpandInfoVo,omitempty"`
	BuyerInfo           PaasWeimobGuideOrderGetListResponseBuyerInfo         `json:"buyerInfo,omitempty"`
	CreateTime          PaasWeimobGuideOrderGetListResponseCreateTime        `json:"createTime,omitempty"`
	FinishTime          PaasWeimobGuideOrderGetListResponseFinishTime        `json:"finishTime,omitempty"`
	ItemList            []PaasWeimobGuideOrderGetListResponseItemList        `json:"itemList,omitempty"`
	PaymentInfo         PaasWeimobGuideOrderGetListResponsePaymentInfo       `json:"paymentInfo,omitempty"`
	UpdateTime          PaasWeimobGuideOrderGetListResponseUpdateTime        `json:"updateTime,omitempty"`
	BizSourceType       int64                                                `json:"bizSourceType,omitempty"`
	BuyerRemark         string                                               `json:"buyerRemark,omitempty"`
	ChannelType         int64                                                `json:"channelType,omitempty"`
	GoodsAmount         int64                                                `json:"goodsAmount,omitempty"`
	GoodsGuideType      int64                                                `json:"goodsGuideType,omitempty"`
	GoodsSourceType     int64                                                `json:"goodsSourceType,omitempty"`
	IsDeleted           int64                                                `json:"isDeleted,omitempty"`
	OrderEntrance       int64                                                `json:"orderEntrance,omitempty"`
	OrderNo             int64                                                `json:"orderNo,omitempty"`
	OrderStatus         int64                                                `json:"orderStatus,omitempty"`
	OrderType           int64                                                `json:"orderType,omitempty"`
	OutOrderNo          int64                                                `json:"outOrderNo,omitempty"`
	PaymentAmount       int64                                                `json:"paymentAmount,omitempty"`
	SaleChannelType     int64                                                `json:"saleChannelType,omitempty"`
	ShouldPayAmount     int64                                                `json:"shouldPayAmount,omitempty"`
	TotalDiscountAmount int64                                                `json:"totalDiscountAmount,omitempty"`
}
type PaasWeimobGuideOrderGetListResponseBuyerExpandInfoVo struct {
	AttributionStoreId    int64  `json:"attributionStoreId,omitempty"`
	AttributionStoreName  string `json:"attributionStoreName,omitempty"`
	PersonalGuiderName    string `json:"personalGuiderName,omitempty"`
	PersonalGuiderStoreId int64  `json:"personalGuiderStoreId,omitempty"`
	PersonalGuiderWid     int64  `json:"personalGuiderWid,omitempty"`
}
type PaasWeimobGuideOrderGetListResponseBuyerInfo struct {
	UserNickname string `json:"userNickname,omitempty"`
	Wid          int64  `json:"wid,omitempty"`
}
type PaasWeimobGuideOrderGetListResponseCreateTime struct {
}
type PaasWeimobGuideOrderGetListResponseFinishTime struct {
}
type PaasWeimobGuideOrderGetListResponseItemList struct {
	BaseDiscountInfo      PaasWeimobGuideOrderGetListResponseBaseDiscountInfo  `json:"baseDiscountInfo,omitempty"`
	GoodsDiscountInfo     PaasWeimobGuideOrderGetListResponseGoodsDiscountInfo `json:"goodsDiscountInfo,omitempty"`
	AvailableStoreId      int64                                                `json:"availableStoreId,omitempty"`
	CloudCustom           string                                               `json:"cloudCustom,omitempty"`
	CostPrice             int64                                                `json:"costPrice,omitempty"`
	CreateStoreId         int64                                                `json:"createStoreId,omitempty"`
	DepositAmount         int64                                                `json:"depositAmount,omitempty"`
	GoodsBelongToType     int64                                                `json:"goodsBelongToType,omitempty"`
	GoodsCategoryId       int64                                                `json:"goodsCategoryId,omitempty"`
	GoodsCode             int64                                                `json:"goodsCode,omitempty"`
	GoodsDistributionType int64                                                `json:"goodsDistributionType,omitempty"`
	GoodsId               int64                                                `json:"goodsId,omitempty"`
	GoodsSellMode         int64                                                `json:"goodsSellMode,omitempty"`
	GoodsSourceType       int64                                                `json:"goodsSourceType,omitempty"`
	GoodsTitle            string                                               `json:"goodsTitle,omitempty"`
	GoodsType             int64                                                `json:"goodsType,omitempty"`
	HasDelivered          bool                                                 `json:"hasDelivered,omitempty"`
	Id                    int64                                                `json:"id,omitempty"`
	ImageUrl              string                                               `json:"imageUrl,omitempty"`
	InvoiceGoodsName      string                                               `json:"invoiceGoodsName,omitempty"`
	ItemRemarkName        string                                               `json:"itemRemarkName,omitempty"`
	MarketPrice           int64                                                `json:"marketPrice,omitempty"`
	OriginalPrice         int64                                                `json:"originalPrice,omitempty"`
	OutItemId             int64                                                `json:"outItemId,omitempty"`
	OuterGoodsId          int64                                                `json:"outerGoodsId,omitempty"`
	OuterMerchantId       int64                                                `json:"outerMerchantId,omitempty"`
	OuterSkuId            int64                                                `json:"outerSkuId,omitempty"`
	PaymentAmount         int64                                                `json:"paymentAmount,omitempty"`
	Point                 int64                                                `json:"point,omitempty"`
	PointDeductRatio      int64                                                `json:"pointDeductRatio,omitempty"`
	Price                 int64                                                `json:"price,omitempty"`
	PrimaryCategory       int64                                                `json:"primaryCategory,omitempty"`
	ProductDetails        string                                               `json:"productDetails,omitempty"`
	ProductType           int64                                                `json:"productType,omitempty"`
	SecondaryCategory     int64                                                `json:"secondaryCategory,omitempty"`
	ShouldPaymentAmount   int64                                                `json:"shouldPaymentAmount,omitempty"`
	SkuAmount             int64                                                `json:"skuAmount,omitempty"`
	SkuBarCode            int64                                                `json:"skuBarCode,omitempty"`
	SkuCode               int64                                                `json:"skuCode,omitempty"`
	SkuId                 int64                                                `json:"skuId,omitempty"`
	SkuName               string                                               `json:"skuName,omitempty"`
	SkuNum                int64                                                `json:"skuNum,omitempty"`
	TaxClassificationCode int64                                                `json:"taxClassificationCode,omitempty"`
	TaxRate               int64                                                `json:"taxRate,omitempty"`
	TotalAmount           int64                                                `json:"totalAmount,omitempty"`
	TotalDiscountAmount   int64                                                `json:"totalDiscountAmount,omitempty"`
	TotalPoint            int64                                                `json:"totalPoint,omitempty"`
	Unit                  string                                               `json:"unit,omitempty"`
	Volume                int64                                                `json:"volume,omitempty"`
	Weight                int64                                                `json:"weight,omitempty"`
}
type PaasWeimobGuideOrderGetListResponseBaseDiscountInfo struct {
	BalanceDiscountAmount          int64 `json:"balanceDiscountAmount,omitempty"`
	CardDiscountAmount             int64 `json:"cardDiscountAmount,omitempty"`
	CouponCodeDiscountAmount       int64 `json:"couponCodeDiscountAmount,omitempty"`
	CouponDiscountAmount           int64 `json:"couponDiscountAmount,omitempty"`
	DepositExpansionDiscountAmount int64 `json:"depositExpansionDiscountAmount,omitempty"`
	GiftCardDiscountAmount         int64 `json:"giftCardDiscountAmount,omitempty"`
	MemberPointsDiscountAmount     int64 `json:"memberPointsDiscountAmount,omitempty"`
	MembershipDiscountAmount       int64 `json:"membershipDiscountAmount,omitempty"`
	MerchantDiscountAmount         int64 `json:"merchantDiscountAmount,omitempty"`
	NynjDiscountAmount             int64 `json:"nynjDiscountAmount,omitempty"`
	PayDiscountAmount              int64 `json:"payDiscountAmount,omitempty"`
	PromotionDiscountAmount        int64 `json:"promotionDiscountAmount,omitempty"`
	TieredPriceDiscountAmount      int64 `json:"tieredPriceDiscountAmount,omitempty"`
	UsedMemberPoints               int64 `json:"usedMemberPoints,omitempty"`
}
type PaasWeimobGuideOrderGetListResponseGoodsDiscountInfo struct {
	CommonDiscountInfo             []PaasWeimobGuideOrderGetListResponseCommonDiscountInfo `json:"commonDiscountInfo,omitempty"`
	BalanceDiscountAmount          int64                                                   `json:"balanceDiscountAmount,omitempty"`
	CardDiscountAmount             int64                                                   `json:"cardDiscountAmount,omitempty"`
	CouponCodeDiscountAmount       int64                                                   `json:"couponCodeDiscountAmount,omitempty"`
	CouponDiscountAmount           int64                                                   `json:"couponDiscountAmount,omitempty"`
	DepositExpansionDiscountAmount int64                                                   `json:"depositExpansionDiscountAmount,omitempty"`
	GiftCardDiscountAmount         int64                                                   `json:"giftCardDiscountAmount,omitempty"`
	MemberPointsDiscountAmount     int64                                                   `json:"memberPointsDiscountAmount,omitempty"`
	MembershipDiscountAmount       int64                                                   `json:"membershipDiscountAmount,omitempty"`
	MerchantDiscountAmount         int64                                                   `json:"merchantDiscountAmount,omitempty"`
	NynjDiscountAmount             int64                                                   `json:"nynjDiscountAmount,omitempty"`
	PayDiscountAmount              int64                                                   `json:"payDiscountAmount,omitempty"`
	PromotionDiscountAmount        int64                                                   `json:"promotionDiscountAmount,omitempty"`
	TieredPriceDiscountAmount      int64                                                   `json:"tieredPriceDiscountAmount,omitempty"`
	UsedMemberPoints               int64                                                   `json:"usedMemberPoints,omitempty"`
}
type PaasWeimobGuideOrderGetListResponseCommonDiscountInfo struct {
	ActivityId         int64  `json:"activityId,omitempty"`
	ActivityType       int64  `json:"activityType,omitempty"`
	AttributionType    int64  `json:"attributionType,omitempty"`
	DiscountAmount     int64  `json:"discountAmount,omitempty"`
	DiscountObjectType int64  `json:"discountObjectType,omitempty"`
	DiscountType       int64  `json:"discountType,omitempty"`
	ExtInfo            string `json:"extInfo,omitempty"`
	ItemId             int64  `json:"itemId,omitempty"`
	Name               string `json:"name,omitempty"`
	OrderNo            int64  `json:"orderNo,omitempty"`
	OutCode            int64  `json:"outCode,omitempty"`
	Participate        int64  `json:"participate,omitempty"`
	Points             int64  `json:"points,omitempty"`
}
type PaasWeimobGuideOrderGetListResponsePaymentInfo struct {
	PaymentItemVoList []PaasWeimobGuideOrderGetListResponsePaymentItemVoList `json:"paymentItemVoList,omitempty"`
	PaymentTime       PaasWeimobGuideOrderGetListResponsePaymentTime2        `json:"paymentTime,omitempty"`
	ChannelTrxNo      int64                                                  `json:"channelTrxNo,omitempty"`
	PaymentId         int64                                                  `json:"paymentId,omitempty"`
	PaymentInfoId     int64                                                  `json:"paymentInfoId,omitempty"`
	PaymentMethodId   string                                                 `json:"paymentMethodId,omitempty"`
	PaymentMethodName string                                                 `json:"paymentMethodName,omitempty"`
	PaymentStatus     int64                                                  `json:"paymentStatus,omitempty"`
	PaymentType       int64                                                  `json:"paymentType,omitempty"`
	PaymentTypeName   string                                                 `json:"paymentTypeName,omitempty"`
	TradeId           int64                                                  `json:"tradeId,omitempty"`
}
type PaasWeimobGuideOrderGetListResponsePaymentItemVoList struct {
	PaymentTime         PaasWeimobGuideOrderGetListResponsePaymentTime `json:"paymentTime,omitempty"`
	ChangeAmount        int64                                          `json:"changeAmount,omitempty"`
	ChannelTrxNo        int64                                          `json:"channelTrxNo,omitempty"`
	ChannelType         int64                                          `json:"channelType,omitempty"`
	CustomPayMethodId   int64                                          `json:"customPayMethodId,omitempty"`
	CustomPayMethodName string                                         `json:"customPayMethodName,omitempty"`
	InteractId          int64                                          `json:"interactId,omitempty"`
	OverseaInfo         string                                         `json:"overseaInfo,omitempty"`
	PartnerMode         int64                                          `json:"partnerMode,omitempty"`
	PaymentAmount       int64                                          `json:"paymentAmount,omitempty"`
	PaymentId           int64                                          `json:"paymentId,omitempty"`
	PaymentMethodId     int64                                          `json:"paymentMethodId,omitempty"`
	PaymentType         int64                                          `json:"paymentType,omitempty"`
	SettlementAmount    int64                                          `json:"settlementAmount,omitempty"`
	TradeId             int64                                          `json:"tradeId,omitempty"`
}
type PaasWeimobGuideOrderGetListResponsePaymentTime struct {
}
type PaasWeimobGuideOrderGetListResponsePaymentTime2 struct {
}
type PaasWeimobGuideOrderGetListResponseUpdateTime struct {
}

func CreatePaasWeimobGuideOrderGetListResponse() *PaasWeimobGuideOrderGetListResponse {
	return &PaasWeimobGuideOrderGetListResponse{}
}

// OnPaasWeimobGuideOrderGetListServiceInvocation
// @id 735
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/735?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询订单列表)
func (s *Service) OnPaasWeimobGuideOrderGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideOrderGetListRequest, PaasWeimobGuideOrderGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideOrderGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideOrderGetListRequest, PaasWeimobGuideOrderGetListResponse](invocation),
		"PaasWeimobGuideOrderGetListService",
		"weimobGuide.order.getList",
	)
	return s
}
