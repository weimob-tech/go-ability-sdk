package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionDiscountUpdate
// @id 1806
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1806?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=限时折扣修改接口)
func (client *Client) PromotionDiscountUpdate(request *PromotionDiscountUpdateRequest) (response *PromotionDiscountUpdateResponse, err error) {
	rpcResponse := CreatePromotionDiscountUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionDiscountUpdateRequest struct {
	*api.BaseRequest

	DiscountGoodsAddVOS    []PromotionDiscountUpdateRequestDiscountGoodsAddVOS `json:"discountGoodsAddVOS,omitempty"`
	BasicInfo              PromotionDiscountUpdateRequestBasicInfo             `json:"basicInfo,omitempty"`
	Title                  string                                              `json:"title,omitempty"`
	StartDate              string                                              `json:"startDate,omitempty"`
	EndDate                string                                              `json:"endDate,omitempty"`
	RepeatType             int64                                               `json:"repeatType,omitempty"`
	RepeatStartInterval    string                                              `json:"repeatStartInterval,omitempty"`
	RepeatEndInterval      string                                              `json:"repeatEndInterval,omitempty"`
	RepeatWeekDays         []int64                                             `json:"repeatWeekDays,omitempty"`
	RepeatDay              int64                                               `json:"repeatDay,omitempty"`
	SyncTag                int64                                               `json:"syncTag,omitempty"`
	SyncTagId              int64                                               `json:"syncTagId,omitempty"`
	LimitType              int64                                               `json:"limitType,omitempty"`
	LimitNum               int64                                               `json:"limitNum,omitempty"`
	OrderAutoClosedTime    int64                                               `json:"orderAutoClosedTime,omitempty"`
	SelectStoreType        int64                                               `json:"selectStoreType,omitempty"`
	SelectStoreIds         []int64                                             `json:"selectStoreIds,omitempty"`
	SelectPeopleType       int64                                               `json:"selectPeopleType,omitempty"`
	SelectPeopleId         int64                                               `json:"selectPeopleId,omitempty"`
	UseOfScene             string                                              `json:"useOfScene,omitempty"`
	CanDeductionType       string                                              `json:"canDeductionType,omitempty"`
	LimitPromotionType     string                                              `json:"limitPromotionType,omitempty"`
	WarmUpType             int64                                               `json:"warmUpType,omitempty"`
	WarmUpDate             string                                              `json:"warmUpDate,omitempty"`
	OriginPriceBuyInWarmUp int64                                               `json:"originPriceBuyInWarmUp,omitempty"`
	TagName                string                                              `json:"tagName,omitempty"`
	ActivityCountdown      int64                                               `json:"activityCountdown,omitempty"`
	ImageUrl               string                                              `json:"imageUrl,omitempty"`
	Description            string                                              `json:"description,omitempty"`
	ActivityId             int64                                               `json:"activityId,omitempty"`
}

type PromotionDiscountUpdateResponse struct {
	ResultVOS  []PromotionDiscountUpdateResponseResultVOS `json:"resultVOS,omitempty"`
	ActivityId string                                     `json:"activityId,omitempty"`
}

func CreatePromotionDiscountUpdateRequest() (request *PromotionDiscountUpdateRequest) {
	request = &PromotionDiscountUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "promotion/discount/update", "apigw")
	request.Method = api.POST
	return
}

func CreatePromotionDiscountUpdateResponse() (response *api.BaseResponse[PromotionDiscountUpdateResponse]) {
	response = api.CreateResponse[PromotionDiscountUpdateResponse](&PromotionDiscountUpdateResponse{})
	return
}

type PromotionDiscountUpdateRequestDiscountGoodsAddVOS struct {
	SkuList              []PromotionDiscountUpdateRequestSkuList `json:"skuList,omitempty"`
	GoodsId              int64                                   `json:"goodsId,omitempty"`
	Sort                 int64                                   `json:"sort,omitempty"`
	AvailableStockNum    int64                                   `json:"availableStockNum,omitempty"`
	DiscountType         int64                                   `json:"discountType,omitempty"`
	IgnoreChangeType     int64                                   `json:"ignoreChangeType,omitempty"`
	InitSubscribeNum     int64                                   `json:"initSubscribeNum,omitempty"`
	InitActivitySalesNum int64                                   `json:"initActivitySalesNum,omitempty"`
}

type PromotionDiscountUpdateRequestSkuList struct {
	SkuId            int64 `json:"skuId,omitempty"`
	ChangeCanSaleNum int64 `json:"changeCanSaleNum,omitempty"`
	Discount         int64 `json:"discount,omitempty"`
}

type PromotionDiscountUpdateRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type PromotionDiscountUpdateResponseResultVOS struct {
	ErrorMsg string  `json:"errorMsg,omitempty"`
	GoodsId  []int64 `json:"goodsId,omitempty"`
}
