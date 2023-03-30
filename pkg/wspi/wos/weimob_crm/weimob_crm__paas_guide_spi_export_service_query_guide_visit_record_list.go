package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCrmPaasGuideSpiExportServiceQueryGuideVisitRecordListService struct {
	handler spi.WosInvocationHandler[PaasWeimobCrmPaasGuideSpiExportQueryGuideVisitRecordListRequest, PaasWeimobCrmPaasGuideSpiExportQueryGuideVisitRecordListResponse]
}

func (s PaasWeimobCrmPaasGuideSpiExportServiceQueryGuideVisitRecordListService) NewRequest() spi.InvocationRequest[PaasWeimobCrmPaasGuideSpiExportQueryGuideVisitRecordListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCrmPaasGuideSpiExportQueryGuideVisitRecordListRequest]{
		Params: &PaasWeimobCrmPaasGuideSpiExportQueryGuideVisitRecordListRequest{},
	}
}

func (s PaasWeimobCrmPaasGuideSpiExportServiceQueryGuideVisitRecordListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCrmPaasGuideSpiExportQueryGuideVisitRecordListRequest],
) (
	spi.InvocationResponse[PaasWeimobCrmPaasGuideSpiExportQueryGuideVisitRecordListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCrmPaasGuideSpiExportQueryGuideVisitRecordListRequest]))
}

type PaasWeimobCrmPaasGuideSpiExportQueryGuideVisitRecordListRequest struct {
}

type PaasWeimobCrmPaasGuideSpiExportQueryGuideVisitRecordListResponse struct {
}

func CreatePaasWeimobCrmPaasGuideSpiExportQueryGuideVisitRecordListResponse() *PaasWeimobCrmPaasGuideSpiExportQueryGuideVisitRecordListResponse {
	return &PaasWeimobCrmPaasGuideSpiExportQueryGuideVisitRecordListResponse{}
}

// OnPaasWeimobCrmPaasGuideSpiExportServiceQueryGuideVisitRecordListServiceInvocation
// @id 487
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/487?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询导购回访信息)
func (s *Service) OnPaasWeimobCrmPaasGuideSpiExportServiceQueryGuideVisitRecordListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCrmPaasGuideSpiExportQueryGuideVisitRecordListRequest, PaasWeimobCrmPaasGuideSpiExportQueryGuideVisitRecordListResponse],
) (service *Service) {
	var invocation = &PaasWeimobCrmPaasGuideSpiExportServiceQueryGuideVisitRecordListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCrmPaasGuideSpiExportQueryGuideVisitRecordListRequest, PaasWeimobCrmPaasGuideSpiExportQueryGuideVisitRecordListResponse](invocation),
		"PaasWeimobCrmPaasGuideSpiExportServiceQueryGuideVisitRecordListService",
		"weimobCrm.PaasGuideSpiExportService.queryGuideVisitRecordList",
	)
	return s
}
