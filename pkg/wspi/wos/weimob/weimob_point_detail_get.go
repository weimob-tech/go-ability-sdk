package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobPointDetailGetService struct {
	handler spi.WosInvocationHandler[PaasWeimobPointDetailGetRequest, PaasWeimobPointDetailGetResponse]
}

func (s PaasWeimobPointDetailGetService) NewRequest() spi.InvocationRequest[PaasWeimobPointDetailGetRequest] {
	return &spi.WosInvocationRequest[PaasWeimobPointDetailGetRequest]{
		Params: &PaasWeimobPointDetailGetRequest{},
	}
}

func (s PaasWeimobPointDetailGetService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobPointDetailGetRequest],
) (
	spi.InvocationResponse[PaasWeimobPointDetailGetResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobPointDetailGetRequest]))
}

type PaasWeimobPointDetailGetRequest struct {
	IsStatistics      bool  `json:"isStatistics,omitempty"`
	Vid               int64 `json:"vid,omitempty"`
	PointPlanId       int64 `json:"pointPlanId,omitempty"`
	Wid               int64 `json:"wid,omitempty"`
	BosId             int64 `json:"bosId,omitempty"`
	ProductInstanceId int64 `json:"productInstanceId,omitempty"`
}

type PaasWeimobPointDetailGetResponse struct {
	AcctId          int64  `json:"acctId,omitempty"`
	AvailablePoint  string `json:"availablePoint,omitempty"`
	TotalPoint      string `json:"totalPoint,omitempty"`
	FrozenPoint     string `json:"frozenPoint,omitempty"`
	TotalIssuePoint string `json:"totalIssuePoint,omitempty"`
}

func CreatePaasWeimobPointDetailGetResponse() *PaasWeimobPointDetailGetResponse {
	return &PaasWeimobPointDetailGetResponse{}
}

// OnPaasWeimobPointDetailGetServiceInvocation
// @id 1102
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1102?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询用户积分详情)
func (s *Service) OnPaasWeimobPointDetailGetServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobPointDetailGetRequest, PaasWeimobPointDetailGetResponse],
) (service *Service) {
	var invocation = &PaasWeimobPointDetailGetService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobPointDetailGetRequest, PaasWeimobPointDetailGetResponse](invocation),
		"PaasWeimobPointDetailGetService",
		"weimob.point.detail.get",
	)
	return s
}
