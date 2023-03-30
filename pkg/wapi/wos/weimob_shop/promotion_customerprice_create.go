package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionCustomerpriceCreate
// @id 1462
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1462?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=客群价创建/更新)
func (client *Client) PromotionCustomerpriceCreate(request *PromotionCustomerpriceCreateRequest) (response *PromotionCustomerpriceCreateResponse, err error) {
	rpcResponse := CreatePromotionCustomerpriceCreateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionCustomerpriceCreateRequest struct {
	*api.BaseRequest

	GoodsAddVoList          []PromotionCustomerpriceCreateRequestGoodsAddVoList `json:"goodsAddVoList,omitempty"`
	BasicInfo               PromotionCustomerpriceCreateRequestBasicInfo        `json:"basicInfo,omitempty"`
	ActivityId              int64                                               `json:"activityId,omitempty"`
	Title                   string                                              `json:"title,omitempty"`
	TimeType                int64                                               `json:"timeType,omitempty"`
	StartDate               int64                                               `json:"startDate,omitempty"`
	EndDate                 int64                                               `json:"endDate,omitempty"`
	SyncTag                 int64                                               `json:"syncTag,omitempty"`
	SyncTagId               int64                                               `json:"syncTagId,omitempty"`
	DefaultDiscount         int64                                               `json:"defaultDiscount,omitempty"`
	DefaultIgnoreChangeType int64                                               `json:"defaultIgnoreChangeType,omitempty"`
	UseOfScene              string                                              `json:"useOfScene,omitempty"`
	SelectStoreType         int64                                               `json:"selectStoreType,omitempty"`
	SelectStoreIds          []int64                                             `json:"selectStoreIds,omitempty"`
	SelectPeopleType        int64                                               `json:"selectPeopleType,omitempty"`
	SelectPeopleId          int64                                               `json:"selectPeopleId,omitempty"`
	CanDeductionType        string                                              `json:"canDeductionType,omitempty"`
	LimitPromotionType      string                                              `json:"limitPromotionType,omitempty"`
	Uuid                    string                                              `json:"uuid,omitempty"`
}

type PromotionCustomerpriceCreateResponse struct {
	ActivityId  int64   `json:"activityId,omitempty"`
	FailGoodsId []int64 `json:"failGoodsId,omitempty"`
}

func CreatePromotionCustomerpriceCreateRequest() (request *PromotionCustomerpriceCreateRequest) {
	request = &PromotionCustomerpriceCreateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "promotion/customerprice/create", "apigw")
	request.Method = api.POST
	return
}

func CreatePromotionCustomerpriceCreateResponse() (response *api.BaseResponse[PromotionCustomerpriceCreateResponse]) {
	response = api.CreateResponse[PromotionCustomerpriceCreateResponse](&PromotionCustomerpriceCreateResponse{})
	return
}

type PromotionCustomerpriceCreateRequestGoodsAddVoList struct {
	SkuLevels        []PromotionCustomerpriceCreateRequestSkuLevels `json:"skuLevels,omitempty"`
	GoodsId          int64                                          `json:"goodsId,omitempty"`
	DiscountType     int64                                          `json:"discountType,omitempty"`
	IgnoreChangeType int64                                          `json:"IgnoreChangeType,omitempty"`
}

type PromotionCustomerpriceCreateRequestSkuLevels struct {
	SkuId    int64 `json:"skuId,omitempty"`
	Discount int64 `json:"discount,omitempty"`
}

type PromotionCustomerpriceCreateRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}
