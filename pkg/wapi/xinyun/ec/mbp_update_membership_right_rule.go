package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpUpdateMembershipRightRule
// @id 1906
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新权益规则信息)
func (client *Client) MbpUpdateMembershipRightRule(request *MbpUpdateMembershipRightRuleRequest) (response *MbpUpdateMembershipRightRuleResponse, err error) {
	rpcResponse := CreateMbpUpdateMembershipRightRuleResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpUpdateMembershipRightRuleRequest struct {
	*api.BaseRequest

	Pid                int64  `json:"pid,omitempty"`
	StoreId            int64  `json:"storeId,omitempty"`
	RuleId             int    `json:"ruleId,omitempty"`
	RuleName           string `json:"ruleName,omitempty"`
	Discount           int64  `json:"discount,omitempty"`
	MemberDiscountType bool   `json:"memberDiscountType,omitempty"`
	IgnoreChangeType   int    `json:"ignoreChangeType,omitempty"`
	IsEdit             int    `json:"isEdit,omitempty"`
}

type MbpUpdateMembershipRightRuleResponse struct {
}

func CreateMbpUpdateMembershipRightRuleRequest() (request *MbpUpdateMembershipRightRuleRequest) {
	request = &MbpUpdateMembershipRightRuleRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "mbp/updateMembershipRightRule", "api")
	request.Method = api.POST
	return
}

func CreateMbpUpdateMembershipRightRuleResponse() (response *api.BaseResponse[MbpUpdateMembershipRightRuleResponse]) {
	response = api.CreateResponse[MbpUpdateMembershipRightRuleResponse](&MbpUpdateMembershipRightRuleResponse{})
	return
}
