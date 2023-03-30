package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionUpdateDiscountGoods
// @id 2627
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=限时折扣商品更新)
func (client *Client) PromotionUpdateDiscountGoods(request *PromotionUpdateDiscountGoodsRequest) (response *PromotionUpdateDiscountGoodsResponse, err error) {
	rpcResponse := CreatePromotionUpdateDiscountGoodsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionUpdateDiscountGoodsRequest struct {
	*api.BaseRequest

	GoodsAddVOList []PromotionUpdateDiscountGoodsRequestGoodsAddVOList `json:"goodsAddVOList,omitempty"`
	DiscountId     int64                                               `json:"discountId,omitempty"`
}

type PromotionUpdateDiscountGoodsResponse struct {
	FailResult []PromotionUpdateDiscountGoodsResponseFailResult `json:"failResult,omitempty"`
}

func CreatePromotionUpdateDiscountGoodsRequest() (request *PromotionUpdateDiscountGoodsRequest) {
	request = &PromotionUpdateDiscountGoodsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "promotion/updateDiscountGoods", "api")
	request.Method = api.POST
	return
}

func CreatePromotionUpdateDiscountGoodsResponse() (response *api.BaseResponse[PromotionUpdateDiscountGoodsResponse]) {
	response = api.CreateResponse[PromotionUpdateDiscountGoodsResponse](&PromotionUpdateDiscountGoodsResponse{})
	return
}

type PromotionUpdateDiscountGoodsRequestGoodsAddVOList struct {
	SkuList               []PromotionUpdateDiscountGoodsRequestSkuList `json:"skuList,omitempty"`
	GoodsId               int64                                        `json:"goodsId,omitempty"`
	DiscountSort          int                                          `json:"discountSort,omitempty"`
	DiscountStoreLimitNum int                                          `json:"discountStoreLimitNum,omitempty"`
	DiscountType          int                                          `json:"discountType,omitempty"`
	IgnoreChangeType      int                                          `json:"ignoreChangeType,omitempty"`
	InitSubscribeNum      int                                          `json:"initSubscribeNum,omitempty"`
	InitActivitySalesNum  int                                          `json:"initActivitySalesNum,omitempty"`
}

type PromotionUpdateDiscountGoodsRequestSkuList struct {
	SkuId            int64   `json:"skuId,omitempty"`
	ChangeCanSaleNum int     `json:"changeCanSaleNum,omitempty"`
	Discount         float64 `json:"discount,omitempty"`
}

type PromotionUpdateDiscountGoodsResponseFailResult struct {
	GoodsId int64  `json:"goodsId,omitempty"`
	Message string `json:"message,omitempty"`
}
