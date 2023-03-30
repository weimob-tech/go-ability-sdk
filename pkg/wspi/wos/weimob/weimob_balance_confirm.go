package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobBalanceConfirmService struct {
	handler spi.WosInvocationHandler[PaasWeimobBalanceConfirmRequest, PaasWeimobBalanceConfirmBool]
}

func (s PaasWeimobBalanceConfirmService) NewRequest() spi.InvocationRequest[PaasWeimobBalanceConfirmRequest] {
	return &spi.WosInvocationRequest[PaasWeimobBalanceConfirmRequest]{
		Params: &PaasWeimobBalanceConfirmRequest{},
	}
}

func (s PaasWeimobBalanceConfirmService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobBalanceConfirmRequest],
) (
	spi.InvocationResponse[PaasWeimobBalanceConfirmBool],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobBalanceConfirmRequest]))
}

type PaasWeimobBalanceConfirmRequest struct {
	RequestId         string `json:"requestId,omitempty"`
	RequestType       string `json:"requestType,omitempty"`
	UserTransId       int64  `json:"userTransId,omitempty"`
	ChangeType        string `json:"changeType,omitempty"`
	Remark            string `json:"remark,omitempty"`
	OccurVid          int64  `json:"occurVid,omitempty"`
	Amount            string `json:"amount,omitempty"`
	OutTransNo        string `json:"outTransNo,omitempty"`
	BosId             int64  `json:"bosId,omitempty"`
	Uid               int64  `json:"uid,omitempty"`
	Vid               int64  `json:"vid,omitempty"`
	OperatorWid       int64  `json:"operatorWid,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	PlanId            int64  `json:"planId,omitempty"`
}

type PaasWeimobBalanceConfirmBool bool

func CreatePaasWeimobBalanceConfirmBool() *PaasWeimobBalanceConfirmBool {
	var br PaasWeimobBalanceConfirmBool
	return &br
}

// OnPaasWeimobBalanceConfirmServiceInvocation
// @id 1115
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1115?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=余额锁定确认)
func (s *Service) OnPaasWeimobBalanceConfirmServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobBalanceConfirmRequest, PaasWeimobBalanceConfirmBool],
) (service *Service) {
	var invocation = &PaasWeimobBalanceConfirmService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobBalanceConfirmRequest, PaasWeimobBalanceConfirmBool](invocation),
		"PaasWeimobBalanceConfirmService",
		"weimob.balance.confirm",
	)
	return s
}
