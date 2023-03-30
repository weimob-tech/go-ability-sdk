package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionQueryCombinationGoodsList
// @id 1790
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询优惠套装活动商品列表)
func (client *Client) PromotionQueryCombinationGoodsList(request *PromotionQueryCombinationGoodsListRequest) (response *PromotionQueryCombinationGoodsListResponse, err error) {
	rpcResponse := CreatePromotionQueryCombinationGoodsListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionQueryCombinationGoodsListRequest struct {
	*api.BaseRequest

	ActivityId  int64   `json:"activityId,omitempty"`
	GroupIdList []int64 `json:"groupIdList,omitempty"`
	GoodsIdList []int64 `json:"goodsIdList,omitempty"`
}

type PromotionQueryCombinationGoodsListResponse struct {
}

func CreatePromotionQueryCombinationGoodsListRequest() (request *PromotionQueryCombinationGoodsListRequest) {
	request = &PromotionQueryCombinationGoodsListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "promotion/queryCombinationGoodsList", "api")
	request.Method = api.POST
	return
}

func CreatePromotionQueryCombinationGoodsListResponse() (response *api.BaseResponse[PromotionQueryCombinationGoodsListResponse]) {
	response = api.CreateResponse[PromotionQueryCombinationGoodsListResponse](&PromotionQueryCombinationGoodsListResponse{})
	return
}
