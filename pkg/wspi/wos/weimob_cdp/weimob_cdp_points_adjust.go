package weimob_cdp

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCdpPointsAdjustService struct {
	handler spi.WosInvocationHandler[PaasWeimobCdpPointsAdjustRequest, PaasWeimobCdpPointsAdjustResponse]
}

func (s PaasWeimobCdpPointsAdjustService) NewRequest() spi.InvocationRequest[PaasWeimobCdpPointsAdjustRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCdpPointsAdjustRequest]{
		Params: &PaasWeimobCdpPointsAdjustRequest{},
	}
}

func (s PaasWeimobCdpPointsAdjustService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCdpPointsAdjustRequest],
) (
	spi.InvocationResponse[PaasWeimobCdpPointsAdjustResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCdpPointsAdjustRequest]))
}

type PaasWeimobCdpPointsAdjustRequest struct {
	OneCrmAdjustInfoList []PaasWeimobCdpPointsAdjustRequestOneCrmAdjustInfoList `json:"oneCrmAdjustInfoList,omitempty"`
	RequestId            string                                                 `json:"requestId,omitempty"`
	RequestType          string                                                 `json:"requestType,omitempty"`
	ChangePoint          int64                                                  `json:"changePoint,omitempty"`
	AdjustType           string                                                 `json:"adjustType,omitempty"`
	ChangeReason         string                                                 `json:"changeReason,omitempty"`
	OperatorWid          int64                                                  `json:"operatorWid,omitempty"`
	OperateStoreVid      int64                                                  `json:"operateStoreVid,omitempty"`
	SourceVid            int64                                                  `json:"sourceVid,omitempty"`
	ChangeType           string                                                 `json:"changeType,omitempty"`
}
type PaasWeimobCdpPointsAdjustRequestOneCrmAdjustInfoList struct {
	Vid               int64  `json:"vid,omitempty"`
	UserType          int64  `json:"userType,omitempty"`
	PointPlanId       int64  `json:"pointPlanId,omitempty"`
	Wid               int64  `json:"wid,omitempty"`
	BosId             int64  `json:"bosId,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	Tcode             string `json:"tcode,omitempty"`
}

type PaasWeimobCdpPointsAdjustResponse struct {
	List   []PaasWeimobCdpPointsAdjustResponselist `json:"list,omitempty"`
	Result bool                                    `json:"result,omitempty"`
}
type PaasWeimobCdpPointsAdjustResponselist struct {
	Status  int64  `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Wid     int64  `json:"wid,omitempty"`
	WidName string `json:"widName,omitempty"`
}

func CreatePaasWeimobCdpPointsAdjustResponse() *PaasWeimobCdpPointsAdjustResponse {
	return &PaasWeimobCdpPointsAdjustResponse{}
}

// OnPaasWeimobCdpPointsAdjustServiceInvocation
// @id 886
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/886?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=赠送积分（治理）)
func (s *Service) OnPaasWeimobCdpPointsAdjustServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCdpPointsAdjustRequest, PaasWeimobCdpPointsAdjustResponse],
) (service *Service) {
	var invocation = &PaasWeimobCdpPointsAdjustService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCdpPointsAdjustRequest, PaasWeimobCdpPointsAdjustResponse](invocation),
		"PaasWeimobCdpPointsAdjustService",
		"weimobCdp.points.adjust",
	)
	return s
}
