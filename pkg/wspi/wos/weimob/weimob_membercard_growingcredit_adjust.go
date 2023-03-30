package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobMembercardGrowingcreditAdjustService struct {
	handler spi.WosInvocationHandler[PaasWeimobMembercardGrowingcreditAdjustRequest, PaasWeimobMembercardGrowingcreditAdjustResponse]
}

func (s PaasWeimobMembercardGrowingcreditAdjustService) NewRequest() spi.InvocationRequest[PaasWeimobMembercardGrowingcreditAdjustRequest] {
	return &spi.WosInvocationRequest[PaasWeimobMembercardGrowingcreditAdjustRequest]{
		Params: &PaasWeimobMembercardGrowingcreditAdjustRequest{},
	}
}

func (s PaasWeimobMembercardGrowingcreditAdjustService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobMembercardGrowingcreditAdjustRequest],
) (
	spi.InvocationResponse[PaasWeimobMembercardGrowingcreditAdjustResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobMembercardGrowingcreditAdjustRequest]))
}

type PaasWeimobMembercardGrowingcreditAdjustRequest struct {
	SourceObjectList []PaasWeimobMembercardGrowingcreditAdjustRequestSourceObjectList `json:"sourceObjectList,omitempty"`
	Wid              int64                                                            `json:"wid,omitempty"`
	MembershipPlanId int64                                                            `json:"membershipPlanId,omitempty"`
	ChangeGrowth     int64                                                            `json:"changeGrowth,omitempty"`
	Reason           string                                                           `json:"reason,omitempty"`
	OperatorChannel  int64                                                            `json:"operatorChannel,omitempty"`
	OperatorWid      int64                                                            `json:"operatorWid,omitempty"`
	RequestId        string                                                           `json:"requestId,omitempty"`
	RequestType      int64                                                            `json:"requestType,omitempty"`
}
type PaasWeimobMembercardGrowingcreditAdjustRequestSourceObjectList struct {
	SourceOpenId string `json:"sourceOpenId,omitempty"`
	SourceAppId  string `json:"sourceAppId,omitempty"`
	Source       int64  `json:"source,omitempty"`
}

type PaasWeimobMembercardGrowingcreditAdjustResponse struct {
	GrowthBillId int64 `json:"growthBillId,omitempty"`
}

func CreatePaasWeimobMembercardGrowingcreditAdjustResponse() *PaasWeimobMembercardGrowingcreditAdjustResponse {
	return &PaasWeimobMembercardGrowingcreditAdjustResponse{}
}

// OnPaasWeimobMembercardGrowingcreditAdjustServiceInvocation
// @id 1097
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1097?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=调整会员卡成长值)
func (s *Service) OnPaasWeimobMembercardGrowingcreditAdjustServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobMembercardGrowingcreditAdjustRequest, PaasWeimobMembercardGrowingcreditAdjustResponse],
) (service *Service) {
	var invocation = &PaasWeimobMembercardGrowingcreditAdjustService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobMembercardGrowingcreditAdjustRequest, PaasWeimobMembercardGrowingcreditAdjustResponse](invocation),
		"PaasWeimobMembercardGrowingcreditAdjustService",
		"weimob.membercard.growingcredit.adjust",
	)
	return s
}
