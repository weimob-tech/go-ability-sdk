package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobShopCouponPaasCalcCouponService struct {
	handler spi.WosInvocationHandler[PaasWeimobShopCouponPaasCalcCouponRequest, PaasWeimobShopCouponPaasCalcCouponResponse]
}

func (s PaasWeimobShopCouponPaasCalcCouponService) NewRequest() spi.InvocationRequest[PaasWeimobShopCouponPaasCalcCouponRequest] {
	return &spi.WosInvocationRequest[PaasWeimobShopCouponPaasCalcCouponRequest]{
		Params: &PaasWeimobShopCouponPaasCalcCouponRequest{},
	}
}

func (s PaasWeimobShopCouponPaasCalcCouponService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobShopCouponPaasCalcCouponRequest],
) (
	spi.InvocationResponse[PaasWeimobShopCouponPaasCalcCouponResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobShopCouponPaasCalcCouponRequest]))
}

type PaasWeimobShopCouponPaasCalcCouponRequest struct {
	BasicInfo             PaasWeimobShopCouponPaasCalcCouponRequestBasicInfo               `json:"basicInfo,omitempty"`
	ExtendInfo            PaasWeimobShopCouponPaasCalcCouponRequestExtendInfo              `json:"extendInfo,omitempty"`
	GoodsTradeMetas       []PaasWeimobShopCouponPaasCalcCouponRequestGoodsTradeMetas       `json:"goodsTradeMetas,omitempty"`
	InputUserDiscountList []PaasWeimobShopCouponPaasCalcCouponRequestInputUserDiscountList `json:"inputUserDiscountList,omitempty"`
	BosId                 int64                                                            `json:"bosId,omitempty"`
	ProductId             int64                                                            `json:"productId,omitempty"`
	ProductInstanceId     int64                                                            `json:"productInstanceId,omitempty"`
	Tcode                 string                                                           `json:"tcode,omitempty"`
	Wid                   int64                                                            `json:"wid,omitempty"`
	Vid                   int64                                                            `json:"vid,omitempty"`
	VidType               int64                                                            `json:"vidType,omitempty"`
}
type PaasWeimobShopCouponPaasCalcCouponRequestBasicInfo struct {
	BosId             int64 `json:"bosId,omitempty"`
	ProductId         int64 `json:"productId,omitempty"`
	ProductInstanceId int64 `json:"productInstanceId,omitempty"`
	ProductVersionId  int64 `json:"productVersionId,omitempty"`
	Vid               int64 `json:"vid,omitempty"`
	VidType           int64 `json:"vidType,omitempty"`
}
type PaasWeimobShopCouponPaasCalcCouponRequestExtendInfo struct {
	GlobalTicket string `json:"globalTicket,omitempty"`
	SalesChannel int64  `json:"salesChannel,omitempty"`
	Source       int64  `json:"source,omitempty"`
}
type PaasWeimobShopCouponPaasCalcCouponRequestGoodsTradeMetas struct {
	BasicGoodsInfo        PaasWeimobShopCouponPaasCalcCouponRequestBasicGoodsInfo          `json:"basicGoodsInfo,omitempty"`
	DiscountInfoList      []PaasWeimobShopCouponPaasCalcCouponRequestDiscountInfoList      `json:"discountInfoList,omitempty"`
	InputItemDiscountList []PaasWeimobShopCouponPaasCalcCouponRequestInputItemDiscountList `json:"inputItemDiscountList,omitempty"`
	Mutex                 PaasWeimobShopCouponPaasCalcCouponRequestMutex                   `json:"mutex,omitempty"`
	TradeGoodsInfo        PaasWeimobShopCouponPaasCalcCouponRequestTradeGoodsInfo          `json:"tradeGoodsInfo,omitempty"`
}
type PaasWeimobShopCouponPaasCalcCouponRequestBasicGoodsInfo struct {
	PriceInfo     []PaasWeimobShopCouponPaasCalcCouponRequestPriceInfo `json:"priceInfo,omitempty"`
	GoodsCode     string                                               `json:"goodsCode,omitempty"`
	GoodsId       int64                                                `json:"goodsId,omitempty"`
	Num           int64                                                `json:"num,omitempty"`
	SellModelType int64                                                `json:"sellModelType,omitempty"`
	SkuCode       string                                               `json:"skuCode,omitempty"`
	SkuId         int64                                                `json:"skuId,omitempty"`
	Vid           int64                                                `json:"vid,omitempty"`
	VidType       int64                                                `json:"vidType,omitempty"`
}
type PaasWeimobShopCouponPaasCalcCouponRequestPriceInfo struct {
	Price string `json:"price,omitempty"`
	Type  int64  `json:"type,omitempty"`
}
type PaasWeimobShopCouponPaasCalcCouponRequestDiscountInfoList struct {
	DiscountId     string `json:"discountId,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	DiscountLevel  int64  `json:"discountLevel,omitempty"`
	Order          int64  `json:"order,omitempty"`
	CheckAmount    string `json:"checkAmount,omitempty"`
	DiscountAmount string `json:"discountAmount,omitempty"`
}
type PaasWeimobShopCouponPaasCalcCouponRequestInputItemDiscountList struct {
	DiscountId     string `json:"discountId,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	DiscountStatus int64  `json:"discountStatus,omitempty"`
	DiscountLevel  int64  `json:"discountLevel,omitempty"`
}
type PaasWeimobShopCouponPaasCalcCouponRequestMutex struct {
	CustomDiscountType     int64   `json:"customDiscountType,omitempty"`
	CustomDiscountTypeList []int64 `json:"customDiscountTypeList,omitempty"`
}
type PaasWeimobShopCouponPaasCalcCouponRequestTradeGoodsInfo struct {
	ItemId         int64  `json:"itemId,omitempty"`
	TradeGoodsType int64  `json:"tradeGoodsType,omitempty"`
	TradePrice     string `json:"tradePrice,omitempty"`
	TradePriceType int64  `json:"tradePriceType,omitempty"`
}
type PaasWeimobShopCouponPaasCalcCouponRequestInputUserDiscountList struct {
	DiscountId       string `json:"discountId,omitempty"`
	DiscountType     int64  `json:"discountType,omitempty"`
	Order            int64  `json:"order,omitempty"`
	DiscountStatus   int64  `json:"discountStatus,omitempty"`
	CouponTemplateId string `json:"couponTemplateId,omitempty"`
	Name             string `json:"name,omitempty"`
}

type PaasWeimobShopCouponPaasCalcCouponResponse struct {
	CalcRuleList    []PaasWeimobShopCouponPaasCalcCouponResponseCalcRuleList    `json:"calcRuleList,omitempty"`
	GoodsTradeMetas []PaasWeimobShopCouponPaasCalcCouponResponseGoodsTradeMetas `json:"goodsTradeMetas,omitempty"`
	GlobalTicket    string                                                      `json:"globalTicket,omitempty"`
	OutTicket       string                                                      `json:"outTicket,omitempty"`
}
type PaasWeimobShopCouponPaasCalcCouponResponseCalcRuleList struct {
	Mutex              PaasWeimobShopCouponPaasCalcCouponResponseMutex     `json:"mutex,omitempty"`
	RuleGroup          PaasWeimobShopCouponPaasCalcCouponResponseRuleGroup `json:"ruleGroup,omitempty"`
	UseInfo            PaasWeimobShopCouponPaasCalcCouponResponseUseInfo   `json:"useInfo,omitempty"`
	UseStartTime       string                                              `json:"useStartTime,omitempty"`
	CalculateType      int64                                               `json:"calculateType,omitempty"`
	DiscountAmount     string                                              `json:"discountAmount,omitempty"`
	UseEndTime         string                                              `json:"useEndTime,omitempty"`
	Vid                int64                                               `json:"vid,omitempty"`
	SubName            string                                              `json:"subName,omitempty"`
	CouponType         int64                                               `json:"couponType,omitempty"`
	InstructionContent string                                              `json:"instructionContent,omitempty"`
	Name               string                                              `json:"name,omitempty"`
	DiscountType       int64                                               `json:"discountType,omitempty"`
	ConsumerCode       string                                              `json:"consumerCode,omitempty"`
	DiscountId         string                                              `json:"discountId,omitempty"`
	VidType            int64                                               `json:"vidType,omitempty"`
}
type PaasWeimobShopCouponPaasCalcCouponResponseMutex struct {
	CustomDiscountTypeList []int64 `json:"customDiscountTypeList,omitempty"`
	CustomDiscountType     int64   `json:"customDiscountType,omitempty"`
}
type PaasWeimobShopCouponPaasCalcCouponResponseRuleGroup struct {
	Superposable   PaasWeimobShopCouponPaasCalcCouponResponseSuperposable `json:"superposable,omitempty"`
	ResultValue    string                                                 `json:"resultValue,omitempty"`
	ConditionValue int64                                                  `json:"conditionValue,omitempty"`
	ConditionType  int64                                                  `json:"conditionType,omitempty"`
	ResultType     int64                                                  `json:"resultType,omitempty"`
}
type PaasWeimobShopCouponPaasCalcCouponResponseSuperposable struct {
	SuperposableCouponNum int64 `json:"superposableCouponNum,omitempty"`
}
type PaasWeimobShopCouponPaasCalcCouponResponseUseInfo struct {
	Valid          int64  `json:"valid,omitempty"`
	DisabledReason string `json:"disabledReason,omitempty"`
	Use            int64  `json:"use,omitempty"`
	Enable         int64  `json:"enable,omitempty"`
	Order          int64  `json:"order,omitempty"`
}
type PaasWeimobShopCouponPaasCalcCouponResponseGoodsTradeMetas struct {
	TradeGoodsInfo   PaasWeimobShopCouponPaasCalcCouponResponseTradeGoodsInfo     `json:"tradeGoodsInfo,omitempty"`
	DiscountInfoList []PaasWeimobShopCouponPaasCalcCouponResponseDiscountInfoList `json:"discountInfoList,omitempty"`
	Mutex            PaasWeimobShopCouponPaasCalcCouponResponseMutex2             `json:"mutex,omitempty"`
}
type PaasWeimobShopCouponPaasCalcCouponResponseTradeGoodsInfo struct {
	ItemId         int64  `json:"itemId,omitempty"`
	TradeGoodsType int64  `json:"tradeGoodsType,omitempty"`
	TradePriceType int64  `json:"tradePriceType,omitempty"`
	TradePrice     string `json:"tradePrice,omitempty"`
}
type PaasWeimobShopCouponPaasCalcCouponResponseDiscountInfoList struct {
	CheckAmount    string `json:"checkAmount,omitempty"`
	DiscountAmount string `json:"discountAmount,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	DiscountId     string `json:"discountId,omitempty"`
	Order          int64  `json:"order,omitempty"`
}
type PaasWeimobShopCouponPaasCalcCouponResponseMutex2 struct {
	CustomDiscountTypeList []int64 `json:"customDiscountTypeList,omitempty"`
	CustomDiscountType     int64   `json:"customDiscountType,omitempty"`
}

func CreatePaasWeimobShopCouponPaasCalcCouponResponse() *PaasWeimobShopCouponPaasCalcCouponResponse {
	return &PaasWeimobShopCouponPaasCalcCouponResponse{}
}

// OnPaasWeimobShopCouponPaasCalcCouponServiceInvocation
// @id 581
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/581?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=券批价)
func (s *Service) OnPaasWeimobShopCouponPaasCalcCouponServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobShopCouponPaasCalcCouponRequest, PaasWeimobShopCouponPaasCalcCouponResponse],
) (service *Service) {
	var invocation = &PaasWeimobShopCouponPaasCalcCouponService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobShopCouponPaasCalcCouponRequest, PaasWeimobShopCouponPaasCalcCouponResponse](invocation),
		"PaasWeimobShopCouponPaasCalcCouponService",
		"weimobShop.coupon.paasCalcCoupon",
	)
	return s
}
