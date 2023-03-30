package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobPointRefundService struct {
	handler spi.WosInvocationHandler[PaasWeimobPointRefundRequest, PaasWeimobPointRefundResponse]
}

func (s PaasWeimobPointRefundService) NewRequest() spi.InvocationRequest[PaasWeimobPointRefundRequest] {
	return &spi.WosInvocationRequest[PaasWeimobPointRefundRequest]{
		Params: &PaasWeimobPointRefundRequest{},
	}
}

func (s PaasWeimobPointRefundService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobPointRefundRequest],
) (
	spi.InvocationResponse[PaasWeimobPointRefundResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobPointRefundRequest]))
}

type PaasWeimobPointRefundRequest struct {
	RefundRequestId   string `json:"refundRequestId,omitempty"`
	RefundRequestType string `json:"refundRequestType,omitempty"`
	Point             string `json:"point,omitempty"`
	UserTransId       int64  `json:"userTransId,omitempty"`
	RequestId         string `json:"requestId,omitempty"`
	RequestType       string `json:"requestType,omitempty"`
	OutTransNo        string `json:"outTransNo,omitempty"`
	ChangeType        string `json:"changeType,omitempty"`
	Remark            string `json:"remark,omitempty"`
	OccurVid          int64  `json:"occurVid,omitempty"`
	Vid               int64  `json:"vid,omitempty"`
	OperatorWid       int64  `json:"operatorWid,omitempty"`
	UserType          int64  `json:"userType,omitempty"`
	ChannelType       string `json:"channelType,omitempty"`
	PointPlanId       int64  `json:"pointPlanId,omitempty"`
	Wid               int64  `json:"wid,omitempty"`
	BosId             int64  `json:"bosId,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	Tcode             string `json:"tcode,omitempty"`
}

type PaasWeimobPointRefundResponse struct {
	RefundTransId int64 `json:"refundTransId,omitempty"`
}

func CreatePaasWeimobPointRefundResponse() *PaasWeimobPointRefundResponse {
	return &PaasWeimobPointRefundResponse{}
}

// OnPaasWeimobPointRefundServiceInvocation
// @id 1112
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1112?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=回退积分)
func (s *Service) OnPaasWeimobPointRefundServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobPointRefundRequest, PaasWeimobPointRefundResponse],
) (service *Service) {
	var invocation = &PaasWeimobPointRefundService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobPointRefundRequest, PaasWeimobPointRefundResponse](invocation),
		"PaasWeimobPointRefundService",
		"weimob.point.refund",
	)
	return s
}
