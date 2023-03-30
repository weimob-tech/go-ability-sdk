package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobShopTradeOrderAuditService struct {
	handler spi.WosInvocationHandler[PaasWeimobShopTradeOrderAuditRequest, PaasWeimobShopTradeOrderAuditResponse]
}

func (s PaasWeimobShopTradeOrderAuditService) NewRequest() spi.InvocationRequest[PaasWeimobShopTradeOrderAuditRequest] {
	return &spi.WosInvocationRequest[PaasWeimobShopTradeOrderAuditRequest]{
		Params: &PaasWeimobShopTradeOrderAuditRequest{},
	}
}

func (s PaasWeimobShopTradeOrderAuditService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobShopTradeOrderAuditRequest],
) (
	spi.InvocationResponse[PaasWeimobShopTradeOrderAuditResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobShopTradeOrderAuditRequest]))
}

type PaasWeimobShopTradeOrderAuditRequest struct {
	OrderList     []PaasWeimobShopTradeOrderAuditRequestOrderList `json:"orderList,omitempty"`
	BasicInfo     PaasWeimobShopTradeOrderAuditRequestBasicInfo   `json:"basicInfo,omitempty"`
	ExtendInfo    PaasWeimobShopTradeOrderAuditRequestExtendInfo  `json:"extendInfo,omitempty"`
	Wid           int64                                           `json:"wid,omitempty"`
	ParentOrderNo int64                                           `json:"parentOrderNo,omitempty"`
}
type PaasWeimobShopTradeOrderAuditRequestOrderList struct {
	GoodsList       []PaasWeimobShopTradeOrderAuditRequestGoodsList       `json:"goodsList,omitempty"`
	PaymentInfo     PaasWeimobShopTradeOrderAuditRequestPaymentInfo       `json:"paymentInfo,omitempty"`
	CustomFieldList []PaasWeimobShopTradeOrderAuditRequestCustomFieldList `json:"customFieldList,omitempty"`
	OrderNo         int64                                                 `json:"orderNo,omitempty"`
	Vid             int64                                                 `json:"vid,omitempty"`
	VidName         string                                                `json:"vidName,omitempty"`
}
type PaasWeimobShopTradeOrderAuditRequestGoodsList struct {
	GoodsAbilityList []PaasWeimobShopTradeOrderAuditRequestGoodsAbilityList `json:"goodsAbilityList,omitempty"`
	GoodsExtMap      PaasWeimobShopTradeOrderAuditRequestGoodsExtMap        `json:"goodsExtMap,omitempty"`
	ItemProductInfo  PaasWeimobShopTradeOrderAuditRequestItemProductInfo    `json:"itemProductInfo,omitempty"`
	UseActivityList  []PaasWeimobShopTradeOrderAuditRequestUseActivityList  `json:"useActivityList,omitempty"`
	ItemMessageList  []PaasWeimobShopTradeOrderAuditRequestItemMessageList  `json:"itemMessageList,omitempty"`
	BuyNum           int64                                                  `json:"buyNum,omitempty"`
	GoodsId          int64                                                  `json:"goodsId,omitempty"`
	GoodsType        int64                                                  `json:"goodsType,omitempty"`
	ItemId           int64                                                  `json:"itemId,omitempty"`
	SkuId            int64                                                  `json:"skuId,omitempty"`
	SkuSalePrice     string                                                 `json:"skuSalePrice,omitempty"`
	SoldType         int64                                                  `json:"soldType,omitempty"`
	SubGoodsType     int64                                                  `json:"subGoodsType,omitempty"`
	TradeGoodsType   int64                                                  `json:"tradeGoodsType,omitempty"`
	SkuCode          string                                                 `json:"skuCode,omitempty"`
	SkuMarketPrice   string                                                 `json:"skuMarketPrice,omitempty"`
}
type PaasWeimobShopTradeOrderAuditRequestGoodsAbilityList struct {
	AbilityDetailInfo PaasWeimobShopTradeOrderAuditRequestAbilityDetailInfo `json:"abilityDetailInfo,omitempty"`
	AbilityCode       string                                                `json:"abilityCode,omitempty"`
	AbilityType       int64                                                 `json:"abilityType,omitempty"`
	BizId             string                                                `json:"bizId,omitempty"`
}
type PaasWeimobShopTradeOrderAuditRequestAbilityDetailInfo struct {
}
type PaasWeimobShopTradeOrderAuditRequestGoodsExtMap struct {
}
type PaasWeimobShopTradeOrderAuditRequestItemProductInfo struct {
	ProductDetailList []PaasWeimobShopTradeOrderAuditRequestProductDetailList `json:"productDetailList,omitempty"`
}
type PaasWeimobShopTradeOrderAuditRequestProductDetailList struct {
	WarehouseList []PaasWeimobShopTradeOrderAuditRequestWarehouseList `json:"warehouseList,omitempty"`
}
type PaasWeimobShopTradeOrderAuditRequestWarehouseList struct {
	WarehouseId   int64  `json:"warehouseId,omitempty"`
	WarehouseName string `json:"warehouseName,omitempty"`
	WarehouseType string `json:"warehouseType,omitempty"`
}
type PaasWeimobShopTradeOrderAuditRequestUseActivityList struct {
	ActivityId         string `json:"activityId,omitempty"`
	ActivityName       string `json:"activityName,omitempty"`
	ActivityType       int64  `json:"activityType,omitempty"`
	DiscountAmount     int64  `json:"discountAmount,omitempty"`
	Level              int64  `json:"level,omitempty"`
	GroupKey           string `json:"groupKey,omitempty"`
	DiscountBizOrderId string `json:"discountBizOrderId,omitempty"`
}
type PaasWeimobShopTradeOrderAuditRequestItemMessageList struct {
	Key   string `json:"key,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
type PaasWeimobShopTradeOrderAuditRequestPaymentInfo struct {
	DeductibleAssets    []PaasWeimobShopTradeOrderAuditRequestDeductibleAssets `json:"deductibleAssets,omitempty"`
	PaymentAmount       string                                                 `json:"paymentAmount,omitempty"`
	ShouldPaymentAmount string                                                 `json:"shouldPaymentAmount,omitempty"`
	TotalAmount         string                                                 `json:"totalAmount,omitempty"`
	TotalDiscountAmount string                                                 `json:"totalDiscountAmount,omitempty"`
}
type PaasWeimobShopTradeOrderAuditRequestDeductibleAssets struct {
	Amount string `json:"amount,omitempty"`
	Type   int64  `json:"type,omitempty"`
}
type PaasWeimobShopTradeOrderAuditRequestCustomFieldList struct {
	Key   string `json:"key,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
type PaasWeimobShopTradeOrderAuditRequestBasicInfo struct {
	MerchantId              int64 `json:"merchantId,omitempty"`
	BosId                   int64 `json:"bosId,omitempty"`
	SourceProductId         int64 `json:"sourceProductId,omitempty"`
	SourceProductVersionId  int64 `json:"sourceProductVersionId,omitempty"`
	SourceProductInstanceId int64 `json:"sourceProductInstanceId,omitempty"`
	Vid                     int64 `json:"vid,omitempty"`
	VidType                 int64 `json:"vidType,omitempty"`
}
type PaasWeimobShopTradeOrderAuditRequestExtendInfo struct {
	Source       int64  `json:"source,omitempty"`
	SalesChannel int64  `json:"salesChannel,omitempty"`
	TradeTrackId string `json:"tradeTrackId,omitempty"`
}

type PaasWeimobShopTradeOrderAuditResponse struct {
	CommitInfo PaasWeimobShopTradeOrderAuditResponseCommitInfo `json:"commitInfo,omitempty"`
}
type PaasWeimobShopTradeOrderAuditResponseCommitInfo struct {
	AllowCommit    bool `json:"allowCommit,omitempty"`
	UnCommitReason bool `json:"unCommitReason,omitempty"`
}

func CreatePaasWeimobShopTradeOrderAuditResponse() *PaasWeimobShopTradeOrderAuditResponse {
	return &PaasWeimobShopTradeOrderAuditResponse{}
}

// OnPaasWeimobShopTradeOrderAuditServiceInvocation
// @id 1367
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1367?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=订单审核)
func (s *Service) OnPaasWeimobShopTradeOrderAuditServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobShopTradeOrderAuditRequest, PaasWeimobShopTradeOrderAuditResponse],
) (service *Service) {
	var invocation = &PaasWeimobShopTradeOrderAuditService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobShopTradeOrderAuditRequest, PaasWeimobShopTradeOrderAuditResponse](invocation),
		"PaasWeimobShopTradeOrderAuditService",
		"weimobShop.trade.order.audit",
	)
	return s
}
