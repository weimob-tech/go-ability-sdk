package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionFulldiscountCreate
// @id 1870
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1870?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=满减满折活动保存)
func (client *Client) PromotionFulldiscountCreate(request *PromotionFulldiscountCreateRequest) (response *PromotionFulldiscountCreateResponse, err error) {
	rpcResponse := CreatePromotionFulldiscountCreateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionFulldiscountCreateRequest struct {
	*api.BaseRequest

	BasicInfo          PromotionFulldiscountCreateRequestBasicInfo `json:"basicInfo,omitempty"`
	RuleBizVO          PromotionFulldiscountCreateRequestRuleBizVO `json:"ruleBizVO,omitempty"`
	CanDeductionType   string                                      `json:"canDeductionType,omitempty"`
	Description        string                                      `json:"description,omitempty"`
	EndDate            int64                                       `json:"endDate,omitempty"`
	Uuid               string                                      `json:"uuid,omitempty"`
	FitPeople          int64                                       `json:"fitPeople,omitempty"`
	ImageUrl           string                                      `json:"imageUrl,omitempty"`
	LimitPromotionType string                                      `json:"limitPromotionType,omitempty"`
	MarkingType        int64                                       `json:"markingType,omitempty"`
	SelectStoreIds     []int64                                     `json:"selectStoreIds,omitempty"`
	SelectStoreType    int64                                       `json:"selectStoreType,omitempty"`
	StartDate          int64                                       `json:"startDate,omitempty"`
	TimeType           int64                                       `json:"timeType,omitempty"`
	Title              string                                      `json:"title,omitempty"`
	TopicStyleType     int64                                       `json:"topicStyleType,omitempty"`
	UseOfScene         string                                      `json:"useOfScene,omitempty"`
	TagRuleId          int64                                       `json:"tagRuleId,omitempty"`
	FitPeopleRuleId    int64                                       `json:"fitPeopleRuleId,omitempty"`
	SelectGoodsType    int64                                       `json:"selectGoodsType,omitempty"`
	SelectGoodsIds     []int64                                     `json:"selectGoodsIds,omitempty"`
	ExcludeGoodsIds    []int64                                     `json:"excludeGoodsIds,omitempty"`
}

type PromotionFulldiscountCreateResponse struct {
	ActivityId int64  `json:"activityId,omitempty"`
	Success    bool   `json:"success,omitempty"`
	Msg        string `json:"msg,omitempty"`
}

func CreatePromotionFulldiscountCreateRequest() (request *PromotionFulldiscountCreateRequest) {
	request = &PromotionFulldiscountCreateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "promotion/fulldiscount/create", "apigw")
	request.Method = api.POST
	return
}

func CreatePromotionFulldiscountCreateResponse() (response *api.BaseResponse[PromotionFulldiscountCreateResponse]) {
	response = api.CreateResponse[PromotionFulldiscountCreateResponse](&PromotionFulldiscountCreateResponse{})
	return
}

type PromotionFulldiscountCreateRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type PromotionFulldiscountCreateRequestRuleBizVO struct {
	RuleGroupDataList []PromotionFulldiscountCreateRequestRuleGroupDataList `json:"ruleGroupDataList,omitempty"`
}

type PromotionFulldiscountCreateRequestRuleGroupDataList struct {
	PromotionContentVOS []PromotionFulldiscountCreateRequestPromotionContentVOS `json:"promotionContentVOS,omitempty"`
	CalculateType       int64                                                   `json:"calculateType,omitempty"`
	ConditionType       int64                                                   `json:"conditionType,omitempty"`
	EffectRange         int64                                                   `json:"effectRange,omitempty"`
	GroupPriority       int64                                                   `json:"groupPriority,omitempty"`
}

type PromotionFulldiscountCreateRequestPromotionContentVOS struct {
	ConditionOperator int64  `json:"conditionOperator,omitempty"`
	ConditionValue    string `json:"conditionValue,omitempty"`
	LevelPriority     int64  `json:"levelPriority,omitempty"`
	ResultType        int64  `json:"resultType,omitempty"`
	ResultValue       int64  `json:"resultValue,omitempty"`
}
