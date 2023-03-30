package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCouponPresentReceiveService struct {
	handler spi.WosInvocationHandler[PaasWeimobCouponPresentReceiveRequest, PaasWeimobCouponPresentReceiveBool]
}

func (s PaasWeimobCouponPresentReceiveService) NewRequest() spi.InvocationRequest[PaasWeimobCouponPresentReceiveRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCouponPresentReceiveRequest]{
		Params: &PaasWeimobCouponPresentReceiveRequest{},
	}
}

func (s PaasWeimobCouponPresentReceiveService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCouponPresentReceiveRequest],
) (
	spi.InvocationResponse[PaasWeimobCouponPresentReceiveBool],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCouponPresentReceiveRequest]))
}

type PaasWeimobCouponPresentReceiveRequest struct {
	FromSourceObjectList   []PaasWeimobCouponPresentReceiveRequestFromSourceObjectList   `json:"fromSourceObjectList,omitempty"`
	TargetSourceObjectList []PaasWeimobCouponPresentReceiveRequestTargetSourceObjectList `json:"targetSourceObjectList,omitempty"`
	FromWid                int64                                                         `json:"fromWid,omitempty"`
	TargetWid              int64                                                         `json:"targetWid,omitempty"`
	Code                   string                                                        `json:"code,omitempty"`
}
type PaasWeimobCouponPresentReceiveRequestFromSourceObjectList struct {
	Source       int64  `json:"source,omitempty"`
	SourceOpenId string `json:"sourceOpenId,omitempty"`
	SourceAppId  string `json:"sourceAppId,omitempty"`
}
type PaasWeimobCouponPresentReceiveRequestTargetSourceObjectList struct {
	Source       int64  `json:"source,omitempty"`
	SourceOpenId string `json:"sourceOpenId,omitempty"`
	SourceAppId  string `json:"sourceAppId,omitempty"`
}

type PaasWeimobCouponPresentReceiveBool bool

func CreatePaasWeimobCouponPresentReceiveBool() *PaasWeimobCouponPresentReceiveBool {
	var br PaasWeimobCouponPresentReceiveBool
	return &br
}

// OnPaasWeimobCouponPresentReceiveServiceInvocation
// @id 1069
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1069?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=转赠成功)
func (s *Service) OnPaasWeimobCouponPresentReceiveServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCouponPresentReceiveRequest, PaasWeimobCouponPresentReceiveBool],
) (service *Service) {
	var invocation = &PaasWeimobCouponPresentReceiveService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCouponPresentReceiveRequest, PaasWeimobCouponPresentReceiveBool](invocation),
		"PaasWeimobCouponPresentReceiveService",
		"weimob.coupon.present.receive",
	)
	return s
}
