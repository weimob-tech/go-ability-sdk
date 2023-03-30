package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionXjxzGet
// @id 1453
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1453?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询第X件X折详情)
func (client *Client) PromotionXjxzGet(request *PromotionXjxzGetRequest) (response *PromotionXjxzGetResponse, err error) {
	rpcResponse := CreatePromotionXjxzGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionXjxzGetRequest struct {
	*api.BaseRequest

	BasicInfo       PromotionXjxzGetRequestBasicInfo `json:"basicInfo,omitempty"`
	ActivityType    int64                            `json:"activityType,omitempty"`
	ActivitySubType int64                            `json:"activitySubType,omitempty"`
	ActivityId      int64                            `json:"activityId,omitempty"`
}

type PromotionXjxzGetResponse struct {
	ResultPolymerBizVOList []PromotionXjxzGetResponseResultPolymerBizVOList `json:"resultPolymerBizVOList,omitempty"`
	CycleTimeList          PromotionXjxzGetResponseCycleTimeList            `json:"cycleTimeList,omitempty"`
	IncludeGoodsBizId      []int64                                          `json:"includeGoodsBizId,omitempty"`
	ExcludeGoodsBizId      []int64                                          `json:"excludeGoodsBizId,omitempty"`
	TopicStyleType         int64                                            `json:"topicStyleType,omitempty"`
	PromotionImage         string                                           `json:"promotionImage,omitempty"`
	PromotionDescription   string                                           `json:"promotionDescription,omitempty"`
	PromotionTag           int64                                            `json:"promotionTag,omitempty"`
	PromotionTagName       string                                           `json:"promotionTagName,omitempty"`
	ShowLinePrice          bool                                             `json:"showLinePrice,omitempty"`
	VerifyPhoneNumber      int64                                            `json:"verifyPhoneNumber,omitempty"`
	ConditionType          int64                                            `json:"conditionType,omitempty"`
	CalculateType          int64                                            `json:"calculateType,omitempty"`
	EachLimitNum           int64                                            `json:"eachLimitNum,omitempty"`
	ActivityId             int64                                            `json:"activityId,omitempty"`
	ActivityType           int64                                            `json:"activityType,omitempty"`
	PromotionName          string                                           `json:"promotionName,omitempty"`
	StartTime              int64                                            `json:"startTime,omitempty"`
	EndTime                int64                                            `json:"endTime,omitempty"`
	MarkingType            int64                                            `json:"markingType,omitempty"`
	TagRuleId              int64                                            `json:"tagRuleId,omitempty"`
	SelectPeopleType       int64                                            `json:"selectPeopleType,omitempty"`
	SelectPeopleRuleId     int64                                            `json:"selectPeopleRuleId,omitempty"`
	ApplicationGoodsType   int64                                            `json:"applicationGoodsType,omitempty"`
	CheckType              int64                                            `json:"checkType,omitempty"`
	SubIds                 []int64                                          `json:"subIds,omitempty"`
	SelectStoreType        int64                                            `json:"selectStoreType,omitempty"`
	SelectStoreIds         []int64                                          `json:"selectStoreIds,omitempty"`
	TradeScene             []int64                                          `json:"tradeScene,omitempty"`
	AvailableDeduct        []int64                                          `json:"availableDeduct,omitempty"`
	AvailablePromotion     []int64                                          `json:"availablePromotion,omitempty"`
	TimeType               int64                                            `json:"timeType,omitempty"`
}

func CreatePromotionXjxzGetRequest() (request *PromotionXjxzGetRequest) {
	request = &PromotionXjxzGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "promotion/xjxz/get", "apigw")
	request.Method = api.POST
	return
}

func CreatePromotionXjxzGetResponse() (response *api.BaseResponse[PromotionXjxzGetResponse]) {
	response = api.CreateResponse[PromotionXjxzGetResponse](&PromotionXjxzGetResponse{})
	return
}

type PromotionXjxzGetRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type PromotionXjxzGetResponseResultPolymerBizVOList struct {
	ResultLimitInfo       PromotionXjxzGetResponseResultLimitInfo         `json:"resultLimitInfo,omitempty"`
	PromotionSkuBizVOList []PromotionXjxzGetResponsePromotionSkuBizVOList `json:"promotionSkuBizVOList,omitempty"`
	ResultType            int64                                           `json:"resultType,omitempty"`
	ResultValue           int64                                           `json:"resultValue,omitempty"`
	ConditionOperator     int64                                           `json:"conditionOperator,omitempty"`
	ConditionValue        int64                                           `json:"conditionValue,omitempty"`
	MaxAmount             int64                                           `json:"maxAmount,omitempty"`
	ItemType              int64                                           `json:"itemType,omitempty"`
	ItemId                int64                                           `json:"itemId,omitempty"`
	LimitDistrictType     int64                                           `json:"limitDistrictType,omitempty"`
	District              string                                          `json:"district,omitempty"`
}

type PromotionXjxzGetResponseResultLimitInfo struct {
	ChildLimitList []PromotionXjxzGetResponseChildLimitList `json:"childLimitList,omitempty"`
	NodeId         string                                   `json:"nodeId,omitempty"`
	Path           string                                   `json:"path,omitempty"`
	LimitBizId     int64                                    `json:"limitBizId,omitempty"`
	LimitBizType   int64                                    `json:"limitBizType,omitempty"`
	LimitType      int64                                    `json:"limitType,omitempty"`
	EachLimitNum   int64                                    `json:"eachLimitNum,omitempty"`
	TotalLimitNum  int64                                    `json:"totalLimitNum,omitempty"`
	EachSoldNum    int64                                    `json:"eachSoldNum,omitempty"`
	TotalSoldNum   int64                                    `json:"totalSoldNum,omitempty"`
	CanBuyNum      int64                                    `json:"canBuyNum,omitempty"`
}

type PromotionXjxzGetResponseChildLimitList struct {
	NodeId         string  `json:"nodeId,omitempty"`
	Path           string  `json:"path,omitempty"`
	LimitBizId     int64   `json:"limitBizId,omitempty"`
	LimitBizType   int64   `json:"limitBizType,omitempty"`
	LimitType      int64   `json:"limitType,omitempty"`
	EachLimitNum   int64   `json:"eachLimitNum,omitempty"`
	TotalLimitNum  int64   `json:"totalLimitNum,omitempty"`
	EachSoldNum    int64   `json:"eachSoldNum,omitempty"`
	TotalSoldNum   int64   `json:"totalSoldNum,omitempty"`
	CanBuyNum      int64   `json:"canBuyNum,omitempty"`
	ChildLimitList []int64 `json:"childLimitList,omitempty"`
}

type PromotionXjxzGetResponsePromotionSkuBizVOList struct {
	GoodsId       int64 `json:"goodsId,omitempty"`
	SkuId         int64 `json:"skuId,omitempty"`
	ActivityPrice int64 `json:"activityPrice,omitempty"`
	CanBuyNum     int64 `json:"canBuyNum,omitempty"`
}

type PromotionXjxzGetResponseCycleTimeList struct {
	RepeatType          int64  `json:"repeatType,omitempty"`
	RepeatDay           string `json:"repeatDay,omitempty"`
	RepeatStartInterval string `json:"repeatStartInterval,omitempty"`
	RepeatEndInterval   string `json:"repeatEndInterval,omitempty"`
	RepeatGmt           int64  `json:"repeatGmt,omitempty"`
}
