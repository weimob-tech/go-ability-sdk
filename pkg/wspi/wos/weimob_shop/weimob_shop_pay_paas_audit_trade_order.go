package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobShopPayPaasAuditTradeOrderService struct {
	handler spi.WosInvocationHandler[PaasWeimobShopPayPaasAuditTradeOrderRequest, PaasWeimobShopPayPaasAuditTradeOrderResponse]
}

func (s PaasWeimobShopPayPaasAuditTradeOrderService) NewRequest() spi.InvocationRequest[PaasWeimobShopPayPaasAuditTradeOrderRequest] {
	return &spi.WosInvocationRequest[PaasWeimobShopPayPaasAuditTradeOrderRequest]{
		Params: &PaasWeimobShopPayPaasAuditTradeOrderRequest{},
	}
}

func (s PaasWeimobShopPayPaasAuditTradeOrderService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobShopPayPaasAuditTradeOrderRequest],
) (
	spi.InvocationResponse[PaasWeimobShopPayPaasAuditTradeOrderResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobShopPayPaasAuditTradeOrderRequest]))
}

type PaasWeimobShopPayPaasAuditTradeOrderRequest struct {
	BasicInfo                    PaasWeimobShopPayPaasAuditTradeOrderRequestBasicInfo                      `json:"basicInfo,omitempty"`
	ConfirmOrderShopGroupReqList []PaasWeimobShopPayPaasAuditTradeOrderRequestConfirmOrderShopGroupReqList `json:"confirmOrderShopGroupReqList,omitempty"`
	ProductId                    int64                                                                     `json:"productId,omitempty"`
	ProductInstanceId            int64                                                                     `json:"productInstanceId,omitempty"`
	TargetProductId              int64                                                                     `json:"targetProductId,omitempty"`
	TargetProductInstanceId      int64                                                                     `json:"targetProductInstanceId,omitempty"`
	Tcode                        string                                                                    `json:"tcode,omitempty"`
	MerchantId                   int64                                                                     `json:"merchantId,omitempty"`
}
type PaasWeimobShopPayPaasAuditTradeOrderRequestBasicInfo struct {
	MerchantId              int64 `json:"merchantId,omitempty"`
	BosId                   int64 `json:"bosId,omitempty"`
	SourceProductId         int64 `json:"sourceProductId,omitempty"`
	SourceProductVersionId  int64 `json:"sourceProductVersionId,omitempty"`
	SourceProductInstanceId int64 `json:"sourceProductInstanceId,omitempty"`
}
type PaasWeimobShopPayPaasAuditTradeOrderRequestConfirmOrderShopGroupReqList struct {
	GoodsReqList  []PaasWeimobShopPayPaasAuditTradeOrderRequestGoodsReqList `json:"goodsReqList,omitempty"`
	PaymentInfo   PaasWeimobShopPayPaasAuditTradeOrderRequestPaymentInfo    `json:"paymentInfo,omitempty"`
	BosId         int64                                                     `json:"bosId,omitempty"`
	MerchantId    int64                                                     `json:"merchantId,omitempty"`
	MerchantName  string                                                    `json:"merchantName,omitempty"`
	Wid           int64                                                     `json:"wid,omitempty"`
	Vid           int64                                                     `json:"vid,omitempty"`
	VidName       string                                                    `json:"vidName,omitempty"`
	ChannelType   int64                                                     `json:"channelType,omitempty"`
	OrderNo       int64                                                     `json:"orderNo,omitempty"`
	ParentOrderNo int64                                                     `json:"parentOrderNo,omitempty"`
}
type PaasWeimobShopPayPaasAuditTradeOrderRequestGoodsReqList struct {
	GoodsAbilityReq  PaasWeimobShopPayPaasAuditTradeOrderRequestGoodsAbilityReq    `json:"goodsAbilityReq,omitempty"`
	UseActivityList  []PaasWeimobShopPayPaasAuditTradeOrderRequestUseActivityList  `json:"useActivityList,omitempty"`
	ItemProductInfo  PaasWeimobShopPayPaasAuditTradeOrderRequestItemProductInfo    `json:"itemProductInfo,omitempty"`
	GoodsExtMap      PaasWeimobShopPayPaasAuditTradeOrderRequestGoodsExtMap        `json:"goodsExtMap,omitempty"`
	GoodsAbilityList []PaasWeimobShopPayPaasAuditTradeOrderRequestGoodsAbilityList `json:"goodsAbilityList,omitempty"`
	ItemId           int64                                                         `json:"itemId,omitempty"`
	GoodsId          int64                                                         `json:"goodsId,omitempty"`
	SkuId            int64                                                         `json:"skuId,omitempty"`
	BuyNum           int64                                                         `json:"buyNum,omitempty"`
	SkuSalePrice     int64                                                         `json:"skuSalePrice,omitempty"`
	SkuMarketPrice   string                                                        `json:"skuMarketPrice,omitempty"`
	TradeGoodsType   int64                                                         `json:"tradeGoodsType,omitempty"`
	GoodsType        int64                                                         `json:"goodsType,omitempty"`
	SubGoodsType     int64                                                         `json:"subGoodsType,omitempty"`
	SoldType         int64                                                         `json:"soldType,omitempty"`
}
type PaasWeimobShopPayPaasAuditTradeOrderRequestGoodsAbilityReq struct {
	CycleInfoReq PaasWeimobShopPayPaasAuditTradeOrderRequestCycleInfoReq `json:"cycleInfoReq,omitempty"`
}
type PaasWeimobShopPayPaasAuditTradeOrderRequestCycleInfoReq struct {
	CycleBizId string `json:"cycleBizId,omitempty"`
}
type PaasWeimobShopPayPaasAuditTradeOrderRequestUseActivityList struct {
	ActivityId         string `json:"activityId,omitempty"`
	ActivityName       string `json:"activityName,omitempty"`
	ActivityType       int64  `json:"activityType,omitempty"`
	DiscountAmount     string `json:"discountAmount,omitempty"`
	DiscountBizOrderId string `json:"discountBizOrderId,omitempty"`
	GroupKey           string `json:"groupKey,omitempty"`
	Level              int64  `json:"level,omitempty"`
	TemplateId         int64  `json:"templateId,omitempty"`
}
type PaasWeimobShopPayPaasAuditTradeOrderRequestItemProductInfo struct {
	ProductDetailList []PaasWeimobShopPayPaasAuditTradeOrderRequestProductDetailList `json:"productDetailList,omitempty"`
}
type PaasWeimobShopPayPaasAuditTradeOrderRequestProductDetailList struct {
	WarehouseList []PaasWeimobShopPayPaasAuditTradeOrderRequestWarehouseList `json:"warehouseList,omitempty"`
}
type PaasWeimobShopPayPaasAuditTradeOrderRequestWarehouseList struct {
	WarehouseId   string `json:"warehouseId,omitempty"`
	WarehouseName string `json:"warehouseName,omitempty"`
	WarehouseType string `json:"warehouseType,omitempty"`
}
type PaasWeimobShopPayPaasAuditTradeOrderRequestGoodsExtMap struct {
}
type PaasWeimobShopPayPaasAuditTradeOrderRequestGoodsAbilityList struct {
	AbilityDetailInfo PaasWeimobShopPayPaasAuditTradeOrderRequestAbilityDetailInfo `json:"abilityDetailInfo,omitempty"`
	AbilityCode       string                                                       `json:"abilityCode,omitempty"`
	AbilityType       int64                                                        `json:"abilityType,omitempty"`
	BizId             string                                                       `json:"bizId,omitempty"`
}
type PaasWeimobShopPayPaasAuditTradeOrderRequestAbilityDetailInfo struct {
}
type PaasWeimobShopPayPaasAuditTradeOrderRequestPaymentInfo struct {
	DeductibleAssetList []PaasWeimobShopPayPaasAuditTradeOrderRequestDeductibleAssetList `json:"DeductibleAssetList,omitempty"`
	PaymentAmount       string                                                           `json:"paymentAmount,omitempty"`
	TotalAmount         string                                                           `json:"totalAmount,omitempty"`
	ShouldPaymentAmount string                                                           `json:"shouldPaymentAmount,omitempty"`
	TotalDiscountAmount string                                                           `json:"totalDiscountAmount,omitempty"`
}
type PaasWeimobShopPayPaasAuditTradeOrderRequestDeductibleAssetList struct {
	Type   int64  `json:"type,omitempty"`
	Amount string `json:"amount,omitempty"`
}

type PaasWeimobShopPayPaasAuditTradeOrderResponse struct {
	CommitInfo PaasWeimobShopPayPaasAuditTradeOrderResponseCommitInfo `json:"commitInfo,omitempty"`
}
type PaasWeimobShopPayPaasAuditTradeOrderResponseCommitInfo struct {
	CommitInfo     string `json:"commitInfo,omitempty"`
	UnCommitReason string `json:"unCommitReason	,omitempty"`
}

func CreatePaasWeimobShopPayPaasAuditTradeOrderResponse() *PaasWeimobShopPayPaasAuditTradeOrderResponse {
	return &PaasWeimobShopPayPaasAuditTradeOrderResponse{}
}

// OnPaasWeimobShopPayPaasAuditTradeOrderServiceInvocation
// @id 1169
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1169?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=订单审核)
func (s *Service) OnPaasWeimobShopPayPaasAuditTradeOrderServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobShopPayPaasAuditTradeOrderRequest, PaasWeimobShopPayPaasAuditTradeOrderResponse],
) (service *Service) {
	var invocation = &PaasWeimobShopPayPaasAuditTradeOrderService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobShopPayPaasAuditTradeOrderRequest, PaasWeimobShopPayPaasAuditTradeOrderResponse](invocation),
		"PaasWeimobShopPayPaasAuditTradeOrderService",
		"weimobShop.pay.paasAuditTradeOrder",
	)
	return s
}
