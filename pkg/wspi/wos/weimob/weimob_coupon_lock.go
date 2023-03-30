package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCouponLockService struct {
	handler spi.WosInvocationHandler[PaasWeimobCouponLockRequest, PaasWeimobCouponLockBool]
}

func (s PaasWeimobCouponLockService) NewRequest() spi.InvocationRequest[PaasWeimobCouponLockRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCouponLockRequest]{
		Params: &PaasWeimobCouponLockRequest{},
	}
}

func (s PaasWeimobCouponLockService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCouponLockRequest],
) (
	spi.InvocationResponse[PaasWeimobCouponLockBool],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCouponLockRequest]))
}

type PaasWeimobCouponLockRequest struct {
	CouponList         []PaasWeimobCouponLockRequestCouponList         `json:"couponList,omitempty"`
	StoreGoodsInfoDTOS []PaasWeimobCouponLockRequestStoreGoodsInfoDTOS `json:"storeGoodsInfoDTOS,omitempty"`
	SourceObjectList   []PaasWeimobCouponLockRequestSourceObjectList   `json:"sourceObjectList,omitempty"`
	Wid                int64                                           `json:"wid,omitempty"`
}
type PaasWeimobCouponLockRequestCouponList struct {
	Amount   int64  `json:"amount,omitempty"`
	Code     string `json:"code,omitempty"`
	OrderNo  string `json:"orderNo,omitempty"`
	UseScene int64  `json:"useScene,omitempty"`
}
type PaasWeimobCouponLockRequestStoreGoodsInfoDTOS struct {
	GoodsInfos  []PaasWeimobCouponLockRequestGoodsInfos `json:"goodsInfos,omitempty"`
	VidNodeList PaasWeimobCouponLockRequestVidNodeList  `json:"vidNodeList,omitempty"`
}
type PaasWeimobCouponLockRequestGoodsInfos struct {
	Skus        []PaasWeimobCouponLockRequestSkus `json:"skus,omitempty"`
	CategoryIds []int64                           `json:"categoryIds,omitempty"`
	GroupIds    []int64                           `json:"groupIds,omitempty"`
	Id          int64                             `json:"id,omitempty"`
}
type PaasWeimobCouponLockRequestSkus struct {
	Id       int64 `json:"id,omitempty"`
	Num      int64 `json:"num,omitempty"`
	Price    int64 `json:"price,omitempty"`
	Selected bool  `json:"selected,omitempty"`
}
type PaasWeimobCouponLockRequestVidNodeList struct {
	ParentVids []PaasWeimobCouponLockRequestParentVids `json:"parentVids,omitempty"`
	BosId      int64                                   `json:"bosId,omitempty"`
	Vid        int64                                   `json:"vid,omitempty"`
	VidType    int64                                   `json:"vidType,omitempty"`
}
type PaasWeimobCouponLockRequestParentVids struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}
type PaasWeimobCouponLockRequestSourceObjectList struct {
	Source       int64  `json:"source,omitempty"`
	SourceOpenId string `json:"sourceOpenId,omitempty"`
	SourceAppId  string `json:"sourceAppId,omitempty"`
}

type PaasWeimobCouponLockBool bool

func CreatePaasWeimobCouponLockBool() *PaasWeimobCouponLockBool {
	var br PaasWeimobCouponLockBool
	return &br
}

// OnPaasWeimobCouponLockServiceInvocation
// @id 1066
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1066?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=锁定优惠券)
func (s *Service) OnPaasWeimobCouponLockServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCouponLockRequest, PaasWeimobCouponLockBool],
) (service *Service) {
	var invocation = &PaasWeimobCouponLockService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCouponLockRequest, PaasWeimobCouponLockBool](invocation),
		"PaasWeimobCouponLockService",
		"weimob.coupon.lock",
	)
	return s
}
