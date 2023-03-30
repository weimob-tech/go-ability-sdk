package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionFulldiscountUpdate
// @id 1871
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1871?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=满减满折编辑接口)
func (client *Client) PromotionFulldiscountUpdate(request *PromotionFulldiscountUpdateRequest) (response *PromotionFulldiscountUpdateResponse, err error) {
	rpcResponse := CreatePromotionFulldiscountUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionFulldiscountUpdateRequest struct {
	*api.BaseRequest

	BasicInfo          PromotionFulldiscountUpdateRequestBasicInfo        `json:"basicInfo,omitempty"`
	RuleBizVO          PromotionFulldiscountUpdateRequestRuleBizVO        `json:"ruleBizVO,omitempty"`
	SelectGoodsBos     []PromotionFulldiscountUpdateRequestSelectGoodsBos `json:"selectGoodsBos,omitempty"`
	ActivityId         int64                                              `json:"activityId,omitempty"`
	CanDeductionType   string                                             `json:"canDeductionType,omitempty"`
	Description        string                                             `json:"description,omitempty"`
	EndDate            int64                                              `json:"endDate,omitempty"`
	Uuid               string                                             `json:"uuid,omitempty"`
	FitPeople          int64                                              `json:"fitPeople,omitempty"`
	ImageUrl           string                                             `json:"imageUrl,omitempty"`
	LimitPromotionType string                                             `json:"limitPromotionType,omitempty"`
	MarkingType        int64                                              `json:"markingType,omitempty"`
	SelectStoreIds     []int64                                            `json:"selectStoreIds,omitempty"`
	SelectStoreType    int64                                              `json:"selectStoreType,omitempty"`
	StartDate          int64                                              `json:"startDate,omitempty"`
	TimeType           int64                                              `json:"timeType,omitempty"`
	Title              string                                             `json:"title,omitempty"`
	TopicStyleType     int64                                              `json:"topicStyleType,omitempty"`
	UseOfScene         string                                             `json:"useOfScene,omitempty"`
	TagRuleId          int64                                              `json:"tagRuleId,omitempty"`
	FitPeopleRuleId    int64                                              `json:"fitPeopleRuleId,omitempty"`
}

type PromotionFulldiscountUpdateResponse struct {
	Success bool   `json:"success,omitempty"`
	Msg     string `json:"msg,omitempty"`
}

func CreatePromotionFulldiscountUpdateRequest() (request *PromotionFulldiscountUpdateRequest) {
	request = &PromotionFulldiscountUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "promotion/fulldiscount/update", "apigw")
	request.Method = api.POST
	return
}

func CreatePromotionFulldiscountUpdateResponse() (response *api.BaseResponse[PromotionFulldiscountUpdateResponse]) {
	response = api.CreateResponse[PromotionFulldiscountUpdateResponse](&PromotionFulldiscountUpdateResponse{})
	return
}

type PromotionFulldiscountUpdateRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type PromotionFulldiscountUpdateRequestRuleBizVO struct {
	RuleGroupDataList []PromotionFulldiscountUpdateRequestRuleGroupDataList `json:"ruleGroupDataList,omitempty"`
}

type PromotionFulldiscountUpdateRequestRuleGroupDataList struct {
	PromotionContentVOS []PromotionFulldiscountUpdateRequestPromotionContentVOS `json:"promotionContentVOS,omitempty"`
	CalculateType       int64                                                   `json:"calculateType,omitempty"`
	ConditionType       int64                                                   `json:"conditionType,omitempty"`
	EffectRange         int64                                                   `json:"effectRange,omitempty"`
	GroupPriority       int64                                                   `json:"groupPriority,omitempty"`
}

type PromotionFulldiscountUpdateRequestPromotionContentVOS struct {
	ConditionOperator int64  `json:"conditionOperator,omitempty"`
	ConditionValue    string `json:"conditionValue,omitempty"`
	LevelPriority     int64  `json:"levelPriority,omitempty"`
	ResultType        int64  `json:"resultType,omitempty"`
	ResultValue       int64  `json:"resultValue,omitempty"`
}

type PromotionFulldiscountUpdateRequestSelectGoodsBos struct {
	BizId          int64 `json:"bizId,omitempty"`
	CheckScopeType int64 `json:"checkScopeType,omitempty"`
	Join           int64 `json:"join,omitempty"`
	Scope          int64 `json:"scope,omitempty"`
}
