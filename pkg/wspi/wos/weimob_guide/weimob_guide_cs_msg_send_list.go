package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideCsMsgSendListService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideCsMsgSendListRequest, PaasWeimobGuideCsMsgSendListResponse]
}

func (s PaasWeimobGuideCsMsgSendListService) NewRequest() spi.InvocationRequest[PaasWeimobGuideCsMsgSendListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideCsMsgSendListRequest]{
		Params: &PaasWeimobGuideCsMsgSendListRequest{},
	}
}

func (s PaasWeimobGuideCsMsgSendListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideCsMsgSendListRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideCsMsgSendListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideCsMsgSendListRequest]))
}

type PaasWeimobGuideCsMsgSendListRequest struct {
	CustomerList     []PaasWeimobGuideCsMsgSendListRequestCustomerList `json:"customerList,omitempty"`
	ServicerId       int64                                             `json:"servicerId,omitempty"`
	MsgType          string                                            `json:"msgType,omitempty"`
	Content          string                                            `json:"content,omitempty"`
	BusinessId       int64                                             `json:"businessId,omitempty"`
	MsgSourceChannel int64                                             `json:"msgSourceChannel,omitempty"`
}
type PaasWeimobGuideCsMsgSendListRequestCustomerList struct {
	CustomerId        int64 `json:"customerId,omitempty"`
	Vid               int64 `json:"vid,omitempty"`
	UserSourceChannel int64 `json:"userSourceChannel,omitempty"`
}

type PaasWeimobGuideCsMsgSendListResponse struct {
	FailList     []PaasWeimobGuideCsMsgSendListResponseFailList `json:"failList,omitempty"`
	SuccessCount int64                                          `json:"successCount,omitempty"`
}
type PaasWeimobGuideCsMsgSendListResponseFailList struct {
	UserId     int64  `json:"userId,omitempty"`
	FailReason string `json:"failReason,omitempty"`
}

func CreatePaasWeimobGuideCsMsgSendListResponse() *PaasWeimobGuideCsMsgSendListResponse {
	return &PaasWeimobGuideCsMsgSendListResponse{}
}

// OnPaasWeimobGuideCsMsgSendListServiceInvocation
// @id 732
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/732?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=发送客服消息)
func (s *Service) OnPaasWeimobGuideCsMsgSendListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideCsMsgSendListRequest, PaasWeimobGuideCsMsgSendListResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideCsMsgSendListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideCsMsgSendListRequest, PaasWeimobGuideCsMsgSendListResponse](invocation),
		"PaasWeimobGuideCsMsgSendListService",
		"weimobGuide.cs.msg.sendList",
	)
	return s
}
