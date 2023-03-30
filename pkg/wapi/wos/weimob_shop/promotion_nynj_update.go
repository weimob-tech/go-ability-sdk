package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionNynjUpdate
// @id 1868
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1868?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=N元N件修改接口)
func (client *Client) PromotionNynjUpdate(request *PromotionNynjUpdateRequest) (response *PromotionNynjUpdateResponse, err error) {
	rpcResponse := CreatePromotionNynjUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionNynjUpdateRequest struct {
	*api.BaseRequest

	BasicInfo           PromotionNynjUpdateRequestBasicInfo        `json:"basicInfo,omitempty"`
	RuleBizVO           PromotionNynjUpdateRequestRuleBizVO        `json:"ruleBizVO,omitempty"`
	SelectGoodsBos      []PromotionNynjUpdateRequestSelectGoodsBos `json:"selectGoodsBos,omitempty"`
	ActivityId          int64                                      `json:"activityId,omitempty"`
	CanDeductionType    string                                     `json:"canDeductionType,omitempty"`
	Description         string                                     `json:"description,omitempty"`
	EndDate             string                                     `json:"endDate,omitempty"`
	FitPeople           int64                                      `json:"fitPeople,omitempty"`
	Uuid                string                                     `json:"uuid,omitempty"`
	ImageUrl            string                                     `json:"imageUrl,omitempty"`
	LimitNum            int64                                      `json:"limitNum,omitempty"`
	LimitPromotionType  string                                     `json:"limitPromotionType,omitempty"`
	MarkingType         int64                                      `json:"markingType,omitempty"`
	SelectStoreIds      []int64                                    `json:"selectStoreIds,omitempty"`
	SelectStoreType     int64                                      `json:"selectStoreType,omitempty"`
	StartDate           string                                     `json:"startDate,omitempty"`
	TimeType            int64                                      `json:"timeType,omitempty"`
	Title               string                                     `json:"title,omitempty"`
	TopicStyleType      int64                                      `json:"topicStyleType,omitempty"`
	UseOfScene          string                                     `json:"useOfScene,omitempty"`
	TagRuleId           int64                                      `json:"tagRuleId,omitempty"`
	FitPeopleRuleId     int64                                      `json:"fitPeopleRuleId	,omitempty"`
	RepeatType          int64                                      `json:"repeatType,omitempty"`
	RepeatStartInterval string                                     `json:"repeatStartInterval,omitempty"`
	RepeatEndInterval   string                                     `json:"repeatEndInterval,omitempty"`
	RepeatWeekDays      []int64                                    `json:"repeatWeekDays,omitempty"`
	RepeatDays          string                                     `json:"repeatDays,omitempty"`
}

type PromotionNynjUpdateResponse struct {
	Success bool   `json:"success,omitempty"`
	Msg     string `json:"msg,omitempty"`
}

func CreatePromotionNynjUpdateRequest() (request *PromotionNynjUpdateRequest) {
	request = &PromotionNynjUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "promotion/nynj/update", "apigw")
	request.Method = api.POST
	return
}

func CreatePromotionNynjUpdateResponse() (response *api.BaseResponse[PromotionNynjUpdateResponse]) {
	response = api.CreateResponse[PromotionNynjUpdateResponse](&PromotionNynjUpdateResponse{})
	return
}

type PromotionNynjUpdateRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type PromotionNynjUpdateRequestRuleBizVO struct {
	RuleGroupDataList []PromotionNynjUpdateRequestRuleGroupDataList `json:"ruleGroupDataList,omitempty"`
}

type PromotionNynjUpdateRequestRuleGroupDataList struct {
	PromotionContentVOS []PromotionNynjUpdateRequestPromotionContentVOS `json:"promotionContentVOS,omitempty"`
	CalculateType       int64                                           `json:"calculateType,omitempty"`
	ConditionType       int64                                           `json:"conditionType,omitempty"`
	EffectRange         int64                                           `json:"effectRange,omitempty"`
	GroupPriority       int64                                           `json:"groupPriority,omitempty"`
}

type PromotionNynjUpdateRequestPromotionContentVOS struct {
	ConditionOperator int64  `json:"conditionOperator,omitempty"`
	ConditionValue    string `json:"conditionValue,omitempty"`
	LevelPriority     int64  `json:"levelPriority,omitempty"`
	ResultType        int64  `json:"resultType,omitempty"`
	ResultValue       int64  `json:"resultValue,omitempty"`
}

type PromotionNynjUpdateRequestSelectGoodsBos struct {
	CheckScopeType int64 `json:"checkScopeType,omitempty"`
	Join           int64 `json:"join,omitempty"`
	Scope          int64 `json:"scope,omitempty"`
	BizId          int64 `json:"bizId,omitempty"`
}
