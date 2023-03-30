package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCouponPresentConfirmService struct {
	handler spi.WosInvocationHandler[PaasWeimobCouponPresentConfirmRequest, PaasWeimobCouponPresentConfirmBool]
}

func (s PaasWeimobCouponPresentConfirmService) NewRequest() spi.InvocationRequest[PaasWeimobCouponPresentConfirmRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCouponPresentConfirmRequest]{
		Params: &PaasWeimobCouponPresentConfirmRequest{},
	}
}

func (s PaasWeimobCouponPresentConfirmService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCouponPresentConfirmRequest],
) (
	spi.InvocationResponse[PaasWeimobCouponPresentConfirmBool],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCouponPresentConfirmRequest]))
}

type PaasWeimobCouponPresentConfirmRequest struct {
	SourceObjectList []PaasWeimobCouponPresentConfirmRequestSourceObjectList `json:"sourceObjectList,omitempty"`
	FromWid          int64                                                   `json:"fromWid,omitempty"`
	Code             string                                                  `json:"code,omitempty"`
}
type PaasWeimobCouponPresentConfirmRequestSourceObjectList struct {
	Source       int64  `json:"source,omitempty"`
	SourceOpenId string `json:"sourceOpenId,omitempty"`
	SourceAppId  string `json:"sourceAppId,omitempty"`
}

type PaasWeimobCouponPresentConfirmBool bool

func CreatePaasWeimobCouponPresentConfirmBool() *PaasWeimobCouponPresentConfirmBool {
	var br PaasWeimobCouponPresentConfirmBool
	return &br
}

// OnPaasWeimobCouponPresentConfirmServiceInvocation
// @id 1068
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1068?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=发起转赠)
func (s *Service) OnPaasWeimobCouponPresentConfirmServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCouponPresentConfirmRequest, PaasWeimobCouponPresentConfirmBool],
) (service *Service) {
	var invocation = &PaasWeimobCouponPresentConfirmService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCouponPresentConfirmRequest, PaasWeimobCouponPresentConfirmBool](invocation),
		"PaasWeimobCouponPresentConfirmService",
		"weimob.coupon.present.confirm",
	)
	return s
}
