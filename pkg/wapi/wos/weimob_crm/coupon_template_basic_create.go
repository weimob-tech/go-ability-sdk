package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponTemplateBasicCreate
// @id 2090
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2090?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=新建优惠券基础模板)
func (client *Client) CouponTemplateBasicCreate(request *CouponTemplateBasicCreateRequest) (response *CouponTemplateBasicCreateResponse, err error) {
	rpcResponse := CreateCouponTemplateBasicCreateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponTemplateBasicCreateRequest struct {
	*api.BaseRequest

	BaseInfo   CouponTemplateBasicCreateRequestBaseInfo   `json:"baseInfo,omitempty"`
	ReduceRule CouponTemplateBasicCreateRequestReduceRule `json:"reduceRule,omitempty"`
	SendRule   CouponTemplateBasicCreateRequestSendRule   `json:"sendRule,omitempty"`
	UseRule    CouponTemplateBasicCreateRequestUseRule    `json:"useRule,omitempty"`
	CreateInfo CouponTemplateBasicCreateRequestCreateInfo `json:"createInfo,omitempty"`
	Vid        int64                                      `json:"vid,omitempty"`
	VidType    int64                                      `json:"vidType,omitempty"`
}

type CouponTemplateBasicCreateResponse struct {
	BasicCouponTemplateId int64 `json:"basicCouponTemplateId,omitempty"`
}

func CreateCouponTemplateBasicCreateRequest() (request *CouponTemplateBasicCreateRequest) {
	request = &CouponTemplateBasicCreateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/template/basic/create", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponTemplateBasicCreateResponse() (response *api.BaseResponse[CouponTemplateBasicCreateResponse]) {
	response = api.CreateResponse[CouponTemplateBasicCreateResponse](&CouponTemplateBasicCreateResponse{})
	return
}

type CouponTemplateBasicCreateRequestBaseInfo struct {
	CouponType    int64  `json:"couponType,omitempty"`
	PlatformType  int64  `json:"platformType,omitempty"`
	CouponSubType int64  `json:"couponSubType,omitempty"`
	Name          string `json:"name,omitempty"`
	BelongVid     int64  `json:"belongVid,omitempty"`
	MerchantName  string `json:"merchantName,omitempty"`
}

type CouponTemplateBasicCreateRequestReduceRule struct {
	ReducePriceType    int64  `json:"reducePriceType,omitempty"`
	CashTicketAmt      string `json:"cashTicketAmt,omitempty"`
	PricingType        int64  `json:"pricingType,omitempty"`
	CouponCostType     int64  `json:"couponCostType,omitempty"`
	CouponCost         string `json:"couponCost,omitempty"`
	CostOfOrganization string `json:"costOfOrganization,omitempty"`
	CostOfStore        int64  `json:"costOfStore,omitempty"`
	Discount           string `json:"discount,omitempty"`
	DiscountLimitType  int64  `json:"discountLimitType,omitempty"`
	DiscountLimitValue string `json:"discountLimitValue,omitempty"`
	MaxReducePrice     string `json:"maxReducePrice,omitempty"`
}

type CouponTemplateBasicCreateRequestSendRule struct {
	SendStartDate         CouponTemplateBasicCreateRequestSendStartDate `json:"sendStartDate,omitempty"`
	SendEndDate           CouponTemplateBasicCreateRequestSendEndDate   `json:"sendEndDate,omitempty"`
	SendTimeType          int64                                         `json:"sendTimeType,omitempty"`
	IsPayment             bool                                          `json:"isPayment,omitempty"`
	CustomerDirectReceive bool                                          `json:"customerDirectReceive,omitempty"`
	MerchantPublish       bool                                          `json:"merchantPublish,omitempty"`
	ActivityPublish       bool                                          `json:"activityPublish,omitempty"`
	EnterpriseAssistant   bool                                          `json:"enterpriseAssistant,omitempty"`
}

type CouponTemplateBasicCreateRequestSendStartDate struct {
}

type CouponTemplateBasicCreateRequestSendEndDate struct {
}

type CouponTemplateBasicCreateRequestUseRule struct {
	SuperpositionRule CouponTemplateBasicCreateRequestSuperpositionRule `json:"superpositionRule,omitempty"`
	CanUseDiscount    CouponTemplateBasicCreateRequestCanUseDiscount    `json:"canUseDiscount,omitempty"`
	AllSceneDTO       CouponTemplateBasicCreateRequestAllSceneDTO       `json:"allSceneDTO,omitempty"`
	AcceptTypeDTO     CouponTemplateBasicCreateRequestAcceptTypeDTO     `json:"acceptTypeDTO,omitempty"`
	HasUseThreshold   bool                                              `json:"hasUseThreshold,omitempty"`
	CanCashTicket     bool                                              `json:"canCashTicket,omitempty"`
	GoodsCashTicket   string                                            `json:"goodsCashTicket,omitempty"`
	OrderDeductList   []int64                                           `json:"orderDeductList,omitempty"`
	IsAllUseScene     bool                                              `json:"isAllUseScene,omitempty"`
	CanGoodsNumber    bool                                              `json:"canGoodsNumber,omitempty"`
	MinGoodsNumber    int64                                             `json:"minGoodsNumber,omitempty"`
	MaxGoodsNumber    int64                                             `json:"maxGoodsNumber,omitempty"`
}

type CouponTemplateBasicCreateRequestSuperpositionRule struct {
	IsSuperposition        bool  `json:"isSuperposition,omitempty"`
	SameSuperposition      bool  `json:"sameSuperposition,omitempty"`
	DifferentSuperposition bool  `json:"differentSuperposition,omitempty"`
	Limit                  int64 `json:"limit,omitempty"`
}

type CouponTemplateBasicCreateRequestCanUseDiscount struct {
	CanUseWithOtherDiscount bool    `json:"canUseWithOtherDiscount,omitempty"`
	ShoppingMallDiscount    []int64 `json:"shoppingMallDiscount,omitempty"`
}

type CouponTemplateBasicCreateRequestAllSceneDTO struct {
	ShoppingMallSceneList []int64 `json:"shoppingMallSceneList,omitempty"`
}

type CouponTemplateBasicCreateRequestAcceptTypeDTO struct {
	ExcludeGoodsIds   CouponTemplateBasicCreateRequestExcludeGoodsIds `json:"excludeGoodsIds,omitempty"`
	AcceptStoreType   int64                                           `json:"acceptStoreType,omitempty"`
	AcceptStoreIds    []int64                                         `json:"acceptStoreIds,omitempty"`
	AcceptGoodsType   int64                                           `json:"acceptGoodsType,omitempty"`
	AcceptGoodsIds    []int64                                         `json:"acceptGoodsIds,omitempty"`
	ExistExcludeGoods bool                                            `json:"existExcludeGoods,omitempty"`
	IncludeStoreGoods bool                                            `json:"includeStoreGoods,omitempty"`
	ExistExcludeStore bool                                            `json:"existExcludeStore,omitempty"`
	ExcludeStoreIds   []int64                                         `json:"excludeStoreIds,omitempty"`
}

type CouponTemplateBasicCreateRequestExcludeGoodsIds struct {
}

type CouponTemplateBasicCreateRequestCreateInfo struct {
	OperatorId int64  `json:"operatorId,omitempty"`
	Source     string `json:"source,omitempty"`
}
