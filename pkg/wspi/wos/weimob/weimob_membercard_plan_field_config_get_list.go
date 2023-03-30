package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobMembercardPlanFieldConfigGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobMembercardPlanFieldConfigGetListRequest, PaasWeimobMembercardPlanFieldConfigGetListResponse]
}

func (s PaasWeimobMembercardPlanFieldConfigGetListService) NewRequest() spi.InvocationRequest[PaasWeimobMembercardPlanFieldConfigGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobMembercardPlanFieldConfigGetListRequest]{
		Params: &PaasWeimobMembercardPlanFieldConfigGetListRequest{},
	}
}

func (s PaasWeimobMembercardPlanFieldConfigGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobMembercardPlanFieldConfigGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobMembercardPlanFieldConfigGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobMembercardPlanFieldConfigGetListRequest]))
}

type PaasWeimobMembercardPlanFieldConfigGetListRequest struct {
	MembershipPlanId int64 `json:"membershipPlanId,omitempty"`
}

type PaasWeimobMembercardPlanFieldConfigGetListResponse struct {
	ChannelFields []PaasWeimobMembercardPlanFieldConfigGetListResponseChannelFields `json:"channelFields,omitempty"`
}
type PaasWeimobMembercardPlanFieldConfigGetListResponseChannelFields struct {
	Fields []PaasWeimobMembercardPlanFieldConfigGetListResponseFields `json:"fields,omitempty"`
}
type PaasWeimobMembercardPlanFieldConfigGetListResponseFields struct {
	OptionList        []PaasWeimobMembercardPlanFieldConfigGetListResponseOptionList `json:"optionList,omitempty"`
	FieldType         int64                                                          `json:"fieldType,omitempty"`
	ValueType         int64                                                          `json:"valueType,omitempty"`
	Name              string                                                         `json:"name,omitempty"`
	ShowName          string                                                         `json:"showName,omitempty"`
	Tips              string                                                         `json:"tips,omitempty"`
	IsRequired        int64                                                          `json:"isRequired,omitempty"`
	IsValueModifiable int64                                                          `json:"isValueModifiable,omitempty"`
	Status            int64                                                          `json:"status,omitempty"`
	Sort              int64                                                          `json:"sort,omitempty"`
	Key               string                                                         `json:"key,omitempty"`
}
type PaasWeimobMembercardPlanFieldConfigGetListResponseOptionList struct {
	OptionName  string `json:"optionName,omitempty"`
	OptionValue string `json:"optionValue,omitempty"`
}

func CreatePaasWeimobMembercardPlanFieldConfigGetListResponse() *PaasWeimobMembercardPlanFieldConfigGetListResponse {
	return &PaasWeimobMembercardPlanFieldConfigGetListResponse{}
}

// OnPaasWeimobMembercardPlanFieldConfigGetListServiceInvocation
// @id 1099
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1099?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询会员方案开卡字段)
func (s *Service) OnPaasWeimobMembercardPlanFieldConfigGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobMembercardPlanFieldConfigGetListRequest, PaasWeimobMembercardPlanFieldConfigGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobMembercardPlanFieldConfigGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobMembercardPlanFieldConfigGetListRequest, PaasWeimobMembercardPlanFieldConfigGetListResponse](invocation),
		"PaasWeimobMembercardPlanFieldConfigGetListService",
		"weimob.membercard.plan.field.config.getList",
	)
	return s
}
