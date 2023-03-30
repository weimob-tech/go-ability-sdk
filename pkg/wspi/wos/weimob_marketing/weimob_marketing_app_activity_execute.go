package weimob_marketing

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobMarketingAppActivityExecuteService struct {
	handler spi.WosInvocationHandler[PaasWeimobMarketingAppActivityExecuteRequest, PaasWeimobMarketingAppActivityExecuteResponse]
}

func (s PaasWeimobMarketingAppActivityExecuteService) NewRequest() spi.InvocationRequest[PaasWeimobMarketingAppActivityExecuteRequest] {
	return &spi.WosInvocationRequest[PaasWeimobMarketingAppActivityExecuteRequest]{
		Params: &PaasWeimobMarketingAppActivityExecuteRequest{},
	}
}

func (s PaasWeimobMarketingAppActivityExecuteService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobMarketingAppActivityExecuteRequest],
) (
	spi.InvocationResponse[PaasWeimobMarketingAppActivityExecuteResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobMarketingAppActivityExecuteRequest]))
}

type PaasWeimobMarketingAppActivityExecuteRequest struct {
	ActivityId        int64  `json:"activityId,omitempty"`
	BosId             int64  `json:"bosId,omitempty"`
	ProductId         int64  `json:"productId,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	CurrentBosId      int64  `json:"currentBosId,omitempty"`
	CurrentVid        int64  `json:"currentVid,omitempty"`
	Wid               int64  `json:"wid,omitempty"`
	Action            string `json:"action,omitempty"`
	Id                int64  `json:"id,omitempty"`
}

type PaasWeimobMarketingAppActivityExecuteResponse struct {
}

func CreatePaasWeimobMarketingAppActivityExecuteResponse() *PaasWeimobMarketingAppActivityExecuteResponse {
	return &PaasWeimobMarketingAppActivityExecuteResponse{}
}

// OnPaasWeimobMarketingAppActivityExecuteServiceInvocation
// @id 940
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/940?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=操作外部活动)
func (s *Service) OnPaasWeimobMarketingAppActivityExecuteServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobMarketingAppActivityExecuteRequest, PaasWeimobMarketingAppActivityExecuteResponse],
) (service *Service) {
	var invocation = &PaasWeimobMarketingAppActivityExecuteService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobMarketingAppActivityExecuteRequest, PaasWeimobMarketingAppActivityExecuteResponse](invocation),
		"PaasWeimobMarketingAppActivityExecuteService",
		"weimobMarketing.app.activity.execute",
	)
	return s
}
