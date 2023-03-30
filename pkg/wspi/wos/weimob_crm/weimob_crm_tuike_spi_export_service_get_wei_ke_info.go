package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCrmTuikeSpiExportServiceGetWeiKeInfoService struct {
	handler spi.WosInvocationHandler[PaasWeimobCrmTuikeSpiExportGetWeiKeInfoRequest, PaasWeimobCrmTuikeSpiExportGetWeiKeInfoResponse]
}

func (s PaasWeimobCrmTuikeSpiExportServiceGetWeiKeInfoService) NewRequest() spi.InvocationRequest[PaasWeimobCrmTuikeSpiExportGetWeiKeInfoRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCrmTuikeSpiExportGetWeiKeInfoRequest]{
		Params: &PaasWeimobCrmTuikeSpiExportGetWeiKeInfoRequest{},
	}
}

func (s PaasWeimobCrmTuikeSpiExportServiceGetWeiKeInfoService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCrmTuikeSpiExportGetWeiKeInfoRequest],
) (
	spi.InvocationResponse[PaasWeimobCrmTuikeSpiExportGetWeiKeInfoResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCrmTuikeSpiExportGetWeiKeInfoRequest]))
}

type PaasWeimobCrmTuikeSpiExportGetWeiKeInfoRequest struct {
}

type PaasWeimobCrmTuikeSpiExportGetWeiKeInfoResponse struct {
}

func CreatePaasWeimobCrmTuikeSpiExportGetWeiKeInfoResponse() *PaasWeimobCrmTuikeSpiExportGetWeiKeInfoResponse {
	return &PaasWeimobCrmTuikeSpiExportGetWeiKeInfoResponse{}
}

// OnPaasWeimobCrmTuikeSpiExportServiceGetWeiKeInfoServiceInvocation
// @id 486
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/486?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询微客信息)
func (s *Service) OnPaasWeimobCrmTuikeSpiExportServiceGetWeiKeInfoServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCrmTuikeSpiExportGetWeiKeInfoRequest, PaasWeimobCrmTuikeSpiExportGetWeiKeInfoResponse],
) (service *Service) {
	var invocation = &PaasWeimobCrmTuikeSpiExportServiceGetWeiKeInfoService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCrmTuikeSpiExportGetWeiKeInfoRequest, PaasWeimobCrmTuikeSpiExportGetWeiKeInfoResponse](invocation),
		"PaasWeimobCrmTuikeSpiExportServiceGetWeiKeInfoService",
		"weimobCrm.tuikeSpiExportService.getWeiKeInfo",
	)
	return s
}
