package xinyun

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasAcquireCouponCodeService struct {
	handler spi.XinyunInvocationHandler[PaasAcquireCouponCodeRequest, PaasAcquireCouponCodeResponse]
}

func (s PaasAcquireCouponCodeService) NewRequest() spi.InvocationRequest[PaasAcquireCouponCodeRequest] {
	return &spi.XinyunInvocationRequest[PaasAcquireCouponCodeRequest]{
		Params: &PaasAcquireCouponCodeRequest{},
	}
}

func (s PaasAcquireCouponCodeService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasAcquireCouponCodeRequest],
) (
	spi.InvocationResponse[PaasAcquireCouponCodeResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.XinyunInvocationRequest[PaasAcquireCouponCodeRequest]))
}

type PaasAcquireCouponCodeRequest struct {
	SourceObjectList    []PaasAcquireCouponCodeRequestSourceObjectList  `json:"sourceObjectList,omitempty"`
	ReceiveCouponList   []PaasAcquireCouponCodeRequestReceiveCouponList `json:"receiveCouponList,omitempty"`
	Pid                 int64                                           `json:"pid,omitempty"`
	Wid                 int64                                           `json:"wid,omitempty"`
	CreateUserOnNoExist bool                                            `json:"createUserOnNoExist,omitempty"`
	AcquireChannel      int64                                           `json:"acquireChannel,omitempty"`
	AcquireChannelId    string                                          `json:"acquireChannelId,omitempty"`
	StoreId             int64                                           `json:"storeId,omitempty"`
	ThirdId             string                                          `json:"thirdId,omitempty"`
	RequestId           string                                          `json:"requestId,omitempty"`
	Operator            string                                          `json:"operator,omitempty"`
	GlobalTicket        string                                          `json:"globalTicket,omitempty"`
}
type PaasAcquireCouponCodeRequestSourceObjectList struct {
	SourceOpenId string `json:"sourceOpenId,omitempty"`
	SourceAppId  string `json:"sourceAppId,omitempty"`
	Source       int64  `json:"source,omitempty"`
}
type PaasAcquireCouponCodeRequestReceiveCouponList struct {
	CouponTemplateId int64 `json:"couponTemplateId,omitempty"`
	Number           int64 `json:"number,omitempty"`
	RepoId           int64 `json:"repoId,omitempty"`
}

type PaasAcquireCouponCodeResponse struct {
	CouponList []PaasAcquireCouponCodeResponseCouponList `json:"couponList,omitempty"`
	HasCode    bool                                      `json:"hasCode,omitempty"`
}
type PaasAcquireCouponCodeResponseCouponList struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	Code             string `json:"code,omitempty"`
	Wid              int64  `json:"wid,omitempty"`
	ErrorCode        int64  `json:"errorCode,omitempty"`
	ErrMsg           string `json:"errMsg,omitempty"`
}

func CreatePaasAcquireCouponCodeResponse() *PaasAcquireCouponCodeResponse {
	return &PaasAcquireCouponCodeResponse{}
}

// OnPaasAcquireCouponCodeServiceInvocation
// @id 3841
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=获取外部优惠券码)
func (s *Service) OnPaasAcquireCouponCodeServiceInvocation(
	handler spi.XinyunInvocationHandler[PaasAcquireCouponCodeRequest, PaasAcquireCouponCodeResponse],
) (service *Service) {
	var invocation = &PaasAcquireCouponCodeService{handler}
	s.Register(
		spi.WrapInvoker[PaasAcquireCouponCodeRequest, PaasAcquireCouponCodeResponse](invocation),
		"PaasAcquireCouponCodeService",
		"sAcquireCouponCode",
	)
	return s
}
