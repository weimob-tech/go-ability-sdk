package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCouponFreezeService struct {
	handler spi.WosInvocationHandler[PaasWeimobCouponFreezeRequest, PaasWeimobCouponFreezeBool]
}

func (s PaasWeimobCouponFreezeService) NewRequest() spi.InvocationRequest[PaasWeimobCouponFreezeRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCouponFreezeRequest]{
		Params: &PaasWeimobCouponFreezeRequest{},
	}
}

func (s PaasWeimobCouponFreezeService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCouponFreezeRequest],
) (
	spi.InvocationResponse[PaasWeimobCouponFreezeBool],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCouponFreezeRequest]))
}

type PaasWeimobCouponFreezeRequest struct {
	OperateCoupons   []PaasWeimobCouponFreezeRequestOperateCoupons   `json:"operateCoupons,omitempty"`
	SourceObjectList []PaasWeimobCouponFreezeRequestSourceObjectList `json:"sourceObjectList,omitempty"`
	Wid              int64                                           `json:"wid,omitempty"`
}
type PaasWeimobCouponFreezeRequestOperateCoupons struct {
	Code    string `json:"code,omitempty"`
	OrderNo string `json:"orderNo,omitempty"`
}
type PaasWeimobCouponFreezeRequestSourceObjectList struct {
	Source       int64  `json:"source,omitempty"`
	SourceOpenId string `json:"sourceOpenId,omitempty"`
	SourceAppId  string `json:"sourceAppId,omitempty"`
}

type PaasWeimobCouponFreezeBool bool

func CreatePaasWeimobCouponFreezeBool() *PaasWeimobCouponFreezeBool {
	var br PaasWeimobCouponFreezeBool
	return &br
}

// OnPaasWeimobCouponFreezeServiceInvocation
// @id 1086
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1086?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=优惠券冻结)
func (s *Service) OnPaasWeimobCouponFreezeServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCouponFreezeRequest, PaasWeimobCouponFreezeBool],
) (service *Service) {
	var invocation = &PaasWeimobCouponFreezeService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCouponFreezeRequest, PaasWeimobCouponFreezeBool](invocation),
		"PaasWeimobCouponFreezeService",
		"weimob.coupon.freeze",
	)
	return s
}
