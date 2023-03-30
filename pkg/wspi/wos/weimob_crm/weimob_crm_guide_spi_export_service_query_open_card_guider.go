package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCrmGuideSpiExportServiceQueryOpenCardGuiderService struct {
	handler spi.WosInvocationHandler[PaasWeimobCrmGuideSpiExportQueryOpenCardGuiderRequest, PaasWeimobCrmGuideSpiExportQueryOpenCardGuiderResponse]
}

func (s PaasWeimobCrmGuideSpiExportServiceQueryOpenCardGuiderService) NewRequest() spi.InvocationRequest[PaasWeimobCrmGuideSpiExportQueryOpenCardGuiderRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCrmGuideSpiExportQueryOpenCardGuiderRequest]{
		Params: &PaasWeimobCrmGuideSpiExportQueryOpenCardGuiderRequest{},
	}
}

func (s PaasWeimobCrmGuideSpiExportServiceQueryOpenCardGuiderService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCrmGuideSpiExportQueryOpenCardGuiderRequest],
) (
	spi.InvocationResponse[PaasWeimobCrmGuideSpiExportQueryOpenCardGuiderResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCrmGuideSpiExportQueryOpenCardGuiderRequest]))
}

type PaasWeimobCrmGuideSpiExportQueryOpenCardGuiderRequest struct {
}

type PaasWeimobCrmGuideSpiExportQueryOpenCardGuiderResponse struct {
}

func CreatePaasWeimobCrmGuideSpiExportQueryOpenCardGuiderResponse() *PaasWeimobCrmGuideSpiExportQueryOpenCardGuiderResponse {
	return &PaasWeimobCrmGuideSpiExportQueryOpenCardGuiderResponse{}
}

// OnPaasWeimobCrmGuideSpiExportServiceQueryOpenCardGuiderServiceInvocation
// @id 488
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/488?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询开卡导购信息)
func (s *Service) OnPaasWeimobCrmGuideSpiExportServiceQueryOpenCardGuiderServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCrmGuideSpiExportQueryOpenCardGuiderRequest, PaasWeimobCrmGuideSpiExportQueryOpenCardGuiderResponse],
) (service *Service) {
	var invocation = &PaasWeimobCrmGuideSpiExportServiceQueryOpenCardGuiderService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCrmGuideSpiExportQueryOpenCardGuiderRequest, PaasWeimobCrmGuideSpiExportQueryOpenCardGuiderResponse](invocation),
		"PaasWeimobCrmGuideSpiExportServiceQueryOpenCardGuiderService",
		"weimobCrm.guideSpiExportService.queryOpenCardGuider",
	)
	return s
}
