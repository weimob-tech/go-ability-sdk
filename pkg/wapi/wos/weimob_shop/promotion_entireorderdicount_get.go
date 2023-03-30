package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionEntireorderdicountGet
// @id 1447
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1447?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=B端获取买单优惠活动详情)
func (client *Client) PromotionEntireorderdicountGet(request *PromotionEntireorderdicountGetRequest) (response *PromotionEntireorderdicountGetResponse, err error) {
	rpcResponse := CreatePromotionEntireorderdicountGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionEntireorderdicountGetRequest struct {
	*api.BaseRequest

	BasicInfo    PromotionEntireorderdicountGetRequestBasicInfo `json:"basicInfo,omitempty"`
	ActivityId   int64                                          `json:"activityId,omitempty"`
	ActivityType int64                                          `json:"activityType,omitempty"`
}

type PromotionEntireorderdicountGetResponse struct {
	ActivityId         int64  `json:"activityId,omitempty"`
	ActivityType       int64  `json:"activityType,omitempty"`
	Discount           int64  `json:"discount,omitempty"`
	DiscountType       int64  `json:"discountType,omitempty"`
	EndTime            string `json:"endTime,omitempty"`
	FitGoods           int64  `json:"fitGoods,omitempty"`
	FitPeople          int64  `json:"fitPeople,omitempty"`
	FitScene           string `json:"fitScene,omitempty"`
	FitVid             int64  `json:"fitVid,omitempty"`
	HaveManagement     bool   `json:"haveManagement,omitempty"`
	OriginatorType     int64  `json:"originatorType,omitempty"`
	SelectPeopleRuleId int64  `json:"selectPeopleRuleId,omitempty"`
	StartTime          string `json:"startTime,omitempty"`
	TagFlag            int64  `json:"tagFlag,omitempty"`
	TagRuleId          int64  `json:"tagRuleId,omitempty"`
	TimeType           int64  `json:"timeType,omitempty"`
	Title              string `json:"title,omitempty"`
	WipeZeroType       int64  `json:"wipeZeroType,omitempty"`
}

func CreatePromotionEntireorderdicountGetRequest() (request *PromotionEntireorderdicountGetRequest) {
	request = &PromotionEntireorderdicountGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "promotion/entireorderdicount/get", "apigw")
	request.Method = api.POST
	return
}

func CreatePromotionEntireorderdicountGetResponse() (response *api.BaseResponse[PromotionEntireorderdicountGetResponse]) {
	response = api.CreateResponse[PromotionEntireorderdicountGetResponse](&PromotionEntireorderdicountGetResponse{})
	return
}

type PromotionEntireorderdicountGetRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}
