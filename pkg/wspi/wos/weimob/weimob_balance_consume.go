package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobBalanceConsumeService struct {
	handler spi.WosInvocationHandler[PaasWeimobBalanceConsumeRequest, PaasWeimobBalanceConsumeResponse]
}

func (s PaasWeimobBalanceConsumeService) NewRequest() spi.InvocationRequest[PaasWeimobBalanceConsumeRequest] {
	return &spi.WosInvocationRequest[PaasWeimobBalanceConsumeRequest]{
		Params: &PaasWeimobBalanceConsumeRequest{},
	}
}

func (s PaasWeimobBalanceConsumeService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobBalanceConsumeRequest],
) (
	spi.InvocationResponse[PaasWeimobBalanceConsumeResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobBalanceConsumeRequest]))
}

type PaasWeimobBalanceConsumeRequest struct {
	ChangeType        string `json:"changeType,omitempty"`
	RequestId         string `json:"requestId,omitempty"`
	RequestType       string `json:"requestType,omitempty"`
	OutTransNo        string `json:"outTransNo,omitempty"`
	OutTransType      string `json:"outTransType,omitempty"`
	Amount            string `json:"amount,omitempty"`
	OccurVid          int64  `json:"occurVid,omitempty"`
	LockExpireTime    int64  `json:"lockExpireTime,omitempty"`
	Remark            string `json:"remark,omitempty"`
	BosId             int64  `json:"bosId,omitempty"`
	Uid               int64  `json:"uid,omitempty"`
	Vid               int64  `json:"vid,omitempty"`
	OperatorWid       int64  `json:"operatorWid,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	PlanId            int64  `json:"planId,omitempty"`
}

type PaasWeimobBalanceConsumeResponse struct {
	BosId       int64  `json:"bosId,omitempty"`
	Wid         int64  `json:"wid,omitempty"`
	Amount      string `json:"amount,omitempty"`
	UserTransId int64  `json:"userTransId,omitempty"`
}

func CreatePaasWeimobBalanceConsumeResponse() *PaasWeimobBalanceConsumeResponse {
	return &PaasWeimobBalanceConsumeResponse{}
}

// OnPaasWeimobBalanceConsumeServiceInvocation
// @id 1116
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1116?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=消耗余额)
func (s *Service) OnPaasWeimobBalanceConsumeServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobBalanceConsumeRequest, PaasWeimobBalanceConsumeResponse],
) (service *Service) {
	var invocation = &PaasWeimobBalanceConsumeService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobBalanceConsumeRequest, PaasWeimobBalanceConsumeResponse](invocation),
		"PaasWeimobBalanceConsumeService",
		"weimob.balance.consume",
	)
	return s
}
