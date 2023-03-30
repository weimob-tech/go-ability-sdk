package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCrmGuideSpiExportServiceGetGuiderInfoByCustomerWidService struct {
	handler spi.WosInvocationHandler[PaasWeimobCrmGuideSpiExportGetGuiderInfoByCustomerWidRequest, PaasWeimobCrmGuideSpiExportGetGuiderInfoByCustomerWidResponse]
}

func (s PaasWeimobCrmGuideSpiExportServiceGetGuiderInfoByCustomerWidService) NewRequest() spi.InvocationRequest[PaasWeimobCrmGuideSpiExportGetGuiderInfoByCustomerWidRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCrmGuideSpiExportGetGuiderInfoByCustomerWidRequest]{
		Params: &PaasWeimobCrmGuideSpiExportGetGuiderInfoByCustomerWidRequest{},
	}
}

func (s PaasWeimobCrmGuideSpiExportServiceGetGuiderInfoByCustomerWidService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCrmGuideSpiExportGetGuiderInfoByCustomerWidRequest],
) (
	spi.InvocationResponse[PaasWeimobCrmGuideSpiExportGetGuiderInfoByCustomerWidResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCrmGuideSpiExportGetGuiderInfoByCustomerWidRequest]))
}

type PaasWeimobCrmGuideSpiExportGetGuiderInfoByCustomerWidRequest struct {
}

type PaasWeimobCrmGuideSpiExportGetGuiderInfoByCustomerWidResponse struct {
}

func CreatePaasWeimobCrmGuideSpiExportGetGuiderInfoByCustomerWidResponse() *PaasWeimobCrmGuideSpiExportGetGuiderInfoByCustomerWidResponse {
	return &PaasWeimobCrmGuideSpiExportGetGuiderInfoByCustomerWidResponse{}
}

// OnPaasWeimobCrmGuideSpiExportServiceGetGuiderInfoByCustomerWidServiceInvocation
// @id 484
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/484?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=依据用户wid查询导购信息)
func (s *Service) OnPaasWeimobCrmGuideSpiExportServiceGetGuiderInfoByCustomerWidServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCrmGuideSpiExportGetGuiderInfoByCustomerWidRequest, PaasWeimobCrmGuideSpiExportGetGuiderInfoByCustomerWidResponse],
) (service *Service) {
	var invocation = &PaasWeimobCrmGuideSpiExportServiceGetGuiderInfoByCustomerWidService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCrmGuideSpiExportGetGuiderInfoByCustomerWidRequest, PaasWeimobCrmGuideSpiExportGetGuiderInfoByCustomerWidResponse](invocation),
		"PaasWeimobCrmGuideSpiExportServiceGetGuiderInfoByCustomerWidService",
		"weimobCrm.guideSpiExportService.getGuiderInfoByCustomerWid",
	)
	return s
}
