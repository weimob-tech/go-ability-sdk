package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembercardPrivilegeRuleGet
// @id 2247
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2247?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询权益规则信息)
func (client *Client) MembercardPrivilegeRuleGet(request *MembercardPrivilegeRuleGetRequest) (response *MembercardPrivilegeRuleGetResponse, err error) {
	rpcResponse := CreateMembercardPrivilegeRuleGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembercardPrivilegeRuleGetRequest struct {
	*api.BaseRequest

	RuleId int64 `json:"ruleId,omitempty"`
}

type MembercardPrivilegeRuleGetResponse struct {
	RuleContent   MembercardPrivilegeRuleGetResponseRuleContent `json:"ruleContent,omitempty"`
	PrivilegeId   int64                                         `json:"privilegeId,omitempty"`
	PrivilegeName string                                        `json:"privilegeName,omitempty"`
	PrivilegeType int64                                         `json:"privilegeType,omitempty"`
	Remark        string                                        `json:"remark,omitempty"`
	RuleId        int64                                         `json:"ruleId,omitempty"`
	RuleName      string                                        `json:"ruleName,omitempty"`
}

func CreateMembercardPrivilegeRuleGetRequest() (request *MembercardPrivilegeRuleGetRequest) {
	request = &MembercardPrivilegeRuleGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "membercard/privilege/rule/get", "apigw")
	request.Method = api.POST
	return
}

func CreateMembercardPrivilegeRuleGetResponse() (response *api.BaseResponse[MembercardPrivilegeRuleGetResponse]) {
	response = api.CreateResponse[MembercardPrivilegeRuleGetResponse](&MembercardPrivilegeRuleGetResponse{})
	return
}

type MembercardPrivilegeRuleGetResponseRuleContent struct {
	FreeShippingConditionVo MembercardPrivilegeRuleGetResponseFreeShippingConditionVo `json:"freeShippingConditionVo,omitempty"`
	DateList                []int                                                     `json:"dateList,omitempty"`
	Discount                int64                                                     `json:"discount,omitempty"`
	ConsumeRuleIdList       []int64                                                   `json:"consumeRuleIdList,omitempty"`
	Multiple                int64                                                     `json:"multiple,omitempty"`
	WeekList                []int64                                                   `json:"weekList,omitempty"`
	TimeType                int64                                                     `json:"timeType,omitempty"`
	MonthList               []int64                                                   `json:"monthList,omitempty"`
	CouponIds               []int64                                                   `json:"couponIds,omitempty"`
}

type MembercardPrivilegeRuleGetResponseFreeShippingConditionVo struct {
	AreaList   []MembercardPrivilegeRuleGetResponseAreaList `json:"areaList,omitempty"`
	RuleType   int64                                        `json:"ruleType,omitempty"`
	LimitNum   int64                                        `json:"limitNum,omitempty"`
	LimitMoney int64                                        `json:"limitMoney,omitempty"`
}

type MembercardPrivilegeRuleGetResponseAreaList struct {
	Children []MembercardPrivilegeRuleGetResponseChildren `json:"children,omitempty"`
	Code     string                                       `json:"code,omitempty"`
	CheckAll bool                                         `json:"checkAll,omitempty"`
	Name     string                                       `json:"name,omitempty"`
}

type MembercardPrivilegeRuleGetResponseChildren struct {
}

type MembercardPrivilegeRuleGetResponseDateList struct {
	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
}
