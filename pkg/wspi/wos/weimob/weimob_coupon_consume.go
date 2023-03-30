package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCouponConsumeService struct {
	handler spi.WosInvocationHandler[PaasWeimobCouponConsumeRequest, PaasWeimobCouponConsumeBool]
}

func (s PaasWeimobCouponConsumeService) NewRequest() spi.InvocationRequest[PaasWeimobCouponConsumeRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCouponConsumeRequest]{
		Params: &PaasWeimobCouponConsumeRequest{},
	}
}

func (s PaasWeimobCouponConsumeService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCouponConsumeRequest],
) (
	spi.InvocationResponse[PaasWeimobCouponConsumeBool],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCouponConsumeRequest]))
}

type PaasWeimobCouponConsumeRequest struct {
	OperateCoupons  []PaasWeimobCouponConsumeRequestOperateCoupons `json:"operateCoupons,omitempty"`
	OrderAmount     int64                                          `json:"orderAmount,omitempty"`
	Operator        int64                                          `json:"operator,omitempty"`
	UseResourceType int64                                          `json:"useResourceType,omitempty"`
	UseScene        int64                                          `json:"useScene,omitempty"`
}
type PaasWeimobCouponConsumeRequestOperateCoupons struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	OrderNo          string `json:"orderNo,omitempty"`
	Code             string `json:"code,omitempty"`
	Amount           int64  `json:"amount,omitempty"`
}

type PaasWeimobCouponConsumeBool bool

func CreatePaasWeimobCouponConsumeBool() *PaasWeimobCouponConsumeBool {
	var br PaasWeimobCouponConsumeBool
	return &br
}

// OnPaasWeimobCouponConsumeServiceInvocation
// @id 1078
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1078?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=核销优惠券)
func (s *Service) OnPaasWeimobCouponConsumeServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCouponConsumeRequest, PaasWeimobCouponConsumeBool],
) (service *Service) {
	var invocation = &PaasWeimobCouponConsumeService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCouponConsumeRequest, PaasWeimobCouponConsumeBool](invocation),
		"PaasWeimobCouponConsumeService",
		"weimob.coupon.consume",
	)
	return s
}
