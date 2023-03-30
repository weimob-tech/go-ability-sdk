package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobBalanceDetailGetService struct {
	handler spi.WosInvocationHandler[PaasWeimobBalanceDetailGetRequest, PaasWeimobBalanceDetailGetResponse]
}

func (s PaasWeimobBalanceDetailGetService) NewRequest() spi.InvocationRequest[PaasWeimobBalanceDetailGetRequest] {
	return &spi.WosInvocationRequest[PaasWeimobBalanceDetailGetRequest]{
		Params: &PaasWeimobBalanceDetailGetRequest{},
	}
}

func (s PaasWeimobBalanceDetailGetService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobBalanceDetailGetRequest],
) (
	spi.InvocationResponse[PaasWeimobBalanceDetailGetResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobBalanceDetailGetRequest]))
}

type PaasWeimobBalanceDetailGetRequest struct {
	BosId             int64 `json:"bosId,omitempty"`
	Uid               int64 `json:"uid,omitempty"`
	Vid               int64 `json:"vid,omitempty"`
	ProductInstanceId int64 `json:"productInstanceId,omitempty"`
	IsStatistics      bool  `json:"isStatistics,omitempty"`
}

type PaasWeimobBalanceDetailGetResponse struct {
	AvailableAmount    string `json:"availableAmount,omitempty"`
	TotalAmount        string `json:"totalAmount,omitempty"`
	FrozenAmount       string `json:"frozenAmount,omitempty"`
	TotalIssueAmount   string `json:"totalIssueAmount,omitempty"`
	TotalCapitalAmount string `json:"totalCapitalAmount,omitempty"`
	TotalGrantsAmount  string `json:"totalGrantsAmount,omitempty"`
	RechargeTimes      int64  `json:"rechargeTimes,omitempty"`
}

func CreatePaasWeimobBalanceDetailGetResponse() *PaasWeimobBalanceDetailGetResponse {
	return &PaasWeimobBalanceDetailGetResponse{}
}

// OnPaasWeimobBalanceDetailGetServiceInvocation
// @id 1105
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1105?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询用户余额详情)
func (s *Service) OnPaasWeimobBalanceDetailGetServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobBalanceDetailGetRequest, PaasWeimobBalanceDetailGetResponse],
) (service *Service) {
	var invocation = &PaasWeimobBalanceDetailGetService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobBalanceDetailGetRequest, PaasWeimobBalanceDetailGetResponse](invocation),
		"PaasWeimobBalanceDetailGetService",
		"weimob.balance.detail.get",
	)
	return s
}
