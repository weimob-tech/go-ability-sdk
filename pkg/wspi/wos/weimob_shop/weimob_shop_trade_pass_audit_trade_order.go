package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobShopTradePassAuditTradeOrderService struct {
	handler spi.WosInvocationHandler[PaasWeimobShopTradePassAuditTradeOrderRequest, PaasWeimobShopTradePassAuditTradeOrderResponse]
}

func (s PaasWeimobShopTradePassAuditTradeOrderService) NewRequest() spi.InvocationRequest[PaasWeimobShopTradePassAuditTradeOrderRequest] {
	return &spi.WosInvocationRequest[PaasWeimobShopTradePassAuditTradeOrderRequest]{
		Params: &PaasWeimobShopTradePassAuditTradeOrderRequest{},
	}
}

func (s PaasWeimobShopTradePassAuditTradeOrderService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobShopTradePassAuditTradeOrderRequest],
) (
	spi.InvocationResponse[PaasWeimobShopTradePassAuditTradeOrderResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobShopTradePassAuditTradeOrderRequest]))
}

type PaasWeimobShopTradePassAuditTradeOrderRequest struct {
	BasicInfo                    PaasWeimobShopTradePassAuditTradeOrderRequestBasicInfo                      `json:"basicInfo,omitempty"`
	ConfirmOrderShopGroupReqList []PaasWeimobShopTradePassAuditTradeOrderRequestConfirmOrderShopGroupReqList `json:"confirmOrderShopGroupReqList,omitempty"`
	ProductId                    int64                                                                       `json:"productId,omitempty"`
	ProductInstanceId            int64                                                                       `json:"productInstanceId,omitempty"`
	TargetProductId              int64                                                                       `json:"targetProductId,omitempty"`
	TargetProductInstanceId      int64                                                                       `json:"targetProductInstanceId,omitempty"`
	Tcode                        string                                                                      `json:"tcode,omitempty"`
	MerchantId                   int64                                                                       `json:"merchantId,omitempty"`
}
type PaasWeimobShopTradePassAuditTradeOrderRequestBasicInfo struct {
	MerchantId              int64 `json:"merchantId,omitempty"`
	BosId                   int64 `json:"bosId,omitempty"`
	SourceProductId         int64 `json:"sourceProductId,omitempty"`
	SourceProductVersionId  int64 `json:"sourceProductVersionId,omitempty"`
	SourceProductInstanceId int64 `json:"sourceProductInstanceId,omitempty"`
}
type PaasWeimobShopTradePassAuditTradeOrderRequestConfirmOrderShopGroupReqList struct {
	GoodsReqList  []PaasWeimobShopTradePassAuditTradeOrderRequestGoodsReqList `json:"goodsReqList,omitempty"`
	PaymentInfo   PaasWeimobShopTradePassAuditTradeOrderRequestPaymentInfo    `json:"paymentInfo,omitempty"`
	BosId         int64                                                       `json:"bosId,omitempty"`
	MerchantId    int64                                                       `json:"merchantId,omitempty"`
	MerchantName  string                                                      `json:"merchantName,omitempty"`
	Wid           int64                                                       `json:"wid,omitempty"`
	Vid           int64                                                       `json:"vid,omitempty"`
	VidName       string                                                      `json:"vidName,omitempty"`
	ChannelType   int64                                                       `json:"channelType,omitempty"`
	OrderNo       int64                                                       `json:"orderNo,omitempty"`
	ParentOrderNo int64                                                       `json:"parentOrderNo,omitempty"`
}
type PaasWeimobShopTradePassAuditTradeOrderRequestGoodsReqList struct {
	GoodsAbilityReq  PaasWeimobShopTradePassAuditTradeOrderRequestGoodsAbilityReq    `json:"goodsAbilityReq,omitempty"`
	UseActivityList  []PaasWeimobShopTradePassAuditTradeOrderRequestUseActivityList  `json:"useActivityList,omitempty"`
	ItemProductInfo  PaasWeimobShopTradePassAuditTradeOrderRequestItemProductInfo    `json:"itemProductInfo,omitempty"`
	GoodsExtMap      PaasWeimobShopTradePassAuditTradeOrderRequestGoodsExtMap        `json:"goodsExtMap,omitempty"`
	GoodsAbilityList []PaasWeimobShopTradePassAuditTradeOrderRequestGoodsAbilityList `json:"goodsAbilityList,omitempty"`
	ItemId           int64                                                           `json:"itemId,omitempty"`
	GoodsId          int64                                                           `json:"goodsId,omitempty"`
	SkuId            int64                                                           `json:"skuId,omitempty"`
	BuyNum           int64                                                           `json:"buyNum,omitempty"`
	SkuSalePrice     int64                                                           `json:"skuSalePrice,omitempty"`
	SkuMarketPrice   string                                                          `json:"skuMarketPrice,omitempty"`
	TradeGoodsType   int64                                                           `json:"tradeGoodsType,omitempty"`
	GoodsType        int64                                                           `json:"goodsType,omitempty"`
	SubGoodsType     int64                                                           `json:"subGoodsType,omitempty"`
	SoldType         int64                                                           `json:"soldType,omitempty"`
}
type PaasWeimobShopTradePassAuditTradeOrderRequestGoodsAbilityReq struct {
	CycleInfoReq PaasWeimobShopTradePassAuditTradeOrderRequestCycleInfoReq `json:"cycleInfoReq,omitempty"`
}
type PaasWeimobShopTradePassAuditTradeOrderRequestCycleInfoReq struct {
	CycleBizId string `json:"cycleBizId,omitempty"`
}
type PaasWeimobShopTradePassAuditTradeOrderRequestUseActivityList struct {
	ActivityId         string `json:"activityId,omitempty"`
	ActivityName       string `json:"activityName,omitempty"`
	ActivityType       int64  `json:"activityType,omitempty"`
	DiscountAmount     string `json:"discountAmount,omitempty"`
	DiscountBizOrderId string `json:"discountBizOrderId,omitempty"`
	GroupKey           string `json:"groupKey,omitempty"`
	Level              int64  `json:"level,omitempty"`
	TemplateId         int64  `json:"templateId,omitempty"`
}
type PaasWeimobShopTradePassAuditTradeOrderRequestItemProductInfo struct {
	ProductDetailList []PaasWeimobShopTradePassAuditTradeOrderRequestProductDetailList `json:"productDetailList,omitempty"`
}
type PaasWeimobShopTradePassAuditTradeOrderRequestProductDetailList struct {
	WarehouseList []PaasWeimobShopTradePassAuditTradeOrderRequestWarehouseList `json:"warehouseList,omitempty"`
}
type PaasWeimobShopTradePassAuditTradeOrderRequestWarehouseList struct {
	WarehouseId   int64  `json:"warehouseId,omitempty"`
	WarehouseName string `json:"warehouseName,omitempty"`
	WarehouseType string `json:"warehouseType,omitempty"`
}
type PaasWeimobShopTradePassAuditTradeOrderRequestGoodsExtMap struct {
}
type PaasWeimobShopTradePassAuditTradeOrderRequestGoodsAbilityList struct {
	AbilityDetailInfo PaasWeimobShopTradePassAuditTradeOrderRequestAbilityDetailInfo `json:"abilityDetailInfo,omitempty"`
	AbilityCode       string                                                         `json:"abilityCode,omitempty"`
	AbilityType       int64                                                          `json:"abilityType,omitempty"`
	BizId             string                                                         `json:"bizId,omitempty"`
}
type PaasWeimobShopTradePassAuditTradeOrderRequestAbilityDetailInfo struct {
}
type PaasWeimobShopTradePassAuditTradeOrderRequestPaymentInfo struct {
	DeductibleAssetList []PaasWeimobShopTradePassAuditTradeOrderRequestDeductibleAssetList `json:"DeductibleAssetList,omitempty"`
	PaymentAmount       string                                                             `json:"paymentAmount,omitempty"`
	TotalAmount         string                                                             `json:"totalAmount,omitempty"`
	ShouldPaymentAmount string                                                             `json:"shouldPaymentAmount,omitempty"`
	TotalDiscountAmount string                                                             `json:"totalDiscountAmount,omitempty"`
}
type PaasWeimobShopTradePassAuditTradeOrderRequestDeductibleAssetList struct {
	Type   int64  `json:"type,omitempty"`
	Amount string `json:"amount,omitempty"`
}

type PaasWeimobShopTradePassAuditTradeOrderResponse struct {
	CommitInfo PaasWeimobShopTradePassAuditTradeOrderResponseCommitInfo `json:"commitInfo,omitempty"`
}
type PaasWeimobShopTradePassAuditTradeOrderResponseCommitInfo struct {
	CommitInfo     bool   `json:"commitInfo,omitempty"`
	UnCommitReason string `json:"unCommitReason	,omitempty"`
}

func CreatePaasWeimobShopTradePassAuditTradeOrderResponse() *PaasWeimobShopTradePassAuditTradeOrderResponse {
	return &PaasWeimobShopTradePassAuditTradeOrderResponse{}
}

// OnPaasWeimobShopTradePassAuditTradeOrderServiceInvocation
// @id 1025
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1025?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=订单审核)
func (s *Service) OnPaasWeimobShopTradePassAuditTradeOrderServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobShopTradePassAuditTradeOrderRequest, PaasWeimobShopTradePassAuditTradeOrderResponse],
) (service *Service) {
	var invocation = &PaasWeimobShopTradePassAuditTradeOrderService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobShopTradePassAuditTradeOrderRequest, PaasWeimobShopTradePassAuditTradeOrderResponse](invocation),
		"PaasWeimobShopTradePassAuditTradeOrderService",
		"weimobShop.trade.passAuditTradeOrder",
	)
	return s
}
