package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionQueryDiscountGoodsList
// @id 1788
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询限时折扣商品列表)
func (client *Client) PromotionQueryDiscountGoodsList(request *PromotionQueryDiscountGoodsListRequest) (response *PromotionQueryDiscountGoodsListResponse, err error) {
	rpcResponse := CreatePromotionQueryDiscountGoodsListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionQueryDiscountGoodsListRequest struct {
	*api.BaseRequest

	DiscountId  int64 `json:"discountId,omitempty"`
	PageSize    int   `json:"pageSize,omitempty"`
	PageNum     int   `json:"pageNum,omitempty"`
	GoodsIdList int64 `json:"goodsIdList,omitempty"`
}

type PromotionQueryDiscountGoodsListResponse struct {
}

func CreatePromotionQueryDiscountGoodsListRequest() (request *PromotionQueryDiscountGoodsListRequest) {
	request = &PromotionQueryDiscountGoodsListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "promotion/queryDiscountGoodsList", "api")
	request.Method = api.POST
	return
}

func CreatePromotionQueryDiscountGoodsListResponse() (response *api.BaseResponse[PromotionQueryDiscountGoodsListResponse]) {
	response = api.CreateResponse[PromotionQueryDiscountGoodsListResponse](&PromotionQueryDiscountGoodsListResponse{})
	return
}
