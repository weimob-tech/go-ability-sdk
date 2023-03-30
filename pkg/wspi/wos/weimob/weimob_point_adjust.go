package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobPointAdjustService struct {
	handler spi.WosInvocationHandler[PaasWeimobPointAdjustRequest, PaasWeimobPointAdjustResponse]
}

func (s PaasWeimobPointAdjustService) NewRequest() spi.InvocationRequest[PaasWeimobPointAdjustRequest] {
	return &spi.WosInvocationRequest[PaasWeimobPointAdjustRequest]{
		Params: &PaasWeimobPointAdjustRequest{},
	}
}

func (s PaasWeimobPointAdjustService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobPointAdjustRequest],
) (
	spi.InvocationResponse[PaasWeimobPointAdjustResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobPointAdjustRequest]))
}

type PaasWeimobPointAdjustRequest struct {
	AdjustInfoReqList []PaasWeimobPointAdjustRequestAdjustInfoReqList `json:"adjustInfoReqList,omitempty"`
	RequestId         string                                          `json:"requestId,omitempty"`
	RequestType       string                                          `json:"requestType,omitempty"`
	Point             string                                          `json:"point,omitempty"`
	AdjustType        string                                          `json:"adjustType,omitempty"`
	Remark            string                                          `json:"remark,omitempty"`
	OperatorWid       string                                          `json:"operatorWid,omitempty"`
	OperateStoreVid   int64                                           `json:"operateStoreVid,omitempty"`
	SourceVid         int64                                           `json:"sourceVid,omitempty"`
	OutTransNo        string                                          `json:"outTransNo,omitempty"`
	OutTransType      string                                          `json:"outTransType,omitempty"`
	ChangeType        string                                          `json:"changeType,omitempty"`
	EffectiveTime     int64                                           `json:"effectiveTime,omitempty"`
	ExpireTime        int64                                           `json:"expireTime,omitempty"`
	ChangeTime        int64                                           `json:"changeTime,omitempty"`
}
type PaasWeimobPointAdjustRequestAdjustInfoReqList struct {
	Vid               int64  `json:"vid,omitempty"`
	BosId             int64  `json:"bosId,omitempty"`
	Wid               int64  `json:"wid,omitempty"`
	WidName           string `json:"widName,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	PlanId            int64  `json:"planId,omitempty"`
}

type PaasWeimobPointAdjustResponse struct {
	List   []PaasWeimobPointAdjustResponselist `json:"list,omitempty"`
	Result bool                                `json:"result,omitempty"`
}
type PaasWeimobPointAdjustResponselist struct {
	Status  int64  `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Wid     int64  `json:"wid,omitempty"`
	WidName string `json:"widName,omitempty"`
}

func CreatePaasWeimobPointAdjustResponse() *PaasWeimobPointAdjustResponse {
	return &PaasWeimobPointAdjustResponse{}
}

// OnPaasWeimobPointAdjustServiceInvocation
// @id 1093
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1093?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=积分调账)
func (s *Service) OnPaasWeimobPointAdjustServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobPointAdjustRequest, PaasWeimobPointAdjustResponse],
) (service *Service) {
	var invocation = &PaasWeimobPointAdjustService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobPointAdjustRequest, PaasWeimobPointAdjustResponse](invocation),
		"PaasWeimobPointAdjustService",
		"weimob.point.adjust",
	)
	return s
}
