package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionGiftGet
// @id 1452
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1452?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询订单满赠活动详情)
func (client *Client) PromotionGiftGet(request *PromotionGiftGetRequest) (response *PromotionGiftGetResponse, err error) {
	rpcResponse := CreatePromotionGiftGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionGiftGetRequest struct {
	*api.BaseRequest

	BasicInfo  PromotionGiftGetRequestBasicInfo `json:"basicInfo,omitempty"`
	ActivityId int64                            `json:"activityId,omitempty"`
}

type PromotionGiftGetResponse struct {
	ResultPolymerBizVOList []PromotionGiftGetResponseResultPolymerBizVOList `json:"resultPolymerBizVOList,omitempty"`
	FullGiftMethod         PromotionGiftGetResponseFullGiftMethod           `json:"fullGiftMethod,omitempty"`
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
}

func CreatePromotionGiftGetRequest() (request *PromotionGiftGetRequest) {
	request = &PromotionGiftGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "promotion/gift/get", "apigw")
	request.Method = api.POST
	return
}

func CreatePromotionGiftGetResponse() (response *api.BaseResponse[PromotionGiftGetResponse]) {
	response = api.CreateResponse[PromotionGiftGetResponse](&PromotionGiftGetResponse{})
	return
}

type PromotionGiftGetRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type PromotionGiftGetResponseResultPolymerBizVOList struct {
	ResultLimitInfo       PromotionGiftGetResponseResultLimitInfo         `json:"resultLimitInfo,omitempty"`
	PromotionSkuBizVOList []PromotionGiftGetResponsePromotionSkuBizVOList `json:"promotionSkuBizVOList,omitempty"`
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

type PromotionGiftGetResponseResultLimitInfo struct {
	ChildLimitList []PromotionGiftGetResponseChildLimitList `json:"childLimitList,omitempty"`
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

type PromotionGiftGetResponseChildLimitList struct {
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

type PromotionGiftGetResponsePromotionSkuBizVOList struct {
	GoodsId       int64 `json:"goodsId,omitempty"`
	SkuId         int64 `json:"skuId,omitempty"`
	ActivityPrice int64 `json:"activityPrice,omitempty"`
	CanBuyNum     int64 `json:"canBuyNum,omitempty"`
}

type PromotionGiftGetResponseFullGiftMethod struct {
	RuleList      []PromotionGiftGetResponseRuleList `json:"ruleList,omitempty"`
	ConditionType int64                              `json:"conditionType,omitempty"`
	CalculateType int64                              `json:"calculateType,omitempty"`
	GlobalGroupId int64                              `json:"globalGroupId,omitempty"`
	GroupPriority int64                              `json:"groupPriority,omitempty"`
}

type PromotionGiftGetResponseRuleList struct {
	GoodsList         []PromotionGiftGetResponseGoodsList `json:"goodsList,omitempty"`
	GlobalLevelId     int64                               `json:"globalLevelId,omitempty"`
	LevelPriority     int64                               `json:"levelPriority,omitempty"`
	FullGiftCondition int64                               `json:"fullGiftCondition,omitempty"`
	FullGiftLimit     int64                               `json:"fullGiftLimit,omitempty"`
	FullGiftTotal     int64                               `json:"fullGiftTotal,omitempty"`
}

type PromotionGiftGetResponseGoodsList struct {
	GoodsId     int64 `json:"goodsId,omitempty"`
	JoinSkuId   int64 `json:"joinSkuId,omitempty"`
	SingleLimit int64 `json:"singleLimit,omitempty"`
}
