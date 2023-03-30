package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionRedemptionGet
// @id 1923
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1923?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询单品换购活动)
func (client *Client) PromotionRedemptionGet(request *PromotionRedemptionGetRequest) (response *PromotionRedemptionGetResponse, err error) {
	rpcResponse := CreatePromotionRedemptionGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionRedemptionGetRequest struct {
	*api.BaseRequest

	BasicInfo  PromotionRedemptionGetRequestBasicInfo `json:"basicInfo,omitempty"`
	ActivityId int64                                  `json:"activityId,omitempty"`
}

type PromotionRedemptionGetResponse struct {
	ResultPolymerBizVOList          []PromotionRedemptionGetResponseResultPolymerBizVOList `json:"resultPolymerBizVOList,omitempty"`
	ExchangePurchaseSetting         PromotionRedemptionGetResponseExchangePurchaseSetting  `json:"exchangePurchaseSetting,omitempty"`
	RedemptionEveryOrderExchangeNum int64                                                  `json:"redemptionEveryOrderExchangeNum,omitempty"`
	EachLimitNum                    int64                                                  `json:"eachLimitNum,omitempty"`
	TopicStyleType                  int64                                                  `json:"topicStyleType,omitempty"`
	PromotionImage                  string                                                 `json:"promotionImage,omitempty"`
	PromotionDescription            string                                                 `json:"promotionDescription,omitempty"`
	PromotionTag                    int64                                                  `json:"promotionTag,omitempty"`
	PromotionTagName                string                                                 `json:"promotionTagName,omitempty"`
	ShowLinePrice                   bool                                                   `json:"showLinePrice,omitempty"`
	VerifyPhoneNumber               int64                                                  `json:"verifyPhoneNumber,omitempty"`
	ConditionType                   int64                                                  `json:"conditionType,omitempty"`
	CalculateType                   int64                                                  `json:"calculateType,omitempty"`
	ActivityId                      int64                                                  `json:"activityId,omitempty"`
	ActivityType                    int64                                                  `json:"activityType,omitempty"`
	PromotionName                   string                                                 `json:"promotionName,omitempty"`
	StartTime                       int64                                                  `json:"startTime,omitempty"`
	EndTime                         int64                                                  `json:"endTime,omitempty"`
	MarkingType                     int64                                                  `json:"markingType,omitempty"`
	TagRuleId                       int64                                                  `json:"tagRuleId,omitempty"`
	SelectPeopleType                int64                                                  `json:"selectPeopleType,omitempty"`
	SelectPeopleRuleId              int64                                                  `json:"selectPeopleRuleId,omitempty"`
	ApplicationGoodsType            int64                                                  `json:"applicationGoodsType,omitempty"`
	CheckType                       int64                                                  `json:"checkType,omitempty"`
	SubIds                          []int64                                                `json:"subIds,omitempty"`
	SelectStoreType                 int64                                                  `json:"selectStoreType,omitempty"`
	SelectStoreIds                  []int64                                                `json:"selectStoreIds,omitempty"`
	TradeScene                      []int64                                                `json:"tradeScene,omitempty"`
	AvailableDeduct                 []int64                                                `json:"availableDeduct,omitempty"`
	AvailablePromotion              []int64                                                `json:"availablePromotion,omitempty"`
}

func CreatePromotionRedemptionGetRequest() (request *PromotionRedemptionGetRequest) {
	request = &PromotionRedemptionGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "promotion/redemption/get", "apigw")
	request.Method = api.POST
	return
}

func CreatePromotionRedemptionGetResponse() (response *api.BaseResponse[PromotionRedemptionGetResponse]) {
	response = api.CreateResponse[PromotionRedemptionGetResponse](&PromotionRedemptionGetResponse{})
	return
}

type PromotionRedemptionGetRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type PromotionRedemptionGetResponseResultPolymerBizVOList struct {
	ResultLimitInfo       PromotionRedemptionGetResponseResultLimitInfo         `json:"resultLimitInfo,omitempty"`
	PromotionSkuBizVOList []PromotionRedemptionGetResponsePromotionSkuBizVOList `json:"promotionSkuBizVOList,omitempty"`
	ResultType            int64                                                 `json:"resultType,omitempty"`
	ResultValue           int64                                                 `json:"resultValue,omitempty"`
	ConditionOperator     int64                                                 `json:"conditionOperator,omitempty"`
	ConditionValue        int64                                                 `json:"conditionValue,omitempty"`
	MaxAmount             int64                                                 `json:"maxAmount,omitempty"`
	ItemType              int64                                                 `json:"itemType,omitempty"`
	ItemId                int64                                                 `json:"itemId,omitempty"`
	LimitDistrictType     int64                                                 `json:"limitDistrictType,omitempty"`
	District              string                                                `json:"district,omitempty"`
}

type PromotionRedemptionGetResponseResultLimitInfo struct {
	NodeId        string `json:"nodeId,omitempty"`
	Path          string `json:"path,omitempty"`
	LimitBizId    int64  `json:"limitBizId,omitempty"`
	LimitBizType  int64  `json:"limitBizType,omitempty"`
	LimitType     int64  `json:"limitType,omitempty"`
	EachLimitNum  int64  `json:"eachLimitNum,omitempty"`
	TotalLimitNum int64  `json:"totalLimitNum,omitempty"`
	EachSoldNum   int64  `json:"eachSoldNum,omitempty"`
	TotalSoldNum  int64  `json:"totalSoldNum,omitempty"`
	CanBuyNum     int64  `json:"canBuyNum,omitempty"`
}

type PromotionRedemptionGetResponsePromotionSkuBizVOList struct {
	GoodsId       int64 `json:"goodsId,omitempty"`
	SkuId         int64 `json:"skuId,omitempty"`
	ActivityPrice int64 `json:"activityPrice,omitempty"`
	CanBuyNum     int64 `json:"canBuyNum,omitempty"`
}

type PromotionRedemptionGetResponseExchangePurchaseSetting struct {
	RuleList      []PromotionRedemptionGetResponseRuleList `json:"ruleList,omitempty"`
	GlobalGroupId int64                                    `json:"globalGroupId,omitempty"`
	GroupPriority int64                                    `json:"groupPriority,omitempty"`
	ConditionType int64                                    `json:"conditionType,omitempty"`
}

type PromotionRedemptionGetResponseRuleList struct {
	GoodsList     []PromotionRedemptionGetResponseGoodsList `json:"goodsList,omitempty"`
	GlobalLevelId int64                                     `json:"globalLevelId,omitempty"`
	LevelPriority int64                                     `json:"levelPriority,omitempty"`
	PurchaseLimit string                                    `json:"purchaseLimit,omitempty"`
}

type PromotionRedemptionGetResponseGoodsList struct {
	SkuList []PromotionRedemptionGetResponseSkuList `json:"skuList,omitempty"`
	GoodsId int64                                   `json:"goodsId,omitempty"`
}

type PromotionRedemptionGetResponseSkuList struct {
	SkuId         int64 `json:"skuId,omitempty"`
	SalePrice     int64 `json:"salePrice,omitempty"`
	ExchangePrice int64 `json:"exchangePrice,omitempty"`
	ExchangeTotal int64 `json:"exchangeTotal,omitempty"`
	ExchangeLimit int64 `json:"exchangeLimit,omitempty"`
}
