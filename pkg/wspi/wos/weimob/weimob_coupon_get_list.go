package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCouponGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobCouponGetListRequest, PaasWeimobCouponGetListResponse]
}

func (s PaasWeimobCouponGetListService) NewRequest() spi.InvocationRequest[PaasWeimobCouponGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCouponGetListRequest]{
		Params: &PaasWeimobCouponGetListRequest{},
	}
}

func (s PaasWeimobCouponGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCouponGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobCouponGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCouponGetListRequest]))
}

type PaasWeimobCouponGetListRequest struct {
	SourceObjectList []PaasWeimobCouponGetListRequestSourceObjectList `json:"sourceObjectList,omitempty"`
	Codes            []int64                                          `json:"codes,omitempty"`
	Wid              int64                                            `json:"wid,omitempty"`
	RequestId        string                                           `json:"requestId,omitempty"`
}
type PaasWeimobCouponGetListRequestSourceObjectList struct {
	Source       int64  `json:"source,omitempty"`
	SourceOpenId string `json:"sourceOpenId,omitempty"`
	SourceAppId  string `json:"sourceAppId,omitempty"`
}

type PaasWeimobCouponGetListResponse struct {
	ValidDate          PaasWeimobCouponGetListResponseValidDate     `json:"validDate,omitempty"`
	AcceptTypeDTO      PaasWeimobCouponGetListResponseAcceptTypeDTO `json:"acceptTypeDTO,omitempty"`
	CouponTemplateId   int64                                        `json:"couponTemplateId,omitempty"`
	Code               string                                       `json:"code,omitempty"`
	Name               string                                       `json:"name,omitempty"`
	CouponType         int64                                        `json:"couponType,omitempty"`
	Desc               string                                       `json:"desc,omitempty"`
	Explain            string                                       `json:"explain,omitempty"`
	UseThreshold       int64                                        `json:"useThreshold,omitempty"`
	PublishTime        int64                                        `json:"publishTime,omitempty"`
	ExpireTime         int64                                        `json:"expireTime,omitempty"`
	Status             int64                                        `json:"status,omitempty"`
	Stock              int64                                        `json:"stock,omitempty"`
	UserTakeLimit      int64                                        `json:"userTakeLimit,omitempty"`
	PublishCount       int64                                        `json:"publishCount,omitempty"`
	CashTicketAmt      int64                                        `json:"cashTicketAmt,omitempty"`
	Discount           int64                                        `json:"discount,omitempty"`
	DiscountLimitType  int64                                        `json:"discountLimitType,omitempty"`
	DiscountLimitValue int64                                        `json:"discountLimitValue,omitempty"`
}
type PaasWeimobCouponGetListResponseValidDate struct {
	UseTimeType  int64 `json:"useTimeType,omitempty"`
	UseStartTime int64 `json:"useStartTime,omitempty"`
	UseEndTime   int64 `json:"useEndTime,omitempty"`
}
type PaasWeimobCouponGetListResponseAcceptTypeDTO struct {
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

func CreatePaasWeimobCouponGetListResponse() *PaasWeimobCouponGetListResponse {
	return &PaasWeimobCouponGetListResponse{}
}

// OnPaasWeimobCouponGetListServiceInvocation
// @id 1080
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1080?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=根据code查优惠券信息)
func (s *Service) OnPaasWeimobCouponGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCouponGetListRequest, PaasWeimobCouponGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobCouponGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCouponGetListRequest, PaasWeimobCouponGetListResponse](invocation),
		"PaasWeimobCouponGetListService",
		"weimob.coupon.getList",
	)
	return s
}
