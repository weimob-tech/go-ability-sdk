package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionNynjGet
// @id 1451
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1451?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询N元N件活动详情)
func (client *Client) PromotionNynjGet(request *PromotionNynjGetRequest) (response *PromotionNynjGetResponse, err error) {
	rpcResponse := CreatePromotionNynjGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionNynjGetRequest struct {
	*api.BaseRequest

	BasicInfo       PromotionNynjGetRequestBasicInfo `json:"basicInfo,omitempty"`
	ActivityType    int64                            `json:"activityType,omitempty"`
	ActivitySubType int64                            `json:"activitySubType,omitempty"`
	ActivityId      int64                            `json:"activityId,omitempty"`
}

type PromotionNynjGetResponse struct {
	ResultPolymerBizVOList []PromotionNynjGetResponseResultPolymerBizVOList `json:"resultPolymerBizVOList,omitempty"`
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

func CreatePromotionNynjGetRequest() (request *PromotionNynjGetRequest) {
	request = &PromotionNynjGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "promotion/nynj/get", "apigw")
	request.Method = api.POST
	return
}

func CreatePromotionNynjGetResponse() (response *api.BaseResponse[PromotionNynjGetResponse]) {
	response = api.CreateResponse[PromotionNynjGetResponse](&PromotionNynjGetResponse{})
	return
}

type PromotionNynjGetRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type PromotionNynjGetResponseResultPolymerBizVOList struct {
	ResultLimitInfo       PromotionNynjGetResponseResultLimitInfo         `json:"resultLimitInfo,omitempty"`
	PromotionSkuBizVOList []PromotionNynjGetResponsePromotionSkuBizVOList `json:"promotionSkuBizVOList,omitempty"`
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

type PromotionNynjGetResponseResultLimitInfo struct {
	ChildLimitList []PromotionNynjGetResponseChildLimitList `json:"childLimitList,omitempty"`
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

type PromotionNynjGetResponseChildLimitList struct {
	ChildLimitList []PromotionNynjGetResponseChildLimitList2 `json:"childLimitList,omitempty"`
	NodeId         string                                    `json:"nodeId,omitempty"`
	Path           string                                    `json:"path,omitempty"`
	LimitBizId     int64                                     `json:"limitBizId,omitempty"`
	LimitBizType   int64                                     `json:"limitBizType,omitempty"`
	LimitType      int64                                     `json:"limitType,omitempty"`
	EachLimitNum   int64                                     `json:"eachLimitNum,omitempty"`
	TotalLimitNum  int64                                     `json:"totalLimitNum,omitempty"`
	EachSoldNum    int64                                     `json:"eachSoldNum,omitempty"`
	TotalSoldNum   int64                                     `json:"totalSoldNum,omitempty"`
	CanBuyNum      int64                                     `json:"canBuyNum,omitempty"`
}

type PromotionNynjGetResponseChildLimitList2 struct {
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

type PromotionNynjGetResponsePromotionSkuBizVOList struct {
	GoodsId       int64 `json:"goodsId,omitempty"`
	SkuId         int64 `json:"skuId,omitempty"`
	ActivityPrice int64 `json:"activityPrice,omitempty"`
	CanBuyNum     int64 `json:"canBuyNum,omitempty"`
}
