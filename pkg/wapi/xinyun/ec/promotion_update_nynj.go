package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionUpdateNynj
// @id 1604
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新N元N件活动)
func (client *Client) PromotionUpdateNynj(request *PromotionUpdateNynjRequest) (response *PromotionUpdateNynjResponse, err error) {
	rpcResponse := CreatePromotionUpdateNynjResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionUpdateNynjRequest struct {
	*api.BaseRequest

	ActivityId          int64   `json:"activityId,omitempty"`
	PromotionName       string  `json:"promotionName,omitempty"`
	StartTime           int     `json:"startTime,omitempty"`
	EndTime             int     `json:"endTime,omitempty"`
	ConditionValue      float64 `json:"conditionValue,omitempty"`
	ActionValue         float64 `json:"actionValue,omitempty"`
	LimitNum            int     `json:"limitNum,omitempty"`
	DeductOverlay       []int   `json:"deductOverlay,omitempty"`
	PromotionOverlay    []int   `json:"promotionOverlay,omitempty"`
	SelectGoodsType     int     `json:"selectGoodsType,omitempty"`
	SelectGoodsIds      []int64 `json:"selectGoodsIds,omitempty"`
	ExcludeGoodsIds     []int64 `json:"excludeGoodsIds,omitempty"`
	SelectStoreType     int     `json:"selectStoreType,omitempty"`
	SelectStoreIds      []int64 `json:"selectStoreIds,omitempty"`
	SelectTradeSceneIds []int   `json:"selectTradeSceneIds,omitempty"`
	TopicPageType       int     `json:"topicPageType,omitempty"`
	PromotionTagId      int64   `json:"promotionTagId,omitempty"`
	ImageUrl            string  `json:"imageUrl,omitempty"`
	Description         string  `json:"description,omitempty"`
}

type PromotionUpdateNynjResponse struct {
}

func CreatePromotionUpdateNynjRequest() (request *PromotionUpdateNynjRequest) {
	request = &PromotionUpdateNynjRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "promotion/updateNynj", "api")
	request.Method = api.POST
	return
}

func CreatePromotionUpdateNynjResponse() (response *api.BaseResponse[PromotionUpdateNynjResponse]) {
	response = api.CreateResponse[PromotionUpdateNynjResponse](&PromotionUpdateNynjResponse{})
	return
}
