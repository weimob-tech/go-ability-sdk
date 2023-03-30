package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobBalanceFreezeService struct {
	handler spi.WosInvocationHandler[PaasWeimobBalanceFreezeRequest, PaasWeimobBalanceFreezeResponse]
}

func (s PaasWeimobBalanceFreezeService) NewRequest() spi.InvocationRequest[PaasWeimobBalanceFreezeRequest] {
	return &spi.WosInvocationRequest[PaasWeimobBalanceFreezeRequest]{
		Params: &PaasWeimobBalanceFreezeRequest{},
	}
}

func (s PaasWeimobBalanceFreezeService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobBalanceFreezeRequest],
) (
	spi.InvocationResponse[PaasWeimobBalanceFreezeResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobBalanceFreezeRequest]))
}

type PaasWeimobBalanceFreezeRequest struct {
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
	BalancePlanId     int64  `json:"balancePlanId,omitempty"`
}

type PaasWeimobBalanceFreezeResponse struct {
	BosId       int64 `json:"bosId,omitempty"`
	Wid         int64 `json:"wid,omitempty"`
	Amount      int64 `json:"amount,omitempty"`
	UserTransId int64 `json:"userTransId,omitempty"`
}

func CreatePaasWeimobBalanceFreezeResponse() *PaasWeimobBalanceFreezeResponse {
	return &PaasWeimobBalanceFreezeResponse{}
}

// OnPaasWeimobBalanceFreezeServiceInvocation
// @id 1119
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1119?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=冻结余额)
func (s *Service) OnPaasWeimobBalanceFreezeServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobBalanceFreezeRequest, PaasWeimobBalanceFreezeResponse],
) (service *Service) {
	var invocation = &PaasWeimobBalanceFreezeService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobBalanceFreezeRequest, PaasWeimobBalanceFreezeResponse](invocation),
		"PaasWeimobBalanceFreezeService",
		"weimob.balance.freeze",
	)
	return s
}
