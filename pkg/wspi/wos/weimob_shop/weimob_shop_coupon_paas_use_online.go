package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobShopCouponPaasUseOnlineService struct {
	handler spi.WosInvocationHandler[PaasWeimobShopCouponPaasUseOnlineRequest, PaasWeimobShopCouponPaasUseOnlineBool]
}

func (s PaasWeimobShopCouponPaasUseOnlineService) NewRequest() spi.InvocationRequest[PaasWeimobShopCouponPaasUseOnlineRequest] {
	return &spi.WosInvocationRequest[PaasWeimobShopCouponPaasUseOnlineRequest]{
		Params: &PaasWeimobShopCouponPaasUseOnlineRequest{},
	}
}

func (s PaasWeimobShopCouponPaasUseOnlineService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobShopCouponPaasUseOnlineRequest],
) (
	spi.InvocationResponse[PaasWeimobShopCouponPaasUseOnlineBool],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobShopCouponPaasUseOnlineRequest]))
}

type PaasWeimobShopCouponPaasUseOnlineRequest struct {
	CouponList        []PaasWeimobShopCouponPaasUseOnlineRequestCouponList `json:"couponList,omitempty"`
	Wid               int64                                                `json:"wid,omitempty"`
	OrderAmount       int64                                                `json:"orderAmount,omitempty"`
	MerchantId        int64                                                `json:"merchantId,omitempty"`
	BosId             int64                                                `json:"bosId,omitempty"`
	ProductId         int64                                                `json:"productId,omitempty"`
	ProductInstanceId int64                                                `json:"productInstanceId,omitempty"`
	Tcode             string                                               `json:"tcode,omitempty"`
	Vid               int64                                                `json:"vid,omitempty"`
	VidType           int64                                                `json:"vidType,omitempty"`
}
type PaasWeimobShopCouponPaasUseOnlineRequestCouponList struct {
	Amount           int64  `json:"amount,omitempty"`
	Code             string `json:"code,omitempty"`
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	OrderNo          int64  `json:"orderNo,omitempty"`
	UseScene         string `json:"useScene,omitempty"`
	Vid              int64  `json:"vid,omitempty"`
}

type PaasWeimobShopCouponPaasUseOnlineBool bool

func CreatePaasWeimobShopCouponPaasUseOnlineBool() *PaasWeimobShopCouponPaasUseOnlineBool {
	var br PaasWeimobShopCouponPaasUseOnlineBool
	return &br
}

// OnPaasWeimobShopCouponPaasUseOnlineServiceInvocation
// @id 578
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/578?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=核销优惠券)
func (s *Service) OnPaasWeimobShopCouponPaasUseOnlineServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobShopCouponPaasUseOnlineRequest, PaasWeimobShopCouponPaasUseOnlineBool],
) (service *Service) {
	var invocation = &PaasWeimobShopCouponPaasUseOnlineService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobShopCouponPaasUseOnlineRequest, PaasWeimobShopCouponPaasUseOnlineBool](invocation),
		"PaasWeimobShopCouponPaasUseOnlineService",
		"weimobShop.coupon.paasUseOnline",
	)
	return s
}
