package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionDiscountGet
// @id 1928
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1928?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询限时折扣详情)
func (client *Client) PromotionDiscountGet(request *PromotionDiscountGetRequest) (response *PromotionDiscountGetResponse, err error) {
	rpcResponse := CreatePromotionDiscountGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionDiscountGetRequest struct {
	*api.BaseRequest

	BasicInfo  PromotionDiscountGetRequestBasicInfo `json:"basicInfo,omitempty"`
	ActivityId int64                                `json:"activityId,omitempty"`
}

type PromotionDiscountGetResponse struct {
	CycleTimeList              []PromotionDiscountGetResponseCycleTimeList `json:"cycleTimeList,omitempty"`
	TimeType                   int64                                       `json:"timeType,omitempty"`
	OverLimitCanOriginPriceBuy int64                                       `json:"overLimitCanOriginPriceBuy,omitempty"`
	EachLimitNum               int64                                       `json:"eachLimitNum,omitempty"`
	OrderCloseTime             int64                                       `json:"orderCloseTime,omitempty"`
	ScopeGoodsIdList           []int64                                     `json:"scopeGoodsIdList,omitempty"`
	PrepareType                int64                                       `json:"prepareType,omitempty"`
	PrepareTime                int64                                       `json:"prepareTime,omitempty"`
	IsPreheatCanOriginPriceBuy int64                                       `json:"isPreheatCanOriginPriceBuy,omitempty"`
	PromotionTag               int64                                       `json:"promotionTag,omitempty"`
	PromotionTagName           string                                      `json:"promotionTagName,omitempty"`
	PromotionDescription       string                                      `json:"promotionDescription,omitempty"`
	PromotionImage             string                                      `json:"promotionImage,omitempty"`
	PromotionStatus            int64                                       `json:"promotionStatus,omitempty"`
	ActivityId                 int64                                       `json:"activityId,omitempty"`
	ActivityType               int64                                       `json:"activityType,omitempty"`
	PromotionName              string                                      `json:"promotionName,omitempty"`
	StartTime                  int64                                       `json:"startTime,omitempty"`
	EndTime                    int64                                       `json:"endTime,omitempty"`
	MarkingType                int64                                       `json:"markingType,omitempty"`
	TagRuleId                  int64                                       `json:"tagRuleId,omitempty"`
	SelectPeopleType           int64                                       `json:"selectPeopleType,omitempty"`
	SelectPeopleRuleId         int64                                       `json:"selectPeopleRuleId,omitempty"`
	ApplicationGoodsType       int64                                       `json:"applicationGoodsType,omitempty"`
	CheckType                  int64                                       `json:"checkType,omitempty"`
	SubIds                     []int64                                     `json:"subIds,omitempty"`
	SelectStoreType            int64                                       `json:"selectStoreType,omitempty"`
	SelectStoreIds             []int64                                     `json:"selectStoreIds,omitempty"`
	TradeScene                 []int64                                     `json:"tradeScene,omitempty"`
	AvailableDeduct            []int64                                     `json:"availableDeduct,omitempty"`
	AvailablePromotion         []int64                                     `json:"availablePromotion,omitempty"`
}

func CreatePromotionDiscountGetRequest() (request *PromotionDiscountGetRequest) {
	request = &PromotionDiscountGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "promotion/discount/get", "apigw")
	request.Method = api.POST
	return
}

func CreatePromotionDiscountGetResponse() (response *api.BaseResponse[PromotionDiscountGetResponse]) {
	response = api.CreateResponse[PromotionDiscountGetResponse](&PromotionDiscountGetResponse{})
	return
}

type PromotionDiscountGetRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type PromotionDiscountGetResponseCycleTimeList struct {
	RepeatType          int64  `json:"repeatType,omitempty"`
	RepeatDay           string `json:"repeatDay,omitempty"`
	RepeatStartInterval string `json:"repeatStartInterval,omitempty"`
	RepeatEndInterval   string `json:"repeatEndInterval,omitempty"`
}
