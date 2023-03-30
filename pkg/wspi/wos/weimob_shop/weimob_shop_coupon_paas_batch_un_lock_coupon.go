package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobShopCouponPaasBatchUnLockCouponService struct {
	handler spi.WosInvocationHandler[PaasWeimobShopCouponPaasBatchUnLockCouponRequest, PaasWeimobShopCouponPaasBatchUnLockCouponBool]
}

func (s PaasWeimobShopCouponPaasBatchUnLockCouponService) NewRequest() spi.InvocationRequest[PaasWeimobShopCouponPaasBatchUnLockCouponRequest] {
	return &spi.WosInvocationRequest[PaasWeimobShopCouponPaasBatchUnLockCouponRequest]{
		Params: &PaasWeimobShopCouponPaasBatchUnLockCouponRequest{},
	}
}

func (s PaasWeimobShopCouponPaasBatchUnLockCouponService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobShopCouponPaasBatchUnLockCouponRequest],
) (
	spi.InvocationResponse[PaasWeimobShopCouponPaasBatchUnLockCouponBool],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobShopCouponPaasBatchUnLockCouponRequest]))
}

type PaasWeimobShopCouponPaasBatchUnLockCouponRequest struct {
	CouponList        []PaasWeimobShopCouponPaasBatchUnLockCouponRequestCouponList `json:"couponList,omitempty"`
	Wid               int64                                                        `json:"wid,omitempty"`
	MerchantId        int64                                                        `json:"merchantId,omitempty"`
	VidType           int64                                                        `json:"vidType,omitempty"`
	Vid               int64                                                        `json:"vid,omitempty"`
	BosId             int64                                                        `json:"bosId,omitempty"`
	ProductInstanceId int64                                                        `json:"productInstanceId,omitempty"`
	Tcode             string                                                       `json:"tcode,omitempty"`
}
type PaasWeimobShopCouponPaasBatchUnLockCouponRequestCouponList struct {
	BosId            int64  `json:"bosId,omitempty"`
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	OrderNo          string `json:"orderNo,omitempty"`
	Code             int64  `json:"code,omitempty"`
	Amount           int64  `json:"amount,omitempty"`
	UseScene         int64  `json:"useScene,omitempty"`
}

type PaasWeimobShopCouponPaasBatchUnLockCouponBool bool

func CreatePaasWeimobShopCouponPaasBatchUnLockCouponBool() *PaasWeimobShopCouponPaasBatchUnLockCouponBool {
	var br PaasWeimobShopCouponPaasBatchUnLockCouponBool
	return &br
}

// OnPaasWeimobShopCouponPaasBatchUnLockCouponServiceInvocation
// @id 579
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/579?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=解锁用户优惠券)
func (s *Service) OnPaasWeimobShopCouponPaasBatchUnLockCouponServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobShopCouponPaasBatchUnLockCouponRequest, PaasWeimobShopCouponPaasBatchUnLockCouponBool],
) (service *Service) {
	var invocation = &PaasWeimobShopCouponPaasBatchUnLockCouponService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobShopCouponPaasBatchUnLockCouponRequest, PaasWeimobShopCouponPaasBatchUnLockCouponBool](invocation),
		"PaasWeimobShopCouponPaasBatchUnLockCouponService",
		"weimobShop.coupon.paasBatchUnLockCoupon",
	)
	return s
}
