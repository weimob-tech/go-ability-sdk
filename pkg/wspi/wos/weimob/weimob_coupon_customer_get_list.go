package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCouponCustomerGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobCouponCustomerGetListRequest, PaasWeimobCouponCustomerGetListResponse]
}

func (s PaasWeimobCouponCustomerGetListService) NewRequest() spi.InvocationRequest[PaasWeimobCouponCustomerGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCouponCustomerGetListRequest]{
		Params: &PaasWeimobCouponCustomerGetListRequest{},
	}
}

func (s PaasWeimobCouponCustomerGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCouponCustomerGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobCouponCustomerGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCouponCustomerGetListRequest]))
}

type PaasWeimobCouponCustomerGetListRequest struct {
	SourceObjectList []PaasWeimobCouponCustomerGetListRequestSourceObjectList `json:"sourceObjectList,omitempty"`
	Wid              int64                                                    `json:"wid,omitempty"`
	StatusRange      []int64                                                  `json:"statusRange,omitempty"`
	PageNum          int64                                                    `json:"pageNum,omitempty"`
	PageSize         int64                                                    `json:"pageSize,omitempty"`
}
type PaasWeimobCouponCustomerGetListRequestSourceObjectList struct {
	Source       int64  `json:"source,omitempty"`
	SourceOpenId string `json:"sourceOpenId,omitempty"`
	SourceAppId  string `json:"sourceAppId,omitempty"`
}

type PaasWeimobCouponCustomerGetListResponse struct {
	PageList   []PaasWeimobCouponCustomerGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                             `json:"pageNum,omitempty"`
	PageSize   int64                                             `json:"pageSize,omitempty"`
	TotalCount int64                                             `json:"totalCount,omitempty"`
}
type PaasWeimobCouponCustomerGetListResponsePageList struct {
	ValidDate          PaasWeimobCouponCustomerGetListResponseValidDate     `json:"validDate,omitempty"`
	AcceptTypeDTO      PaasWeimobCouponCustomerGetListResponseAcceptTypeDTO `json:"acceptTypeDTO,omitempty"`
	Code               string                                               `json:"code,omitempty"`
	PublishTime        int64                                                `json:"publishTime,omitempty"`
	AcquireVid         int64                                                `json:"acquireVid,omitempty"`
	PlatformType       int64                                                `json:"platformType,omitempty"`
	CouponSubType      int64                                                `json:"couponSubType,omitempty"`
	CouponTemplateId   int64                                                `json:"couponTemplateId,omitempty"`
	CouponType         int64                                                `json:"couponType,omitempty"`
	ReducePriceType    int64                                                `json:"reducePriceType,omitempty"`
	DiscountLimitType  int64                                                `json:"discountLimitType,omitempty"`
	DiscountLimitValue int64                                                `json:"discountLimitValue,omitempty"`
	Status             int64                                                `json:"status,omitempty"`
	Name               string                                               `json:"name,omitempty"`
	DiscountAmount     int64                                                `json:"discountAmount,omitempty"`
	UseThreshold       int64                                                `json:"useThreshold,omitempty"`
	MaxGoodsCount      int64                                                `json:"maxGoodsCount,omitempty"`
	MinGoodsCount      int64                                                `json:"minGoodsCount,omitempty"`
	CanGift            bool                                                 `json:"canGift,omitempty"`
	CanShare           bool                                                 `json:"canShare,omitempty"`
	Desc               string                                               `json:"desc,omitempty"`
	Explain            string                                               `json:"explain,omitempty"`
	PricingType        int64                                                `json:"pricingType,omitempty"`
	IsAllUseScene      bool                                                 `json:"isAllUseScene,omitempty"`
	AllSceneList       []int64                                              `json:"allSceneList,omitempty"`
	UserTakeLimit      int64                                                `json:"userTakeLimit,omitempty"`
	FreightReduceType  int64                                                `json:"freightReduceType,omitempty"`
	UsingDate          int64                                                `json:"usingDate,omitempty"`
}
type PaasWeimobCouponCustomerGetListResponseValidDate struct {
	UseStartTime   int64 `json:"useStartTime,omitempty"`
	UseEndTime     int64 `json:"useEndTime,omitempty"`
	UseTimeType    int64 `json:"useTimeType,omitempty"`
	UseDelayDays   int64 `json:"useDelayDays,omitempty"`
	ValidDays      int64 `json:"validDays,omitempty"`
	SubUseTimeType int64 `json:"subUseTimeType,omitempty"`
}
type PaasWeimobCouponCustomerGetListResponseAcceptTypeDTO struct {
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

func CreatePaasWeimobCouponCustomerGetListResponse() *PaasWeimobCouponCustomerGetListResponse {
	return &PaasWeimobCouponCustomerGetListResponse{}
}

// OnPaasWeimobCouponCustomerGetListServiceInvocation
// @id 1085
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1085?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询用户优惠券列表)
func (s *Service) OnPaasWeimobCouponCustomerGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCouponCustomerGetListRequest, PaasWeimobCouponCustomerGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobCouponCustomerGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCouponCustomerGetListRequest, PaasWeimobCouponCustomerGetListResponse](invocation),
		"PaasWeimobCouponCustomerGetListService",
		"weimob.coupon.customer.getList",
	)
	return s
}
