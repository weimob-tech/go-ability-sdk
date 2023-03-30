package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCouponTemplateGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobCouponTemplateGetListRequest, PaasWeimobCouponTemplateGetListResponse]
}

func (s PaasWeimobCouponTemplateGetListService) NewRequest() spi.InvocationRequest[PaasWeimobCouponTemplateGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCouponTemplateGetListRequest]{
		Params: &PaasWeimobCouponTemplateGetListRequest{},
	}
}

func (s PaasWeimobCouponTemplateGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCouponTemplateGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobCouponTemplateGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCouponTemplateGetListRequest]))
}

type PaasWeimobCouponTemplateGetListRequest struct {
	CouponTemplateIds    []int64 `json:"couponTemplateIds,omitempty"`
	Useful               bool    `json:"useful,omitempty"`
	QueryUserReceiveInfo bool    `json:"queryUserReceiveInfo,omitempty"`
	Wid                  int64   `json:"wid,omitempty"`
	ProductInstanceId    int64   `json:"productInstanceId,omitempty"`
}

type PaasWeimobCouponTemplateGetListResponse struct {
	ValidDate                 PaasWeimobCouponTemplateGetListResponseValidDate     `json:"validDate,omitempty"`
	AcceptTypeDTO             PaasWeimobCouponTemplateGetListResponseAcceptTypeDTO `json:"acceptTypeDTO,omitempty"`
	CouponTemplateId          int64                                                `json:"couponTemplateId,omitempty"`
	Name                      string                                               `json:"name,omitempty"`
	CouponType                int64                                                `json:"couponType,omitempty"`
	CouponSubType             int64                                                `json:"couponSubType,omitempty"`
	Status                    int64                                                `json:"status,omitempty"`
	PlatformType              int64                                                `json:"platformType,omitempty"`
	Stock                     int64                                                `json:"stock,omitempty"`
	Explain                   string                                               `json:"explain,omitempty"`
	ReducePriceType           int64                                                `json:"reducePriceType,omitempty"`
	FreightReduceType         int64                                                `json:"freightReduceType,omitempty"`
	CashTicketAmt             int64                                                `json:"cashTicketAmt,omitempty"`
	MinCashTicketAmt          int64                                                `json:"minCashTicketAmt,omitempty"`
	MaxCashTicketAmt          int64                                                `json:"maxCashTicketAmt,omitempty"`
	Discount                  int64                                                `json:"discount,omitempty"`
	DiscountLimitType         int64                                                `json:"discountLimitType,omitempty"`
	DiscountLimitValue        int64                                                `json:"discountLimitValue,omitempty"`
	CashTicketCondition       int64                                                `json:"cashTicketCondition,omitempty"`
	UserTakeLimit             int64                                                `json:"userTakeLimit,omitempty"`
	HasUseThreshold           int64                                                `json:"hasUseThreshold,omitempty"`
	CanCashTicket             bool                                                 `json:"canCashTicket,omitempty"`
	GoodsCashTicket           int64                                                `json:"goodsCashTicket,omitempty"`
	CanGoodsNumber            bool                                                 `json:"canGoodsNumber,omitempty"`
	MinGoodsNumber            int64                                                `json:"minGoodsNumber,omitempty"`
	MaxGoodsNumber            int64                                                `json:"maxGoodsNumber,omitempty"`
	CanStoreLaunch            bool                                                 `json:"canStoreLaunch,omitempty"`
	IsEnable                  bool                                                 `json:"isEnable,omitempty"`
	Desc                      string                                               `json:"desc,omitempty"`
	RecommendStartTime        int64                                                `json:"recommendStartTime,omitempty"`
	RecommendEndTime          int64                                                `json:"recommendEndTime,omitempty"`
	ImageUrl                  string                                               `json:"imageUrl,omitempty"`
	IsPayment                 bool                                                 `json:"isPayment,omitempty"`
	CanGiveFriend             bool                                                 `json:"canGiveFriend,omitempty"`
	CodeGenerateType          int64                                                `json:"codeGenerateType,omitempty"`
	CustomEndNum              int64                                                `json:"customEndNum,omitempty"`
	CustomStartNum            int64                                                `json:"customStartNum,omitempty"`
	FixLeft                   string                                               `json:"fixLeft,omitempty"`
	FixRight                  string                                               `json:"fixRight,omitempty"`
	CreateTime                int64                                                `json:"createTime,omitempty"`
	IsAlipayCouponRecommended bool                                                 `json:"isAlipayCouponRecommended,omitempty"`
	RemainStock               int64                                                `json:"remainStock,omitempty"`
	Outer                     bool                                                 `json:"outer,omitempty"`
	ActivityPublish           bool                                                 `json:"activityPublish,omitempty"`
	CustomerDirectReceive     bool                                                 `json:"customerDirectReceive,omitempty"`
}
type PaasWeimobCouponTemplateGetListResponseValidDate struct {
}
type PaasWeimobCouponTemplateGetListResponseAcceptTypeDTO struct {
}

func CreatePaasWeimobCouponTemplateGetListResponse() *PaasWeimobCouponTemplateGetListResponse {
	return &PaasWeimobCouponTemplateGetListResponse{}
}

// OnPaasWeimobCouponTemplateGetListServiceInvocation
// @id 1079
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1079?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=批量查询券模板信息)
func (s *Service) OnPaasWeimobCouponTemplateGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCouponTemplateGetListRequest, PaasWeimobCouponTemplateGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobCouponTemplateGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCouponTemplateGetListRequest, PaasWeimobCouponTemplateGetListResponse](invocation),
		"PaasWeimobCouponTemplateGetListService",
		"weimob.coupon.template.getList",
	)
	return s
}
