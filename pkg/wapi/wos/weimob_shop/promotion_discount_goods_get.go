package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionDiscountGoodsGet
// @id 1927
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1927?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询限时折扣商品详情)
func (client *Client) PromotionDiscountGoodsGet(request *PromotionDiscountGoodsGetRequest) (response *PromotionDiscountGoodsGetResponse, err error) {
	rpcResponse := CreatePromotionDiscountGoodsGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionDiscountGoodsGetRequest struct {
	*api.BaseRequest

	QueryParameter PromotionDiscountGoodsGetRequestQueryParameter `json:"queryParameter,omitempty"`
	BasicInfo      PromotionDiscountGoodsGetRequestBasicInfo      `json:"basicInfo,omitempty"`
	PageNum        int64                                          `json:"pageNum,omitempty"`
	PageSize       int64                                          `json:"pageSize,omitempty"`
}

type PromotionDiscountGoodsGetResponse struct {
	PageList   []PromotionDiscountGoodsGetResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                       `json:"pageNum,omitempty"`
	PageSize   int64                                       `json:"pageSize,omitempty"`
	TotalCount int64                                       `json:"totalCount,omitempty"`
}

func CreatePromotionDiscountGoodsGetRequest() (request *PromotionDiscountGoodsGetRequest) {
	request = &PromotionDiscountGoodsGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "promotion/discount/goods/get", "apigw")
	request.Method = api.POST
	return
}

func CreatePromotionDiscountGoodsGetResponse() (response *api.BaseResponse[PromotionDiscountGoodsGetResponse]) {
	response = api.CreateResponse[PromotionDiscountGoodsGetResponse](&PromotionDiscountGoodsGetResponse{})
	return
}

type PromotionDiscountGoodsGetRequestQueryParameter struct {
	ActivityId   int64   `json:"activityId,omitempty"`
	ActivityType int64   `json:"activityType,omitempty"`
	GoodsIdList  []int64 `json:"goodsIdList,omitempty"`
}

type PromotionDiscountGoodsGetRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type PromotionDiscountGoodsGetResponsePageList struct {
	ItemList     []PromotionDiscountGoodsGetResponseItemList `json:"itemList,omitempty"`
	GoodsId      int64                                       `json:"goodsId,omitempty"`
	ItemType     int64                                       `json:"itemType,omitempty"`
	TruncType    int64                                       `json:"truncType,omitempty"`
	EachLimitNum int64                                       `json:"eachLimitNum,omitempty"`
	Sort         int64                                       `json:"sort,omitempty"`
	Follower     int64                                       `json:"follower,omitempty"`
	Sold         int64                                       `json:"sold,omitempty"`
}

type PromotionDiscountGoodsGetResponseItemList struct {
	SkuId                     int64 `json:"skuId,omitempty"`
	ItemId                    int64 `json:"itemId,omitempty"`
	ResultType                int64 `json:"resultType,omitempty"`
	ResultValue               int64 `json:"resultValue,omitempty"`
	AvailableActivityStockNum int64 `json:"availableActivityStockNum,omitempty"`
	ActivityStockNum          int64 `json:"activityStockNum,omitempty"`
}
