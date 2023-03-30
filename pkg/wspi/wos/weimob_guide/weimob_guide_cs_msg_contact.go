package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideCsMsgContactService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideCsMsgContactRequest, PaasWeimobGuideCsMsgContactResponse]
}

func (s PaasWeimobGuideCsMsgContactService) NewRequest() spi.InvocationRequest[PaasWeimobGuideCsMsgContactRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideCsMsgContactRequest]{
		Params: &PaasWeimobGuideCsMsgContactRequest{},
	}
}

func (s PaasWeimobGuideCsMsgContactService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideCsMsgContactRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideCsMsgContactResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideCsMsgContactRequest]))
}

type PaasWeimobGuideCsMsgContactRequest struct {
	CustomerWid       int64 `json:"customerWid,omitempty"`
	GuiderWid         int64 `json:"guiderWid,omitempty"`
	Terminal          int64 `json:"terminal,omitempty"`
	UserSourceChannel int64 `json:"userSourceChannel,omitempty"`
}

type PaasWeimobGuideCsMsgContactResponse struct {
	DialogId int64 `json:"dialogId,omitempty"`
}

func CreatePaasWeimobGuideCsMsgContactResponse() *PaasWeimobGuideCsMsgContactResponse {
	return &PaasWeimobGuideCsMsgContactResponse{}
}

// OnPaasWeimobGuideCsMsgContactServiceInvocation
// @id 908
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/908?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=客服联系客户)
func (s *Service) OnPaasWeimobGuideCsMsgContactServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideCsMsgContactRequest, PaasWeimobGuideCsMsgContactResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideCsMsgContactService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideCsMsgContactRequest, PaasWeimobGuideCsMsgContactResponse](invocation),
		"PaasWeimobGuideCsMsgContactService",
		"weimobGuide.cs.msg.contact",
	)
	return s
}
