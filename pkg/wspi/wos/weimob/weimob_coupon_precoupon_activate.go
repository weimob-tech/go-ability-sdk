package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCouponPrecouponActivateService struct {
	handler spi.WosInvocationHandler[PaasWeimobCouponPrecouponActivateRequest, PaasWeimobCouponPrecouponActivateResponse]
}

func (s PaasWeimobCouponPrecouponActivateService) NewRequest() spi.InvocationRequest[PaasWeimobCouponPrecouponActivateRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCouponPrecouponActivateRequest]{
		Params: &PaasWeimobCouponPrecouponActivateRequest{},
	}
}

func (s PaasWeimobCouponPrecouponActivateService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCouponPrecouponActivateRequest],
) (
	spi.InvocationResponse[PaasWeimobCouponPrecouponActivateResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCouponPrecouponActivateRequest]))
}

type PaasWeimobCouponPrecouponActivateRequest struct {
	RevealCouponList []PaasWeimobCouponPrecouponActivateRequestRevealCouponList `json:"revealCouponList,omitempty"`
	SourceObjectList []PaasWeimobCouponPrecouponActivateRequestSourceObjectList `json:"sourceObjectList,omitempty"`
	ApplyNo          string                                                     `json:"applyNo,omitempty"`
	Wid              int64                                                      `json:"wid,omitempty"`
}
type PaasWeimobCouponPrecouponActivateRequestRevealCouponList struct {
	Code             string `json:"code,omitempty"`
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
}
type PaasWeimobCouponPrecouponActivateRequestSourceObjectList struct {
	Source       int64  `json:"source,omitempty"`
	SourceOpenId string `json:"sourceOpenId,omitempty"`
	SourceAppId  string `json:"sourceAppId,omitempty"`
}

type PaasWeimobCouponPrecouponActivateResponse struct {
	CouponFailList    []PaasWeimobCouponPrecouponActivateResponseCouponFailList    `json:"couponFailList,omitempty"`
	CouponSuccessList []PaasWeimobCouponPrecouponActivateResponseCouponSuccessList `json:"couponSuccessList,omitempty"`
	FailCodes         []int64                                                      `json:"failCodes,omitempty"`
	SuccessCount      int64                                                        `json:"successCount,omitempty"`
	FailedCount       int64                                                        `json:"failedCount,omitempty"`
}
type PaasWeimobCouponPrecouponActivateResponseCouponFailList struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	Code             string `json:"code,omitempty"`
	ErrCode          string `json:"errCode,omitempty"`
	ErrMsg           string `json:"errMsg,omitempty"`
}
type PaasWeimobCouponPrecouponActivateResponseCouponSuccessList struct {
	Code             string `json:"code,omitempty"`
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
}

func CreatePaasWeimobCouponPrecouponActivateResponse() *PaasWeimobCouponPrecouponActivateResponse {
	return &PaasWeimobCouponPrecouponActivateResponse{}
}

// OnPaasWeimobCouponPrecouponActivateServiceInvocation
// @id 1082
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1082?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=预发券领取)
func (s *Service) OnPaasWeimobCouponPrecouponActivateServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCouponPrecouponActivateRequest, PaasWeimobCouponPrecouponActivateResponse],
) (service *Service) {
	var invocation = &PaasWeimobCouponPrecouponActivateService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCouponPrecouponActivateRequest, PaasWeimobCouponPrecouponActivateResponse](invocation),
		"PaasWeimobCouponPrecouponActivateService",
		"weimob.coupon.precoupon.activate",
	)
	return s
}
