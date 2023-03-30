package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponTemplateBasicUpdate
// @id 2091
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2091?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新优惠券基础模板)
func (client *Client) CouponTemplateBasicUpdate(request *CouponTemplateBasicUpdateRequest) (response *CouponTemplateBasicUpdateResponse, err error) {
	rpcResponse := CreateCouponTemplateBasicUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponTemplateBasicUpdateRequest struct {
	*api.BaseRequest

	BaseInfo   CouponTemplateBasicUpdateRequestBaseInfo   `json:"baseInfo,omitempty"`
	ReduceRule CouponTemplateBasicUpdateRequestReduceRule `json:"reduceRule,omitempty"`
	SendRule   CouponTemplateBasicUpdateRequestSendRule   `json:"sendRule,omitempty"`
	UseRule    CouponTemplateBasicUpdateRequestUseRule    `json:"useRule,omitempty"`
	CreateInfo CouponTemplateBasicUpdateRequestCreateInfo `json:"createInfo,omitempty"`
	Vid        int64                                      `json:"vid,omitempty"`
	VidType    int64                                      `json:"vidType,omitempty"`
}

type CouponTemplateBasicUpdateResponse struct {
}

func CreateCouponTemplateBasicUpdateRequest() (request *CouponTemplateBasicUpdateRequest) {
	request = &CouponTemplateBasicUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/template/basic/update", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponTemplateBasicUpdateResponse() (response *api.BaseResponse[CouponTemplateBasicUpdateResponse]) {
	response = api.CreateResponse[CouponTemplateBasicUpdateResponse](&CouponTemplateBasicUpdateResponse{})
	return
}

type CouponTemplateBasicUpdateRequestBaseInfo struct {
	CouponSubType         CouponTemplateBasicUpdateRequestCouponSubType `json:"couponSubType,omitempty"`
	CouponType            int64                                         `json:"couponType,omitempty"`
	PlatformType          int64                                         `json:"platformType,omitempty"`
	Name                  string                                        `json:"name,omitempty"`
	BelongVid             int64                                         `json:"belongVid,omitempty"`
	BasicCouponTemplateId int64                                         `json:"basicCouponTemplateId,omitempty"`
}

type CouponTemplateBasicUpdateRequestCouponSubType struct {
}

type CouponTemplateBasicUpdateRequestReduceRule struct {
	ReducePriceType    int64  `json:"reducePriceType,omitempty"`
	CashTicketAmt      int64  `json:"cashTicketAmt,omitempty"`
	PricingType        int64  `json:"pricingType,omitempty"`
	CouponCostType     int64  `json:"couponCostType,omitempty"`
	CouponCost         int64  `json:"couponCost,omitempty"`
	CostOfOrganization string `json:"costOfOrganization,omitempty"`
	CostOfStore        int64  `json:"costOfStore,omitempty"`
	Discount           string `json:"discount,omitempty"`
	DiscountLimitType  int64  `json:"discountLimitType,omitempty"`
	DiscountLimitValue string `json:"discountLimitValue,omitempty"`
}

type CouponTemplateBasicUpdateRequestSendRule struct {
	SendStartDate         CouponTemplateBasicUpdateRequestSendStartDate `json:"sendStartDate,omitempty"`
	SendEndDate           CouponTemplateBasicUpdateRequestSendEndDate   `json:"sendEndDate,omitempty"`
	SendTimeType          int64                                         `json:"sendTimeType,omitempty"`
	IsPayment             bool                                          `json:"isPayment,omitempty"`
	CustomerDirectReceive bool                                          `json:"customerDirectReceive,omitempty"`
	MerchantPublish       bool                                          `json:"merchantPublish,omitempty"`
	ActivityPublish       bool                                          `json:"activityPublish,omitempty"`
	EnterpriseAssistant   bool                                          `json:"enterpriseAssistant,omitempty"`
}

type CouponTemplateBasicUpdateRequestSendStartDate struct {
}

type CouponTemplateBasicUpdateRequestSendEndDate struct {
}

type CouponTemplateBasicUpdateRequestUseRule struct {
	SuperpositionRule CouponTemplateBasicUpdateRequestSuperpositionRule `json:"superpositionRule,omitempty"`
	CanUseDiscount    CouponTemplateBasicUpdateRequestCanUseDiscount    `json:"canUseDiscount,omitempty"`
	AllSceneDTO       CouponTemplateBasicUpdateRequestAllSceneDTO       `json:"allSceneDTO,omitempty"`
	AcceptTypeDTO     CouponTemplateBasicUpdateRequestAcceptTypeDTO     `json:"acceptTypeDTO,omitempty"`
	HasUseThreshold   bool                                              `json:"hasUseThreshold,omitempty"`
	CanCashTicket     bool                                              `json:"canCashTicket,omitempty"`
	GoodsCashTicket   string                                            `json:"goodsCashTicket,omitempty"`
	OrderDeductList   []int64                                           `json:"orderDeductList,omitempty"`
	IsAllUseScene     bool                                              `json:"isAllUseScene,omitempty"`
}

type CouponTemplateBasicUpdateRequestSuperpositionRule struct {
	IsSuperposition        bool  `json:"isSuperposition,omitempty"`
	SameSuperposition      bool  `json:"sameSuperposition,omitempty"`
	DifferentSuperposition bool  `json:"differentSuperposition,omitempty"`
	Limit                  int64 `json:"limit,omitempty"`
}

type CouponTemplateBasicUpdateRequestCanUseDiscount struct {
	CanUseWithOtherDiscount bool    `json:"canUseWithOtherDiscount,omitempty"`
	ShoppingMallDiscount    []int64 `json:"shoppingMallDiscount,omitempty"`
}

type CouponTemplateBasicUpdateRequestAllSceneDTO struct {
	ShoppingMallSceneList []int64 `json:"shoppingMallSceneList,omitempty"`
}

type CouponTemplateBasicUpdateRequestAcceptTypeDTO struct {
	ExcludeGoodsIds   CouponTemplateBasicUpdateRequestExcludeGoodsIds `json:"excludeGoodsIds,omitempty"`
	AcceptStoreType   int64                                           `json:"acceptStoreType,omitempty"`
	AcceptStoreIds    []int64                                         `json:"acceptStoreIds,omitempty"`
	AcceptGoodsType   int64                                           `json:"acceptGoodsType,omitempty"`
	AcceptGoodsIds    []int64                                         `json:"acceptGoodsIds,omitempty"`
	ExistExcludeGoods bool                                            `json:"existExcludeGoods,omitempty"`
	IncludeStoreGoods bool                                            `json:"includeStoreGoods,omitempty"`
}

type CouponTemplateBasicUpdateRequestExcludeGoodsIds struct {
}

type CouponTemplateBasicUpdateRequestCreateInfo struct {
	OperatorId int64 `json:"operatorId,omitempty"`
}
