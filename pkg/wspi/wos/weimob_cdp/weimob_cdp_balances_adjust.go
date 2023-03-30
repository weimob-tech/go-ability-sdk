package weimob_cdp

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCdpBalancesAdjustService struct {
	handler spi.WosInvocationHandler[PaasWeimobCdpBalancesAdjustRequest, PaasWeimobCdpBalancesAdjustResponse]
}

func (s PaasWeimobCdpBalancesAdjustService) NewRequest() spi.InvocationRequest[PaasWeimobCdpBalancesAdjustRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCdpBalancesAdjustRequest]{
		Params: &PaasWeimobCdpBalancesAdjustRequest{},
	}
}

func (s PaasWeimobCdpBalancesAdjustService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCdpBalancesAdjustRequest],
) (
	spi.InvocationResponse[PaasWeimobCdpBalancesAdjustResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCdpBalancesAdjustRequest]))
}

type PaasWeimobCdpBalancesAdjustRequest struct {
	OneCrmAdjustInfoList []PaasWeimobCdpBalancesAdjustRequestOneCrmAdjustInfoList `json:"oneCrmAdjustInfoList,omitempty"`
	SettleRule           PaasWeimobCdpBalancesAdjustRequestSettleRule             `json:"settleRule,omitempty"`
	RequestId            string                                                   `json:"requestId,omitempty"`
	RequestType          string                                                   `json:"requestType,omitempty"`
	ChangCapitalAmount   int64                                                    `json:"changCapitalAmount,omitempty"`
	ChangeGrantsAmount   int64                                                    `json:"changeGrantsAmount,omitempty"`
	AdjustType           string                                                   `json:"adjustType,omitempty"`
	ChangeReason         string                                                   `json:"changeReason,omitempty"`
	SourceVid            int64                                                    `json:"sourceVid,omitempty"`
	BosId                int64                                                    `json:"bosId,omitempty"`
	ProductInstanceId    int64                                                    `json:"productInstanceId,omitempty"`
	OperatorWid          int64                                                    `json:"operatorWid,omitempty"`
	OperateStoreVid      int64                                                    `json:"operateStoreVid,omitempty"`
}
type PaasWeimobCdpBalancesAdjustRequestOneCrmAdjustInfoList struct {
	Vid               int64 `json:"vid,omitempty"`
	Wid               int64 `json:"wid,omitempty"`
	BosId             int64 `json:"bosId,omitempty"`
	ProductInstanceId int64 `json:"productInstanceId,omitempty"`
	BalancePlanId     int64 `json:"balancePlanId,omitempty"`
}
type PaasWeimobCdpBalancesAdjustRequestSettleRule struct {
	Vid  int64 `json:"vid,omitempty"`
	Rate int64 `json:"rate,omitempty"`
}

type PaasWeimobCdpBalancesAdjustResponse struct {
	List   []PaasWeimobCdpBalancesAdjustResponselist `json:"list,omitempty"`
	Result bool                                      `json:"result,omitempty"`
}
type PaasWeimobCdpBalancesAdjustResponselist struct {
	Status  int64  `json:"status,omitempty"`
	Wid     int64  `json:"wid,omitempty"`
	Message string `json:"message,omitempty"`
	WidName string `json:"widName,omitempty"`
}

func CreatePaasWeimobCdpBalancesAdjustResponse() *PaasWeimobCdpBalancesAdjustResponse {
	return &PaasWeimobCdpBalancesAdjustResponse{}
}

// OnPaasWeimobCdpBalancesAdjustServiceInvocation
// @id 885
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/885?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=余额发放（治理）)
func (s *Service) OnPaasWeimobCdpBalancesAdjustServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCdpBalancesAdjustRequest, PaasWeimobCdpBalancesAdjustResponse],
) (service *Service) {
	var invocation = &PaasWeimobCdpBalancesAdjustService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCdpBalancesAdjustRequest, PaasWeimobCdpBalancesAdjustResponse](invocation),
		"PaasWeimobCdpBalancesAdjustService",
		"weimobCdp.balances.adjust",
	)
	return s
}
