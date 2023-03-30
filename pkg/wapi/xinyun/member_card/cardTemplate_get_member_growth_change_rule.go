package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CardTemplateGetMemberGrowthChangeRule
// @id 977
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询成长值增长规则)
func (client *Client) CardTemplateGetMemberGrowthChangeRule(request *CardTemplateGetMemberGrowthChangeRuleRequest) (response *CardTemplateGetMemberGrowthChangeRuleResponse, err error) {
	rpcResponse := CreateCardTemplateGetMemberGrowthChangeRuleResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CardTemplateGetMemberGrowthChangeRuleRequest struct {
	*api.BaseRequest

	Type int `json:"type,omitempty"`
}

type CardTemplateGetMemberGrowthChangeRuleResponse struct {
}

func CreateCardTemplateGetMemberGrowthChangeRuleRequest() (request *CardTemplateGetMemberGrowthChangeRuleRequest) {
	request = &CardTemplateGetMemberGrowthChangeRuleRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "cardTemplate/getMemberGrowthChangeRule", "api")
	request.Method = api.POST
	return
}

func CreateCardTemplateGetMemberGrowthChangeRuleResponse() (response *api.BaseResponse[CardTemplateGetMemberGrowthChangeRuleResponse]) {
	response = api.CreateResponse[CardTemplateGetMemberGrowthChangeRuleResponse](&CardTemplateGetMemberGrowthChangeRuleResponse{})
	return
}
