package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionCustomerpriceGoodsGet
// @id 1460
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1460?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=客群价活动商品列表查询)
func (client *Client) PromotionCustomerpriceGoodsGet(request *PromotionCustomerpriceGoodsGetRequest) (response *PromotionCustomerpriceGoodsGetResponse, err error) {
	rpcResponse := CreatePromotionCustomerpriceGoodsGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionCustomerpriceGoodsGetRequest struct {
	*api.BaseRequest

	QueryParameter PromotionCustomerpriceGoodsGetRequestQueryParameter `json:"queryParameter,omitempty"`
	BasicInfo      PromotionCustomerpriceGoodsGetRequestBasicInfo      `json:"basicInfo,omitempty"`
	PageNum        int64                                               `json:"pageNum,omitempty"`
	PageSize       int64                                               `json:"pageSize,omitempty"`
}

type PromotionCustomerpriceGoodsGetResponse struct {
	PageList   []PromotionCustomerpriceGoodsGetResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                            `json:"pageNum,omitempty"`
	PageSize   int64                                            `json:"pageSize,omitempty"`
	TotalCount int64                                            `json:"totalCount,omitempty"`
}

func CreatePromotionCustomerpriceGoodsGetRequest() (request *PromotionCustomerpriceGoodsGetRequest) {
	request = &PromotionCustomerpriceGoodsGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "promotion/customerprice/goods/get", "apigw")
	request.Method = api.POST
	return
}

func CreatePromotionCustomerpriceGoodsGetResponse() (response *api.BaseResponse[PromotionCustomerpriceGoodsGetResponse]) {
	response = api.CreateResponse[PromotionCustomerpriceGoodsGetResponse](&PromotionCustomerpriceGoodsGetResponse{})
	return
}

type PromotionCustomerpriceGoodsGetRequestQueryParameter struct {
	ActivityId   int64   `json:"activityId,omitempty"`
	ActivityType int64   `json:"activityType,omitempty"`
	GoodsIdList  []int64 `json:"goodsIdList,omitempty"`
}

type PromotionCustomerpriceGoodsGetRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type PromotionCustomerpriceGoodsGetResponsePageList struct {
	ItemList  []PromotionCustomerpriceGoodsGetResponseItemList `json:"itemList,omitempty"`
	GoodsId   int64                                            `json:"goodsId,omitempty"`
	ItemType  int64                                            `json:"itemType,omitempty"`
	TruncType int64                                            `json:"truncType,omitempty"`
}

type PromotionCustomerpriceGoodsGetResponseItemList struct {
	SkuId       int64 `json:"skuId,omitempty"`
	ItemId      int64 `json:"itemId,omitempty"`
	ResultType  int64 `json:"resultType,omitempty"`
	ResultValue int64 `json:"resultValue,omitempty"`
}
