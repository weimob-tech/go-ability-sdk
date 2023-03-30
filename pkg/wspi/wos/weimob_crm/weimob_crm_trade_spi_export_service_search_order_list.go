package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCrmTradeSpiExportServiceSearchOrderListService struct {
	handler spi.WosInvocationHandler[PaasWeimobCrmTradeSpiExportSearchOrderListRequest, PaasWeimobCrmTradeSpiExportSearchOrderListResponse]
}

func (s PaasWeimobCrmTradeSpiExportServiceSearchOrderListService) NewRequest() spi.InvocationRequest[PaasWeimobCrmTradeSpiExportSearchOrderListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCrmTradeSpiExportSearchOrderListRequest]{
		Params: &PaasWeimobCrmTradeSpiExportSearchOrderListRequest{},
	}
}

func (s PaasWeimobCrmTradeSpiExportServiceSearchOrderListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCrmTradeSpiExportSearchOrderListRequest],
) (
	spi.InvocationResponse[PaasWeimobCrmTradeSpiExportSearchOrderListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCrmTradeSpiExportSearchOrderListRequest]))
}

type PaasWeimobCrmTradeSpiExportSearchOrderListRequest struct {
}

type PaasWeimobCrmTradeSpiExportSearchOrderListResponse struct {
}

func CreatePaasWeimobCrmTradeSpiExportSearchOrderListResponse() *PaasWeimobCrmTradeSpiExportSearchOrderListResponse {
	return &PaasWeimobCrmTradeSpiExportSearchOrderListResponse{}
}

// OnPaasWeimobCrmTradeSpiExportServiceSearchOrderListServiceInvocation
// @id 483
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/483?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询订单数据)
func (s *Service) OnPaasWeimobCrmTradeSpiExportServiceSearchOrderListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCrmTradeSpiExportSearchOrderListRequest, PaasWeimobCrmTradeSpiExportSearchOrderListResponse],
) (service *Service) {
	var invocation = &PaasWeimobCrmTradeSpiExportServiceSearchOrderListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCrmTradeSpiExportSearchOrderListRequest, PaasWeimobCrmTradeSpiExportSearchOrderListResponse](invocation),
		"PaasWeimobCrmTradeSpiExportServiceSearchOrderListService",
		"weimobCrm.tradeSpiExportService.searchOrderList",
	)
	return s
}
