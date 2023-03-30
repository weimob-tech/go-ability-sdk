package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCouponChannelUnbindService struct {
	handler spi.WosInvocationHandler[PaasWeimobCouponChannelUnbindRequest, PaasWeimobCouponChannelUnbindBool]
}

func (s PaasWeimobCouponChannelUnbindService) NewRequest() spi.InvocationRequest[PaasWeimobCouponChannelUnbindRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCouponChannelUnbindRequest]{
		Params: &PaasWeimobCouponChannelUnbindRequest{},
	}
}

func (s PaasWeimobCouponChannelUnbindService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCouponChannelUnbindRequest],
) (
	spi.InvocationResponse[PaasWeimobCouponChannelUnbindBool],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCouponChannelUnbindRequest]))
}

type PaasWeimobCouponChannelUnbindRequest struct {
	BindChannel       int64   `json:"bindChannel,omitempty"`
	BindChannelId     int64   `json:"bindChannelId,omitempty"`
	CouponTemplateIds []int64 `json:"couponTemplateIds,omitempty"`
}

type PaasWeimobCouponChannelUnbindBool bool

func CreatePaasWeimobCouponChannelUnbindBool() *PaasWeimobCouponChannelUnbindBool {
	var br PaasWeimobCouponChannelUnbindBool
	return &br
}

// OnPaasWeimobCouponChannelUnbindServiceInvocation
// @id 1084
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1084?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=优惠券模板解绑)
func (s *Service) OnPaasWeimobCouponChannelUnbindServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCouponChannelUnbindRequest, PaasWeimobCouponChannelUnbindBool],
) (service *Service) {
	var invocation = &PaasWeimobCouponChannelUnbindService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCouponChannelUnbindRequest, PaasWeimobCouponChannelUnbindBool](invocation),
		"PaasWeimobCouponChannelUnbindService",
		"weimob.coupon.channel.unbind",
	)
	return s
}
