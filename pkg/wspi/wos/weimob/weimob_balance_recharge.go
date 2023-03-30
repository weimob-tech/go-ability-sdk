package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobBalanceRechargeService struct {
	handler spi.WosInvocationHandler[PaasWeimobBalanceRechargeRequest, PaasWeimobBalanceRechargeResponse]
}

func (s PaasWeimobBalanceRechargeService) NewRequest() spi.InvocationRequest[PaasWeimobBalanceRechargeRequest] {
	return &spi.WosInvocationRequest[PaasWeimobBalanceRechargeRequest]{
		Params: &PaasWeimobBalanceRechargeRequest{},
	}
}

func (s PaasWeimobBalanceRechargeService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobBalanceRechargeRequest],
) (
	spi.InvocationResponse[PaasWeimobBalanceRechargeResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobBalanceRechargeRequest]))
}

type PaasWeimobBalanceRechargeRequest struct {
	SettleRule        []PaasWeimobBalanceRechargeRequestSettleRule `json:"settleRule,omitempty"`
	RequestId         string                                       `json:"requestId,omitempty"`
	RequestType       string                                       `json:"requestType,omitempty"`
	OutTransNo        string                                       `json:"outTransNo,omitempty"`
	OutTransType      string                                       `json:"outTransType,omitempty"`
	CapitalAmount     string                                       `json:"capitalAmount,omitempty"`
	GrantsAmount      string                                       `json:"grantsAmount,omitempty"`
	ChangeType        string                                       `json:"changeType,omitempty"`
	OccurVid          int64                                        `json:"occurVid,omitempty"`
	Remark            string                                       `json:"remark,omitempty"`
	SourceVid         int64                                        `json:"sourceVid,omitempty"`
	BosId             int64                                        `json:"bosId,omitempty"`
	Uid               int64                                        `json:"uid,omitempty"`
	Vid               int64                                        `json:"vid,omitempty"`
	OperatorWid       int64                                        `json:"operatorWid,omitempty"`
	ProductInstanceId int64                                        `json:"productInstanceId,omitempty"`
	PlanId            int64                                        `json:"planId,omitempty"`
}
type PaasWeimobBalanceRechargeRequestSettleRule struct {
	Vid  int64 `json:"vid,omitempty"`
	Rate int64 `json:"rate,omitempty"`
}

type PaasWeimobBalanceRechargeResponse struct {
	UserTransNo string `json:"userTransNo,omitempty"`
	AfterAmount string `json:"afterAmount,omitempty"`
	Result      bool   `json:"result,omitempty"`
}

func CreatePaasWeimobBalanceRechargeResponse() *PaasWeimobBalanceRechargeResponse {
	return &PaasWeimobBalanceRechargeResponse{}
}

// OnPaasWeimobBalanceRechargeServiceInvocation
// @id 1095
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1095?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=余额充值)
func (s *Service) OnPaasWeimobBalanceRechargeServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobBalanceRechargeRequest, PaasWeimobBalanceRechargeResponse],
) (service *Service) {
	var invocation = &PaasWeimobBalanceRechargeService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobBalanceRechargeRequest, PaasWeimobBalanceRechargeResponse](invocation),
		"PaasWeimobBalanceRechargeService",
		"weimob.balance.recharge",
	)
	return s
}
