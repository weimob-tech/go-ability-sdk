package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionGiftUpdate
// @id 1808
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1808?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=订单满赠修改接口)
func (client *Client) PromotionGiftUpdate(request *PromotionGiftUpdateRequest) (response *PromotionGiftUpdateResponse, err error) {
	rpcResponse := CreatePromotionGiftUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionGiftUpdateRequest struct {
	*api.BaseRequest

	GiftMarketRuleLevelVOS []PromotionGiftUpdateRequestGiftMarketRuleLevelVOS `json:"giftMarketRuleLevelVOS,omitempty"`
	BasicInfo              PromotionGiftUpdateRequestBasicInfo                `json:"basicInfo,omitempty"`
	ActivityId             int64                                              `json:"activityId,omitempty"`
	Title                  string                                             `json:"title,omitempty"`
	StartDate              string                                             `json:"startDate,omitempty"`
	EndDate                string                                             `json:"endDate,omitempty"`
	SyncTag                int64                                              `json:"syncTag,omitempty"`
	SyncTagId              int64                                              `json:"syncTagId,omitempty"`
	SelectStoreType        int64                                              `json:"selectStoreType,omitempty"`
	SelectStoreIds         []int64                                            `json:"selectStoreIds,omitempty"`
	SelectPeopleType       int64                                              `json:"selectPeopleType,omitempty"`
	SelectPeopleId         int64                                              `json:"selectPeopleId,omitempty"`
	SelectGoodsType        int64                                              `json:"selectGoodsType,omitempty"`
	SelectGoodsIds         []int64                                            `json:"selectGoodsIds,omitempty"`
	Join                   int64                                              `json:"join,omitempty"`
	UseOfScene             string                                             `json:"useOfScene,omitempty"`
	CanDeductionType       string                                             `json:"canDeductionType,omitempty"`
	LimitPromotionType     string                                             `json:"limitPromotionType,omitempty"`
	ProjectPageStyle       int64                                              `json:"projectPageStyle,omitempty"`
	ImageUrl               string                                             `json:"imageUrl,omitempty"`
	Description            string                                             `json:"description,omitempty"`
	FullGiftFactor         int64                                              `json:"fullGiftFactor,omitempty"`
	FullGiftType           int64                                              `json:"fullGiftType,omitempty"`
	LimitNum               int64                                              `json:"limitNum,omitempty"`
}

type PromotionGiftUpdateResponse struct {
	ActivityId int64 `json:"activityId,omitempty"`
}

func CreatePromotionGiftUpdateRequest() (request *PromotionGiftUpdateRequest) {
	request = &PromotionGiftUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "promotion/gift/update", "apigw")
	request.Method = api.POST
	return
}

func CreatePromotionGiftUpdateResponse() (response *api.BaseResponse[PromotionGiftUpdateResponse]) {
	response = api.CreateResponse[PromotionGiftUpdateResponse](&PromotionGiftUpdateResponse{})
	return
}

type PromotionGiftUpdateRequestGiftMarketRuleLevelVOS struct {
	GiftMarketGoodsVOS []PromotionGiftUpdateRequestGiftMarketGoodsVOS `json:"giftMarketGoodsVOS,omitempty"`
	FullGiftCondition  int64                                          `json:"fullGiftCondition,omitempty"`
	FullGiftLimit      int64                                          `json:"fullGiftLimit,omitempty"`
	LevelPriority      int64                                          `json:"levelPriority,omitempty"`
}

type PromotionGiftUpdateRequestGiftMarketGoodsVOS struct {
	GoodsId     int64 `json:"goodsId,omitempty"`
	SkuId       int64 `json:"skuId,omitempty"`
	SingleLimit int64 `json:"singleLimit,omitempty"`
}

type PromotionGiftUpdateRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}
