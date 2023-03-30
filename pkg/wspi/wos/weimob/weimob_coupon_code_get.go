package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCouponCodeGetService struct {
	handler spi.WosInvocationHandler[PaasWeimobCouponCodeGetRequest, PaasWeimobCouponCodeGetResponse]
}

func (s PaasWeimobCouponCodeGetService) NewRequest() spi.InvocationRequest[PaasWeimobCouponCodeGetRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCouponCodeGetRequest]{
		Params: &PaasWeimobCouponCodeGetRequest{},
	}
}

func (s PaasWeimobCouponCodeGetService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCouponCodeGetRequest],
) (
	spi.InvocationResponse[PaasWeimobCouponCodeGetResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCouponCodeGetRequest]))
}

type PaasWeimobCouponCodeGetRequest struct {
	CouponList []PaasWeimobCouponCodeGetRequestCouponList `json:"couponList,omitempty"`
	RequestId  string                                     `json:"requestId,omitempty"`
	Channel    int64                                      `json:"channel,omitempty"`
	Scene      int64                                      `json:"scene,omitempty"`
	SceneId    string                                     `json:"sceneId,omitempty"`
}
type PaasWeimobCouponCodeGetRequestCouponList struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	Num              int64  `json:"num,omitempty"`
	ThirdId          string `json:"thirdId,omitempty"`
	RepoId           int64  `json:"repoId,omitempty"`
}

type PaasWeimobCouponCodeGetResponse struct {
	TemplateCodes PaasWeimobCouponCodeGetResponseTemplateCodes `json:"templateCodes,omitempty"`
	Wid           int64                                        `json:"wid,omitempty"`
}
type PaasWeimobCouponCodeGetResponseTemplateCodes struct {
}

func CreatePaasWeimobCouponCodeGetResponse() *PaasWeimobCouponCodeGetResponse {
	return &PaasWeimobCouponCodeGetResponse{}
}

// OnPaasWeimobCouponCodeGetServiceInvocation
// @id 1073
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1073?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=生成外部优惠券码)
func (s *Service) OnPaasWeimobCouponCodeGetServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCouponCodeGetRequest, PaasWeimobCouponCodeGetResponse],
) (service *Service) {
	var invocation = &PaasWeimobCouponCodeGetService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCouponCodeGetRequest, PaasWeimobCouponCodeGetResponse](invocation),
		"PaasWeimobCouponCodeGetService",
		"weimob.coupon.code.get",
	)
	return s
}
