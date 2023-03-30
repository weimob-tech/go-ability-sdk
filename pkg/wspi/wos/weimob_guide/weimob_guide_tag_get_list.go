package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideTagGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideTagGetListRequest, PaasWeimobGuideTagGetListResponse]
}

func (s PaasWeimobGuideTagGetListService) NewRequest() spi.InvocationRequest[PaasWeimobGuideTagGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideTagGetListRequest]{
		Params: &PaasWeimobGuideTagGetListRequest{},
	}
}

func (s PaasWeimobGuideTagGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideTagGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideTagGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideTagGetListRequest]))
}

type PaasWeimobGuideTagGetListRequest struct {
}

type PaasWeimobGuideTagGetListResponse struct {
}

func CreatePaasWeimobGuideTagGetListResponse() *PaasWeimobGuideTagGetListResponse {
	return &PaasWeimobGuideTagGetListResponse{}
}

// OnPaasWeimobGuideTagGetListServiceInvocation
// @id 742
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/742?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询标签信息)
func (s *Service) OnPaasWeimobGuideTagGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideTagGetListRequest, PaasWeimobGuideTagGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideTagGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideTagGetListRequest, PaasWeimobGuideTagGetListResponse](invocation),
		"PaasWeimobGuideTagGetListService",
		"weimobGuide.tag.getList",
	)
	return s
}
