package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpGenerateCrowdRuleId
// @id 2129
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=生成人群弹窗规则)
func (client *Client) MbpGenerateCrowdRuleId(request *MbpGenerateCrowdRuleIdRequest) (response *MbpGenerateCrowdRuleIdResponse, err error) {
	rpcResponse := CreateMbpGenerateCrowdRuleIdResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpGenerateCrowdRuleIdRequest struct {
	*api.BaseRequest

	Customer MbpGenerateCrowdRuleIdRequestCustomer `json:"customer,omitempty"`
	Pid      int64                                 `json:"pid,omitempty"`
}

type MbpGenerateCrowdRuleIdResponse struct {
	RuleId int64 `json:"ruleId,omitempty"`
}

func CreateMbpGenerateCrowdRuleIdRequest() (request *MbpGenerateCrowdRuleIdRequest) {
	request = &MbpGenerateCrowdRuleIdRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "mbp/generateCrowdRuleId", "api")
	request.Method = api.POST
	return
}

func CreateMbpGenerateCrowdRuleIdResponse() (response *api.BaseResponse[MbpGenerateCrowdRuleIdResponse]) {
	response = api.CreateResponse[MbpGenerateCrowdRuleIdResponse](&MbpGenerateCrowdRuleIdResponse{})
	return
}

type MbpGenerateCrowdRuleIdRequestCustomer struct {
	MembershipCardList []MbpGenerateCrowdRuleIdRequestMembershipCardList `json:"membershipCardList,omitempty"`
	RelationType       int                                               `json:"relationType,omitempty"`
	IsMembership       bool                                              `json:"isMembership,omitempty"`
}

type MbpGenerateCrowdRuleIdRequestMembershipCardList struct {
	RankList       []MbpGenerateCrowdRuleIdRequestRankList `json:"rankList,omitempty"`
	CardTemplateId int64                                   `json:"cardTemplateId,omitempty"`
	CardType       int                                     `json:"cardType,omitempty"`
}

type MbpGenerateCrowdRuleIdRequestRankList struct {
	RankId   int64  `json:"rankId,omitempty"`
	RankName string `json:"rankName,omitempty"`
}
