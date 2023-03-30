package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CardTemplateGetMemberCardPointUseRule
// @id 979
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取积分使用规则)
func (client *Client) CardTemplateGetMemberCardPointUseRule(request *CardTemplateGetPointUseRuleRequest) (response *CardTemplateGetPointUseRuleResponse, err error) {
	rpcResponse := CreateCardTemplateGetPointUseRuleResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CardTemplateGetPointUseRuleRequest struct {
	*api.BaseRequest

	Type int `json:"type,omitempty"`
}

type CardTemplateGetPointUseRuleResponse struct {
}

func CreateCardTemplateGetPointUseRuleRequest() (request *CardTemplateGetPointUseRuleRequest) {
	request = &CardTemplateGetPointUseRuleRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "cardTemplate/getMemberCardPointUseRule", "api")
	request.Method = api.POST
	return
}

func CreateCardTemplateGetPointUseRuleResponse() (response *api.BaseResponse[CardTemplateGetPointUseRuleResponse]) {
	response = api.CreateResponse[CardTemplateGetPointUseRuleResponse](&CardTemplateGetPointUseRuleResponse{})
	return
}
