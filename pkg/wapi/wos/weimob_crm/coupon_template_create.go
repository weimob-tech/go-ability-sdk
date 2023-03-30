package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponTemplateCreate
// @id 2092
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2092?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=新建优惠券)
func (client *Client) CouponTemplateCreate(request *CouponTemplateCreateRequest) (response *CouponTemplateCreateResponse, err error) {
	rpcResponse := CreateCouponTemplateCreateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponTemplateCreateRequest struct {
	*api.BaseRequest

	BaseInfo     CouponTemplateCreateRequestBaseInfo     `json:"baseInfo,omitempty"`
	Setting      CouponTemplateCreateRequestSetting      `json:"setting,omitempty"`
	SendRule     CouponTemplateCreateRequestSendRule     `json:"sendRule,omitempty"`
	UseRule      CouponTemplateCreateRequestUseRule      `json:"useRule,omitempty"`
	OtherSetting CouponTemplateCreateRequestOtherSetting `json:"otherSetting,omitempty"`
	CreateInfo   CouponTemplateCreateRequestCreateInfo   `json:"createInfo,omitempty"`
	ReduceRule   CouponTemplateCreateRequestReduceRule   `json:"reduceRule,omitempty"`
	Vid          int64                                   `json:"vid,omitempty"`
	VidType      int64                                   `json:"vidType,omitempty"`
}

type CouponTemplateCreateResponse struct {
	CouponTemplateIds []int64 `json:"couponTemplateIds,omitempty"`
}

func CreateCouponTemplateCreateRequest() (request *CouponTemplateCreateRequest) {
	request = &CouponTemplateCreateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/template/create", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponTemplateCreateResponse() (response *api.BaseResponse[CouponTemplateCreateResponse]) {
	response = api.CreateResponse[CouponTemplateCreateResponse](&CouponTemplateCreateResponse{})
	return
}

type CouponTemplateCreateRequestBaseInfo struct {
	Name                      string  `json:"name,omitempty"`
	MerchantName              string  `json:"merchantName,omitempty"`
	BasicCouponTemplateIdList []int64 `json:"basicCouponTemplateIdList,omitempty"`
}

type CouponTemplateCreateRequestSetting struct {
	Stock          int64 `json:"stock,omitempty"`
	StockUnlimited bool  `json:"stockUnlimited,omitempty"`
}

type CouponTemplateCreateRequestSendRule struct {
	Crowd                 CouponTemplateCreateRequestCrowd `json:"crowd,omitempty"`
	HasUserTakeLimit      bool                             `json:"hasUserTakeLimit,omitempty"`
	SendTimeType          int64                            `json:"sendTimeType,omitempty"`
	SendStartDate         int64                            `json:"sendStartDate,omitempty"`
	SendEndDate           int64                            `json:"sendEndDate,omitempty"`
	IsPayment             bool                             `json:"isPayment,omitempty"`
	CustomerDirectReceive bool                             `json:"customerDirectReceive,omitempty"`
	RecommendStartTime    int64                            `json:"recommendStartTime,omitempty"`
	RecommendEndTime      int64                            `json:"recommendEndTime,omitempty"`
	ActivityPublish       bool                             `json:"activityPublish,omitempty"`
	UserTakeLimit         int64                            `json:"userTakeLimit,omitempty"`
	CanGift               bool                             `json:"canGift,omitempty"`
	EnterpriseAssistant   bool                             `json:"enterpriseAssistant,omitempty"`
	IsAcceptAllCrowd      bool                             `json:"isAcceptAllCrowd,omitempty"`
	HasGiftLimit          bool                             `json:"hasGiftLimit,omitempty"`
	GiftCountLimit        int64                            `json:"giftCountLimit,omitempty"`
	MerchantPublish       bool                             `json:"merchantPublish,omitempty"`
	ShoppingPublish       bool                             `json:"shoppingPublish,omitempty"`
	CustomerListPublish   bool                             `json:"customerListPublish,omitempty"`
	ServicePublish        bool                             `json:"servicePublish,omitempty"`
	StoreLimitNum         int64                            `json:"storeLimitNum,omitempty"`
	CanShare              bool                             `json:"canShare,omitempty"`
	CanStoreLaunch        bool                             `json:"canStoreLaunch,omitempty"`
}

type CouponTemplateCreateRequestCrowd struct {
	RuleContent  string `json:"ruleContent,omitempty"`
	SelectedData string `json:"selectedData,omitempty"`
}

type CouponTemplateCreateRequestUseRule struct {
	UseValidDate      CouponTemplateCreateRequestUseValidDate      `json:"useValidDate,omitempty"`
	CycleUseRule      []CouponTemplateCreateRequestCycleUseRule    `json:"cycleUseRule,omitempty"`
	AcceptTypeDTO     CouponTemplateCreateRequestAcceptTypeDTO     `json:"acceptTypeDTO,omitempty"`
	AllSceneDTO       CouponTemplateCreateRequestAllSceneDTO       `json:"allSceneDTO,omitempty"`
	SuperpositionRule CouponTemplateCreateRequestSuperpositionRule `json:"superpositionRule,omitempty"`
	CanUseDiscount    CouponTemplateCreateRequestCanUseDiscount    `json:"canUseDiscount,omitempty"`
	OrderDeductList   []int64                                      `json:"orderDeductList,omitempty"`
	IsAllTimeUse      bool                                         `json:"isAllTimeUse,omitempty"`
	CanGoodsNumber    bool                                         `json:"canGoodsNumber,omitempty"`
	MinGoodsNumber    int64                                        `json:"minGoodsNumber,omitempty"`
	MaxGoodsNumber    int64                                        `json:"maxGoodsNumber,omitempty"`
	IsAllUseScene     bool                                         `json:"isAllUseScene,omitempty"`
}

type CouponTemplateCreateRequestUseValidDate struct {
	UseTimeType    int64 `json:"useTimeType,omitempty"`
	UseStartTime   int64 `json:"useStartTime,omitempty"`
	UseEndTime     int64 `json:"useEndTime,omitempty"`
	ValidDays      int64 `json:"validDays,omitempty"`
	UseDelayDays   int64 `json:"useDelayDays,omitempty"`
	SubUseTimeType int64 `json:"subUseTimeType,omitempty"`
}

type CouponTemplateCreateRequestCycleUseRule struct {
	TimeSegment []CouponTemplateCreateRequestTimeSegment `json:"timeSegment,omitempty"`
	WeekDay     []int64                                  `json:"weekDay,omitempty"`
}

type CouponTemplateCreateRequestTimeSegment struct {
	BeginHour   int64 `json:"beginHour,omitempty"`
	BeginMinute int64 `json:"beginMinute,omitempty"`
	EndHour     int64 `json:"endHour,omitempty"`
	EndMinute   int64 `json:"endMinute,omitempty"`
}

type CouponTemplateCreateRequestAcceptTypeDTO struct {
	AcceptStoreType   int64   `json:"acceptStoreType,omitempty"`
	AcceptStoreIds    []int64 `json:"acceptStoreIds,omitempty"`
	AcceptGoodsType   int64   `json:"acceptGoodsType,omitempty"`
	AcceptGoodsIds    []int64 `json:"acceptGoodsIds,omitempty"`
	ExistExcludeGoods bool    `json:"existExcludeGoods,omitempty"`
	ExcludeGoodsIds   []int64 `json:"excludeGoodsIds,omitempty"`
	ExistExcludeStore bool    `json:"existExcludeStore,omitempty"`
	ExcludeStoreIds   []int64 `json:"excludeStoreIds,omitempty"`
	IncludeStoreGoods bool    `json:"includeStoreGoods,omitempty"`
}

type CouponTemplateCreateRequestAllSceneDTO struct {
	ShoppingMallSceneList []int64 `json:"shoppingMallSceneList,omitempty"`
}

type CouponTemplateCreateRequestSuperpositionRule struct {
	IsSuperposition        bool  `json:"isSuperposition,omitempty"`
	SameSuperposition      bool  `json:"sameSuperposition,omitempty"`
	DifferentSuperposition bool  `json:"differentSuperposition,omitempty"`
	Limit                  int64 `json:"limit,omitempty"`
}

type CouponTemplateCreateRequestCanUseDiscount struct {
	CanUseWithOtherDiscount bool    `json:"canUseWithOtherDiscount,omitempty"`
	ShoppingMallDiscount    []int64 `json:"shoppingMallDiscount,omitempty"`
}

type CouponTemplateCreateRequestOtherSetting struct {
	MemberTagIds    []CouponTemplateCreateRequestMemberTagIds `json:"memberTagIds,omitempty"`
	Explain         string                                    `json:"explain,omitempty"`
	IsOpenMessage   bool                                      `json:"isOpenMessage,omitempty"`
	MessageStockNum int64                                     `json:"messageStockNum,omitempty"`
	HasMemberTag    bool                                      `json:"hasMemberTag,omitempty"`
}

type CouponTemplateCreateRequestMemberTagIds struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CouponTemplateCreateRequestCreateInfo struct {
	OperatorId int64  `json:"operatorId,omitempty"`
	Source     string `json:"source,omitempty"`
}

type CouponTemplateCreateRequestReduceRule struct {
	PricingType int64 `json:"pricingType,omitempty"`
}
