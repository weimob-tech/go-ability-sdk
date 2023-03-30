package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembercardPrivilegeRuleUpdate
// @id 1748
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1748?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新权益规则)
func (client *Client) MembercardPrivilegeRuleUpdate(request *MembercardPrivilegeRuleUpdateRequest) (response *MembercardPrivilegeRuleUpdateResponse, err error) {
	rpcResponse := CreateMembercardPrivilegeRuleUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembercardPrivilegeRuleUpdateRequest struct {
	*api.BaseRequest

	RuleName           string `json:"ruleName,omitempty"`
	Discount           int64  `json:"discount,omitempty"`
	MemberDiscountType int64  `json:"memberDiscountType,omitempty"`
	IgnoreChangeType   int64  `json:"ignoreChangeType,omitempty"`
	Wid                int64  `json:"wid,omitempty"`
	VidType            int64  `json:"vidType,omitempty"`
	Vid                int64  `json:"vid,omitempty"`
	RuleId             int64  `json:"ruleId,omitempty"`
}

type MembercardPrivilegeRuleUpdateResponse struct {
	RuleId int64 `json:"ruleId,omitempty"`
}

func CreateMembercardPrivilegeRuleUpdateRequest() (request *MembercardPrivilegeRuleUpdateRequest) {
	request = &MembercardPrivilegeRuleUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "membercard/privilege/rule/update", "apigw")
	request.Method = api.POST
	return
}

func CreateMembercardPrivilegeRuleUpdateResponse() (response *api.BaseResponse[MembercardPrivilegeRuleUpdateResponse]) {
	response = api.CreateResponse[MembercardPrivilegeRuleUpdateResponse](&MembercardPrivilegeRuleUpdateResponse{})
	return
}
