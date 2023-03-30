package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobBalanceRefundService struct {
	handler spi.WosInvocationHandler[PaasWeimobBalanceRefundRequest, PaasWeimobBalanceRefundResponse]
}

func (s PaasWeimobBalanceRefundService) NewRequest() spi.InvocationRequest[PaasWeimobBalanceRefundRequest] {
	return &spi.WosInvocationRequest[PaasWeimobBalanceRefundRequest]{
		Params: &PaasWeimobBalanceRefundRequest{},
	}
}

func (s PaasWeimobBalanceRefundService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobBalanceRefundRequest],
) (
	spi.InvocationResponse[PaasWeimobBalanceRefundResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobBalanceRefundRequest]))
}

type PaasWeimobBalanceRefundRequest struct {
	RefundRequestId   string `json:"refundRequestId,omitempty"`
	RefundRequestType string `json:"refundRequestType,omitempty"`
	RequestId         string `json:"requestId,omitempty"`
	RequestType       string `json:"requestType,omitempty"`
	UserTransId       int64  `json:"userTransId,omitempty"`
	OutTransNo        string `json:"outTransNo,omitempty"`
	Amount            string `json:"amount,omitempty"`
	ChangeType        string `json:"changeType,omitempty"`
	Remark            string `json:"remark,omitempty"`
	OccurVid          int64  `json:"occurVid,omitempty"`
	BosId             int64  `json:"bosId,omitempty"`
	Uid               int64  `json:"uid,omitempty"`
	Vid               int64  `json:"vid,omitempty"`
	OperatorWid       int64  `json:"operatorWid,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	ChannelType       int64  `json:"channelType,omitempty"`
	PlanId            int64  `json:"planId,omitempty"`
}

type PaasWeimobBalanceRefundResponse struct {
	RefundTransId int64 `json:"refundTransId,omitempty"`
}

func CreatePaasWeimobBalanceRefundResponse() *PaasWeimobBalanceRefundResponse {
	return &PaasWeimobBalanceRefundResponse{}
}

// OnPaasWeimobBalanceRefundServiceInvocation
// @id 1117
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1117?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=回退余额)
func (s *Service) OnPaasWeimobBalanceRefundServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobBalanceRefundRequest, PaasWeimobBalanceRefundResponse],
) (service *Service) {
	var invocation = &PaasWeimobBalanceRefundService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobBalanceRefundRequest, PaasWeimobBalanceRefundResponse](invocation),
		"PaasWeimobBalanceRefundService",
		"weimob.balance.refund",
	)
	return s
}
