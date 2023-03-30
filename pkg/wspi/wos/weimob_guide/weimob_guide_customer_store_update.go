package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideCustomerStoreUpdateService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideCustomerStoreUpdateRequest, PaasWeimobGuideCustomerStoreUpdateResponse]
}

func (s PaasWeimobGuideCustomerStoreUpdateService) NewRequest() spi.InvocationRequest[PaasWeimobGuideCustomerStoreUpdateRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideCustomerStoreUpdateRequest]{
		Params: &PaasWeimobGuideCustomerStoreUpdateRequest{},
	}
}

func (s PaasWeimobGuideCustomerStoreUpdateService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideCustomerStoreUpdateRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideCustomerStoreUpdateResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideCustomerStoreUpdateRequest]))
}

type PaasWeimobGuideCustomerStoreUpdateRequest struct {
}

type PaasWeimobGuideCustomerStoreUpdateResponse struct {
}

func CreatePaasWeimobGuideCustomerStoreUpdateResponse() *PaasWeimobGuideCustomerStoreUpdateResponse {
	return &PaasWeimobGuideCustomerStoreUpdateResponse{}
}

// OnPaasWeimobGuideCustomerStoreUpdateServiceInvocation
// @id 745
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/745?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新客户归属门店)
func (s *Service) OnPaasWeimobGuideCustomerStoreUpdateServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideCustomerStoreUpdateRequest, PaasWeimobGuideCustomerStoreUpdateResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideCustomerStoreUpdateService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideCustomerStoreUpdateRequest, PaasWeimobGuideCustomerStoreUpdateResponse](invocation),
		"PaasWeimobGuideCustomerStoreUpdateService",
		"weimobGuide.customer.store.update",
	)
	return s
}
