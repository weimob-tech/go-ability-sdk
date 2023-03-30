package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCrmTuikeSpiExportServiceGetTuikeLevelInfoService struct {
	handler spi.WosInvocationHandler[PaasWeimobCrmTuikeSpiExportGetTuikeLevelInfoRequest, PaasWeimobCrmTuikeSpiExportGetTuikeLevelInfoResponse]
}

func (s PaasWeimobCrmTuikeSpiExportServiceGetTuikeLevelInfoService) NewRequest() spi.InvocationRequest[PaasWeimobCrmTuikeSpiExportGetTuikeLevelInfoRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCrmTuikeSpiExportGetTuikeLevelInfoRequest]{
		Params: &PaasWeimobCrmTuikeSpiExportGetTuikeLevelInfoRequest{},
	}
}

func (s PaasWeimobCrmTuikeSpiExportServiceGetTuikeLevelInfoService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCrmTuikeSpiExportGetTuikeLevelInfoRequest],
) (
	spi.InvocationResponse[PaasWeimobCrmTuikeSpiExportGetTuikeLevelInfoResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCrmTuikeSpiExportGetTuikeLevelInfoRequest]))
}

type PaasWeimobCrmTuikeSpiExportGetTuikeLevelInfoRequest struct {
}

type PaasWeimobCrmTuikeSpiExportGetTuikeLevelInfoResponse struct {
}

func CreatePaasWeimobCrmTuikeSpiExportGetTuikeLevelInfoResponse() *PaasWeimobCrmTuikeSpiExportGetTuikeLevelInfoResponse {
	return &PaasWeimobCrmTuikeSpiExportGetTuikeLevelInfoResponse{}
}

// OnPaasWeimobCrmTuikeSpiExportServiceGetTuikeLevelInfoServiceInvocation
// @id 485
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/485?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=批量查询微客等级信息)
func (s *Service) OnPaasWeimobCrmTuikeSpiExportServiceGetTuikeLevelInfoServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCrmTuikeSpiExportGetTuikeLevelInfoRequest, PaasWeimobCrmTuikeSpiExportGetTuikeLevelInfoResponse],
) (service *Service) {
	var invocation = &PaasWeimobCrmTuikeSpiExportServiceGetTuikeLevelInfoService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCrmTuikeSpiExportGetTuikeLevelInfoRequest, PaasWeimobCrmTuikeSpiExportGetTuikeLevelInfoResponse](invocation),
		"PaasWeimobCrmTuikeSpiExportServiceGetTuikeLevelInfoService",
		"weimobCrm.tuikeSpiExportService.getTuikeLevelInfo",
	)
	return s
}
