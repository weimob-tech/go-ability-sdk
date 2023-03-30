package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobBalanceAdjustService struct {
	handler spi.WosInvocationHandler[PaasWeimobBalanceAdjustRequest, PaasWeimobBalanceAdjustResponse]
}

func (s PaasWeimobBalanceAdjustService) NewRequest() spi.InvocationRequest[PaasWeimobBalanceAdjustRequest] {
	return &spi.WosInvocationRequest[PaasWeimobBalanceAdjustRequest]{
		Params: &PaasWeimobBalanceAdjustRequest{},
	}
}

func (s PaasWeimobBalanceAdjustService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobBalanceAdjustRequest],
) (
	spi.InvocationResponse[PaasWeimobBalanceAdjustResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobBalanceAdjustRequest]))
}

type PaasWeimobBalanceAdjustRequest struct {
	AdjustInfoReqList []PaasWeimobBalanceAdjustRequestAdjustInfoReqList `json:"adjustInfoReqList,omitempty"`
	RequestId         string                                            `json:"requestId,omitempty"`
	RequestType       string                                            `json:"requestType,omitempty"`
	AdjustType        string                                            `json:"adjustType,omitempty"`
	Remark            string                                            `json:"remark,omitempty"`
	OperatorWid       string                                            `json:"operatorWid,omitempty"`
	CapitalAmount     string                                            `json:"capitalAmount,omitempty"`
	GrantsAmount      string                                            `json:"grantsAmount,omitempty"`
	OperateStoreVid   int64                                             `json:"operateStoreVid,omitempty"`
	SourceVid         int64                                             `json:"sourceVid,omitempty"`
}
type PaasWeimobBalanceAdjustRequestAdjustInfoReqList struct {
	Vid               int64  `json:"vid,omitempty"`
	BosId             int64  `json:"bosId,omitempty"`
	Wid               int64  `json:"wid,omitempty"`
	WidName           string `json:"widName,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	PlanId            int64  `json:"planId,omitempty"`
}

type PaasWeimobBalanceAdjustResponse struct {
	List   []PaasWeimobBalanceAdjustResponselist `json:"list,omitempty"`
	Result bool                                  `json:"result,omitempty"`
}
type PaasWeimobBalanceAdjustResponselist struct {
	Status  int64  `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Wid     int64  `json:"wid,omitempty"`
	WidName string `json:"widName,omitempty"`
}

func CreatePaasWeimobBalanceAdjustResponse() *PaasWeimobBalanceAdjustResponse {
	return &PaasWeimobBalanceAdjustResponse{}
}

// OnPaasWeimobBalanceAdjustServiceInvocation
// @id 1094
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1094?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=余额调账)
func (s *Service) OnPaasWeimobBalanceAdjustServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobBalanceAdjustRequest, PaasWeimobBalanceAdjustResponse],
) (service *Service) {
	var invocation = &PaasWeimobBalanceAdjustService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobBalanceAdjustRequest, PaasWeimobBalanceAdjustResponse](invocation),
		"PaasWeimobBalanceAdjustService",
		"weimob.balance.adjust",
	)
	return s
}
