package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpAddMembershipRightRule
// @id 1905
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=添加权益规则信息)
func (client *Client) MbpAddMembershipRightRule(request *MbpAddMembershipRightRuleRequest) (response *MbpAddMembershipRightRuleResponse, err error) {
	rpcResponse := CreateMbpAddMembershipRightRuleResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpAddMembershipRightRuleRequest struct {
	*api.BaseRequest

	Pid                int64  `json:"pid,omitempty"`
	StoreId            int64  `json:"storeId,omitempty"`
	RightType          string `json:"rightType,omitempty"`
	RuleName           string `json:"ruleName,omitempty"`
	Discount           int64  `json:"discount,omitempty"`
	MemberDiscountType bool   `json:"memberDiscountType,omitempty"`
	IgnoreChangeType   int    `json:"ignoreChangeType,omitempty"`
	IsEdit             int    `json:"isEdit,omitempty"`
}

type MbpAddMembershipRightRuleResponse struct {
}

func CreateMbpAddMembershipRightRuleRequest() (request *MbpAddMembershipRightRuleRequest) {
	request = &MbpAddMembershipRightRuleRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "mbp/addMembershipRightRule", "api")
	request.Method = api.POST
	return
}

func CreateMbpAddMembershipRightRuleResponse() (response *api.BaseResponse[MbpAddMembershipRightRuleResponse]) {
	response = api.CreateResponse[MbpAddMembershipRightRuleResponse](&MbpAddMembershipRightRuleResponse{})
	return
}
