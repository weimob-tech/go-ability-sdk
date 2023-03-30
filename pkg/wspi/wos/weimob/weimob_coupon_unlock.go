package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCouponUnlockService struct {
	handler spi.WosInvocationHandler[PaasWeimobCouponUnlockRequest, PaasWeimobCouponUnlockBool]
}

func (s PaasWeimobCouponUnlockService) NewRequest() spi.InvocationRequest[PaasWeimobCouponUnlockRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCouponUnlockRequest]{
		Params: &PaasWeimobCouponUnlockRequest{},
	}
}

func (s PaasWeimobCouponUnlockService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCouponUnlockRequest],
) (
	spi.InvocationResponse[PaasWeimobCouponUnlockBool],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCouponUnlockRequest]))
}

type PaasWeimobCouponUnlockRequest struct {
	CouponList       []PaasWeimobCouponUnlockRequestCouponList       `json:"couponList,omitempty"`
	SourceObjectList []PaasWeimobCouponUnlockRequestSourceObjectList `json:"sourceObjectList,omitempty"`
	OrderAmount      int64                                           `json:"orderAmount,omitempty"`
	Wid              int64                                           `json:"wid,omitempty"`
}
type PaasWeimobCouponUnlockRequestCouponList struct {
	Wid              int64  `json:"wid,omitempty"`
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	OrderNo          string `json:"orderNo,omitempty"`
	Code             string `json:"code,omitempty"`
}
type PaasWeimobCouponUnlockRequestSourceObjectList struct {
	Source       int64  `json:"source,omitempty"`
	SourceOpenId string `json:"sourceOpenId,omitempty"`
	SourceAppId  string `json:"sourceAppId,omitempty"`
}

type PaasWeimobCouponUnlockBool bool

func CreatePaasWeimobCouponUnlockBool() *PaasWeimobCouponUnlockBool {
	var br PaasWeimobCouponUnlockBool
	return &br
}

// OnPaasWeimobCouponUnlockServiceInvocation
// @id 1071
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1071?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=解锁优惠券)
func (s *Service) OnPaasWeimobCouponUnlockServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCouponUnlockRequest, PaasWeimobCouponUnlockBool],
) (service *Service) {
	var invocation = &PaasWeimobCouponUnlockService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCouponUnlockRequest, PaasWeimobCouponUnlockBool](invocation),
		"PaasWeimobCouponUnlockService",
		"weimob.coupon.unlock",
	)
	return s
}
