package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideCouponDetailGetService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideCouponDetailGetRequest, PaasWeimobGuideCouponDetailGetResponse]
}

func (s PaasWeimobGuideCouponDetailGetService) NewRequest() spi.InvocationRequest[PaasWeimobGuideCouponDetailGetRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideCouponDetailGetRequest]{
		Params: &PaasWeimobGuideCouponDetailGetRequest{},
	}
}

func (s PaasWeimobGuideCouponDetailGetService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideCouponDetailGetRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideCouponDetailGetResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideCouponDetailGetRequest]))
}

type PaasWeimobGuideCouponDetailGetRequest struct {
	Wid               int64   `json:"wid,omitempty"`
	CouponTemplateIds []int64 `json:"couponTemplateIds,omitempty"`
}

type PaasWeimobGuideCouponDetailGetResponse struct {
	RecommendStartTime PaasWeimobGuideCouponDetailGetResponseRecommendStartTime `json:"recommendStartTime,omitempty"`
	RecommendEndTime   PaasWeimobGuideCouponDetailGetResponseRecommendEndTime   `json:"recommendEndTime,omitempty"`
	ValidDate          PaasWeimobGuideCouponDetailGetResponseValidDate          `json:"validDate,omitempty"`
	CouponTemplateId   int64                                                    `json:"couponTemplateId,omitempty"`
	Name               string                                                   `json:"name,omitempty"`
	Status             int64                                                    `json:"status,omitempty"`
	ImageUrl           string                                                   `json:"imageUrl,omitempty"`
	PlatformType       int64                                                    `json:"platformType,omitempty"`
	CouponType         int64                                                    `json:"couponType,omitempty"`
	Stock              int64                                                    `json:"stock,omitempty"`
}
type PaasWeimobGuideCouponDetailGetResponseRecommendStartTime struct {
}
type PaasWeimobGuideCouponDetailGetResponseRecommendEndTime struct {
}
type PaasWeimobGuideCouponDetailGetResponseValidDate struct {
	UseStartTime  PaasWeimobGuideCouponDetailGetResponseUseStartTime `json:"useStartTime,omitempty"`
	UseEndTime    PaasWeimobGuideCouponDetailGetResponseUseEndTime   `json:"useEndTime,omitempty"`
	ValidDays     int64                                              `json:"validDays,omitempty"`
	ValidDateDesc string                                             `json:"validDateDesc,omitempty"`
}
type PaasWeimobGuideCouponDetailGetResponseUseStartTime struct {
}
type PaasWeimobGuideCouponDetailGetResponseUseEndTime struct {
}

func CreatePaasWeimobGuideCouponDetailGetResponse() *PaasWeimobGuideCouponDetailGetResponse {
	return &PaasWeimobGuideCouponDetailGetResponse{}
}

// OnPaasWeimobGuideCouponDetailGetServiceInvocation
// @id 729
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/729?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询优惠券信息)
func (s *Service) OnPaasWeimobGuideCouponDetailGetServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideCouponDetailGetRequest, PaasWeimobGuideCouponDetailGetResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideCouponDetailGetService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideCouponDetailGetRequest, PaasWeimobGuideCouponDetailGetResponse](invocation),
		"PaasWeimobGuideCouponDetailGetService",
		"weimobGuide.coupon.detail.get",
	)
	return s
}
