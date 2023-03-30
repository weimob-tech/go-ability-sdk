package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerFieldConfigUpdate
// @id 1746
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1746?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新客户资料字段设置)
func (client *Client) CustomerFieldConfigUpdate(request *CustomerFieldConfigUpdateRequest) (response *CustomerFieldConfigUpdateResponse, err error) {
	rpcResponse := CreateCustomerFieldConfigUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerFieldConfigUpdateRequest struct {
	*api.BaseRequest

	GroupList []CustomerFieldConfigUpdateRequestGroupList `json:"groupList,omitempty"`
	Vid       int64                                       `json:"vid,omitempty"`
}

type CustomerFieldConfigUpdateResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateCustomerFieldConfigUpdateRequest() (request *CustomerFieldConfigUpdateRequest) {
	request = &CustomerFieldConfigUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "customer/field/config/update", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerFieldConfigUpdateResponse() (response *api.BaseResponse[CustomerFieldConfigUpdateResponse]) {
	response = api.CreateResponse[CustomerFieldConfigUpdateResponse](&CustomerFieldConfigUpdateResponse{})
	return
}

type CustomerFieldConfigUpdateRequestGroupList struct {
	FieldList     []CustomerFieldConfigUpdateRequestFieldList `json:"fieldList,omitempty"`
	GroupId       int64                                       `json:"groupId,omitempty"`
	GroupName     string                                      `json:"groupName,omitempty"`
	GroupKey      string                                      `json:"groupKey,omitempty"`
	GroupType     int64                                       `json:"groupType,omitempty"`
	Sort          int64                                       `json:"sort,omitempty"`
	IsFieldRepeat int64                                       `json:"isFieldRepeat,omitempty"`
	MaxLimitGroup int64                                       `json:"maxLimitGroup,omitempty"`
	GroupNum      int64                                       `json:"groupNum,omitempty"`
}

type CustomerFieldConfigUpdateRequestFieldList struct {
	OptionList        []CustomerFieldConfigUpdateRequestOptionList `json:"optionList,omitempty"`
	RuleList          []CustomerFieldConfigUpdateRequestRuleList   `json:"ruleList,omitempty"`
	GroupId           int64                                        `json:"groupId,omitempty"`
	FieldId           int64                                        `json:"fieldId,omitempty"`
	FieldType         int64                                        `json:"fieldType,omitempty"`
	Name              string                                       `json:"name,omitempty"`
	Key               string                                       `json:"key,omitempty"`
	ValueType         int64                                        `json:"valueType,omitempty"`
	Tips              string                                       `json:"tips,omitempty"`
	IsRequired        int64                                        `json:"isRequired,omitempty"`
	Sort              int64                                        `json:"sort,omitempty"`
	Status            int64                                        `json:"status,omitempty"`
	IsValueUnique     int64                                        `json:"isValueUnique,omitempty"`
	IsValueModifiable int64                                        `json:"isValueModifiable,omitempty"`
	DefaultValue      string                                       `json:"defaultValue,omitempty"`
	Value             string                                       `json:"value,omitempty"`
	IsChange          bool                                         `json:"isChange,omitempty"`
}

type CustomerFieldConfigUpdateRequestOptionList struct {
	FieldId     int64  `json:"fieldId,omitempty"`
	OptionName  string `json:"optionName,omitempty"`
	OptionValue string `json:"optionValue,omitempty"`
}

type CustomerFieldConfigUpdateRequestRuleList struct {
	FieldId   int64  `json:"fieldId,omitempty"`
	RuleType  int64  `json:"ruleType,omitempty"`
	Max       int64  `json:"max,omitempty"`
	Min       string `json:"min,omitempty"`
	LimitType int64  `json:"limitType,omitempty"`
}
