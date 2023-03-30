package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCouponUnfreezeService struct {
	handler spi.WosInvocationHandler[PaasWeimobCouponUnfreezeRequest, PaasWeimobCouponUnfreezeBool]
}

func (s PaasWeimobCouponUnfreezeService) NewRequest() spi.InvocationRequest[PaasWeimobCouponUnfreezeRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCouponUnfreezeRequest]{
		Params: &PaasWeimobCouponUnfreezeRequest{},
	}
}

func (s PaasWeimobCouponUnfreezeService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCouponUnfreezeRequest],
) (
	spi.InvocationResponse[PaasWeimobCouponUnfreezeBool],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCouponUnfreezeRequest]))
}

type PaasWeimobCouponUnfreezeRequest struct {
	OperateCoupons   []PaasWeimobCouponUnfreezeRequestOperateCoupons   `json:"operateCoupons,omitempty"`
	SourceObjectList []PaasWeimobCouponUnfreezeRequestSourceObjectList `json:"sourceObjectList,omitempty"`
	Wid              int64                                             `json:"wid,omitempty"`
}
type PaasWeimobCouponUnfreezeRequestOperateCoupons struct {
	Code    string `json:"code,omitempty"`
	OrderNo string `json:"orderNo,omitempty"`
}
type PaasWeimobCouponUnfreezeRequestSourceObjectList struct {
	Source       int64  `json:"source,omitempty"`
	SourceOpenId string `json:"sourceOpenId,omitempty"`
	SourceAppId  string `json:"sourceAppId,omitempty"`
}

type PaasWeimobCouponUnfreezeBool bool

func CreatePaasWeimobCouponUnfreezeBool() *PaasWeimobCouponUnfreezeBool {
	var br PaasWeimobCouponUnfreezeBool
	return &br
}

// OnPaasWeimobCouponUnfreezeServiceInvocation
// @id 1087
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1087?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=优惠券解冻)
func (s *Service) OnPaasWeimobCouponUnfreezeServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCouponUnfreezeRequest, PaasWeimobCouponUnfreezeBool],
) (service *Service) {
	var invocation = &PaasWeimobCouponUnfreezeService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCouponUnfreezeRequest, PaasWeimobCouponUnfreezeBool](invocation),
		"PaasWeimobCouponUnfreezeService",
		"weimob.coupon.unfreeze",
	)
	return s
}
