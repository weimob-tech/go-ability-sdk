package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCustomerFieldConfigGetService struct {
	handler spi.WosInvocationHandler[PaasWeimobCustomerFieldConfigGetRequest, PaasWeimobCustomerFieldConfigGetResponse]
}

func (s PaasWeimobCustomerFieldConfigGetService) NewRequest() spi.InvocationRequest[PaasWeimobCustomerFieldConfigGetRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCustomerFieldConfigGetRequest]{
		Params: &PaasWeimobCustomerFieldConfigGetRequest{},
	}
}

func (s PaasWeimobCustomerFieldConfigGetService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCustomerFieldConfigGetRequest],
) (
	spi.InvocationResponse[PaasWeimobCustomerFieldConfigGetResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCustomerFieldConfigGetRequest]))
}

type PaasWeimobCustomerFieldConfigGetRequest struct {
}

type PaasWeimobCustomerFieldConfigGetResponse struct {
	GroupList []PaasWeimobCustomerFieldConfigGetResponseGroupList `json:"groupList,omitempty"`
}
type PaasWeimobCustomerFieldConfigGetResponseGroupList struct {
	FieldList     []PaasWeimobCustomerFieldConfigGetResponseFieldList `json:"fieldList,omitempty"`
	GroupId       int64                                               `json:"groupId,omitempty"`
	GroupKey      string                                              `json:"groupKey,omitempty"`
	GroupName     string                                              `json:"groupName,omitempty"`
	GroupNum      int64                                               `json:"groupNum,omitempty"`
	GroupType     int64                                               `json:"groupType,omitempty"`
	IsFieldRepeat int64                                               `json:"isFieldRepeat,omitempty"`
	MaxLimitGroup int64                                               `json:"maxLimitGroup,omitempty"`
	Sort          int64                                               `json:"sort,omitempty"`
}
type PaasWeimobCustomerFieldConfigGetResponseFieldList struct {
	OptionList        []PaasWeimobCustomerFieldConfigGetResponseOptionList `json:"optionList,omitempty"`
	RuleList          []PaasWeimobCustomerFieldConfigGetResponseRuleList   `json:"ruleList,omitempty"`
	FieldId           int64                                                `json:"fieldId,omitempty"`
	FieldType         int64                                                `json:"fieldType,omitempty"`
	GroupId           int64                                                `json:"groupId,omitempty"`
	IsChange          bool                                                 `json:"isChange,omitempty"`
	IsRequired        int64                                                `json:"isRequired,omitempty"`
	IsValueModifiable int64                                                `json:"isValueModifiable,omitempty"`
	IsValueUnique     int64                                                `json:"isValueUnique,omitempty"`
	Key               string                                               `json:"key,omitempty"`
	Name              string                                               `json:"name,omitempty"`
	Sort              int64                                                `json:"sort,omitempty"`
	Status            int64                                                `json:"status,omitempty"`
	Tips              string                                               `json:"tips,omitempty"`
	Value             string                                               `json:"value,omitempty"`
	ValueType         int64                                                `json:"valueType,omitempty"`
}
type PaasWeimobCustomerFieldConfigGetResponseOptionList struct {
	FieldId     int64  `json:"fieldId,omitempty"`
	OptionName  string `json:"optionName,omitempty"`
	OptionValue string `json:"optionValue,omitempty"`
}
type PaasWeimobCustomerFieldConfigGetResponseRuleList struct {
	LimitType PaasWeimobCustomerFieldConfigGetResponseLimitType `json:"limitType,omitempty"`
	Max       int64                                             `json:"max,omitempty"`
	Min       int64                                             `json:"min,omitempty"`
	RuleType  int64                                             `json:"ruleType,omitempty"`
	FieldId   int64                                             `json:"fieldId,omitempty"`
}
type PaasWeimobCustomerFieldConfigGetResponseLimitType struct {
}

func CreatePaasWeimobCustomerFieldConfigGetResponse() *PaasWeimobCustomerFieldConfigGetResponse {
	return &PaasWeimobCustomerFieldConfigGetResponse{}
}

// OnPaasWeimobCustomerFieldConfigGetServiceInvocation
// @id 1091
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1091?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询客户资料字段设置)
func (s *Service) OnPaasWeimobCustomerFieldConfigGetServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCustomerFieldConfigGetRequest, PaasWeimobCustomerFieldConfigGetResponse],
) (service *Service) {
	var invocation = &PaasWeimobCustomerFieldConfigGetService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCustomerFieldConfigGetRequest, PaasWeimobCustomerFieldConfigGetResponse](invocation),
		"PaasWeimobCustomerFieldConfigGetService",
		"weimob.customer.field.config.get",
	)
	return s
}
