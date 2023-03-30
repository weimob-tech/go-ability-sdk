package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionCustomerpriceGet
// @id 1461
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1461?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=客群价详情查询)
func (client *Client) PromotionCustomerpriceGet(request *PromotionCustomerpriceGetRequest) (response *PromotionCustomerpriceGetResponse, err error) {
	rpcResponse := CreatePromotionCustomerpriceGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionCustomerpriceGetRequest struct {
	*api.BaseRequest

	BasicInfo    PromotionCustomerpriceGetRequestBasicInfo `json:"basicInfo,omitempty"`
	ActivityId   int64                                     `json:"activityId,omitempty"`
	ActivityType int64                                     `json:"activityType,omitempty"`
}

type PromotionCustomerpriceGetResponse struct {
	CycleTimeVos                  []PromotionCustomerpriceGetResponseCycleTimeVos `json:"cycleTimeVos,omitempty"`
	TimeType                      int64                                           `json:"timeType,omitempty"`
	CustomersPriceDefaultDiscount int64                                           `json:"customersPriceDefaultDiscount,omitempty"`
	CustomersPriceTruncType       int64                                           `json:"customersPriceTruncType,omitempty"`
	ActivityId                    int64                                           `json:"activityId,omitempty"`
	ActivityType                  int64                                           `json:"activityType,omitempty"`
	PromotionName                 string                                          `json:"promotionName,omitempty"`
	StartTime                     int64                                           `json:"startTime,omitempty"`
	EndTime                       int64                                           `json:"endTime,omitempty"`
	MarkingType                   int64                                           `json:"markingType,omitempty"`
	TagRuleId                     int64                                           `json:"tagRuleId,omitempty"`
	SelectPeopleType              int64                                           `json:"selectPeopleType,omitempty"`
	SelectPeopleRuleId            int64                                           `json:"selectPeopleRuleId,omitempty"`
	SelectStoreType               int64                                           `json:"selectStoreType,omitempty"`
	SelectStoreIds                []int64                                         `json:"selectStoreIds,omitempty"`
	TradeScene                    []int64                                         `json:"tradeScene,omitempty"`
	AvailableDeduct               []int64                                         `json:"availableDeduct,omitempty"`
	AvailablePromotion            []int64                                         `json:"availablePromotion,omitempty"`
}

func CreatePromotionCustomerpriceGetRequest() (request *PromotionCustomerpriceGetRequest) {
	request = &PromotionCustomerpriceGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "promotion/customerprice/get", "apigw")
	request.Method = api.POST
	return
}

func CreatePromotionCustomerpriceGetResponse() (response *api.BaseResponse[PromotionCustomerpriceGetResponse]) {
	response = api.CreateResponse[PromotionCustomerpriceGetResponse](&PromotionCustomerpriceGetResponse{})
	return
}

type PromotionCustomerpriceGetRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type PromotionCustomerpriceGetResponseCycleTimeVos struct {
	RepeatType          int64  `json:"repeatType,omitempty"`
	RepeatDay           string `json:"repeatDay,omitempty"`
	RepeatStartInterval string `json:"repeatStartInterval,omitempty"`
	RepeatEndInterval   string `json:"repeatEndInterval,omitempty"`
}
