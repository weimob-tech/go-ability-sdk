package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCouponCustomerSceneCountService struct {
	handler spi.WosInvocationHandler[PaasWeimobCouponCustomerSceneCountRequest, PaasWeimobCouponCustomerSceneCountInt64]
}

func (s PaasWeimobCouponCustomerSceneCountService) NewRequest() spi.InvocationRequest[PaasWeimobCouponCustomerSceneCountRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCouponCustomerSceneCountRequest]{
		Params: &PaasWeimobCouponCustomerSceneCountRequest{},
	}
}

func (s PaasWeimobCouponCustomerSceneCountService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCouponCustomerSceneCountRequest],
) (
	spi.InvocationResponse[PaasWeimobCouponCustomerSceneCountInt64],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCouponCustomerSceneCountRequest]))
}

type PaasWeimobCouponCustomerSceneCountRequest struct {
	SubScene   int64   `json:"subScene,omitempty"`
	SubSceneId string  `json:"subSceneId,omitempty"`
	Status     []int64 `json:"status,omitempty"`
}

type PaasWeimobCouponCustomerSceneCountInt64 int64

func CreatePaasWeimobCouponCustomerSceneCountInt64() *PaasWeimobCouponCustomerSceneCountInt64 {
	var br PaasWeimobCouponCustomerSceneCountInt64
	return &br
}

// OnPaasWeimobCouponCustomerSceneCountServiceInvocation
// @id 1088
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1088?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询券用户数)
func (s *Service) OnPaasWeimobCouponCustomerSceneCountServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCouponCustomerSceneCountRequest, PaasWeimobCouponCustomerSceneCountInt64],
) (service *Service) {
	var invocation = &PaasWeimobCouponCustomerSceneCountService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCouponCustomerSceneCountRequest, PaasWeimobCouponCustomerSceneCountInt64](invocation),
		"PaasWeimobCouponCustomerSceneCountService",
		"weimob.coupon.customer.scene.count",
	)
	return s
}
