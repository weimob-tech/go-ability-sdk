package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobPointConsumeService struct {
	handler spi.WosInvocationHandler[PaasWeimobPointConsumeRequest, PaasWeimobPointConsumeResponse]
}

func (s PaasWeimobPointConsumeService) NewRequest() spi.InvocationRequest[PaasWeimobPointConsumeRequest] {
	return &spi.WosInvocationRequest[PaasWeimobPointConsumeRequest]{
		Params: &PaasWeimobPointConsumeRequest{},
	}
}

func (s PaasWeimobPointConsumeService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobPointConsumeRequest],
) (
	spi.InvocationResponse[PaasWeimobPointConsumeResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobPointConsumeRequest]))
}

type PaasWeimobPointConsumeRequest struct {
	RequestId         string `json:"requestId,omitempty"`
	RequestType       string `json:"requestType,omitempty"`
	Business          string `json:"business,omitempty"`
	OutTransNo        string `json:"outTransNo,omitempty"`
	OutTransType      string `json:"outTransType,omitempty"`
	Point             string `json:"point,omitempty"`
	ChangeType        string `json:"changeType,omitempty"`
	Remark            string `json:"remark,omitempty"`
	OccurVid          int64  `json:"occurVid,omitempty"`
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

type PaasWeimobPointConsumeResponse struct {
	BosId       int64  `json:"bosId,omitempty"`
	Wid         int64  `json:"wid,omitempty"`
	Point       int64  `json:"point,omitempty"`
	ChangeType  string `json:"changeType,omitempty"`
	UserTransId int64  `json:"userTransId,omitempty"`
}

func CreatePaasWeimobPointConsumeResponse() *PaasWeimobPointConsumeResponse {
	return &PaasWeimobPointConsumeResponse{}
}

// OnPaasWeimobPointConsumeServiceInvocation
// @id 1111
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1111?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=消耗积分)
func (s *Service) OnPaasWeimobPointConsumeServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobPointConsumeRequest, PaasWeimobPointConsumeResponse],
) (service *Service) {
	var invocation = &PaasWeimobPointConsumeService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobPointConsumeRequest, PaasWeimobPointConsumeResponse](invocation),
		"PaasWeimobPointConsumeService",
		"weimob.point.consume",
	)
	return s
}
