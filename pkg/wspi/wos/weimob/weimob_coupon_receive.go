package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCouponReceiveService struct {
	handler spi.WosInvocationHandler[PaasWeimobCouponReceiveRequest, PaasWeimobCouponReceiveResponse]
}

func (s PaasWeimobCouponReceiveService) NewRequest() spi.InvocationRequest[PaasWeimobCouponReceiveRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCouponReceiveRequest]{
		Params: &PaasWeimobCouponReceiveRequest{},
	}
}

func (s PaasWeimobCouponReceiveService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCouponReceiveRequest],
) (
	spi.InvocationResponse[PaasWeimobCouponReceiveResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCouponReceiveRequest]))
}

type PaasWeimobCouponReceiveRequest struct {
	CouponNums       []PaasWeimobCouponReceiveRequestCouponNums       `json:"couponNums,omitempty"`
	SourceObjectList []PaasWeimobCouponReceiveRequestSourceObjectList `json:"sourceObjectList,omitempty"`
	Wid              int64                                            `json:"wid,omitempty"`
	Channel          int64                                            `json:"channel,omitempty"`
	Scene            int64                                            `json:"scene,omitempty"`
	SceneId          string                                           `json:"sceneId,omitempty"`
	Operator         string                                           `json:"operator,omitempty"`
	IsPreReceive     bool                                             `json:"isPreReceive,omitempty"`
	NeedLoginCheck   bool                                             `json:"needLoginCheck,omitempty"`
	RequestId        string                                           `json:"requestId,omitempty"`
}
type PaasWeimobCouponReceiveRequestCouponNums struct {
	CouponTemplateId int64 `json:"couponTemplateId,omitempty"`
	Num              int64 `json:"num,omitempty"`
	RepoId           int64 `json:"repoId,omitempty"`
}
type PaasWeimobCouponReceiveRequestSourceObjectList struct {
	Source       int64  `json:"source,omitempty"`
	SourceOpenId string `json:"sourceOpenId,omitempty"`
	SourceAppId  string `json:"sourceAppId,omitempty"`
}

type PaasWeimobCouponReceiveResponse struct {
	CouponResultList []PaasWeimobCouponReceiveResponseCouponResultList `json:"couponResultList,omitempty"`
	SuccessCount     int64                                             `json:"successCount,omitempty"`
	FailedCount      int64                                             `json:"failedCount,omitempty"`
	Status           int64                                             `json:"status,omitempty"`
}
type PaasWeimobCouponReceiveResponseCouponResultList struct {
	CouponTemplateId int64   `json:"couponTemplateId,omitempty"`
	Name             string  `json:"name,omitempty"`
	Wid              int64   `json:"wid,omitempty"`
	IsSuccess        bool    `json:"isSuccess,omitempty"`
	FailureNum       int64   `json:"failureNum,omitempty"`
	ErrCode          string  `json:"errCode,omitempty"`
	ErrMsg           string  `json:"errMsg,omitempty"`
	Codes            []int64 `json:"codes,omitempty"`
	ExpDate          string  `json:"expDate,omitempty"`
}

func CreatePaasWeimobCouponReceiveResponse() *PaasWeimobCouponReceiveResponse {
	return &PaasWeimobCouponReceiveResponse{}
}

// OnPaasWeimobCouponReceiveServiceInvocation
// @id 1072
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1072?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=领取优惠券)
func (s *Service) OnPaasWeimobCouponReceiveServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCouponReceiveRequest, PaasWeimobCouponReceiveResponse],
) (service *Service) {
	var invocation = &PaasWeimobCouponReceiveService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCouponReceiveRequest, PaasWeimobCouponReceiveResponse](invocation),
		"PaasWeimobCouponReceiveService",
		"weimob.coupon.receive",
	)
	return s
}
