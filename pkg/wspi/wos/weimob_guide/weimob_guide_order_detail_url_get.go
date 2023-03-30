package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideOrderDetailUrlGetService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideOrderDetailUrlGetRequest, PaasWeimobGuideOrderDetailUrlGetResponse]
}

func (s PaasWeimobGuideOrderDetailUrlGetService) NewRequest() spi.InvocationRequest[PaasWeimobGuideOrderDetailUrlGetRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideOrderDetailUrlGetRequest]{
		Params: &PaasWeimobGuideOrderDetailUrlGetRequest{},
	}
}

func (s PaasWeimobGuideOrderDetailUrlGetService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideOrderDetailUrlGetRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideOrderDetailUrlGetResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideOrderDetailUrlGetRequest]))
}

type PaasWeimobGuideOrderDetailUrlGetRequest struct {
	PageType string `json:"pageType,omitempty"`
}

type PaasWeimobGuideOrderDetailUrlGetResponse struct {
	UrlLink string `json:"urlLink,omitempty"`
}

func CreatePaasWeimobGuideOrderDetailUrlGetResponse() *PaasWeimobGuideOrderDetailUrlGetResponse {
	return &PaasWeimobGuideOrderDetailUrlGetResponse{}
}

// OnPaasWeimobGuideOrderDetailUrlGetServiceInvocation
// @id 733
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/733?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询订单详情页地址)
func (s *Service) OnPaasWeimobGuideOrderDetailUrlGetServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideOrderDetailUrlGetRequest, PaasWeimobGuideOrderDetailUrlGetResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideOrderDetailUrlGetService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideOrderDetailUrlGetRequest, PaasWeimobGuideOrderDetailUrlGetResponse](invocation),
		"PaasWeimobGuideOrderDetailUrlGetService",
		"weimobGuide.order.detail.url.get",
	)
	return s
}
