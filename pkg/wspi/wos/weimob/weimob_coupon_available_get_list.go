package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCouponAvailableGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobCouponAvailableGetListRequest, PaasWeimobCouponAvailableGetListResponse]
}

func (s PaasWeimobCouponAvailableGetListService) NewRequest() spi.InvocationRequest[PaasWeimobCouponAvailableGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCouponAvailableGetListRequest]{
		Params: &PaasWeimobCouponAvailableGetListRequest{},
	}
}

func (s PaasWeimobCouponAvailableGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCouponAvailableGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobCouponAvailableGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCouponAvailableGetListRequest]))
}

type PaasWeimobCouponAvailableGetListRequest struct {
	GoodsInfos     []PaasWeimobCouponAvailableGetListRequestGoodsInfos  `json:"goodsInfos,omitempty"`
	ParentVidInfo  PaasWeimobCouponAvailableGetListRequestParentVidInfo `json:"parentVidInfo,omitempty"`
	Scene          int64                                                `json:"scene,omitempty"`
	Source         string                                               `json:"source,omitempty"`
	Wid            int64                                                `json:"wid,omitempty"`
	QuerySceneType int64                                                `json:"querySceneType,omitempty"`
}
type PaasWeimobCouponAvailableGetListRequestGoodsInfos struct {
	Goods   []PaasWeimobCouponAvailableGetListRequestGoods `json:"goods,omitempty"`
	VidInfo PaasWeimobCouponAvailableGetListRequestVidInfo `json:"vidInfo,omitempty"`
}
type PaasWeimobCouponAvailableGetListRequestGoods struct {
	Skus        []PaasWeimobCouponAvailableGetListRequestSkus `json:"skus,omitempty"`
	Id          int64                                         `json:"id,omitempty"`
	CategoryIds []int64                                       `json:"categoryIds,omitempty"`
	GroupIds    []int64                                       `json:"groupIds,omitempty"`
	Tags        []int64                                       `json:"tags,omitempty"`
}
type PaasWeimobCouponAvailableGetListRequestSkus struct {
	Id    int64 `json:"id,omitempty"`
	Num   int64 `json:"num,omitempty"`
	Price int64 `json:"price,omitempty"`
}
type PaasWeimobCouponAvailableGetListRequestVidInfo struct {
	ParentVids []PaasWeimobCouponAvailableGetListRequestParentVids `json:"parentVids,omitempty"`
	Vid        int64                                               `json:"vid,omitempty"`
	VidType    int64                                               `json:"vidType,omitempty"`
}
type PaasWeimobCouponAvailableGetListRequestParentVids struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}
type PaasWeimobCouponAvailableGetListRequestParentVidInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type PaasWeimobCouponAvailableGetListResponse struct {
	ResourceUseDTOS            []PaasWeimobCouponAvailableGetListResponseResourceUseDTOS            `json:"resourceUseDTOS,omitempty"`
	CannotSendResourceUseDTOS  []PaasWeimobCouponAvailableGetListResponseCannotSendResourceUseDTOS  `json:"cannotSendResourceUseDTOS,omitempty"`
	UnavailableResourceUseDTOS []PaasWeimobCouponAvailableGetListResponseUnavailableResourceUseDTOS `json:"unavailableResourceUseDTOS,omitempty"`
}
type PaasWeimobCouponAvailableGetListResponseResourceUseDTOS struct {
	Resource           PaasWeimobCouponAvailableGetListResponseResource      `json:"resource,omitempty"`
	SkuInfoList        []PaasWeimobCouponAvailableGetListResponseSkuInfoList `json:"skuInfoList,omitempty"`
	UseRule            PaasWeimobCouponAvailableGetListResponseUseRule       `json:"useRule,omitempty"`
	IsCanSend          bool                                                  `json:"isCanSend,omitempty"`
	Name               string                                                `json:"name,omitempty"`
	ResourceTemplateId int64                                                 `json:"resourceTemplateId,omitempty"`
}
type PaasWeimobCouponAvailableGetListResponseResource struct {
	CanGift            bool   `json:"canGift,omitempty"`
	CanShare           bool   `json:"canShare,omitempty"`
	Code               string `json:"code,omitempty"`
	CreateTime         int64  `json:"createTime,omitempty"`
	Desc               string `json:"desc,omitempty"`
	EndTime            int64  `json:"endTime,omitempty"`
	Name               string `json:"name,omitempty"`
	PromotionType      int64  `json:"promotionType,omitempty"`
	PublishTime        int64  `json:"publishTime,omitempty"`
	ResourceTemplateId int64  `json:"resourceTemplateId,omitempty"`
	SendChannel        int64  `json:"sendChannel,omitempty"`
	StartTime          int64  `json:"startTime,omitempty"`
	Status             int64  `json:"status,omitempty"`
	StockId            int64  `json:"stockId,omitempty"`
	SubName            string `json:"subName,omitempty"`
	SubType            int64  `json:"subType,omitempty"`
	ThresholdType      int64  `json:"thresholdType,omitempty"`
	Type               int64  `json:"type,omitempty"`
	UseAmount          int64  `json:"useAmount,omitempty"`
	UseThreshold       int64  `json:"useThreshold,omitempty"`
	Wid                int64  `json:"wid,omitempty"`
}
type PaasWeimobCouponAvailableGetListResponseSkuInfoList struct {
	GoodsId int64 `json:"goodsId,omitempty"`
	SkuId   int64 `json:"skuId,omitempty"`
	Vid     int64 `json:"vid,omitempty"`
}
type PaasWeimobCouponAvailableGetListResponseUseRule struct {
	LimitPromotionTypes  PaasWeimobCouponAvailableGetListResponseLimitPromotionTypes `json:"limitPromotionTypes,omitempty"`
	CanUseDeductionsList []int64                                                     `json:"canUseDeductionsList,omitempty"`
	ConditionType        int64                                                       `json:"conditionType,omitempty"`
	ConditionValue       int64                                                       `json:"conditionValue,omitempty"`
	CouponCost           int64                                                       `json:"couponCost,omitempty"`
	CouponCostType       int64                                                       `json:"couponCostType,omitempty"`
	FreightAmount        int64                                                       `json:"freightAmount,omitempty"`
	PricingType          int64                                                       `json:"pricingType,omitempty"`
}
type PaasWeimobCouponAvailableGetListResponseLimitPromotionTypes struct {
}
type PaasWeimobCouponAvailableGetListResponseCannotSendResourceUseDTOS struct {
}
type PaasWeimobCouponAvailableGetListResponseUnavailableResourceUseDTOS struct {
}

func CreatePaasWeimobCouponAvailableGetListResponse() *PaasWeimobCouponAvailableGetListResponse {
	return &PaasWeimobCouponAvailableGetListResponse{}
}

// OnPaasWeimobCouponAvailableGetListServiceInvocation
// @id 1065
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1065?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=结算页查询用户可用优惠券)
func (s *Service) OnPaasWeimobCouponAvailableGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCouponAvailableGetListRequest, PaasWeimobCouponAvailableGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobCouponAvailableGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCouponAvailableGetListRequest, PaasWeimobCouponAvailableGetListResponse](invocation),
		"PaasWeimobCouponAvailableGetListService",
		"weimob.coupon.available.getList",
	)
	return s
}
