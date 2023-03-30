package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobPointConfirmService struct {
	handler spi.WosInvocationHandler[PaasWeimobPointConfirmRequest, PaasWeimobPointConfirmBool]
}

func (s PaasWeimobPointConfirmService) NewRequest() spi.InvocationRequest[PaasWeimobPointConfirmRequest] {
	return &spi.WosInvocationRequest[PaasWeimobPointConfirmRequest]{
		Params: &PaasWeimobPointConfirmRequest{},
	}
}

func (s PaasWeimobPointConfirmService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobPointConfirmRequest],
) (
	spi.InvocationResponse[PaasWeimobPointConfirmBool],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobPointConfirmRequest]))
}

type PaasWeimobPointConfirmRequest struct {
	RequestId         string `json:"requestId,omitempty"`
	RequestType       string `json:"requestType,omitempty"`
	UserTransId       int64  `json:"userTransId,omitempty"`
	OutTransNo        string `json:"outTransNo,omitempty"`
	Point             string `json:"point,omitempty"`
	Remark            string `json:"remark,omitempty"`
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

type PaasWeimobPointConfirmBool bool

func CreatePaasWeimobPointConfirmBool() *PaasWeimobPointConfirmBool {
	var br PaasWeimobPointConfirmBool
	return &br
}

// OnPaasWeimobPointConfirmServiceInvocation
// @id 1110
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1110?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=积分锁定确认)
func (s *Service) OnPaasWeimobPointConfirmServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobPointConfirmRequest, PaasWeimobPointConfirmBool],
) (service *Service) {
	var invocation = &PaasWeimobPointConfirmService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobPointConfirmRequest, PaasWeimobPointConfirmBool](invocation),
		"PaasWeimobPointConfirmService",
		"weimob.point.confirm",
	)
	return s
}
