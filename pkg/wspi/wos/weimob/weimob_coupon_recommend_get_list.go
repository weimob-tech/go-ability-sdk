package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCouponRecommendGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobCouponRecommendGetListRequest, PaasWeimobCouponRecommendGetListResponse]
}

func (s PaasWeimobCouponRecommendGetListService) NewRequest() spi.InvocationRequest[PaasWeimobCouponRecommendGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCouponRecommendGetListRequest]{
		Params: &PaasWeimobCouponRecommendGetListRequest{},
	}
}

func (s PaasWeimobCouponRecommendGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCouponRecommendGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobCouponRecommendGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCouponRecommendGetListRequest]))
}

type PaasWeimobCouponRecommendGetListRequest struct {
	ShopGoods        []PaasWeimobCouponRecommendGetListRequestShopGoods        `json:"shopGoods,omitempty"`
	SourceObjectList []PaasWeimobCouponRecommendGetListRequestSourceObjectList `json:"sourceObjectList,omitempty"`
	Wid              int64                                                     `json:"wid,omitempty"`
	PageNum          int64                                                     `json:"pageNum,omitempty"`
	PageSize         int64                                                     `json:"pageSize,omitempty"`
	CouponTypes      []int64                                                   `json:"couponTypes,omitempty"`
	PlatformTypes    []int64                                                   `json:"platformTypes,omitempty"`
	SubTypes         []int64                                                   `json:"subTypes,omitempty"`
	PublishChannel   int64                                                     `json:"publishChannel,omitempty"`
}
type PaasWeimobCouponRecommendGetListRequestShopGoods struct {
	GoodsItems []PaasWeimobCouponRecommendGetListRequestGoodsItems `json:"goodsItems,omitempty"`
	Vid        int64                                               `json:"vid,omitempty"`
	VidType    int64                                               `json:"vidType,omitempty"`
}
type PaasWeimobCouponRecommendGetListRequestGoodsItems struct {
	Skus         []PaasWeimobCouponRecommendGetListRequestSkus `json:"skus,omitempty"`
	Id           int64                                         `json:"id,omitempty"`
	CategoryIds  []int64                                       `json:"categoryIds,omitempty"`
	GroupIds     []int64                                       `json:"groupIds,omitempty"`
	IsStoreGoods bool                                          `json:"isStoreGoods,omitempty"`
}
type PaasWeimobCouponRecommendGetListRequestSkus struct {
	Id    int64 `json:"id,omitempty"`
	Price int64 `json:"price,omitempty"`
	Num   int64 `json:"num,omitempty"`
}
type PaasWeimobCouponRecommendGetListRequestSourceObjectList struct {
	Source       int64  `json:"source,omitempty"`
	SourceOpenId string `json:"sourceOpenId,omitempty"`
	SourceAppId  string `json:"sourceAppId,omitempty"`
}

type PaasWeimobCouponRecommendGetListResponse struct {
	RecommendCouponList []PaasWeimobCouponRecommendGetListResponseRecommendCouponList `json:"recommendCouponList,omitempty"`
}
type PaasWeimobCouponRecommendGetListResponseRecommendCouponList struct {
	ValidDate          PaasWeimobCouponRecommendGetListResponseValidDate     `json:"validDate,omitempty"`
	AcceptTypeDTO      PaasWeimobCouponRecommendGetListResponseAcceptTypeDTO `json:"acceptTypeDTO,omitempty"`
	CouponTemplateId   int64                                                 `json:"couponTemplateId,omitempty"`
	Name               string                                                `json:"name,omitempty"`
	Status             int64                                                 `json:"status,omitempty"`
	PlatformType       int64                                                 `json:"platformType,omitempty"`
	CouponType         int64                                                 `json:"couponType,omitempty"`
	CouponSubType      int64                                                 `json:"couponSubType,omitempty"`
	Explain            string                                                `json:"explain,omitempty"`
	Desc               string                                                `json:"desc,omitempty"`
	UserTakeLimit      int64                                                 `json:"userTakeLimit,omitempty"`
	ReducePriceType    int64                                                 `json:"reducePriceType,omitempty"`
	DiscountAmount     int64                                                 `json:"discountAmount,omitempty"`
	UseThreshold       int64                                                 `json:"useThreshold,omitempty"`
	DiscountLimitType  int64                                                 `json:"discountLimitType,omitempty"`
	DiscountLimitValue int64                                                 `json:"discountLimitValue,omitempty"`
	MinCashTicketAmt   int64                                                 `json:"minCashTicketAmt,omitempty"`
	MaxCashTicketAmt   int64                                                 `json:"maxCashTicketAmt,omitempty"`
	MinGoodsCount      int64                                                 `json:"minGoodsCount,omitempty"`
	MaxGoodsCount      int64                                                 `json:"maxGoodsCount,omitempty"`
	IsPartTimeUse      bool                                                  `json:"isPartTimeUse,omitempty"`
	IsAllUseScene      bool                                                  `json:"isAllUseScene,omitempty"`
	AllSceneList       []int64                                               `json:"allSceneList,omitempty"`
	CanGift            bool                                                  `json:"canGift,omitempty"`
	CanShare           bool                                                  `json:"canShare,omitempty"`
	FreightReduceType  int64                                                 `json:"freightReduceType,omitempty"`
	Outer              bool                                                  `json:"outer,omitempty"`
	ThirdTemplateId    string                                                `json:"thirdTemplateId,omitempty"`
	PricingType        int64                                                 `json:"pricingType,omitempty"`
}
type PaasWeimobCouponRecommendGetListResponseValidDate struct {
	UseEndTime   int64 `json:"useEndTime,omitempty"`
	UseTimeType  int64 `json:"useTimeType,omitempty"`
	UseStartTime int64 `json:"useStartTime,omitempty"`
	UseDelayDays int64 `json:"useDelayDays,omitempty"`
	ValidDays    int64 `json:"validDays,omitempty"`
}
type PaasWeimobCouponRecommendGetListResponseAcceptTypeDTO struct {
	AcceptStoreIds    []int64 `json:"acceptStoreIds,omitempty"`
	ExistExcludeStore bool    `json:"existExcludeStore,omitempty"`
	AcceptGoodsIds    []int64 `json:"acceptGoodsIds,omitempty"`
	AcceptStoreType   int64   `json:"acceptStoreType,omitempty"`
	AcceptGoodsType   int64   `json:"acceptGoodsType,omitempty"`
	IncludeStoreGoods bool    `json:"includeStoreGoods,omitempty"`
	ExistExcludeGoods bool    `json:"existExcludeGoods,omitempty"`
	ExcludeGoodsIds   []int64 `json:"excludeGoodsIds,omitempty"`
	ExcludeStoreIds   []int64 `json:"excludeStoreIds,omitempty"`
}

func CreatePaasWeimobCouponRecommendGetListResponse() *PaasWeimobCouponRecommendGetListResponse {
	return &PaasWeimobCouponRecommendGetListResponse{}
}

// OnPaasWeimobCouponRecommendGetListServiceInvocation
// @id 1077
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1077?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询适用的模板列表-商详推荐券等)
func (s *Service) OnPaasWeimobCouponRecommendGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCouponRecommendGetListRequest, PaasWeimobCouponRecommendGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobCouponRecommendGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCouponRecommendGetListRequest, PaasWeimobCouponRecommendGetListResponse](invocation),
		"PaasWeimobCouponRecommendGetListService",
		"weimob.coupon.recommend.getList",
	)
	return s
}
