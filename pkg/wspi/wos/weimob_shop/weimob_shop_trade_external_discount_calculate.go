package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobShopTradeExternalDiscountCalculateService struct {
	handler spi.WosInvocationHandler[PaasWeimobShopTradeExternalDiscountCalculateRequest, PaasWeimobShopTradeExternalDiscountCalculateResponse]
}

func (s PaasWeimobShopTradeExternalDiscountCalculateService) NewRequest() spi.InvocationRequest[PaasWeimobShopTradeExternalDiscountCalculateRequest] {
	return &spi.WosInvocationRequest[PaasWeimobShopTradeExternalDiscountCalculateRequest]{
		Params: &PaasWeimobShopTradeExternalDiscountCalculateRequest{},
	}
}

func (s PaasWeimobShopTradeExternalDiscountCalculateService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobShopTradeExternalDiscountCalculateRequest],
) (
	spi.InvocationResponse[PaasWeimobShopTradeExternalDiscountCalculateResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobShopTradeExternalDiscountCalculateRequest]))
}

type PaasWeimobShopTradeExternalDiscountCalculateRequest struct {
	ExtendInfo            PaasWeimobShopTradeExternalDiscountCalculateRequestExtendInfo              `json:"extendInfo,omitempty"`
	GoodsTradeMetas       []PaasWeimobShopTradeExternalDiscountCalculateRequestGoodsTradeMetas       `json:"goodsTradeMetas,omitempty"`
	InputUserDiscountList []PaasWeimobShopTradeExternalDiscountCalculateRequestInputUserDiscountList `json:"inputUserDiscountList,omitempty"`
	BizExt                PaasWeimobShopTradeExternalDiscountCalculateRequestBizExt                  `json:"bizExt,omitempty"`
	Wid                   int64                                                                      `json:"wid,omitempty"`
	BosId                 int64                                                                      `json:"bosId,omitempty"`
	ProductId             int64                                                                      `json:"productId,omitempty"`
	ProductInstanceId     int64                                                                      `json:"productInstanceId,omitempty"`
	Vid                   int64                                                                      `json:"vid,omitempty"`
	VidType               int64                                                                      `json:"vidType,omitempty"`
	WxTemplateId          int64                                                                      `json:"wxTemplateId,omitempty"`
}
type PaasWeimobShopTradeExternalDiscountCalculateRequestExtendInfo struct {
	SalesChannel int64  `json:"salesChannel,omitempty"`
	Source       int64  `json:"source,omitempty"`
	TradeTrackId string `json:"tradeTrackId,omitempty"`
	GlobalTicket string `json:"globalTicket,omitempty"`
	TradeScene   int64  `json:"tradeScene,omitempty"`
}
type PaasWeimobShopTradeExternalDiscountCalculateRequestGoodsTradeMetas struct {
	BasicGoodsInfo   PaasWeimobShopTradeExternalDiscountCalculateRequestBasicGoodsInfo     `json:"basicGoodsInfo,omitempty"`
	DiscountInfoList []PaasWeimobShopTradeExternalDiscountCalculateRequestDiscountInfoList `json:"discountInfoList,omitempty"`
	Mutex            PaasWeimobShopTradeExternalDiscountCalculateRequestMutex              `json:"mutex,omitempty"`
	TradeGoodsInfo   PaasWeimobShopTradeExternalDiscountCalculateRequestTradeGoodsInfo     `json:"tradeGoodsInfo,omitempty"`
}
type PaasWeimobShopTradeExternalDiscountCalculateRequestBasicGoodsInfo struct {
	PreSellInfo   PaasWeimobShopTradeExternalDiscountCalculateRequestPreSellInfo `json:"preSellInfo,omitempty"`
	PriceInfo     []PaasWeimobShopTradeExternalDiscountCalculateRequestPriceInfo `json:"priceInfo,omitempty"`
	GoodsId       int64                                                          `json:"goodsId,omitempty"`
	Num           int64                                                          `json:"num,omitempty"`
	SellModelType int64                                                          `json:"sellModelType,omitempty"`
	SkuId         int64                                                          `json:"skuId,omitempty"`
	Vid           int64                                                          `json:"vid,omitempty"`
	VidType       int64                                                          `json:"vidType,omitempty"`
}
type PaasWeimobShopTradeExternalDiscountCalculateRequestPreSellInfo struct {
	Type int64 `json:"type,omitempty"`
}
type PaasWeimobShopTradeExternalDiscountCalculateRequestPriceInfo struct {
	Price string `json:"price,omitempty"`
	Type  int64  `json:"type,omitempty"`
}
type PaasWeimobShopTradeExternalDiscountCalculateRequestDiscountInfoList struct {
	DiscountAmount string `json:"discountAmount,omitempty"`
	DiscountId     string `json:"discountId,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	DiscountLevel  int64  `json:"discountLevel,omitempty"`
	Order          int64  `json:"order,omitempty"`
	CheckAmount    string `json:"checkAmount,omitempty"`
	CostAmount     string `json:"costAmount,omitempty"`
}
type PaasWeimobShopTradeExternalDiscountCalculateRequestMutex struct {
	CustomDiscountType     int64   `json:"customDiscountType,omitempty"`
	CustomDiscountTypeList []int64 `json:"customDiscountTypeList,omitempty"`
}
type PaasWeimobShopTradeExternalDiscountCalculateRequestTradeGoodsInfo struct {
	ItemId         int64  `json:"itemId,omitempty"`
	TradeGoodsType int64  `json:"tradeGoodsType,omitempty"`
	TradePrice     string `json:"tradePrice,omitempty"`
	TradePriceType int64  `json:"tradePriceType,omitempty"`
}
type PaasWeimobShopTradeExternalDiscountCalculateRequestInputUserDiscountList struct {
	DiscountId     string `json:"discountId,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	DeductionType  int64  `json:"deductionType,omitempty"`
	Order          int64  `json:"order,omitempty"`
	DiscountStatus int64  `json:"discountStatus,omitempty"`
	CostAmount     string `json:"costAmount,omitempty"`
}
type PaasWeimobShopTradeExternalDiscountCalculateRequestBizExt struct {
	OuterOrderNo string `json:"outerOrderNo,omitempty"`
	BizSource    string `json:"bizSource,omitempty"`
}

type PaasWeimobShopTradeExternalDiscountCalculateResponse struct {
	GoodsTradeMetas []PaasWeimobShopTradeExternalDiscountCalculateResponseGoodsTradeMetas `json:"goodsTradeMetas,omitempty"`
	GlobalTicket    string                                                                `json:"globalTicket,omitempty"`
	OutTicket       string                                                                `json:"outTicket,omitempty"`
}
type PaasWeimobShopTradeExternalDiscountCalculateResponseGoodsTradeMetas struct {
	DiscountInfoList     []PaasWeimobShopTradeExternalDiscountCalculateResponseDiscountInfoList     `json:"discountInfoList,omitempty"`
	TradeGoodsInfo       PaasWeimobShopTradeExternalDiscountCalculateResponseTradeGoodsInfo         `json:"tradeGoodsInfo,omitempty"`
	FeeTypeDiscountInfos []PaasWeimobShopTradeExternalDiscountCalculateResponseFeeTypeDiscountInfos `json:"feeTypeDiscountInfos,omitempty"`
}
type PaasWeimobShopTradeExternalDiscountCalculateResponseDiscountInfoList struct {
	CheckAmount    string `json:"checkAmount,omitempty"`
	DeductionType  int64  `json:"deductionType,omitempty"`
	DiscountAmount string `json:"discountAmount,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	Order          int64  `json:"order,omitempty"`
	CostAmount     string `json:"costAmount,omitempty"`
	DiscountId     string `json:"discountId,omitempty"`
}
type PaasWeimobShopTradeExternalDiscountCalculateResponseTradeGoodsInfo struct {
	ItemId         int64  `json:"itemId,omitempty"`
	TradeGoodsType int64  `json:"tradeGoodsType,omitempty"`
	TradePrice     string `json:"tradePrice,omitempty"`
	TradePriceType int64  `json:"tradePriceType,omitempty"`
}
type PaasWeimobShopTradeExternalDiscountCalculateResponseFeeTypeDiscountInfos struct {
	Mutex            PaasWeimobShopTradeExternalDiscountCalculateResponseMutex               `json:"mutex,omitempty"`
	DiscountInfoList []PaasWeimobShopTradeExternalDiscountCalculateResponseDiscountInfoList2 `json:"discountInfoList,omitempty"`
	FeeType          int64                                                                   `json:"feeType,omitempty"`
}
type PaasWeimobShopTradeExternalDiscountCalculateResponseMutex struct {
	CustomDiscountType     int64   `json:"customDiscountType,omitempty"`
	CustomDiscountTypeList []int64 `json:"customDiscountTypeList,omitempty"`
}
type PaasWeimobShopTradeExternalDiscountCalculateResponseDiscountInfoList2 struct {
	DiscountAmount string `json:"discountAmount,omitempty"`
	DiscountId     string `json:"discountId,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	Order          int64  `json:"order,omitempty"`
	CheckAmount    string `json:"checkAmount,omitempty"`
	CostAmount     string `json:"costAmount,omitempty"`
}

func CreatePaasWeimobShopTradeExternalDiscountCalculateResponse() *PaasWeimobShopTradeExternalDiscountCalculateResponse {
	return &PaasWeimobShopTradeExternalDiscountCalculateResponse{}
}

// OnPaasWeimobShopTradeExternalDiscountCalculateServiceInvocation
// @id 1490
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1490?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=外部优惠PaaS批价接口)
func (s *Service) OnPaasWeimobShopTradeExternalDiscountCalculateServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobShopTradeExternalDiscountCalculateRequest, PaasWeimobShopTradeExternalDiscountCalculateResponse],
) (service *Service) {
	var invocation = &PaasWeimobShopTradeExternalDiscountCalculateService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobShopTradeExternalDiscountCalculateRequest, PaasWeimobShopTradeExternalDiscountCalculateResponse](invocation),
		"PaasWeimobShopTradeExternalDiscountCalculateService",
		"weimobShop.trade.external.discount.calculate",
	)
	return s
}
