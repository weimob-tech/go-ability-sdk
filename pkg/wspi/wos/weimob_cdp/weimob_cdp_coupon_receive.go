package weimob_cdp

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCdpCouponReceiveService struct {
	handler spi.WosInvocationHandler[PaasWeimobCdpCouponReceiveRequest, PaasWeimobCdpCouponReceiveResponse]
}

func (s PaasWeimobCdpCouponReceiveService) NewRequest() spi.InvocationRequest[PaasWeimobCdpCouponReceiveRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCdpCouponReceiveRequest]{
		Params: &PaasWeimobCdpCouponReceiveRequest{},
	}
}

func (s PaasWeimobCdpCouponReceiveService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCdpCouponReceiveRequest],
) (
	spi.InvocationResponse[PaasWeimobCdpCouponReceiveResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCdpCouponReceiveRequest]))
}

type PaasWeimobCdpCouponReceiveRequest struct {
	CouponNums        []PaasWeimobCdpCouponReceiveRequestCouponNums `json:"couponNums,omitempty"`
	Wid               int64                                         `json:"wid,omitempty"`
	Channel           int64                                         `json:"channel,omitempty"`
	Scene             int64                                         `json:"scene,omitempty"`
	SceneId           int64                                         `json:"sceneId,omitempty"`
	IsPreReceive      bool                                          `json:"isPreReceive,omitempty"`
	BosId             int64                                         `json:"bosId,omitempty"`
	ProductInstanceId int64                                         `json:"productInstanceId,omitempty"`
}
type PaasWeimobCdpCouponReceiveRequestCouponNums struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	Num              int64  `json:"num,omitempty"`
	RequestId        string `json:"requestId,omitempty"`
}

type PaasWeimobCdpCouponReceiveResponse struct {
	CouponResultList []PaasWeimobCdpCouponReceiveResponseCouponResultList `json:"couponResultList,omitempty"`
	SuccessCount     int64                                                `json:"successCount,omitempty"`
	FailedCount      int64                                                `json:"failedCount,omitempty"`
	Status           int64                                                `json:"status,omitempty"`
}
type PaasWeimobCdpCouponReceiveResponseCouponResultList struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	Name             string `json:"name,omitempty"`
	Wid              int64  `json:"wid,omitempty"`
	IsSuccess        bool   `json:"isSuccess,omitempty"`
	FailureNum       int64  `json:"failureNum,omitempty"`
	ErrCode          int64  `json:"errCode,omitempty"`
	ErrMsg           string `json:"errMsg,omitempty"`
	CanReceive       bool   `json:"canReceive,omitempty"`
}

func CreatePaasWeimobCdpCouponReceiveResponse() *PaasWeimobCdpCouponReceiveResponse {
	return &PaasWeimobCdpCouponReceiveResponse{}
}

// OnPaasWeimobCdpCouponReceiveServiceInvocation
// @id 887
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/887?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=优惠券（治理))
func (s *Service) OnPaasWeimobCdpCouponReceiveServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCdpCouponReceiveRequest, PaasWeimobCdpCouponReceiveResponse],
) (service *Service) {
	var invocation = &PaasWeimobCdpCouponReceiveService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCdpCouponReceiveRequest, PaasWeimobCdpCouponReceiveResponse](invocation),
		"PaasWeimobCdpCouponReceiveService",
		"weimobCdp.coupon.receive",
	)
	return s
}
