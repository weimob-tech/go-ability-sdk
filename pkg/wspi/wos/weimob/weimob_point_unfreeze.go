package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobPointUnfreezeService struct {
	handler spi.WosInvocationHandler[PaasWeimobPointUnfreezeRequest, PaasWeimobPointUnfreezeBool]
}

func (s PaasWeimobPointUnfreezeService) NewRequest() spi.InvocationRequest[PaasWeimobPointUnfreezeRequest] {
	return &spi.WosInvocationRequest[PaasWeimobPointUnfreezeRequest]{
		Params: &PaasWeimobPointUnfreezeRequest{},
	}
}

func (s PaasWeimobPointUnfreezeService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobPointUnfreezeRequest],
) (
	spi.InvocationResponse[PaasWeimobPointUnfreezeBool],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobPointUnfreezeRequest]))
}

type PaasWeimobPointUnfreezeRequest struct {
	RequestId         string `json:"requestId,omitempty"`
	RequestType       string `json:"requestType,omitempty"`
	UserTransId       int64  `json:"userTransId,omitempty"`
	OutTransNo        string `json:"outTransNo,omitempty"`
	OccurVid          int64  `json:"occurVid,omitempty"`
	ChangeType        string `json:"changeType,omitempty"`
	Remark            string `json:"remark,omitempty"`
	Point             string `json:"point,omitempty"`
	Vid               int64  `json:"vid,omitempty"`
	OperatorWid       int64  `json:"operatorWid,omitempty"`
	UserType          int64  `json:"userType,omitempty"`
	ChannelType       int64  `json:"channelType,omitempty"`
	PointPlanId       int64  `json:"pointPlanId,omitempty"`
	Wid               int64  `json:"wid,omitempty"`
	BosId             int64  `json:"bosId,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	Tcode             string `json:"tcode,omitempty"`
}

type PaasWeimobPointUnfreezeBool bool

func CreatePaasWeimobPointUnfreezeBool() *PaasWeimobPointUnfreezeBool {
	var br PaasWeimobPointUnfreezeBool
	return &br
}

// OnPaasWeimobPointUnfreezeServiceInvocation
// @id 1113
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1113?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=解冻积分)
func (s *Service) OnPaasWeimobPointUnfreezeServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobPointUnfreezeRequest, PaasWeimobPointUnfreezeBool],
) (service *Service) {
	var invocation = &PaasWeimobPointUnfreezeService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobPointUnfreezeRequest, PaasWeimobPointUnfreezeBool](invocation),
		"PaasWeimobPointUnfreezeService",
		"weimob.point.unfreeze",
	)
	return s
}
