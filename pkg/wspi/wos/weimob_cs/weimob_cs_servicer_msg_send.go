package weimob_cs

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCsServicerMsgSendService struct {
	handler spi.WosInvocationHandler[PaasWeimobCsrMsgSendRequest, PaasWeimobCsrMsgSendResponse]
}

func (s PaasWeimobCsServicerMsgSendService) NewRequest() spi.InvocationRequest[PaasWeimobCsrMsgSendRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCsrMsgSendRequest]{
		Params: &PaasWeimobCsrMsgSendRequest{},
	}
}

func (s PaasWeimobCsServicerMsgSendService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCsrMsgSendRequest],
) (
	spi.InvocationResponse[PaasWeimobCsrMsgSendResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCsrMsgSendRequest]))
}

type PaasWeimobCsrMsgSendRequest struct {
	Ext                     PaasWeimobCsrMsgSendRequestExt `json:"ext,omitempty"`
	OriginProductInstanceId int64                          `json:"originProductInstanceId,omitempty"`
	MsgType                 string                         `json:"msgType,omitempty"`
	Content                 string                         `json:"content,omitempty"`
	Status                  int64                          `json:"status,omitempty"`
}
type PaasWeimobCsrMsgSendRequestExt struct {
	CustomerId int64  `json:"customerId,omitempty"`
	Channel    string `json:"channel,omitempty"`
	ClientId   string `json:"clientId,omitempty"`
	ClientName string `json:"clientName,omitempty"`
}

type PaasWeimobCsrMsgSendResponse struct {
	Ext     PaasWeimobCsrMsgSendResponseExt `json:"ext,omitempty"`
	MsgType string                          `json:"msgType,omitempty"`
	Content string                          `json:"content,omitempty"`
	Status  int64                           `json:"status,omitempty"`
}
type PaasWeimobCsrMsgSendResponseExt struct {
	CustomerId int64  `json:"customerId,omitempty"`
	Channel    string `json:"channel,omitempty"`
	ClientId   string `json:"clientId,omitempty"`
	ClientName string `json:"clientName,omitempty"`
}

func CreatePaasWeimobCsrMsgSendResponse() *PaasWeimobCsrMsgSendResponse {
	return &PaasWeimobCsrMsgSendResponse{}
}

// OnPaasWeimobCsServicerMsgSendServiceInvocation
// @id 852
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/852?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=推送消息和事件v2)
func (s *Service) OnPaasWeimobCsServicerMsgSendServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCsrMsgSendRequest, PaasWeimobCsrMsgSendResponse],
) (service *Service) {
	var invocation = &PaasWeimobCsServicerMsgSendService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCsrMsgSendRequest, PaasWeimobCsrMsgSendResponse](invocation),
		"PaasWeimobCsServicerMsgSendService",
		"weimobCs.servicer.msg.send",
	)
	return s
}
