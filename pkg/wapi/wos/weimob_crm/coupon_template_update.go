package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponTemplateUpdate
// @id 2094
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2094?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新优惠券)
func (client *Client) CouponTemplateUpdate(request *CouponTemplateUpdateRequest) (response *CouponTemplateUpdateResponse, err error) {
	rpcResponse := CreateCouponTemplateUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponTemplateUpdateRequest struct {
	*api.BaseRequest

	BaseInfo     CouponTemplateUpdateRequestBaseInfo     `json:"baseInfo,omitempty"`
	Setting      CouponTemplateUpdateRequestSetting      `json:"setting,omitempty"`
	SendRule     CouponTemplateUpdateRequestSendRule     `json:"sendRule,omitempty"`
	UseRule      CouponTemplateUpdateRequestUseRule      `json:"useRule,omitempty"`
	OtherSetting CouponTemplateUpdateRequestOtherSetting `json:"otherSetting,omitempty"`
	CreateInfo   CouponTemplateUpdateRequestCreateInfo   `json:"createInfo,omitempty"`
	Vid          int64                                   `json:"vid,omitempty"`
	VidType      int64                                   `json:"vidType,omitempty"`
}

type CouponTemplateUpdateResponse struct {
}

func CreateCouponTemplateUpdateRequest() (request *CouponTemplateUpdateRequest) {
	request = &CouponTemplateUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/template/update", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponTemplateUpdateResponse() (response *api.BaseResponse[CouponTemplateUpdateResponse]) {
	response = api.CreateResponse[CouponTemplateUpdateResponse](&CouponTemplateUpdateResponse{})
	return
}

type CouponTemplateUpdateRequestBaseInfo struct {
	CouponTemplateId int64 `json:"couponTemplateId,omitempty"`
}

type CouponTemplateUpdateRequestSetting struct {
	Stock int64 `json:"stock,omitempty"`
}

type CouponTemplateUpdateRequestSendRule struct {
	Crowd                 CouponTemplateUpdateRequestCrowd `json:"crowd,omitempty"`
	HasUserTakeLimit      bool                             `json:"hasUserTakeLimit,omitempty"`
	SendTimeType          int64                            `json:"sendTimeType,omitempty"`
	SendStartDate         int64                            `json:"sendStartDate,omitempty"`
	SendEndDate           int64                            `json:"sendEndDate,omitempty"`
	CustomerDirectReceive bool                             `json:"customerDirectReceive,omitempty"`
	RecommendStartTime    int64                            `json:"recommendStartTime,omitempty"`
	RecommendEndTime      int64                            `json:"recommendEndTime,omitempty"`
	ActivityPublish       bool                             `json:"activityPublish,omitempty"`
	UserTakeLimit         int64                            `json:"userTakeLimit,omitempty"`
	EnterpriseAssistant   bool                             `json:"enterpriseAssistant,omitempty"`
	IsAcceptAllCrowd      bool                             `json:"isAcceptAllCrowd,omitempty"`
	MerchantPublish       bool                             `json:"merchantPublish,omitempty"`
	ShoppingPublish       bool                             `json:"shoppingPublish,omitempty"`
	CustomerListPublish   bool                             `json:"customerListPublish,omitempty"`
	ServicePublish        bool                             `json:"servicePublish,omitempty"`
	StoreLimitNum         int64                            `json:"storeLimitNum,omitempty"`
	CanShare              bool                             `json:"canShare,omitempty"`
	CanStoreLaunch        bool                             `json:"canStoreLaunch,omitempty"`
}

type CouponTemplateUpdateRequestCrowd struct {
	RuleContent  string `json:"ruleContent,omitempty"`
	SelectedData string `json:"selectedData,omitempty"`
}

type CouponTemplateUpdateRequestUseRule struct {
	UseValidDate    CouponTemplateUpdateRequestUseValidDate   `json:"useValidDate,omitempty"`
	CycleUseRule    []CouponTemplateUpdateRequestCycleUseRule `json:"cycleUseRule,omitempty"`
	OrderDeductList []int64                                   `json:"orderDeductList,omitempty"`
	IsAllTimeUse    bool                                      `json:"isAllTimeUse,omitempty"`
}

type CouponTemplateUpdateRequestUseValidDate struct {
	UseEndTime int64 `json:"useEndTime,omitempty"`
}

type CouponTemplateUpdateRequestCycleUseRule struct {
	TimeSegment []CouponTemplateUpdateRequestTimeSegment `json:"timeSegment,omitempty"`
	WeekDay     []int64                                  `json:"weekDay,omitempty"`
}

type CouponTemplateUpdateRequestTimeSegment struct {
	BeginHour   int64 `json:"beginHour,omitempty"`
	BeginMinute int64 `json:"beginMinute,omitempty"`
	EndHour     int64 `json:"endHour,omitempty"`
	EndMinute   int64 `json:"endMinute,omitempty"`
}

type CouponTemplateUpdateRequestOtherSetting struct {
	Explain         string  `json:"explain,omitempty"`
	MemberTagIds    []int64 `json:"memberTagIds,omitempty"`
	IsOpenMessage   bool    `json:"isOpenMessage,omitempty"`
	MessageStockNum int64   `json:"messageStockNum,omitempty"`
}

type CouponTemplateUpdateRequestCreateInfo struct {
	OperatorId int64  `json:"operatorId,omitempty"`
	Source     string `json:"source,omitempty"`
}
