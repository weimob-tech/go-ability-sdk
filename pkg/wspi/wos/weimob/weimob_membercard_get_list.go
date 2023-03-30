package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobMembercardGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobMembercardGetListRequest, PaasWeimobMembercardGetListResponse]
}

func (s PaasWeimobMembercardGetListService) NewRequest() spi.InvocationRequest[PaasWeimobMembercardGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobMembercardGetListRequest]{
		Params: &PaasWeimobMembercardGetListRequest{},
	}
}

func (s PaasWeimobMembercardGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobMembercardGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobMembercardGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobMembercardGetListRequest]))
}

type PaasWeimobMembercardGetListRequest struct {
	SourceObjectList []PaasWeimobMembercardGetListRequestSourceObjectList `json:"sourceObjectList,omitempty"`
	Wid              int64                                                `json:"wid,omitempty"`
}
type PaasWeimobMembercardGetListRequestSourceObjectList struct {
	SourceOpenId string `json:"sourceOpenId,omitempty"`
	SourceAppId  string `json:"sourceAppId,omitempty"`
	Source       int64  `json:"source,omitempty"`
}

type PaasWeimobMembercardGetListResponse struct {
	MembershipPlanId      int64  `json:"membershipPlanId,omitempty"`
	Vid                   int64  `json:"vid,omitempty"`
	MembershipCardChannel int64  `json:"membershipCardChannel,omitempty"`
	MembershipCardScene   int64  `json:"membershipCardScene,omitempty"`
	MembershipCardSceneId int64  `json:"membershipCardSceneId,omitempty"`
	CustomCardNo          string `json:"customCardNo,omitempty"`
	ObtainTime            int64  `json:"obtainTime,omitempty"`
	EffectiveTime         int64  `json:"effectiveTime,omitempty"`
	ExpireTime            int64  `json:"expireTime,omitempty"`
	Status                int64  `json:"status,omitempty"`
	LevelId               int64  `json:"levelId,omitempty"`
	GrowthValue           int64  `json:"growthValue,omitempty"`
}

func CreatePaasWeimobMembercardGetListResponse() *PaasWeimobMembercardGetListResponse {
	return &PaasWeimobMembercardGetListResponse{}
}

// OnPaasWeimobMembercardGetListServiceInvocation
// @id 1101
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1101?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取会员卡列表)
func (s *Service) OnPaasWeimobMembercardGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobMembercardGetListRequest, PaasWeimobMembercardGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobMembercardGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobMembercardGetListRequest, PaasWeimobMembercardGetListResponse](invocation),
		"PaasWeimobMembercardGetListService",
		"weimob.membercard.getList",
	)
	return s
}
