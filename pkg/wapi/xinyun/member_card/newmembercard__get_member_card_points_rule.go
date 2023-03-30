package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NewmembercardGetMemberCardPointsRule
// @id 106
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取积分配置)
func (client *Client) NewmembercardGetMemberCardPointsRule(request *NewmembercardGetPointsRuleRequest) (response *NewmembercardGetPointsRuleResponse, err error) {
	rpcResponse := CreateNewmembercardGetPointsRuleResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NewmembercardGetPointsRuleRequest struct {
	*api.BaseRequest
}

type NewmembercardGetPointsRuleResponse struct {
}

func CreateNewmembercardGetPointsRuleRequest() (request *NewmembercardGetPointsRuleRequest) {
	request = &NewmembercardGetPointsRuleRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "newmembercard/GetMemberCardPointsRule", "api")
	request.Method = api.POST
	return
}

func CreateNewmembercardGetPointsRuleResponse() (response *api.BaseResponse[NewmembercardGetPointsRuleResponse]) {
	response = api.CreateResponse[NewmembercardGetPointsRuleResponse](&NewmembercardGetPointsRuleResponse{})
	return
}
