package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponCustomerAvailableGetList
// @id 2042
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2042?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询用户可用券列表)
func (client *Client) CouponCustomerAvailableGetList(request *CouponCustomerAvailableGetListRequest) (response *CouponCustomerAvailableGetListResponse, err error) {
	rpcResponse := CreateCouponCustomerAvailableGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponCustomerAvailableGetListRequest struct {
	*api.BaseRequest

	GoodsInfos     []CouponCustomerAvailableGetListRequestGoodsInfos `json:"goodsInfos,omitempty"`
	QuerySceneType int64                                             `json:"querySceneType,omitempty"`
	Scene          int64                                             `json:"scene,omitempty"`
	Vid            int64                                             `json:"vid,omitempty"`
	Wid            int64                                             `json:"wid,omitempty"`
	OrderAmount    string                                            `json:"orderAmount,omitempty"`
	VidType        int64                                             `json:"vidType,omitempty"`
}

type CouponCustomerAvailableGetListResponse struct {
	AvailableCoupons             []CouponCustomerAvailableGetListResponseAvailableCoupons             `json:"availableCoupons,omitempty"`
	CannotReceiveCouponTemplates []CouponCustomerAvailableGetListResponseCannotReceiveCouponTemplates `json:"cannotReceiveCouponTemplates,omitempty"`
	UnavailableCoupons           []CouponCustomerAvailableGetListResponseUnavailableCoupons           `json:"unavailableCoupons,omitempty"`
}

func CreateCouponCustomerAvailableGetListRequest() (request *CouponCustomerAvailableGetListRequest) {
	request = &CouponCustomerAvailableGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/customer/available/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponCustomerAvailableGetListResponse() (response *api.BaseResponse[CouponCustomerAvailableGetListResponse]) {
	response = api.CreateResponse[CouponCustomerAvailableGetListResponse](&CouponCustomerAvailableGetListResponse{})
	return
}

type CouponCustomerAvailableGetListRequestGoodsInfos struct {
	Goods   []CouponCustomerAvailableGetListRequestGoods `json:"goods,omitempty"`
	VidInfo CouponCustomerAvailableGetListRequestVidInfo `json:"vidInfo,omitempty"`
}

type CouponCustomerAvailableGetListRequestGoods struct {
	Skus      []CouponCustomerAvailableGetListRequestSkus    `json:"skus,omitempty"`
	BelongVid CouponCustomerAvailableGetListRequestBelongVid `json:"belongVid,omitempty"`
	Id        int64                                          `json:"id,omitempty"`
}

type CouponCustomerAvailableGetListRequestSkus struct {
	Price       string `json:"price,omitempty"`
	Num         int64  `json:"num,omitempty"`
	Id          int64  `json:"id,omitempty"`
	MarketPrice string `json:"marketPrice,omitempty"`
}

type CouponCustomerAvailableGetListRequestBelongVid struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type CouponCustomerAvailableGetListRequestVidInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type CouponCustomerAvailableGetListResponseAvailableCoupons struct {
}

type CouponCustomerAvailableGetListResponseCannotReceiveCouponTemplates struct {
}

type CouponCustomerAvailableGetListResponseUnavailableCoupons struct {
	CanUseDiscountDTO       CouponCustomerAvailableGetListResponseCanUseDiscountDTO    `json:"canUseDiscountDTO,omitempty"`
	ReceiveCouponInfos      []CouponCustomerAvailableGetListResponseReceiveCouponInfos `json:"receiveCouponInfos,omitempty"`
	SkuInfos                []CouponCustomerAvailableGetListResponseSkuInfos           `json:"skuInfos,omitempty"`
	ValidDate               CouponCustomerAvailableGetListResponseValidDate            `json:"validDate,omitempty"`
	BosId                   int64                                                      `json:"bosId,omitempty"`
	CanCashTicket           bool                                                       `json:"canCashTicket,omitempty"`
	CanUseWithOtherDiscount bool                                                       `json:"canUseWithOtherDiscount,omitempty"`
	ConditionType           int64                                                      `json:"conditionType,omitempty"`
	CouponTemplateId        int64                                                      `json:"couponTemplateId,omitempty"`
	CouponType              int64                                                      `json:"couponType,omitempty"`
	Desc                    string                                                     `json:"desc,omitempty"`
	ErrorCode               string                                                     `json:"errorCode,omitempty"`
	ErrorMessage            string                                                     `json:"errorMessage,omitempty"`
	Explain                 string                                                     `json:"explain,omitempty"`
	IsCanSend               bool                                                       `json:"isCanSend,omitempty"`
	MaxCashTicketAmt        string                                                     `json:"maxCashTicketAmt,omitempty"`
	MaxGoodsNumber          int64                                                      `json:"maxGoodsNumber,omitempty"`
	MinCashTicketAmt        string                                                     `json:"minCashTicketAmt,omitempty"`
	MinGoodsNumber          int64                                                      `json:"minGoodsNumber,omitempty"`
	Name                    string                                                     `json:"name,omitempty"`
	OrderDeductList         []int64                                                    `json:"orderDeductList,omitempty"`
	PlatformType            int64                                                      `json:"platformType,omitempty"`
	PricingType             int64                                                      `json:"pricingType,omitempty"`
	ReducePriceType         int64                                                      `json:"reducePriceType,omitempty"`
	SubName                 string                                                     `json:"subName,omitempty"`
	UseConditionValue       int64                                                      `json:"useConditionValue,omitempty"`
	Vid                     int64                                                      `json:"vid,omitempty"`
}

type CouponCustomerAvailableGetListResponseCanUseDiscountDTO struct {
	CanUseWithOtherDiscount bool    `json:"canUseWithOtherDiscount,omitempty"`
	ShoppingMallDiscount    []int64 `json:"shoppingMallDiscount,omitempty"`
}

type CouponCustomerAvailableGetListResponseReceiveCouponInfos struct {
	Code             string `json:"code,omitempty"`
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	DiscountAmount   string `json:"discountAmount,omitempty"`
	ExpireTime       string `json:"expireTime,omitempty"`
	Name             string `json:"name,omitempty"`
	PublishTime      string `json:"publishTime,omitempty"`
	UseThreshold     int64  `json:"useThreshold,omitempty"`
	Wid              int64  `json:"wid,omitempty"`
}

type CouponCustomerAvailableGetListResponseSkuInfos struct {
	GoodsId int64 `json:"goodsId,omitempty"`
	SkuId   int64 `json:"skuId,omitempty"`
	Vid     int64 `json:"vid,omitempty"`
}

type CouponCustomerAvailableGetListResponseValidDate struct {
	UseDelayDays int64  `json:"useDelayDays,omitempty"`
	UseEndTime   string `json:"useEndTime,omitempty"`
	UseStartTime string `json:"useStartTime,omitempty"`
	UseTimeType  int64  `json:"useTimeType,omitempty"`
	ValidDays    int64  `json:"validDays,omitempty"`
}
