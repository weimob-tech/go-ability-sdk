package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCouponTemplateSearchService struct {
	handler spi.WosInvocationHandler[PaasWeimobCouponTemplateSearchRequest, PaasWeimobCouponTemplateSearchResponse]
}

func (s PaasWeimobCouponTemplateSearchService) NewRequest() spi.InvocationRequest[PaasWeimobCouponTemplateSearchRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCouponTemplateSearchRequest]{
		Params: &PaasWeimobCouponTemplateSearchRequest{},
	}
}

func (s PaasWeimobCouponTemplateSearchService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCouponTemplateSearchRequest],
) (
	spi.InvocationResponse[PaasWeimobCouponTemplateSearchResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCouponTemplateSearchRequest]))
}

type PaasWeimobCouponTemplateSearchRequest struct {
	CouponTypes     []int64 `json:"couponTypes,omitempty"`
	Name            string  `json:"name,omitempty"`
	StatusList      []int64 `json:"statusList,omitempty"`
	PlatformType    int64   `json:"platformType,omitempty"`
	BelongVidType   int64   `json:"belongVidType,omitempty"`
	AcceptVidType   int64   `json:"acceptVidType,omitempty"`
	PublishChannels []int64 `json:"publishChannels,omitempty"`
	AcceptVids      []int64 `json:"acceptVids,omitempty"`
	Paths           []int64 `json:"paths,omitempty"`
	PageNum         int64   `json:"pageNum,omitempty"`
	PageSize        int64   `json:"pageSize,omitempty"`
}

type PaasWeimobCouponTemplateSearchResponse struct {
	PageList   []PaasWeimobCouponTemplateSearchResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                            `json:"pageNum,omitempty"`
	PageSize   int64                                            `json:"pageSize,omitempty"`
	TotalCount int64                                            `json:"totalCount,omitempty"`
}
type PaasWeimobCouponTemplateSearchResponsePageList struct {
	ValidDate                 PaasWeimobCouponTemplateSearchResponseValidDate     `json:"validDate,omitempty"`
	AcceptTypeDTO             PaasWeimobCouponTemplateSearchResponseAcceptTypeDTO `json:"acceptTypeDTO,omitempty"`
	BelongVid                 int64                                               `json:"belongVid,omitempty"`
	BelongVidName             string                                              `json:"belongVidName,omitempty"`
	CouponTemplateId          int64                                               `json:"couponTemplateId,omitempty"`
	Name                      string                                              `json:"name,omitempty"`
	CouponType                int64                                               `json:"couponType,omitempty"`
	CouponSubType             int64                                               `json:"couponSubType,omitempty"`
	CouponTypeDesc            string                                              `json:"couponTypeDesc,omitempty"`
	Status                    int64                                               `json:"status,omitempty"`
	PublishChannels           []int64                                             `json:"publishChannels,omitempty"`
	PlatformType              int64                                               `json:"platformType,omitempty"`
	Stock                     int64                                               `json:"stock,omitempty"`
	Explain                   string                                              `json:"explain,omitempty"`
	ReducePriceType           int64                                               `json:"reducePriceType,omitempty"`
	FreightReduceType         int64                                               `json:"freightReduceType,omitempty"`
	CashTicketAmt             int64                                               `json:"cashTicketAmt,omitempty"`
	MinCashTicketAmt          int64                                               `json:"minCashTicketAmt,omitempty"`
	MaxCashTicketAmt          int64                                               `json:"maxCashTicketAmt,omitempty"`
	Discount                  int64                                               `json:"discount,omitempty"`
	DiscountLimitType         int64                                               `json:"discountLimitType,omitempty"`
	DiscountLimitValue        int64                                               `json:"discountLimitValue,omitempty"`
	CashTicketCondition       int64                                               `json:"cashTicketCondition,omitempty"`
	UserTakeLimit             int64                                               `json:"userTakeLimit,omitempty"`
	HasUseThreshold           bool                                                `json:"hasUseThreshold,omitempty"`
	CanCashTicket             bool                                                `json:"canCashTicket,omitempty"`
	GoodsCashTicket           int64                                               `json:"goodsCashTicket,omitempty"`
	CanGoodsNumber            bool                                                `json:"canGoodsNumber,omitempty"`
	MinGoodsNumber            int64                                               `json:"minGoodsNumber,omitempty"`
	MaxGoodsNumber            int64                                               `json:"maxGoodsNumber,omitempty"`
	CanStoreLaunch            bool                                                `json:"canStoreLaunch,omitempty"`
	IsEnable                  bool                                                `json:"isEnable,omitempty"`
	Desc                      string                                              `json:"desc,omitempty"`
	RecommendStartTime        int64                                               `json:"recommendStartTime,omitempty"`
	RecommendEndTime          int64                                               `json:"recommendEndTime,omitempty"`
	ImageUrl                  string                                              `json:"imageUrl,omitempty"`
	ShareImageUrl             string                                              `json:"shareImageUrl,omitempty"`
	StockCreatorMchid         string                                              `json:"stockCreatorMchid,omitempty"`
	ServiceMchid              string                                              `json:"serviceMchid,omitempty"`
	IsPayment                 bool                                                `json:"isPayment,omitempty"`
	CanGiveFriend             bool                                                `json:"canGiveFriend,omitempty"`
	CodeGenerateType          int64                                               `json:"codeGenerateType,omitempty"`
	CustomEndNum              int64                                               `json:"customEndNum,omitempty"`
	CustomStartNum            int64                                               `json:"customStartNum,omitempty"`
	FixLeft                   string                                              `json:"fixLeft,omitempty"`
	FixRight                  string                                              `json:"fixRight,omitempty"`
	CreateTime                int64                                               `json:"createTime,omitempty"`
	IsAlipayCouponRecommended bool                                                `json:"isAlipayCouponRecommended,omitempty"`
	ThirdTemplateId           string                                              `json:"thirdTemplateId,omitempty"`
	RemainStock               int64                                               `json:"remainStock,omitempty"`
	Outer                     bool                                                `json:"outer,omitempty"`
	ActivityPublish           bool                                                `json:"activityPublish,omitempty"`
	CustomerDirectReceive     bool                                                `json:"customerDirectReceive,omitempty"`
	IsAcceptAllCrowd          bool                                                `json:"isAcceptAllCrowd,omitempty"`
}
type PaasWeimobCouponTemplateSearchResponseValidDate struct {
	UseTimeType  int64 `json:"useTimeType,omitempty"`
	UseStartTime int64 `json:"useStartTime,omitempty"`
	UseEndTime   int64 `json:"useEndTime,omitempty"`
	UseDelayDays int64 `json:"useDelayDays,omitempty"`
	ValidDays    int64 `json:"validDays,omitempty"`
}
type PaasWeimobCouponTemplateSearchResponseAcceptTypeDTO struct {
	AcceptStoreType   int64   `json:"acceptStoreType,omitempty"`
	AcceptStoreIds    []int64 `json:"acceptStoreIds,omitempty"`
	AcceptGoodsType   int64   `json:"acceptGoodsType,omitempty"`
	AcceptGoodsIds    []int64 `json:"acceptGoodsIds,omitempty"`
	ExistExcludeGoods bool    `json:"existExcludeGoods,omitempty"`
	ExcludeGoodsIds   []int64 `json:"excludeGoodsIds,omitempty"`
	ExistExcludeStore bool    `json:"existExcludeStore,omitempty"`
	ExcludeStoreIds   []int64 `json:"excludeStoreIds,omitempty"`
	IncludeStoreGoods bool    `json:"includeStoreGoods,omitempty"`
}

func CreatePaasWeimobCouponTemplateSearchResponse() *PaasWeimobCouponTemplateSearchResponse {
	return &PaasWeimobCouponTemplateSearchResponse{}
}

// OnPaasWeimobCouponTemplateSearchServiceInvocation
// @id 1076
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1076?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询商户券模板列表)
func (s *Service) OnPaasWeimobCouponTemplateSearchServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCouponTemplateSearchRequest, PaasWeimobCouponTemplateSearchResponse],
) (service *Service) {
	var invocation = &PaasWeimobCouponTemplateSearchService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCouponTemplateSearchRequest, PaasWeimobCouponTemplateSearchResponse](invocation),
		"PaasWeimobCouponTemplateSearchService",
		"weimob.coupon.template.search",
	)
	return s
}
