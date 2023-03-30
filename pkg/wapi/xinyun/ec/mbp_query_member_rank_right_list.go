package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpQueryMemberRankRightList
// @id 1781
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询会员卡设置的等级权益列表（仅支持多等级会员卡）)
func (client *Client) MbpQueryMemberRankRightList(request *MbpQueryMemberRankRightListRequest) (response *MbpQueryMemberRankRightListResponse, err error) {
	rpcResponse := CreateMbpQueryMemberRankRightListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpQueryMemberRankRightListRequest struct {
	*api.BaseRequest

	MemberCardTemplateId int64  `json:"memberCardTemplateId,omitempty"`
	RightType            string `json:"rightType,omitempty"`
}

type MbpQueryMemberRankRightListResponse struct {
	RankRights []MbpQueryMemberRankRightListResponseRankRights `json:"rankRights,omitempty"`
}

func CreateMbpQueryMemberRankRightListRequest() (request *MbpQueryMemberRankRightListRequest) {
	request = &MbpQueryMemberRankRightListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "mbp/queryMemberRankRightList", "api")
	request.Method = api.POST
	return
}

func CreateMbpQueryMemberRankRightListResponse() (response *api.BaseResponse[MbpQueryMemberRankRightListResponse]) {
	response = api.CreateResponse[MbpQueryMemberRankRightListResponse](&MbpQueryMemberRankRightListResponse{})
	return
}

type MbpQueryMemberRankRightListResponseRankRights struct {
	RightsList []MbpQueryMemberRankRightListResponseRightsList `json:"rightsList,omitempty"`
	RankId     int64                                           `json:"rankId,omitempty"`
	RankName   string                                          `json:"rankName,omitempty"`
}

type MbpQueryMemberRankRightListResponseRightsList struct {
	RightId   string `json:"rightId,omitempty"`
	RightName string `json:"rightName,omitempty"`
	RightType string `json:"rightType,omitempty"`
	RuleId    string `json:"ruleId,omitempty"`
	RuleName  string `json:"ruleName,omitempty"`
}
