package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpDeleteMembershipRightRule
// @id 1907
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除权益规则)
func (client *Client) MbpDeleteMembershipRightRule(request *MbpDeleteMembershipRightRuleRequest) (response *MbpDeleteMembershipRightRuleResponse, err error) {
	rpcResponse := CreateMbpDeleteMembershipRightRuleResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpDeleteMembershipRightRuleRequest struct {
	*api.BaseRequest

	Pid     int64 `json:"pid,omitempty"`
	StoreId int64 `json:"storeId,omitempty"`
	RuleId  int   `json:"ruleId,omitempty"`
}

type MbpDeleteMembershipRightRuleResponse struct {
}

func CreateMbpDeleteMembershipRightRuleRequest() (request *MbpDeleteMembershipRightRuleRequest) {
	request = &MbpDeleteMembershipRightRuleRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "mbp/deleteMembershipRightRule", "api")
	request.Method = api.POST
	return
}

func CreateMbpDeleteMembershipRightRuleResponse() (response *api.BaseResponse[MbpDeleteMembershipRightRuleResponse]) {
	response = api.CreateResponse[MbpDeleteMembershipRightRuleResponse](&MbpDeleteMembershipRightRuleResponse{})
	return
}
