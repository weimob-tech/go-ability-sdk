package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembercardPrivilegeRuleDelete
// @id 1769
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1769?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除权益规则)
func (client *Client) MembercardPrivilegeRuleDelete(request *MembercardPrivilegeRuleDeleteRequest) (response *MembercardPrivilegeRuleDeleteResponse, err error) {
	rpcResponse := CreateMembercardPrivilegeRuleDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembercardPrivilegeRuleDeleteRequest struct {
	*api.BaseRequest

	RuleId  int64 `json:"ruleId,omitempty"`
	Wid     int64 `json:"wid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
	Vid     int64 `json:"vid,omitempty"`
}

type MembercardPrivilegeRuleDeleteResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateMembercardPrivilegeRuleDeleteRequest() (request *MembercardPrivilegeRuleDeleteRequest) {
	request = &MembercardPrivilegeRuleDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "membercard/privilege/rule/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateMembercardPrivilegeRuleDeleteResponse() (response *api.BaseResponse[MembercardPrivilegeRuleDeleteResponse]) {
	response = api.CreateResponse[MembercardPrivilegeRuleDeleteResponse](&MembercardPrivilegeRuleDeleteResponse{})
	return
}
