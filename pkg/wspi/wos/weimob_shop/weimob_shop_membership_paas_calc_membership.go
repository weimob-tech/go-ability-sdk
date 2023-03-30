package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobShopMembershipPaasCalcMembershipService struct {
	handler spi.WosInvocationHandler[PaasWeimobShopMembershipPaasCalcMembershipRequest, PaasWeimobShopMembershipPaasCalcMembershipResponse]
}

func (s PaasWeimobShopMembershipPaasCalcMembershipService) NewRequest() spi.InvocationRequest[PaasWeimobShopMembershipPaasCalcMembershipRequest] {
	return &spi.WosInvocationRequest[PaasWeimobShopMembershipPaasCalcMembershipRequest]{
		Params: &PaasWeimobShopMembershipPaasCalcMembershipRequest{},
	}
}

func (s PaasWeimobShopMembershipPaasCalcMembershipService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobShopMembershipPaasCalcMembershipRequest],
) (
	spi.InvocationResponse[PaasWeimobShopMembershipPaasCalcMembershipResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobShopMembershipPaasCalcMembershipRequest]))
}

type PaasWeimobShopMembershipPaasCalcMembershipRequest struct {
	BasicInfo         PaasWeimobShopMembershipPaasCalcMembershipRequestBasicInfo         `json:"basicInfo,omitempty"`
	ExtendInfo        PaasWeimobShopMembershipPaasCalcMembershipRequestExtendInfo        `json:"extendInfo,omitempty"`
	GoodsTradeMetas   []PaasWeimobShopMembershipPaasCalcMembershipRequestGoodsTradeMetas `json:"goodsTradeMetas,omitempty"`
	ProductId         int64                                                              `json:"productId,omitempty"`
	ProductInstanceId int64                                                              `json:"productInstanceId,omitempty"`
	Tcode             string                                                             `json:"tcode,omitempty"`
	Wid               int64                                                              `json:"wid,omitempty"`
	WxTemplateId      int64                                                              `json:"wxTemplateId,omitempty"`
}
type PaasWeimobShopMembershipPaasCalcMembershipRequestBasicInfo struct {
	BosId             int64  `json:"bosId,omitempty"`
	ProductId         int64  `json:"productId,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	ProductVersionId  string `json:"productVersionId,omitempty"`
	Tcode             string `json:"tcode,omitempty"`
	Vid               int64  `json:"vid,omitempty"`
	VidType           int64  `json:"vidType,omitempty"`
}
type PaasWeimobShopMembershipPaasCalcMembershipRequestExtendInfo struct {
	GlobalTicket string `json:"globalTicket,omitempty"`
	SalesChannel int64  `json:"salesChannel,omitempty"`
	Source       int64  `json:"source,omitempty"`
}
type PaasWeimobShopMembershipPaasCalcMembershipRequestGoodsTradeMetas struct {
	BasicGoodsInfo        PaasWeimobShopMembershipPaasCalcMembershipRequestBasicGoodsInfo          `json:"basicGoodsInfo,omitempty"`
	DiscountInfoList      []PaasWeimobShopMembershipPaasCalcMembershipRequestDiscountInfoList      `json:"discountInfoList,omitempty"`
	InputItemDiscountList []PaasWeimobShopMembershipPaasCalcMembershipRequestInputItemDiscountList `json:"inputItemDiscountList,omitempty"`
	Mutex                 PaasWeimobShopMembershipPaasCalcMembershipRequestMutex                   `json:"mutex,omitempty"`
	PostageDiscountInfo   PaasWeimobShopMembershipPaasCalcMembershipRequestPostageDiscountInfo     `json:"postageDiscountInfo,omitempty"`
	TradeGoodsInfo        PaasWeimobShopMembershipPaasCalcMembershipRequestTradeGoodsInfo          `json:"tradeGoodsInfo,omitempty"`
}
type PaasWeimobShopMembershipPaasCalcMembershipRequestBasicGoodsInfo struct {
	PreSellInfo   PaasWeimobShopMembershipPaasCalcMembershipRequestPreSellInfo `json:"preSellInfo,omitempty"`
	PriceInfo     []PaasWeimobShopMembershipPaasCalcMembershipRequestPriceInfo `json:"priceInfo,omitempty"`
	GoodsCode     string                                                       `json:"goodsCode,omitempty"`
	GoodsId       int64                                                        `json:"goodsId,omitempty"`
	Num           int64                                                        `json:"num,omitempty"`
	SellModelType int64                                                        `json:"sellModelType,omitempty"`
	SkuId         int64                                                        `json:"skuId,omitempty"`
	Vid           int64                                                        `json:"vid,omitempty"`
	VidType       int64                                                        `json:"vidType,omitempty"`
	SkuCode       string                                                       `json:"skuCode,omitempty"`
}
type PaasWeimobShopMembershipPaasCalcMembershipRequestPreSellInfo struct {
	Type int64 `json:"type,omitempty"`
}
type PaasWeimobShopMembershipPaasCalcMembershipRequestPriceInfo struct {
	Price string `json:"price,omitempty"`
	Type  int64  `json:"type,omitempty"`
}
type PaasWeimobShopMembershipPaasCalcMembershipRequestDiscountInfoList struct {
	CheckAmount    string `json:"checkAmount,omitempty"`
	DiscountAmount string `json:"discountAmount,omitempty"`
	DiscountId     int64  `json:"discountId,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	DiscountLevel  int64  `json:"discountLevel,omitempty"`
	DiscountStatus int64  `json:"discountStatus,omitempty"`
	Order          int64  `json:"order,omitempty"`
	CostAmount     string `json:"costAmount,omitempty"`
}
type PaasWeimobShopMembershipPaasCalcMembershipRequestInputItemDiscountList struct {
	RuleGroup      PaasWeimobShopMembershipPaasCalcMembershipRequestRuleGroup `json:"ruleGroup,omitempty"`
	CostAmount     string                                                     `json:"costAmount,omitempty"`
	DiscountId     string                                                     `json:"discountId,omitempty"`
	DiscountStatus int64                                                      `json:"discountStatus,omitempty"`
	DiscountType   int64                                                      `json:"discountType,omitempty"`
	Order          int64                                                      `json:"order,omitempty"`
	DiscountLevel  int64                                                      `json:"discountLevel,omitempty"`
}
type PaasWeimobShopMembershipPaasCalcMembershipRequestRuleGroup struct {
	Superposable   PaasWeimobShopMembershipPaasCalcMembershipRequestSuperposable `json:"superposable,omitempty"`
	ConditionType  int64                                                         `json:"conditionType,omitempty"`
	ConditionValue string                                                        `json:"conditionValue,omitempty"`
	ResultType     int64                                                         `json:"resultType,omitempty"`
	ResultValue    string                                                        `json:"resultValue,omitempty"`
}
type PaasWeimobShopMembershipPaasCalcMembershipRequestSuperposable struct {
	SuperposableCouponNum int64 `json:"superposableCouponNum,omitempty"`
	Superposable          bool  `json:"superposable,omitempty"`
	IdenticalSuperposable bool  `json:"identicalSuperposable,omitempty"`
	DifferentSuperposable bool  `json:"differentSuperposable,omitempty"`
}
type PaasWeimobShopMembershipPaasCalcMembershipRequestMutex struct {
	CustomDiscountType     int64   `json:"customDiscountType,omitempty"`
	CustomDiscountTypeList []int64 `json:"customDiscountTypeList,omitempty"`
}
type PaasWeimobShopMembershipPaasCalcMembershipRequestPostageDiscountInfo struct {
	DiscountInfoList []PaasWeimobShopMembershipPaasCalcMembershipRequestDiscountInfoList2 `json:"discountInfoList,omitempty"`
	Mutex            PaasWeimobShopMembershipPaasCalcMembershipRequestMutex2              `json:"mutex,omitempty"`
}
type PaasWeimobShopMembershipPaasCalcMembershipRequestDiscountInfoList2 struct {
}
type PaasWeimobShopMembershipPaasCalcMembershipRequestMutex2 struct {
	CustomDiscountType     int64   `json:"customDiscountType,omitempty"`
	CustomDiscountTypeList []int64 `json:"customDiscountTypeList,omitempty"`
}
type PaasWeimobShopMembershipPaasCalcMembershipRequestTradeGoodsInfo struct {
	ItemId         int64  `json:"itemId,omitempty"`
	TradeGoodsType int64  `json:"tradeGoodsType,omitempty"`
	TradePrice     string `json:"tradePrice,omitempty"`
	TradePriceType int64  `json:"tradePriceType,omitempty"`
}

type PaasWeimobShopMembershipPaasCalcMembershipResponse struct {
	GoodsTradeMetas PaasWeimobShopMembershipPaasCalcMembershipResponseGoodsTradeMetas `json:"goodsTradeMetas,omitempty"`
	CalcRuleList    []PaasWeimobShopMembershipPaasCalcMembershipResponseCalcRuleList  `json:"calcRuleList,omitempty"`
	GlobalTicket    string                                                            `json:"globalTicket,omitempty"`
	OutTicket       string                                                            `json:"outTicket,omitempty"`
}
type PaasWeimobShopMembershipPaasCalcMembershipResponseGoodsTradeMetas struct {
	TradeGoodsInfo      PaasWeimobShopMembershipPaasCalcMembershipResponseTradeGoodsInfo      `json:"tradeGoodsInfo,omitempty"`
	DiscountInfoList    []PaasWeimobShopMembershipPaasCalcMembershipResponseDiscountInfoList  `json:"discountInfoList,omitempty"`
	PostageDiscountInfo PaasWeimobShopMembershipPaasCalcMembershipResponsePostageDiscountInfo `json:"postageDiscountInfo,omitempty"`
}
type PaasWeimobShopMembershipPaasCalcMembershipResponseTradeGoodsInfo struct {
	ItemId         int64  `json:"itemId,omitempty"`
	TradeGoodsType int64  `json:"tradeGoodsType,omitempty"`
	TradePrice     string `json:"tradePrice,omitempty"`
	TradePriceType int64  `json:"tradePriceType,omitempty"`
}
type PaasWeimobShopMembershipPaasCalcMembershipResponseDiscountInfoList struct {
	CheckAmount    string `json:"checkAmount,omitempty"`
	DiscountAmount string `json:"discountAmount,omitempty"`
	DiscountId     string `json:"discountId,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	DiscountLevel  int64  `json:"discountLevel,omitempty"`
	DiscountStatus int64  `json:"discountStatus,omitempty"`
	Order          int64  `json:"order,omitempty"`
	CostAmount     string `json:"costAmount,omitempty"`
}
type PaasWeimobShopMembershipPaasCalcMembershipResponsePostageDiscountInfo struct {
	Mutex            PaasWeimobShopMembershipPaasCalcMembershipResponseMutex               `json:"mutex,omitempty"`
	DiscountInfoList []PaasWeimobShopMembershipPaasCalcMembershipResponseDiscountInfoList2 `json:"discountInfoList,omitempty"`
}
type PaasWeimobShopMembershipPaasCalcMembershipResponseMutex struct {
	CustomDiscountType     string `json:"customDiscountType,omitempty"`
	CustomDiscountTypeList string `json:"customDiscountTypeList,omitempty"`
}
type PaasWeimobShopMembershipPaasCalcMembershipResponseDiscountInfoList2 struct {
	CheckAmount    string `json:"checkAmount	,omitempty"`
	DiscountAmount string `json:"discountAmount,omitempty"`
	DiscountId     string `json:"discountId,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	DiscountLevel  int64  `json:"discountLevel,omitempty"`
	DiscountStatus int64  `json:"discountStatus,omitempty"`
	Order          int64  `json:"order,omitempty"`
	CostAmount     string `json:"costAmount,omitempty"`
}
type PaasWeimobShopMembershipPaasCalcMembershipResponseCalcRuleList struct {
	Mutex           PaasWeimobShopMembershipPaasCalcMembershipResponseMutex2          `json:"mutex,omitempty"`
	DiscountUseInfo PaasWeimobShopMembershipPaasCalcMembershipResponseDiscountUseInfo `json:"discountUseInfo,omitempty"`
	DiscountId      string                                                            `json:"discountId,omitempty"`
	DiscountType    string                                                            `json:"discountType,omitempty"`
}
type PaasWeimobShopMembershipPaasCalcMembershipResponseMutex2 struct {
	CustomDiscountType     int64   `json:"customDiscountType,omitempty"`
	CustomDiscountTypeList []int64 `json:"customDiscountTypeList,omitempty"`
}
type PaasWeimobShopMembershipPaasCalcMembershipResponseDiscountUseInfo struct {
	UsedDiscountDetailList     []PaasWeimobShopMembershipPaasCalcMembershipResponseUsedDiscountDetailList `json:"usedDiscountDetailList,omitempty"`
	MaxDiscountTotalAmount     string                                                                     `json:"maxDiscountTotalAmount,omitempty"`
	MaxDiscountTotalCostAmount string                                                                     `json:"maxDiscountTotalCostAmount,omitempty"`
	UseDiscountTotalAmount     string                                                                     `json:"useDiscountTotalAmount,omitempty"`
	UseDiscountTotalCostAmount string                                                                     `json:"useDiscountTotalCostAmount,omitempty"`
	UseStatus                  bool                                                                       `json:"useStatus,omitempty"`
	GreyStatus                 bool                                                                       `json:"greyStatus,omitempty"`
	Toast                      string                                                                     `json:"toast,omitempty"`
}
type PaasWeimobShopMembershipPaasCalcMembershipResponseUsedDiscountDetailList struct {
	DiscountAmount string `json:"discountAmount,omitempty"`
	CostAmount     string `json:"costAmount,omitempty"`
	DeductionType  string `json:"deductionType,omitempty"`
}

func CreatePaasWeimobShopMembershipPaasCalcMembershipResponse() *PaasWeimobShopMembershipPaasCalcMembershipResponse {
	return &PaasWeimobShopMembershipPaasCalcMembershipResponse{}
}

// OnPaasWeimobShopMembershipPaasCalcMembershipServiceInvocation
// @id 1167
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1167?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=会员批价能力)
func (s *Service) OnPaasWeimobShopMembershipPaasCalcMembershipServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobShopMembershipPaasCalcMembershipRequest, PaasWeimobShopMembershipPaasCalcMembershipResponse],
) (service *Service) {
	var invocation = &PaasWeimobShopMembershipPaasCalcMembershipService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobShopMembershipPaasCalcMembershipRequest, PaasWeimobShopMembershipPaasCalcMembershipResponse](invocation),
		"PaasWeimobShopMembershipPaasCalcMembershipService",
		"weimobShop.membership.paasCalcMembership",
	)
	return s
}
