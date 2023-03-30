package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionCustomerpriceUpdate
// @id 1459
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1459?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=客群价活动商品更新)
func (client *Client) PromotionCustomerpriceUpdate(request *PromotionCustomerpriceUpdateRequest) (response *PromotionCustomerpriceUpdateResponse, err error) {
	rpcResponse := CreatePromotionCustomerpriceUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionCustomerpriceUpdateRequest struct {
	*api.BaseRequest

	GoodsAddVoList []PromotionCustomerpriceUpdateRequestGoodsAddVoList `json:"goodsAddVoList,omitempty"`
	BasicInfo      PromotionCustomerpriceUpdateRequestBasicInfo        `json:"basicInfo,omitempty"`
	ActivityId     int64                                               `json:"activityId,omitempty"`
}

type PromotionCustomerpriceUpdateResponse struct {
}

func CreatePromotionCustomerpriceUpdateRequest() (request *PromotionCustomerpriceUpdateRequest) {
	request = &PromotionCustomerpriceUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "promotion/customerprice/update", "apigw")
	request.Method = api.POST
	return
}

func CreatePromotionCustomerpriceUpdateResponse() (response *api.BaseResponse[PromotionCustomerpriceUpdateResponse]) {
	response = api.CreateResponse[PromotionCustomerpriceUpdateResponse](&PromotionCustomerpriceUpdateResponse{})
	return
}

type PromotionCustomerpriceUpdateRequestGoodsAddVoList struct {
	SkuLevels        []PromotionCustomerpriceUpdateRequestSkuLevels `json:"skuLevels,omitempty"`
	GoodsId          int64                                          `json:"goodsId,omitempty"`
	DiscountType     int64                                          `json:"discountType,omitempty"`
	IgnoreChangeType int64                                          `json:"IgnoreChangeType,omitempty"`
}

type PromotionCustomerpriceUpdateRequestSkuLevels struct {
	SkuId    int64 `json:"skuId,omitempty"`
	Discount int64 `json:"discount,omitempty"`
}

type PromotionCustomerpriceUpdateRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}
