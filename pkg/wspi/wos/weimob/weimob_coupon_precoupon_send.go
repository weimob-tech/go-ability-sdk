package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCouponPrecouponSendService struct {
	handler spi.WosInvocationHandler[PaasWeimobCouponPrecouponSendRequest, PaasWeimobCouponPrecouponSendResponse]
}

func (s PaasWeimobCouponPrecouponSendService) NewRequest() spi.InvocationRequest[PaasWeimobCouponPrecouponSendRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCouponPrecouponSendRequest]{
		Params: &PaasWeimobCouponPrecouponSendRequest{},
	}
}

func (s PaasWeimobCouponPrecouponSendService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCouponPrecouponSendRequest],
) (
	spi.InvocationResponse[PaasWeimobCouponPrecouponSendResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCouponPrecouponSendRequest]))
}

type PaasWeimobCouponPrecouponSendRequest struct {
	CouponNums       []PaasWeimobCouponPrecouponSendRequestCouponNums       `json:"couponNums,omitempty"`
	SourceObjectList []PaasWeimobCouponPrecouponSendRequestSourceObjectList `json:"sourceObjectList,omitempty"`
	Wid              int64                                                  `json:"wid,omitempty"`
	Scene            int64                                                  `json:"scene,omitempty"`
	SceneId          string                                                 `json:"sceneId,omitempty"`
}
type PaasWeimobCouponPrecouponSendRequestCouponNums struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	Num              int64  `json:"num,omitempty"`
	RequestId        string `json:"requestId,omitempty"`
}
type PaasWeimobCouponPrecouponSendRequestSourceObjectList struct {
	Source       int64  `json:"source,omitempty"`
	SourceOpenId string `json:"sourceOpenId,omitempty"`
	SourceAppId  string `json:"sourceAppId,omitempty"`
}

type PaasWeimobCouponPrecouponSendResponse struct {
	Result  bool    `json:"result,omitempty"`
	State   int64   `json:"state,omitempty"`
	Codes   []int64 `json:"codes,omitempty"`
	ApplyNo string  `json:"applyNo,omitempty"`
}

func CreatePaasWeimobCouponPrecouponSendResponse() *PaasWeimobCouponPrecouponSendResponse {
	return &PaasWeimobCouponPrecouponSendResponse{}
}

// OnPaasWeimobCouponPrecouponSendServiceInvocation
// @id 1081
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1081?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=预发券)
func (s *Service) OnPaasWeimobCouponPrecouponSendServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCouponPrecouponSendRequest, PaasWeimobCouponPrecouponSendResponse],
) (service *Service) {
	var invocation = &PaasWeimobCouponPrecouponSendService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCouponPrecouponSendRequest, PaasWeimobCouponPrecouponSendResponse](invocation),
		"PaasWeimobCouponPrecouponSendService",
		"weimob.coupon.precoupon.send",
	)
	return s
}
