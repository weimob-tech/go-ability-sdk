package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionEntireorderdicountUpdate
// @id 1446
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1446?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=买单优惠保存或更新)
func (client *Client) PromotionEntireorderdicountUpdate(request *PromotionEntireorderdicountUpdateRequest) (response *PromotionEntireorderdicountUpdateResponse, err error) {
	rpcResponse := CreatePromotionEntireorderdicountUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionEntireorderdicountUpdateRequest struct {
	*api.BaseRequest

	BasicInfo       PromotionEntireorderdicountUpdateRequestBasicInfo `json:"basicInfo,omitempty"`
	TagFlag         int64                                             `json:"tagFlag,omitempty"`
	WipeZeroType    int64                                             `json:"wipeZeroType,omitempty"`
	TimeType        int64                                             `json:"timeType,omitempty"`
	Discount        int64                                             `json:"discount,omitempty"`
	Title           string                                            `json:"title,omitempty"`
	Uuid            int64                                             `json:"uuid,omitempty"`
	ActivityId      int64                                             `json:"activityId,omitempty"`
	TagRuleId       int64                                             `json:"tagRuleId,omitempty"`
	FitPeople       int64                                             `json:"fitPeople,omitempty"`
	FitPeopleRuleId int64                                             `json:"fitPeopleRuleId,omitempty"`
	StartTime       string                                            `json:"startTime,omitempty"`
	DiscountType    int64                                             `json:"discountType,omitempty"`
	EndTime         string                                            `json:"endTime,omitempty"`
}

type PromotionEntireorderdicountUpdateResponse struct {
	ActivityId int64 `json:"activityId,omitempty"`
	Success    bool  `json:"success,omitempty"`
}

func CreatePromotionEntireorderdicountUpdateRequest() (request *PromotionEntireorderdicountUpdateRequest) {
	request = &PromotionEntireorderdicountUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "promotion/entireorderdicount/update", "apigw")
	request.Method = api.POST
	return
}

func CreatePromotionEntireorderdicountUpdateResponse() (response *api.BaseResponse[PromotionEntireorderdicountUpdateResponse]) {
	response = api.CreateResponse[PromotionEntireorderdicountUpdateResponse](&PromotionEntireorderdicountUpdateResponse{})
	return
}

type PromotionEntireorderdicountUpdateRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}
