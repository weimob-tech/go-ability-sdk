package weimob_cdp

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCdpAssetAddAmountService struct {
	handler spi.WosInvocationHandler[PaasWeimobCdpAssetAddAmountRequest, PaasWeimobCdpAssetAddAmountResponse]
}

func (s PaasWeimobCdpAssetAddAmountService) NewRequest() spi.InvocationRequest[PaasWeimobCdpAssetAddAmountRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCdpAssetAddAmountRequest]{
		Params: &PaasWeimobCdpAssetAddAmountRequest{},
	}
}

func (s PaasWeimobCdpAssetAddAmountService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCdpAssetAddAmountRequest],
) (
	spi.InvocationResponse[PaasWeimobCdpAssetAddAmountResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCdpAssetAddAmountRequest]))
}

type PaasWeimobCdpAssetAddAmountRequest struct {
}

type PaasWeimobCdpAssetAddAmountResponse struct {
}

func CreatePaasWeimobCdpAssetAddAmountResponse() *PaasWeimobCdpAssetAddAmountResponse {
	return &PaasWeimobCdpAssetAddAmountResponse{}
}

// OnPaasWeimobCdpAssetAddAmountServiceInvocation
// @id 668
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/668?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=addAmount)
func (s *Service) OnPaasWeimobCdpAssetAddAmountServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCdpAssetAddAmountRequest, PaasWeimobCdpAssetAddAmountResponse],
) (service *Service) {
	var invocation = &PaasWeimobCdpAssetAddAmountService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCdpAssetAddAmountRequest, PaasWeimobCdpAssetAddAmountResponse](invocation),
		"PaasWeimobCdpAssetAddAmountService",
		"weimobCdp.asset.addAmount",
	)
	return s
}
