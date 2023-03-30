package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCouponTemplateDetailGetService struct {
	handler spi.WosInvocationHandler[PaasWeimobCouponTemplateDetailGetRequest, PaasWeimobCouponTemplateDetailGetResponse]
}

func (s PaasWeimobCouponTemplateDetailGetService) NewRequest() spi.InvocationRequest[PaasWeimobCouponTemplateDetailGetRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCouponTemplateDetailGetRequest]{
		Params: &PaasWeimobCouponTemplateDetailGetRequest{},
	}
}

func (s PaasWeimobCouponTemplateDetailGetService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCouponTemplateDetailGetRequest],
) (
	spi.InvocationResponse[PaasWeimobCouponTemplateDetailGetResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCouponTemplateDetailGetRequest]))
}

type PaasWeimobCouponTemplateDetailGetRequest struct {
	CouponTemplateId int64 `json:"couponTemplateId,omitempty"`
	Status           int64 `json:"status,omitempty"`
}

type PaasWeimobCouponTemplateDetailGetResponse struct {
	BaseInfo     PaasWeimobCouponTemplateDetailGetResponseBaseInfo     `json:"baseInfo,omitempty"`
	CreateInfo   PaasWeimobCouponTemplateDetailGetResponseCreateInfo   `json:"createInfo,omitempty"`
	OtherSetting PaasWeimobCouponTemplateDetailGetResponseOtherSetting `json:"otherSetting,omitempty"`
	ReduceRule   PaasWeimobCouponTemplateDetailGetResponseReduceRule   `json:"reduceRule,omitempty"`
	SendRule     PaasWeimobCouponTemplateDetailGetResponseSendRule     `json:"sendRule,omitempty"`
	Setting      PaasWeimobCouponTemplateDetailGetResponseSetting      `json:"setting,omitempty"`
	UseRule      PaasWeimobCouponTemplateDetailGetResponseUseRule      `json:"useRule,omitempty"`
	Status       int64                                                 `json:"status,omitempty"`
}
type PaasWeimobCouponTemplateDetailGetResponseBaseInfo struct {
	BelongVid        int64  `json:"belongVid,omitempty"`
	BelongVidName    string `json:"belongVidName,omitempty"`
	CouponDesc       string `json:"couponDesc,omitempty"`
	CouponSubType    int64  `json:"couponSubType,omitempty"`
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	CouponType       int64  `json:"couponType,omitempty"`
	MerchantName     string `json:"merchantName,omitempty"`
	Name             string `json:"name,omitempty"`
	PlatformType     int64  `json:"platformType,omitempty"`
	SubName          string `json:"subName,omitempty"`
}
type PaasWeimobCouponTemplateDetailGetResponseCreateInfo struct {
	CreateTime string `json:"createTime,omitempty"`
	OperatorId int64  `json:"operatorId,omitempty"`
	Source     string `json:"source,omitempty"`
}
type PaasWeimobCouponTemplateDetailGetResponseOtherSetting struct {
	Explain      string  `json:"explain,omitempty"`
	HasMemberTag bool    `json:"hasMemberTag,omitempty"`
	MemberTagIds []int64 `json:"memberTagIds,omitempty"`
}
type PaasWeimobCouponTemplateDetailGetResponseReduceRule struct {
	CashTicketAmt      int64 `json:"cashTicketAmt,omitempty"`
	CostOfOrganization int64 `json:"costOfOrganization,omitempty"`
	CostOfStore        int64 `json:"costOfStore,omitempty"`
	CouponCost         int64 `json:"couponCost,omitempty"`
	CouponCostType     int64 `json:"couponCostType,omitempty"`
	Discount           int64 `json:"discount,omitempty"`
	DiscountLimitType  int64 `json:"discountLimitType,omitempty"`
	DiscountLimitValue int64 `json:"discountLimitValue,omitempty"`
	ExchangeGoodsIds   int64 `json:"exchangeGoodsIds,omitempty"`
	ExchangeGoodsNum   int64 `json:"exchangeGoodsNum,omitempty"`
	ExchangeGoodsType  int64 `json:"exchangeGoodsType,omitempty"`
	FreightReduceType  int64 `json:"freightReduceType,omitempty"`
	MaxCashTicketAmt   int64 `json:"maxCashTicketAmt,omitempty"`
	MaxDiscountNum     int64 `json:"maxDiscountNum,omitempty"`
	MaxReducePrice     int64 `json:"maxReducePrice,omitempty"`
	MinCashTicketAmt   int64 `json:"minCashTicketAmt,omitempty"`
	PricingType        int64 `json:"pricingType,omitempty"`
	ReducePriceType    int64 `json:"reducePriceType,omitempty"`
}
type PaasWeimobCouponTemplateDetailGetResponseSendRule struct {
	Crowd                 PaasWeimobCouponTemplateDetailGetResponseCrowd `json:"crowd,omitempty"`
	ActivityPublish       bool                                           `json:"activityPublish,omitempty"`
	CanGift               bool                                           `json:"canGift,omitempty"`
	CanShare              bool                                           `json:"canShare,omitempty"`
	CanStoreLaunch        bool                                           `json:"canStoreLaunch,omitempty"`
	CustomerDirectReceive bool                                           `json:"customerDirectReceive,omitempty"`
	CustomerListPublish   bool                                           `json:"customerListPublish,omitempty"`
	IsAcceptAllCrowd      bool                                           `json:"isAcceptAllCrowd,omitempty"`
	IsPayment             bool                                           `json:"isPayment,omitempty"`
	MerchantDelivery      bool                                           `json:"merchantDelivery,omitempty"`
	MerchantIssueChannels []int64                                        `json:"merchantIssueChannels,omitempty"`
	MerchantPublish       bool                                           `json:"merchantPublish,omitempty"`
	RecommendEndTime      string                                         `json:"recommendEndTime,omitempty"`
	RecommendStartTime    string                                         `json:"recommendStartTime,omitempty"`
	SendChannels          []int64                                        `json:"sendChannels,omitempty"`
	SendEndDate           string                                         `json:"sendEndDate,omitempty"`
	SendStartDate         string                                         `json:"sendStartDate,omitempty"`
	SendTimeType          int64                                          `json:"sendTimeType,omitempty"`
	ServicePublish        bool                                           `json:"servicePublish,omitempty"`
	UserTakeLimit         int64                                          `json:"userTakeLimit,omitempty"`
}
type PaasWeimobCouponTemplateDetailGetResponseCrowd struct {
	RuleContent string `json:"ruleContent,omitempty"`
	RuleId      int64  `json:"ruleId,omitempty"`
}
type PaasWeimobCouponTemplateDetailGetResponseSetting struct {
	CustomCodeSetting PaasWeimobCouponTemplateDetailGetResponseCustomCodeSetting `json:"customCodeSetting,omitempty"`
	CodeGenerateType  int64                                                      `json:"codeGenerateType,omitempty"`
	Stock             int64                                                      `json:"stock,omitempty"`
}
type PaasWeimobCouponTemplateDetailGetResponseCustomCodeSetting struct {
}
type PaasWeimobCouponTemplateDetailGetResponseUseRule struct {
	AcceptTypeDTO       PaasWeimobCouponTemplateDetailGetResponseAcceptTypeDTO     `json:"acceptTypeDTO,omitempty"`
	SuperpositionRule   PaasWeimobCouponTemplateDetailGetResponseSuperpositionRule `json:"superpositionRule,omitempty"`
	UseValidDate        PaasWeimobCouponTemplateDetailGetResponseUseValidDate      `json:"useValidDate,omitempty"`
	CanCashTicket       bool                                                       `json:"canCashTicket,omitempty"`
	CanGoodsNumber      bool                                                       `json:"canGoodsNumber,omitempty"`
	GoodsCashTicket     int64                                                      `json:"goodsCashTicket,omitempty"`
	HasUseThreshold     bool                                                       `json:"hasUseThreshold,omitempty"`
	ImportGoodsFileName string                                                     `json:"importGoodsFileName,omitempty"`
	ImportGoodsFlag     bool                                                       `json:"importGoodsFlag,omitempty"`
	IsAllTimeUse        bool                                                       `json:"isAllTimeUse,omitempty"`
	IsAllUseScene       bool                                                       `json:"isAllUseScene,omitempty"`
	MaxGoodsNumber      int64                                                      `json:"maxGoodsNumber,omitempty"`
	MinGoodsNumber      int64                                                      `json:"minGoodsNumber,omitempty"`
	OrderDeductList     []int64                                                    `json:"orderDeductList,omitempty"`
}
type PaasWeimobCouponTemplateDetailGetResponseAcceptTypeDTO struct {
	AcceptStoreIds    PaasWeimobCouponTemplateDetailGetResponseAcceptStoreIds `json:"acceptStoreIds,omitempty"`
	AcceptGoodsIds    []int64                                                 `json:"acceptGoodsIds,omitempty"`
	AcceptGoodsType   int64                                                   `json:"acceptGoodsType,omitempty"`
	AcceptStoreType   int64                                                   `json:"acceptStoreType,omitempty"`
	ExistExcludeGoods bool                                                    `json:"existExcludeGoods,omitempty"`
	ExistExcludeStore bool                                                    `json:"existExcludeStore,omitempty"`
}
type PaasWeimobCouponTemplateDetailGetResponseAcceptStoreIds struct {
}
type PaasWeimobCouponTemplateDetailGetResponseSuperpositionRule struct {
	DifferentSuperposition bool  `json:"differentSuperposition,omitempty"`
	IsSuperposition        bool  `json:"isSuperposition,omitempty"`
	Limit                  int64 `json:"limit,omitempty"`
	SameSuperposition      bool  `json:"sameSuperposition,omitempty"`
}
type PaasWeimobCouponTemplateDetailGetResponseUseValidDate struct {
	UseDelayDays  int64  `json:"useDelayDays,omitempty"`
	UseEndTime    string `json:"useEndTime,omitempty"`
	UseStartTime  string `json:"useStartTime,omitempty"`
	UseTimeType   int64  `json:"useTimeType,omitempty"`
	ValidDateDesc string `json:"validDateDesc,omitempty"`
	ValidDays     int64  `json:"validDays,omitempty"`
}

func CreatePaasWeimobCouponTemplateDetailGetResponse() *PaasWeimobCouponTemplateDetailGetResponse {
	return &PaasWeimobCouponTemplateDetailGetResponse{}
}

// OnPaasWeimobCouponTemplateDetailGetServiceInvocation
// @id 1160
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1160?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询优惠券详情)
func (s *Service) OnPaasWeimobCouponTemplateDetailGetServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCouponTemplateDetailGetRequest, PaasWeimobCouponTemplateDetailGetResponse],
) (service *Service) {
	var invocation = &PaasWeimobCouponTemplateDetailGetService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCouponTemplateDetailGetRequest, PaasWeimobCouponTemplateDetailGetResponse](invocation),
		"PaasWeimobCouponTemplateDetailGetService",
		"weimob.coupon.template.detail.get",
	)
	return s
}
