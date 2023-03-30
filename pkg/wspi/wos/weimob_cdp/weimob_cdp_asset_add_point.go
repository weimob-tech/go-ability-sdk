package weimob_cdp

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCdpAssetAddPointService struct {
	handler spi.WosInvocationHandler[PaasWeimobCdpAssetAddPointRequest, PaasWeimobCdpAssetAddPointResponse]
}

func (s PaasWeimobCdpAssetAddPointService) NewRequest() spi.InvocationRequest[PaasWeimobCdpAssetAddPointRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCdpAssetAddPointRequest]{
		Params: &PaasWeimobCdpAssetAddPointRequest{},
	}
}

func (s PaasWeimobCdpAssetAddPointService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCdpAssetAddPointRequest],
) (
	spi.InvocationResponse[PaasWeimobCdpAssetAddPointResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCdpAssetAddPointRequest]))
}

type PaasWeimobCdpAssetAddPointRequest struct {
}

type PaasWeimobCdpAssetAddPointResponse struct {
}

func CreatePaasWeimobCdpAssetAddPointResponse() *PaasWeimobCdpAssetAddPointResponse {
	return &PaasWeimobCdpAssetAddPointResponse{}
}

// OnPaasWeimobCdpAssetAddPointServiceInvocation
// @id 614
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/614?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=addPoint)
func (s *Service) OnPaasWeimobCdpAssetAddPointServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCdpAssetAddPointRequest, PaasWeimobCdpAssetAddPointResponse],
) (service *Service) {
	var invocation = &PaasWeimobCdpAssetAddPointService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCdpAssetAddPointRequest, PaasWeimobCdpAssetAddPointResponse](invocation),
		"PaasWeimobCdpAssetAddPointService",
		"weimobCdp.asset.addPoint",
	)
	return s
}
