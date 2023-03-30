package kld

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// KLDMemberCardGetMemberCardPointsRule
// @id 211
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取积分配置)
func (client *Client) KLDMemberCardGetMemberCardPointsRule(request *KLDMemberCardGetMemberCardPointsRuleRequest) (response *KLDMemberCardGetMemberCardPointsRuleResponse, err error) {
	rpcResponse := CreateKLDMemberCardGetMemberCardPointsRuleResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type KLDMemberCardGetMemberCardPointsRuleRequest struct {
	*api.BaseRequest
}

type KLDMemberCardGetMemberCardPointsRuleResponse struct {
}

func CreateKLDMemberCardGetMemberCardPointsRuleRequest() (request *KLDMemberCardGetMemberCardPointsRuleRequest) {
	request = &KLDMemberCardGetMemberCardPointsRuleRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("kld", "1_0", "KLDMemberCard/GetMemberCardPointsRule", "api")
	request.Method = api.POST
	return
}

func CreateKLDMemberCardGetMemberCardPointsRuleResponse() (response *api.BaseResponse[KLDMemberCardGetMemberCardPointsRuleResponse]) {
	response = api.CreateResponse[KLDMemberCardGetMemberCardPointsRuleResponse](&KLDMemberCardGetMemberCardPointsRuleResponse{})
	return
}
