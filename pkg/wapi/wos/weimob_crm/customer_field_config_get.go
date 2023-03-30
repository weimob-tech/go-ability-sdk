package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerFieldConfigGet
// @id 1508
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1508?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询客户资料字段设置)
func (client *Client) CustomerFieldConfigGet(request *CustomerFieldConfigGetRequest) (response *CustomerFieldConfigGetResponse, err error) {
	rpcResponse := CreateCustomerFieldConfigGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerFieldConfigGetRequest struct {
	*api.BaseRequest

	Vid int64 `json:"vid,omitempty"`
}

type CustomerFieldConfigGetResponse struct {
	GroupList []CustomerFieldConfigGetResponseGroupList `json:"groupList,omitempty"`
}

func CreateCustomerFieldConfigGetRequest() (request *CustomerFieldConfigGetRequest) {
	request = &CustomerFieldConfigGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "customer/field/config/get", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerFieldConfigGetResponse() (response *api.BaseResponse[CustomerFieldConfigGetResponse]) {
	response = api.CreateResponse[CustomerFieldConfigGetResponse](&CustomerFieldConfigGetResponse{})
	return
}

type CustomerFieldConfigGetResponseGroupList struct {
	FieldList     []CustomerFieldConfigGetResponseFieldList `json:"fieldList,omitempty"`
	GroupId       int64                                     `json:"groupId,omitempty"`
	GroupName     string                                    `json:"groupName,omitempty"`
	GroupKey      string                                    `json:"groupKey,omitempty"`
	GroupType     int64                                     `json:"groupType,omitempty"`
	GroupNum      int64                                     `json:"groupNum,omitempty"`
	Sort          int64                                     `json:"sort,omitempty"`
	IsFieldRepeat int64                                     `json:"isFieldRepeat,omitempty"`
	MaxLimitGroup int64                                     `json:"maxLimitGroup,omitempty"`
}

type CustomerFieldConfigGetResponseFieldList struct {
	OptionList        []CustomerFieldConfigGetResponseOptionList `json:"optionList,omitempty"`
	RuleList          []CustomerFieldConfigGetResponseRuleList   `json:"ruleList,omitempty"`
	GroupId           int64                                      `json:"groupId,omitempty"`
	FieldId           int64                                      `json:"fieldId,omitempty"`
	GroupNum          int64                                      `json:"groupNum,omitempty"`
	FieldType         int64                                      `json:"fieldType,omitempty"`
	Name              string                                     `json:"name,omitempty"`
	Key               string                                     `json:"key,omitempty"`
	ValueType         int64                                      `json:"valueType,omitempty"`
	Tips              string                                     `json:"tips,omitempty"`
	IsRequired        int64                                      `json:"isRequired,omitempty"`
	Sort              int64                                      `json:"sort,omitempty"`
	Status            int64                                      `json:"status,omitempty"`
	IsValueUnique     int64                                      `json:"isValueUnique,omitempty"`
	IsValueModifiable int64                                      `json:"isValueModifiable,omitempty"`
	DefaultValue      string                                     `json:"defaultValue,omitempty"`
	Value             string                                     `json:"value,omitempty"`
	IsChange          bool                                       `json:"isChange,omitempty"`
}

type CustomerFieldConfigGetResponseOptionList struct {
	FieldId     int64  `json:"fieldId,omitempty"`
	OptionName  string `json:"optionName,omitempty"`
	OptionValue string `json:"optionValue,omitempty"`
}

type CustomerFieldConfigGetResponseRuleList struct {
	FieldId   int64 `json:"fieldId,omitempty"`
	RuleType  int64 `json:"ruleType,omitempty"`
	Max       int64 `json:"max,omitempty"`
	Min       int64 `json:"min,omitempty"`
	LimitType int64 `json:"limitType,omitempty"`
}
