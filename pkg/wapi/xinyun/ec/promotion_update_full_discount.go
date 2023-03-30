package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionUpdateFullDiscount
// @id 1601
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新满减满折活动)
func (client *Client) PromotionUpdateFullDiscount(request *PromotionUpdateFullDiscountRequest) (response *PromotionUpdateFullDiscountResponse, err error) {
	rpcResponse := CreatePromotionUpdateFullDiscountResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionUpdateFullDiscountRequest struct {
	*api.BaseRequest

	RuleList            []PromotionUpdateFullDiscountRequestRuleList `json:"ruleList,omitempty"`
	ActivityId          int64                                        `json:"activityId,omitempty"`
	PromotionName       string                                       `json:"promotionName,omitempty"`
	StartTime           int                                          `json:"startTime,omitempty"`
	EndTime             int                                          `json:"endTime,omitempty"`
	DeductOverlay       []int                                        `json:"deductOverlay,omitempty"`
	PromotionOverlay    []int                                        `json:"promotionOverlay,omitempty"`
	SelectGoodsType     int                                          `json:"selectGoodsType,omitempty"`
	SelectGoodsIds      []int64                                      `json:"selectGoodsIds,omitempty"`
	ExcludeGoodsIds     []int64                                      `json:"excludeGoodsIds,omitempty"`
	SelectStoreType     int                                          `json:"selectStoreType,omitempty"`
	SelectStoreIds      []int64                                      `json:"selectStoreIds,omitempty"`
	SelectTradeSceneIds []int                                        `json:"selectTradeSceneIds,omitempty"`
	TopicPageType       int                                          `json:"topicPageType,omitempty"`
	PromotionTagId      int64                                        `json:"promotionTagId,omitempty"`
	ImageUrl            string                                       `json:"imageUrl,omitempty"`
	Description         string                                       `json:"description,omitempty"`
}

type PromotionUpdateFullDiscountResponse struct {
}

func CreatePromotionUpdateFullDiscountRequest() (request *PromotionUpdateFullDiscountRequest) {
	request = &PromotionUpdateFullDiscountRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "promotion/updateFullDiscount", "api")
	request.Method = api.POST
	return
}

func CreatePromotionUpdateFullDiscountResponse() (response *api.BaseResponse[PromotionUpdateFullDiscountResponse]) {
	response = api.CreateResponse[PromotionUpdateFullDiscountResponse](&PromotionUpdateFullDiscountResponse{})
	return
}

type PromotionUpdateFullDiscountRequestRuleList struct {
	Level           int     `json:"level,omitempty"`
	GlobalLevelId   int64   `json:"globalLevelId,omitempty"`
	ConditionValue  float64 `json:"conditionValue,omitempty"`
	ActionValue     float64 `json:"actionValue,omitempty"`
	ActionMaxAmount float64 `json:"actionMaxAmount,omitempty"`
}
