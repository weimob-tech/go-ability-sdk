package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCouponCancelService struct {
	handler spi.WosInvocationHandler[PaasWeimobCouponCancelRequest, PaasWeimobCouponCancelBool]
}

func (s PaasWeimobCouponCancelService) NewRequest() spi.InvocationRequest[PaasWeimobCouponCancelRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCouponCancelRequest]{
		Params: &PaasWeimobCouponCancelRequest{},
	}
}

func (s PaasWeimobCouponCancelService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCouponCancelRequest],
) (
	spi.InvocationResponse[PaasWeimobCouponCancelBool],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCouponCancelRequest]))
}

type PaasWeimobCouponCancelRequest struct {
	OperateCoupons   []PaasWeimobCouponCancelRequestOperateCoupons   `json:"operateCoupons,omitempty"`
	SourceObjectList []PaasWeimobCouponCancelRequestSourceObjectList `json:"sourceObjectList,omitempty"`
	UseScene         int64                                           `json:"useScene,omitempty"`
	Wid              int64                                           `json:"wid,omitempty"`
}
type PaasWeimobCouponCancelRequestOperateCoupons struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	OrderNo          string `json:"orderNo,omitempty"`
	Code             string `json:"code,omitempty"`
	Amount           int64  `json:"amount,omitempty"`
}
type PaasWeimobCouponCancelRequestSourceObjectList struct {
	Source       int64  `json:"source,omitempty"`
	SourceOpenId string `json:"sourceOpenId,omitempty"`
	SourceAppId  string `json:"sourceAppId,omitempty"`
}

type PaasWeimobCouponCancelBool bool

func CreatePaasWeimobCouponCancelBool() *PaasWeimobCouponCancelBool {
	var br PaasWeimobCouponCancelBool
	return &br
}

// OnPaasWeimobCouponCancelServiceInvocation
// @id 1075
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1075?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=作废优惠券)
func (s *Service) OnPaasWeimobCouponCancelServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCouponCancelRequest, PaasWeimobCouponCancelBool],
) (service *Service) {
	var invocation = &PaasWeimobCouponCancelService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCouponCancelRequest, PaasWeimobCouponCancelBool](invocation),
		"PaasWeimobCouponCancelService",
		"weimob.coupon.cancel",
	)
	return s
}
