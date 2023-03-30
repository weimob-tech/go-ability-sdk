package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideCustomerUrlGetService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideCustomerUrlGetRequest, PaasWeimobGuideCustomerUrlGetResponse]
}

func (s PaasWeimobGuideCustomerUrlGetService) NewRequest() spi.InvocationRequest[PaasWeimobGuideCustomerUrlGetRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideCustomerUrlGetRequest]{
		Params: &PaasWeimobGuideCustomerUrlGetRequest{},
	}
}

func (s PaasWeimobGuideCustomerUrlGetService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideCustomerUrlGetRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideCustomerUrlGetResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideCustomerUrlGetRequest]))
}

type PaasWeimobGuideCustomerUrlGetRequest struct {
	BosId             int64 `json:"bosId,omitempty"`
	Vid               int64 `json:"vid,omitempty"`
	ProductInstanceId int64 `json:"productInstanceId,omitempty"`
	VidType           int64 `json:"vidType,omitempty"`
	QueryWid          int64 `json:"queryWid,omitempty"`
}

type PaasWeimobGuideCustomerUrlGetResponse struct {
	UrlLink string `json:"urlLink,omitempty"`
}

func CreatePaasWeimobGuideCustomerUrlGetResponse() *PaasWeimobGuideCustomerUrlGetResponse {
	return &PaasWeimobGuideCustomerUrlGetResponse{}
}

// OnPaasWeimobGuideCustomerUrlGetServiceInvocation
// @id 749
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/749?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取客户页面链接)
func (s *Service) OnPaasWeimobGuideCustomerUrlGetServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideCustomerUrlGetRequest, PaasWeimobGuideCustomerUrlGetResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideCustomerUrlGetService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideCustomerUrlGetRequest, PaasWeimobGuideCustomerUrlGetResponse](invocation),
		"PaasWeimobGuideCustomerUrlGetService",
		"weimobGuide.customer.url.get",
	)
	return s
}
