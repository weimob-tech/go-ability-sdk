package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionDiscountCreate
// @id 1807
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1807?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=限时折扣创建)
func (client *Client) PromotionDiscountCreate(request *PromotionDiscountCreateRequest) (response *PromotionDiscountCreateResponse, err error) {
	rpcResponse := CreatePromotionDiscountCreateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionDiscountCreateRequest struct {
	*api.BaseRequest

	DiscountGoodsAddVOS    []PromotionDiscountCreateRequestDiscountGoodsAddVOS `json:"discountGoodsAddVOS,omitempty"`
	BasicInfo              PromotionDiscountCreateRequestBasicInfo             `json:"basicInfo,omitempty"`
	Title                  string                                              `json:"title,omitempty"`
	TimeType               int64                                               `json:"timeType,omitempty"`
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
}

type PromotionDiscountCreateResponse struct {
	ResultVOS  []PromotionDiscountCreateResponseResultVOS `json:"resultVOS,omitempty"`
	ActivityId int64                                      `json:"activityId,omitempty"`
}

func CreatePromotionDiscountCreateRequest() (request *PromotionDiscountCreateRequest) {
	request = &PromotionDiscountCreateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "promotion/discount/create", "apigw")
	request.Method = api.POST
	return
}

func CreatePromotionDiscountCreateResponse() (response *api.BaseResponse[PromotionDiscountCreateResponse]) {
	response = api.CreateResponse[PromotionDiscountCreateResponse](&PromotionDiscountCreateResponse{})
	return
}

type PromotionDiscountCreateRequestDiscountGoodsAddVOS struct {
	SkuList              []PromotionDiscountCreateRequestSkuList `json:"skuList,omitempty"`
	GoodsId              int64                                   `json:"goodsId,omitempty"`
	Sort                 int64                                   `json:"sort,omitempty"`
	AvailableStockNum    int64                                   `json:"availableStockNum,omitempty"`
	DiscountType         int64                                   `json:"discountType,omitempty"`
	IgnoreChangeType     int64                                   `json:"ignoreChangeType,omitempty"`
	InitSubscribeNum     int64                                   `json:"initSubscribeNum,omitempty"`
	InitActivitySalesNum int64                                   `json:"initActivitySalesNum,omitempty"`
}

type PromotionDiscountCreateRequestSkuList struct {
	SkuId            int64 `json:"skuId,omitempty"`
	ChangeCanSaleNum int64 `json:"changeCanSaleNum,omitempty"`
	Discount         int64 `json:"discount,omitempty"`
}

type PromotionDiscountCreateRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type PromotionDiscountCreateResponseResultVOS struct {
	ErrorMsg string  `json:"errorMsg,omitempty"`
	GoodsId  []int64 `json:"goodsId,omitempty"`
}
