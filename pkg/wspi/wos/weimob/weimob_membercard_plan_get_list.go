package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobMembercardPlanGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobMembercardPlanGetListRequest, PaasWeimobMembercardPlanGetListResponse]
}

func (s PaasWeimobMembercardPlanGetListService) NewRequest() spi.InvocationRequest[PaasWeimobMembercardPlanGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobMembercardPlanGetListRequest]{
		Params: &PaasWeimobMembercardPlanGetListRequest{},
	}
}

func (s PaasWeimobMembercardPlanGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobMembercardPlanGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobMembercardPlanGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobMembercardPlanGetListRequest]))
}

type PaasWeimobMembercardPlanGetListRequest struct {
	MembershipPlanId int64   `json:"membershipPlanId,omitempty"`
	CardType         int64   `json:"cardType,omitempty"`
	QueryType        []int64 `json:"queryType,omitempty"`
}

type PaasWeimobMembercardPlanGetListResponse struct {
	LevelList         []PaasWeimobMembercardPlanGetListResponseLevelList `json:"levelList,omitempty"`
	MembershipPlanId  int64                                              `json:"membershipPlanId,omitempty"`
	Name              string                                             `json:"name,omitempty"`
	IsDynamicCode     int64                                              `json:"isDynamicCode,omitempty"`
	CardNumberSetting int64                                              `json:"cardNumberSetting,omitempty"`
	Prefix            string                                             `json:"prefix,omitempty"`
	Postfix           string                                             `json:"postfix,omitempty"`
	CreateType        int64                                              `json:"createType,omitempty"`
	Status            int64                                              `json:"status,omitempty"`
}
type PaasWeimobMembercardPlanGetListResponseLevelList struct {
	LevelId   int64  `json:"levelId,omitempty"`
	LevelName string `json:"levelName,omitempty"`
	RankDown  int64  `json:"rankDown,omitempty"`
	Sort      int64  `json:"sort,omitempty"`
}

func CreatePaasWeimobMembercardPlanGetListResponse() *PaasWeimobMembercardPlanGetListResponse {
	return &PaasWeimobMembercardPlanGetListResponse{}
}

// OnPaasWeimobMembercardPlanGetListServiceInvocation
// @id 1100
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1100?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询会员卡方案列表)
func (s *Service) OnPaasWeimobMembercardPlanGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobMembercardPlanGetListRequest, PaasWeimobMembercardPlanGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobMembercardPlanGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobMembercardPlanGetListRequest, PaasWeimobMembercardPlanGetListResponse](invocation),
		"PaasWeimobMembercardPlanGetListService",
		"weimob.membercard.plan.getList",
	)
	return s
}
