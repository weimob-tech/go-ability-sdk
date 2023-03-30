package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCouponChannelBindService struct {
	handler spi.WosInvocationHandler[PaasWeimobCouponChannelBindRequest, PaasWeimobCouponChannelBindBool]
}

func (s PaasWeimobCouponChannelBindService) NewRequest() spi.InvocationRequest[PaasWeimobCouponChannelBindRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCouponChannelBindRequest]{
		Params: &PaasWeimobCouponChannelBindRequest{},
	}
}

func (s PaasWeimobCouponChannelBindService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCouponChannelBindRequest],
) (
	spi.InvocationResponse[PaasWeimobCouponChannelBindBool],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCouponChannelBindRequest]))
}

type PaasWeimobCouponChannelBindRequest struct {
	BindChannel       int64   `json:"bindChannel,omitempty"`
	BindChannelId     int64   `json:"bindChannelId,omitempty"`
	CouponTemplateIds []int64 `json:"couponTemplateIds,omitempty"`
}

type PaasWeimobCouponChannelBindBool bool

func CreatePaasWeimobCouponChannelBindBool() *PaasWeimobCouponChannelBindBool {
	var br PaasWeimobCouponChannelBindBool
	return &br
}

// OnPaasWeimobCouponChannelBindServiceInvocation
// @id 1083
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1083?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=优惠券模板绑定)
func (s *Service) OnPaasWeimobCouponChannelBindServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCouponChannelBindRequest, PaasWeimobCouponChannelBindBool],
) (service *Service) {
	var invocation = &PaasWeimobCouponChannelBindService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCouponChannelBindRequest, PaasWeimobCouponChannelBindBool](invocation),
		"PaasWeimobCouponChannelBindService",
		"weimob.coupon.channel.bind",
	)
	return s
}
