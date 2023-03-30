package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobBalanceUnfreezeService struct {
	handler spi.WosInvocationHandler[PaasWeimobBalanceUnfreezeRequest, PaasWeimobBalanceUnfreezeBool]
}

func (s PaasWeimobBalanceUnfreezeService) NewRequest() spi.InvocationRequest[PaasWeimobBalanceUnfreezeRequest] {
	return &spi.WosInvocationRequest[PaasWeimobBalanceUnfreezeRequest]{
		Params: &PaasWeimobBalanceUnfreezeRequest{},
	}
}

func (s PaasWeimobBalanceUnfreezeService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobBalanceUnfreezeRequest],
) (
	spi.InvocationResponse[PaasWeimobBalanceUnfreezeBool],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobBalanceUnfreezeRequest]))
}

type PaasWeimobBalanceUnfreezeRequest struct {
	RequestId         string `json:"requestId,omitempty"`
	RequestType       string `json:"requestType,omitempty"`
	UserTransId       int64  `json:"userTransId,omitempty"`
	OutTransNo        string `json:"outTransNo,omitempty"`
	OccurVid          int64  `json:"occurVid,omitempty"`
	Amount            string `json:"amount,omitempty"`
	Remark            string `json:"remark,omitempty"`
	ChangeType        string `json:"changeType,omitempty"`
	BosId             int64  `json:"bosId,omitempty"`
	Uid               int64  `json:"uid,omitempty"`
	Vid               int64  `json:"vid,omitempty"`
	OperatorWid       int64  `json:"operatorWid,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	PlanId            int64  `json:"planId,omitempty"`
}

type PaasWeimobBalanceUnfreezeBool bool

func CreatePaasWeimobBalanceUnfreezeBool() *PaasWeimobBalanceUnfreezeBool {
	var br PaasWeimobBalanceUnfreezeBool
	return &br
}

// OnPaasWeimobBalanceUnfreezeServiceInvocation
// @id 1118
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1118?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=解冻余额)
func (s *Service) OnPaasWeimobBalanceUnfreezeServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobBalanceUnfreezeRequest, PaasWeimobBalanceUnfreezeBool],
) (service *Service) {
	var invocation = &PaasWeimobBalanceUnfreezeService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobBalanceUnfreezeRequest, PaasWeimobBalanceUnfreezeBool](invocation),
		"PaasWeimobBalanceUnfreezeService",
		"weimob.balance.unfreeze",
	)
	return s
}
