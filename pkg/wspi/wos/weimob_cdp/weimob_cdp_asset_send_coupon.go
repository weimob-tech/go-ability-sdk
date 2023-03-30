package weimob_cdp

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCdpAssetSendCouponService struct {
	handler spi.WosInvocationHandler[PaasWeimobCdpAssetSendCouponRequest, PaasWeimobCdpAssetSendCouponResponse]
}

func (s PaasWeimobCdpAssetSendCouponService) NewRequest() spi.InvocationRequest[PaasWeimobCdpAssetSendCouponRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCdpAssetSendCouponRequest]{
		Params: &PaasWeimobCdpAssetSendCouponRequest{},
	}
}

func (s PaasWeimobCdpAssetSendCouponService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCdpAssetSendCouponRequest],
) (
	spi.InvocationResponse[PaasWeimobCdpAssetSendCouponResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCdpAssetSendCouponRequest]))
}

type PaasWeimobCdpAssetSendCouponRequest struct {
}

type PaasWeimobCdpAssetSendCouponResponse struct {
}

func CreatePaasWeimobCdpAssetSendCouponResponse() *PaasWeimobCdpAssetSendCouponResponse {
	return &PaasWeimobCdpAssetSendCouponResponse{}
}

// OnPaasWeimobCdpAssetSendCouponServiceInvocation
// @id 612
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/612?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=sendCoupon)
func (s *Service) OnPaasWeimobCdpAssetSendCouponServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCdpAssetSendCouponRequest, PaasWeimobCdpAssetSendCouponResponse],
) (service *Service) {
	var invocation = &PaasWeimobCdpAssetSendCouponService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCdpAssetSendCouponRequest, PaasWeimobCdpAssetSendCouponResponse](invocation),
		"PaasWeimobCdpAssetSendCouponService",
		"weimobCdp.asset.sendCoupon",
	)
	return s
}
