package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideCustomerConsumeGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideCustomerConsumeGetListRequest, PaasWeimobGuideCustomerConsumeGetListResponse]
}

func (s PaasWeimobGuideCustomerConsumeGetListService) NewRequest() spi.InvocationRequest[PaasWeimobGuideCustomerConsumeGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideCustomerConsumeGetListRequest]{
		Params: &PaasWeimobGuideCustomerConsumeGetListRequest{},
	}
}

func (s PaasWeimobGuideCustomerConsumeGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideCustomerConsumeGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideCustomerConsumeGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideCustomerConsumeGetListRequest]))
}

type PaasWeimobGuideCustomerConsumeGetListRequest struct {
	WidList           []PaasWeimobGuideCustomerConsumeGetListRequestWidList `json:"widList,omitempty"`
	Vid               int64                                                 `json:"vid,omitempty"`
	ProductInstanceId int64                                                 `json:"productInstanceId,omitempty"`
	BosId             int64                                                 `json:"bosId,omitempty"`
	VidType           int64                                                 `json:"vidType	,omitempty"`
}
type PaasWeimobGuideCustomerConsumeGetListRequestWidList struct {
}

type PaasWeimobGuideCustomerConsumeGetListResponse struct {
	Ukey                   string `json:"ukey,omitempty"`
	TotalConsumptionAmount int64  `json:"totalConsumptionAmount,omitempty"`
	TotalConsumptionCount  int64  `json:"totalConsumptionCount,omitempty"`
	LastPayTime            int64  `json:"lastPayTime,omitempty"`
}

func CreatePaasWeimobGuideCustomerConsumeGetListResponse() *PaasWeimobGuideCustomerConsumeGetListResponse {
	return &PaasWeimobGuideCustomerConsumeGetListResponse{}
}

// OnPaasWeimobGuideCustomerConsumeGetListServiceInvocation
// @id 748
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/748?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询客户消费)
func (s *Service) OnPaasWeimobGuideCustomerConsumeGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideCustomerConsumeGetListRequest, PaasWeimobGuideCustomerConsumeGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideCustomerConsumeGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideCustomerConsumeGetListRequest, PaasWeimobGuideCustomerConsumeGetListResponse](invocation),
		"PaasWeimobGuideCustomerConsumeGetListService",
		"weimobGuide.customer.consume.getList",
	)
	return s
}
