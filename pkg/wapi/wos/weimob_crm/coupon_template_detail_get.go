package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponTemplateDetailGet
// @id 1608
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1608?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询优惠券详细信息)
func (client *Client) CouponTemplateDetailGet(request *CouponTemplateDetailGetRequest) (response *CouponTemplateDetailGetResponse, err error) {
	rpcResponse := CreateCouponTemplateDetailGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponTemplateDetailGetRequest struct {
	*api.BaseRequest

	CouponTemplateId int64 `json:"couponTemplateId,omitempty"`
	Status           int64 `json:"status,omitempty"`
}

type CouponTemplateDetailGetResponse struct {
	BaseInfo     CouponTemplateDetailGetResponseBaseInfo     `json:"baseInfo,omitempty"`
	Setting      CouponTemplateDetailGetResponseSetting      `json:"setting,omitempty"`
	ReduceRule   CouponTemplateDetailGetResponseReduceRule   `json:"reduceRule,omitempty"`
	SendRule     CouponTemplateDetailGetResponseSendRule     `json:"sendRule,omitempty"`
	UseRule      CouponTemplateDetailGetResponseUseRule      `json:"useRule,omitempty"`
	OtherSetting CouponTemplateDetailGetResponseOtherSetting `json:"otherSetting,omitempty"`
	CreateInfo   CouponTemplateDetailGetResponseCreateInfo   `json:"createInfo,omitempty"`
	Status       int64                                       `json:"status,omitempty"`
}

func CreateCouponTemplateDetailGetRequest() (request *CouponTemplateDetailGetRequest) {
	request = &CouponTemplateDetailGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/template/detail/get", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponTemplateDetailGetResponse() (response *api.BaseResponse[CouponTemplateDetailGetResponse]) {
	response = api.CreateResponse[CouponTemplateDetailGetResponse](&CouponTemplateDetailGetResponse{})
	return
}

type CouponTemplateDetailGetResponseBaseInfo struct {
	CouponTemplateId int64   `json:"couponTemplateId,omitempty"`
	CouponType       int64   `json:"couponType,omitempty"`
	PlatformType     int64   `json:"platformType,omitempty"`
	Name             string  `json:"name,omitempty"`
	SubName          string  `json:"subName,omitempty"`
	CouponDesc       string  `json:"couponDesc,omitempty"`
	CouponSubType    int64   `json:"couponSubType,omitempty"`
	BelongVid        int64   `json:"belongVid,omitempty"`
	BelongVidName    string  `json:"belongVidName,omitempty"`
	BasicTemplateIds []int64 `json:"basicTemplateIds,omitempty"`
}

type CouponTemplateDetailGetResponseSetting struct {
	CustomCodeSetting CouponTemplateDetailGetResponseCustomCodeSetting `json:"customCodeSetting,omitempty"`
	Stock             int64                                            `json:"stock,omitempty"`
	CodeGenerateType  int64                                            `json:"codeGenerateType,omitempty"`
	StockUnlimited    bool                                             `json:"stockUnlimited,omitempty"`
}

type CouponTemplateDetailGetResponseCustomCodeSetting struct {
	CustomEndNum   int64  `json:"customEndNum,omitempty"`
	CustomStartNum int64  `json:"customStartNum,omitempty"`
	FixLeft        string `json:"fixLeft,omitempty"`
	FixRight       string `json:"fixRight,omitempty"`
}

type CouponTemplateDetailGetResponseReduceRule struct {
	ReducePriceType    int64  `json:"reducePriceType,omitempty"`
	MaxReducePrice     string `json:"maxReducePrice,omitempty"`
	CashTicketAmt      int64  `json:"cashTicketAmt,omitempty"`
	Discount           string `json:"discount,omitempty"`
	DiscountLimitType  int64  `json:"discountLimitType,omitempty"`
	DiscountLimitValue string `json:"discountLimitValue,omitempty"`
	MinCashTicketAmt   string `json:"minCashTicketAmt,omitempty"`
	MaxCashTicketAmt   string `json:"maxCashTicketAmt,omitempty"`
	CouponCostType     int64  `json:"couponCostType,omitempty"`
	CouponCost         string `json:"couponCost,omitempty"`
	PricingType        int64  `json:"pricingType,omitempty"`
	CostOfOrganization int64  `json:"costOfOrganization,omitempty"`
	CostOfStore        int64  `json:"costOfStore,omitempty"`
	FreightReduceType  int64  `json:"freightReduceType,omitempty"`
}

type CouponTemplateDetailGetResponseSendRule struct {
	Crowd                 CouponTemplateDetailGetResponseCrowd `json:"crowd,omitempty"`
	SendTimeType          int64                                `json:"sendTimeType,omitempty"`
	SendStartDate         string                               `json:"sendStartDate,omitempty"`
	SendEndDate           string                               `json:"sendEndDate,omitempty"`
	IsPayment             bool                                 `json:"isPayment,omitempty"`
	SendChannels          []int64                              `json:"sendChannels,omitempty"`
	RecommendStartTime    string                               `json:"recommendStartTime,omitempty"`
	RecommendEndTime      string                               `json:"recommendEndTime,omitempty"`
	IsAcceptAllCrowd      bool                                 `json:"isAcceptAllCrowd,omitempty"`
	UserTakeLimit         int64                                `json:"userTakeLimit,omitempty"`
	CanStoreLaunch        bool                                 `json:"canStoreLaunch,omitempty"`
	CanGift               bool                                 `json:"canGift,omitempty"`
	CanShare              bool                                 `json:"canShare,omitempty"`
	StoreLimitNum         int64                                `json:"storeLimitNum,omitempty"`
	MerchantIssueChannels []int64                              `json:"merchantIssueChannels,omitempty"`
	MerchantPublish       bool                                 `json:"merchantPublish,omitempty"`
	ShoppingPublish       bool                                 `json:"shoppingPublish,omitempty"`
	CustomerListPublish   bool                                 `json:"customerListPublish,omitempty"`
	ServicePublish        bool                                 `json:"servicePublish,omitempty"`
	ActivityPublish       bool                                 `json:"activityPublish,omitempty"`
	CustomerDirectReceive bool                                 `json:"customerDirectReceive,omitempty"`
	ScanByIssue           bool                                 `json:"scanByIssue,omitempty"`
	EnterpriseAssistant   bool                                 `json:"enterpriseAssistant,omitempty"`
	MerchantDelivery      bool                                 `json:"merchantDelivery,omitempty"`
	LimitedTimeDelivery   bool                                 `json:"limitedTimeDelivery,omitempty"`
	IsOpenRecommend       bool                                 `json:"isOpenRecommend,omitempty"`
	HasUserTakeLimit      bool                                 `json:"hasUserTakeLimit,omitempty"`
	HasUserTakeDayLimit   bool                                 `json:"hasUserTakeDayLimit,omitempty"`
	UserTakeDayLimit      int64                                `json:"userTakeDayLimit,omitempty"`
}

type CouponTemplateDetailGetResponseCrowd struct {
	RuleContent  string `json:"ruleContent,omitempty"`
	SelectedData string `json:"selectedData,omitempty"`
	RuleId       int64  `json:"ruleId,omitempty"`
}

type CouponTemplateDetailGetResponseUseRule struct {
	UseValidDate        CouponTemplateDetailGetResponseUseValidDate      `json:"useValidDate,omitempty"`
	AcceptTypeDTO       CouponTemplateDetailGetResponseAcceptTypeDTO     `json:"acceptTypeDTO,omitempty"`
	AllSceneDTO         CouponTemplateDetailGetResponseAllSceneDTO       `json:"allSceneDTO,omitempty"`
	CycleUseRule        []CouponTemplateDetailGetResponseCycleUseRule    `json:"cycleUseRule,omitempty"`
	SuperpositionRule   CouponTemplateDetailGetResponseSuperpositionRule `json:"superpositionRule,omitempty"`
	CanUseDiscount      CouponTemplateDetailGetResponseCanUseDiscount    `json:"canUseDiscount,omitempty"`
	HasUseThreshold     bool                                             `json:"hasUseThreshold,omitempty"`
	CanCashTicket       bool                                             `json:"canCashTicket,omitempty"`
	GoodsCashTicket     string                                           `json:"goodsCashTicket,omitempty"`
	CanGoodsNumber      bool                                             `json:"canGoodsNumber,omitempty"`
	MinGoodsNumber      int64                                            `json:"minGoodsNumber,omitempty"`
	MaxGoodsNumber      int64                                            `json:"maxGoodsNumber,omitempty"`
	IsAllUseScene       bool                                             `json:"isAllUseScene,omitempty"`
	IsAllTimeUse        bool                                             `json:"isAllTimeUse,omitempty"`
	OrderDeductList     []int64                                          `json:"orderDeductList,omitempty"`
	ImportGoodsFlag     bool                                             `json:"importGoodsFlag,omitempty"`
	ImportGoodsFileName string                                           `json:"importGoodsFileName,omitempty"`
}

type CouponTemplateDetailGetResponseUseValidDate struct {
	UseTimeType    int64  `json:"useTimeType,omitempty"`
	UseStartTime   string `json:"useStartTime,omitempty"`
	UseEndTime     string `json:"useEndTime,omitempty"`
	UseDelayDays   int64  `json:"useDelayDays,omitempty"`
	ValidDays      int64  `json:"validDays,omitempty"`
	ValidDateDesc  string `json:"validDateDesc,omitempty"`
	SubUseTimeType int64  `json:"subUseTimeType,omitempty"`
}

type CouponTemplateDetailGetResponseAcceptTypeDTO struct {
	AcceptStoreType   int64   `json:"acceptStoreType,omitempty"`
	AcceptStoreIds    []int64 `json:"acceptStoreIds,omitempty"`
	AcceptGoodsType   int64   `json:"acceptGoodsType,omitempty"`
	AcceptGoodsIds    []int64 `json:"acceptGoodsIds,omitempty"`
	ExistExcludeGoods bool    `json:"existExcludeGoods,omitempty"`
	ExcludeGoodsIds   []int64 `json:"excludeGoodsIds,omitempty"`
	ExcludeStoreIds   []int64 `json:"excludeStoreIds,omitempty"`
	ExistExcludeStore bool    `json:"existExcludeStore,omitempty"`
}

type CouponTemplateDetailGetResponseAllSceneDTO struct {
	ShoppingMallSceneList []int64 `json:"shoppingMallSceneList,omitempty"`
}

type CouponTemplateDetailGetResponseCycleUseRule struct {
	TimeSegment []CouponTemplateDetailGetResponseTimeSegment `json:"timeSegment,omitempty"`
	WeekDay     []int64                                      `json:"weekDay,omitempty"`
}

type CouponTemplateDetailGetResponseTimeSegment struct {
	BeginHour   int64 `json:"beginHour,omitempty"`
	EndHour     int64 `json:"endHour,omitempty"`
	BeginMinute int64 `json:"beginMinute,omitempty"`
	EndMinute   int64 `json:"endMinute,omitempty"`
}

type CouponTemplateDetailGetResponseSuperpositionRule struct {
	IsSuperposition        bool  `json:"isSuperposition,omitempty"`
	SameSuperposition      bool  `json:"sameSuperposition,omitempty"`
	DifferentSuperposition bool  `json:"differentSuperposition,omitempty"`
	Limit                  int64 `json:"limit,omitempty"`
}

type CouponTemplateDetailGetResponseCanUseDiscount struct {
	CanUseWithOtherDiscount bool    `json:"canUseWithOtherDiscount,omitempty"`
	ShoppingMallDiscount    []int64 `json:"shoppingMallDiscount,omitempty"`
	ForHereDiscount         []int64 `json:"forHereDiscount,omitempty"`
	TakeOutDiscount         []int64 `json:"takeOutDiscount,omitempty"`
	WineTourDiscount        []int64 `json:"wineTourDiscount,omitempty"`
}

type CouponTemplateDetailGetResponseOtherSetting struct {
	MemberTagIds    []CouponTemplateDetailGetResponseMemberTagIds `json:"memberTagIds,omitempty"`
	HasMemberTag    bool                                          `json:"hasMemberTag,omitempty"`
	Explain         string                                        `json:"explain,omitempty"`
	IsOpenMessage   bool                                          `json:"isOpenMessage,omitempty"`
	MessageStockNum int64                                         `json:"messageStockNum,omitempty"`
}

type CouponTemplateDetailGetResponseMemberTagIds struct {
	List []CouponTemplateDetailGetResponselist `json:"list,omitempty"`
	Id   int64                                 `json:"id,omitempty"`
	Name string                                `json:"name,omitempty"`
	Type int64                                 `json:"type,omitempty"`
}

type CouponTemplateDetailGetResponselist struct {
	Id         int64  `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	StartValue string `json:"startValue,omitempty"`
	EndValue   string `json:"endValue,omitempty"`
}

type CouponTemplateDetailGetResponseCreateInfo struct {
	CreateTime string `json:"createTime,omitempty"`
	OperatorId int64  `json:"operatorId,omitempty"`
	Source     string `json:"source,omitempty"`
}
