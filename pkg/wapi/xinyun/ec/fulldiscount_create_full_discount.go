package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FulldiscountCreateFullDiscount
// @id 1600
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=创建满减满折活动)
func (client *Client) FulldiscountCreateFullDiscount(request *FulldiscountCreateFullDiscountRequest) (response *FulldiscountCreateFullDiscountResponse, err error) {
	rpcResponse := CreateFulldiscountCreateFullDiscountResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FulldiscountCreateFullDiscountRequest struct {
	*api.BaseRequest

	RuleList            []FulldiscountCreateFullDiscountRequestRuleList `json:"ruleList,omitempty"`
	PromotionName       string                                          `json:"promotionName,omitempty"`
	StartTime           int                                             `json:"startTime,omitempty"`
	EndTime             int                                             `json:"endTime,omitempty"`
	ReduceType          int                                             `json:"reduceType,omitempty"`
	ConditionType       int                                             `json:"conditionType,omitempty"`
	ActionType          int                                             `json:"actionType,omitempty"`
	DeductOverlay       []int                                           `json:"deductOverlay,omitempty"`
	PromotionOverlay    []int                                           `json:"promotionOverlay,omitempty"`
	SelectGoodsType     int                                             `json:"selectGoodsType,omitempty"`
	SelectGoodsIds      []int64                                         `json:"selectGoodsIds,omitempty"`
	ExcludeGoodsIds     []int64                                         `json:"excludeGoodsIds,omitempty"`
	SelectStoreType     int                                             `json:"selectStoreType,omitempty"`
	SelectStoreIds      []int64                                         `json:"selectStoreIds,omitempty"`
	SelectTradeSceneIds []int                                           `json:"selectTradeSceneIds,omitempty"`
	TopicPageType       int                                             `json:"topicPageType,omitempty"`
	PromotionTagId      int64                                           `json:"promotionTagId,omitempty"`
	ImageUrl            string                                          `json:"imageUrl,omitempty"`
	Description         string                                          `json:"description,omitempty"`
}

type FulldiscountCreateFullDiscountResponse struct {
}

func CreateFulldiscountCreateFullDiscountRequest() (request *FulldiscountCreateFullDiscountRequest) {
	request = &FulldiscountCreateFullDiscountRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "fulldiscount/createFullDiscount", "api")
	request.Method = api.POST
	return
}

func CreateFulldiscountCreateFullDiscountResponse() (response *api.BaseResponse[FulldiscountCreateFullDiscountResponse]) {
	response = api.CreateResponse[FulldiscountCreateFullDiscountResponse](&FulldiscountCreateFullDiscountResponse{})
	return
}

type FulldiscountCreateFullDiscountRequestRuleList struct {
	Level           int     `json:"level,omitempty"`
	ConditionValue  float64 `json:"conditionValue,omitempty"`
	ActionValue     float64 `json:"actionValue,omitempty"`
	ActionMaxAmount float64 `json:"actionMaxAmount,omitempty"`
}
