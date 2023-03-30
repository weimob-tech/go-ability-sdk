package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionNynjCreate
// @id 1869
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1869?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=N元N件活动保存)
func (client *Client) PromotionNynjCreate(request *PromotionNynjCreateRequest) (response *PromotionNynjCreateResponse, err error) {
	rpcResponse := CreatePromotionNynjCreateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionNynjCreateRequest struct {
	*api.BaseRequest

	BasicInfo          PromotionNynjCreateRequestBasicInfo `json:"basicInfo,omitempty"`
	RuleBizVO          PromotionNynjCreateRequestRuleBizVO `json:"ruleBizVO,omitempty"`
	CanDeductionType   string                              `json:"canDeductionType,omitempty"`
	Description        string                              `json:"description,omitempty"`
	EndDate            int64                               `json:"endDate,omitempty"`
	FitPeople          int64                               `json:"fitPeople,omitempty"`
	Uuid               string                              `json:"uuid,omitempty"`
	ImageUrl           string                              `json:"imageUrl,omitempty"`
	LimitNum           int64                               `json:"limitNum,omitempty"`
	LimitPromotionType string                              `json:"limitPromotionType,omitempty"`
	MarkingType        int64                               `json:"markingType,omitempty"`
	SelectStoreIds     []int64                             `json:"selectStoreIds,omitempty"`
	SelectStoreType    int64                               `json:"selectStoreType,omitempty"`
	StartDate          int64                               `json:"startDate,omitempty"`
	TimeType           int64                               `json:"timeType,omitempty"`
	Title              string                              `json:"title,omitempty"`
	TopicStyleType     int64                               `json:"topicStyleType,omitempty"`
	UseOfScene         string                              `json:"useOfScene,omitempty"`
	TagRuleId          int64                               `json:"tagRuleId,omitempty"`
	FitPeopleRuleId    int64                               `json:"fitPeopleRuleId,omitempty"`
	SelectGoodsType    int64                               `json:"selectGoodsType,omitempty"`
	SelectGoodsIds     []int64                             `json:"selectGoodsIds,omitempty"`
	ExcludeGoodsIds    []int64                             `json:"excludeGoodsIds,omitempty"`
}

type PromotionNynjCreateResponse struct {
	ActivityId int64  `json:"activityId,omitempty"`
	Success    bool   `json:"success,omitempty"`
	Msg        string `json:"msg,omitempty"`
}

func CreatePromotionNynjCreateRequest() (request *PromotionNynjCreateRequest) {
	request = &PromotionNynjCreateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "promotion/nynj/create", "apigw")
	request.Method = api.POST
	return
}

func CreatePromotionNynjCreateResponse() (response *api.BaseResponse[PromotionNynjCreateResponse]) {
	response = api.CreateResponse[PromotionNynjCreateResponse](&PromotionNynjCreateResponse{})
	return
}

type PromotionNynjCreateRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type PromotionNynjCreateRequestRuleBizVO struct {
	RuleGroupDataList []PromotionNynjCreateRequestRuleGroupDataList `json:"ruleGroupDataList,omitempty"`
}

type PromotionNynjCreateRequestRuleGroupDataList struct {
	PromotionContentVOS []PromotionNynjCreateRequestPromotionContentVOS `json:"promotionContentVOS,omitempty"`
	CalculateType       int64                                           `json:"calculateType,omitempty"`
	ConditionType       int64                                           `json:"conditionType,omitempty"`
	EffectRange         int64                                           `json:"effectRange,omitempty"`
	GroupPriority       int64                                           `json:"groupPriority,omitempty"`
}

type PromotionNynjCreateRequestPromotionContentVOS struct {
	ConditionOperator int64  `json:"conditionOperator,omitempty"`
	ConditionValue    string `json:"conditionValue,omitempty"`
	LevelPriority     int64  `json:"levelPriority,omitempty"`
	ResultType        int64  `json:"resultType,omitempty"`
	ResultValue       int64  `json:"resultValue,omitempty"`
}
