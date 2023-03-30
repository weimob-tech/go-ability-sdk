package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembercardPrivilegeRuleInsert
// @id 1768
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1768?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=添加权益规则)
func (client *Client) MembercardPrivilegeRuleInsert(request *MembercardPrivilegeRuleInsertRequest) (response *MembercardPrivilegeRuleInsertResponse, err error) {
	rpcResponse := CreateMembercardPrivilegeRuleInsertResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembercardPrivilegeRuleInsertRequest struct {
	*api.BaseRequest

	RuleName           string `json:"ruleName,omitempty"`
	Discount           int64  `json:"discount,omitempty"`
	MemberDiscountType int64  `json:"memberDiscountType,omitempty"`
	IgnoreChangeType   int64  `json:"ignoreChangeType,omitempty"`
	Wid                int64  `json:"wid,omitempty"`
	VidType            int64  `json:"vidType,omitempty"`
	Vid                int64  `json:"vid,omitempty"`
	PrivilegeType      string `json:"privilegeType,omitempty"`
}

type MembercardPrivilegeRuleInsertResponse struct {
	RuleId int64 `json:"ruleId,omitempty"`
}

func CreateMembercardPrivilegeRuleInsertRequest() (request *MembercardPrivilegeRuleInsertRequest) {
	request = &MembercardPrivilegeRuleInsertRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "membercard/privilege/rule/insert", "apigw")
	request.Method = api.POST
	return
}

func CreateMembercardPrivilegeRuleInsertResponse() (response *api.BaseResponse[MembercardPrivilegeRuleInsertResponse]) {
	response = api.CreateResponse[MembercardPrivilegeRuleInsertResponse](&MembercardPrivilegeRuleInsertResponse{})
	return
}
