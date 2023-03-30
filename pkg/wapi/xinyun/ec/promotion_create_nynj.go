package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionCreateNynj
// @id 1603
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=创建N元N件活动)
func (client *Client) PromotionCreateNynj(request *PromotionCreateNynjRequest) (response *PromotionCreateNynjResponse, err error) {
	rpcResponse := CreatePromotionCreateNynjResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionCreateNynjRequest struct {
	*api.BaseRequest

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

type PromotionCreateNynjResponse struct {
}

func CreatePromotionCreateNynjRequest() (request *PromotionCreateNynjRequest) {
	request = &PromotionCreateNynjRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "promotion/createNynj", "api")
	request.Method = api.POST
	return
}

func CreatePromotionCreateNynjResponse() (response *api.BaseResponse[PromotionCreateNynjResponse]) {
	response = api.CreateResponse[PromotionCreateNynjResponse](&PromotionCreateNynjResponse{})
	return
}
