package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobPointFreezeService struct {
	handler spi.WosInvocationHandler[PaasWeimobPointFreezeRequest, PaasWeimobPointFreezeResponse]
}

func (s PaasWeimobPointFreezeService) NewRequest() spi.InvocationRequest[PaasWeimobPointFreezeRequest] {
	return &spi.WosInvocationRequest[PaasWeimobPointFreezeRequest]{
		Params: &PaasWeimobPointFreezeRequest{},
	}
}

func (s PaasWeimobPointFreezeService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobPointFreezeRequest],
) (
	spi.InvocationResponse[PaasWeimobPointFreezeResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobPointFreezeRequest]))
}

type PaasWeimobPointFreezeRequest struct {
	ChangeType        string `json:"changeType,omitempty"`
	RequestId         string `json:"requestId,omitempty"`
	RequestType       string `json:"requestType,omitempty"`
	Business          string `json:"business,omitempty"`
	OutTransNo        string `json:"outTransNo,omitempty"`
	OutTransType      string `json:"outTransType,omitempty"`
	Point             string `json:"point,omitempty"`
	OccurVid          int64  `json:"occurVid,omitempty"`
	LockExpireTime    int64  `json:"lockExpireTime,omitempty"`
	ExtInfo           string `json:"extInfo,omitempty"`
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

type PaasWeimobPointFreezeResponse struct {
	BosId       int64  `json:"bosId,omitempty"`
	Wid         int64  `json:"wid,omitempty"`
	Point       string `json:"point,omitempty"`
	ChangeType  string `json:"changeType,omitempty"`
	UserTransId int64  `json:"userTransId,omitempty"`
}

func CreatePaasWeimobPointFreezeResponse() *PaasWeimobPointFreezeResponse {
	return &PaasWeimobPointFreezeResponse{}
}

// OnPaasWeimobPointFreezeServiceInvocation
// @id 1114
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1114?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=锁定积分)
func (s *Service) OnPaasWeimobPointFreezeServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobPointFreezeRequest, PaasWeimobPointFreezeResponse],
) (service *Service) {
	var invocation = &PaasWeimobPointFreezeService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobPointFreezeRequest, PaasWeimobPointFreezeResponse](invocation),
		"PaasWeimobPointFreezeService",
		"weimob.point.freeze",
	)
	return s
}
