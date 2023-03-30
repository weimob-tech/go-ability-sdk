package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcOrderOrderTransferListener struct {
	handler msg.XinyunEventHandler[EcOrderOrderTransferEvent]
}

func (s EcOrderOrderTransferListener) NewMessage() msg.MsgRequest[EcOrderOrderTransferEvent] {
	return &msg.XinyunOpenMessage[EcOrderOrderTransferEvent]{
		MsgBody: &EcOrderOrderTransferEvent{},
	}
}

func (s EcOrderOrderTransferListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcOrderOrderTransferEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcOrderOrderTransferEvent]))
}

type EcOrderOrderTransferEvent struct {
	OrderNo int64 `json:"orderNo,omitempty"`
}

// OnEcOrderOrderTransferEvent
// @id 1174
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=订单转单消息（智慧零售）
func (s *Service) OnEcOrderOrderTransferEvent(handler msg.XinyunEventHandler[EcOrderOrderTransferEvent]) (service *Service) {
	var listener = &EcOrderOrderTransferListener{handler}
	s.Register(msg.WrapListener[EcOrderOrderTransferEvent](listener), "ec_order", "orderTransfer")
	return s
}
