package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCouponCustomerCountService struct {
	handler spi.WosInvocationHandler[PaasWeimobCouponCustomerCountRequest, PaasWeimobCouponCustomerCountResponse]
}

func (s PaasWeimobCouponCustomerCountService) NewRequest() spi.InvocationRequest[PaasWeimobCouponCustomerCountRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCouponCustomerCountRequest]{
		Params: &PaasWeimobCouponCustomerCountRequest{},
	}
}

func (s PaasWeimobCouponCustomerCountService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCouponCustomerCountRequest],
) (
	spi.InvocationResponse[PaasWeimobCouponCustomerCountResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCouponCustomerCountRequest]))
}

type PaasWeimobCouponCustomerCountRequest struct {
	SourceObjectList []PaasWeimobCouponCustomerCountRequestSourceObjectList `json:"sourceObjectList,omitempty"`
	Wid              int64                                                  `json:"wid,omitempty"`
}
type PaasWeimobCouponCustomerCountRequestSourceObjectList struct {
	Source       int64  `json:"source,omitempty"`
	SourceOpenId string `json:"sourceOpenId,omitempty"`
	SourceAppId  string `json:"sourceAppId,omitempty"`
}

type PaasWeimobCouponCustomerCountResponse struct {
	Count      int64 `json:"count,omitempty"`
	UsedNum    int64 `json:"usedNum,omitempty"`
	ExpiredNum int64 `json:"expiredNum,omitempty"`
}

func CreatePaasWeimobCouponCustomerCountResponse() *PaasWeimobCouponCustomerCountResponse {
	return &PaasWeimobCouponCustomerCountResponse{}
}

// OnPaasWeimobCouponCustomerCountServiceInvocation
// @id 1074
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1074?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=统计优惠券数量)
func (s *Service) OnPaasWeimobCouponCustomerCountServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCouponCustomerCountRequest, PaasWeimobCouponCustomerCountResponse],
) (service *Service) {
	var invocation = &PaasWeimobCouponCustomerCountService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCouponCustomerCountRequest, PaasWeimobCouponCustomerCountResponse](invocation),
		"PaasWeimobCouponCustomerCountService",
		"weimob.coupon.customer.count",
	)
	return s
}
