package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembercardPlanGetList
// @id 1510
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1510?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询会员卡方案列表)
func (client *Client) MembercardPlanGetList(request *MembercardPlanGetListRequest) (response *MembercardPlanGetListResponse, err error) {
	rpcResponse := CreateMembercardPlanGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembercardPlanGetListRequest struct {
	*api.BaseRequest

	MembershipPlanId int64   `json:"membershipPlanId,omitempty"`
	CardType         int64   `json:"cardType,omitempty"`
	Vid              int64   `json:"vid,omitempty"`
	QueryType        []int64 `json:"queryType,omitempty"`
}

type MembercardPlanGetListResponse struct {
	LevelInfos        []MembercardPlanGetListResponseLevelInfos      `json:"levelInfos,omitempty"`
	DynamicSettings   []MembercardPlanGetListResponseDynamicSettings `json:"dynamicSettings,omitempty"`
	RightInfo         []MembercardPlanGetListResponseRightInfo2      `json:"rightInfo,omitempty"`
	BosId             int64                                          `json:"bosId,omitempty"`
	Name              string                                         `json:"name,omitempty"`
	GrowthName        string                                         `json:"growthName,omitempty"`
	MembershipPlanId  int64                                          `json:"membershipPlanId,omitempty"`
	ProductInstanceId int64                                          `json:"productInstanceId,omitempty"`
	CardType          int64                                          `json:"cardType,omitempty"`
	LevelRuleType     int64                                          `json:"levelRuleType,omitempty"`
	Status            int64                                          `json:"status,omitempty"`
	CardFreeType      int64                                          `json:"cardFreeType,omitempty"`
	CardExplain       string                                         `json:"cardExplain,omitempty"`
}

func CreateMembercardPlanGetListRequest() (request *MembercardPlanGetListRequest) {
	request = &MembercardPlanGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "membercard/plan/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateMembercardPlanGetListResponse() (response *api.BaseResponse[MembercardPlanGetListResponse]) {
	response = api.CreateResponse[MembercardPlanGetListResponse](&MembercardPlanGetListResponse{})
	return
}

type MembercardPlanGetListResponseLevelInfos struct {
	RightInfo []MembercardPlanGetListResponseRightInfo `json:"rightInfo,omitempty"`
	LevelId   int64                                    `json:"levelId,omitempty"`
	LevelName string                                   `json:"levelName,omitempty"`
}

type MembercardPlanGetListResponseRightInfo struct {
	PrivilegeId   int64  `json:"privilegeId,omitempty"`
	PrivilegeName string `json:"privilegeName,omitempty"`
	PrivilegeType int    `json:"privilegeType,omitempty"`
	Remark        string `json:"remark,omitempty"`
	RuleId        int64  `json:"ruleId,omitempty"`
	RuleName      string `json:"ruleName,omitempty"`
}

type MembercardPlanGetListResponseDynamicSettings struct {
	ValidType  int64 `json:"validType,omitempty"`
	ValidValue int64 `json:"validValue	,omitempty"`
	Sort       int64 `json:"sort,omitempty"`
}

type MembercardPlanGetListResponseRightInfo2 struct {
	PrivilegeId   int64  `json:"privilegeId,omitempty"`
	PrivilegeName string `json:"privilegeName,omitempty"`
	PrivilegeType int    `json:"privilegeType,omitempty"`
	Remark        string `json:"remark,omitempty"`
	RuleId        int64  `json:"ruleId,omitempty"`
	RuleName      string `json:"ruleName,omitempty"`
}
