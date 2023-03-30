package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcOrderUpdateOrderAmountListener struct {
	handler msg.XinyunEventHandler[EcOrderUpdateOrderAmountEvent]
}

func (s EcOrderUpdateOrderAmountListener) NewMessage() msg.MsgRequest[EcOrderUpdateOrderAmountEvent] {
	return &msg.XinyunOpenMessage[EcOrderUpdateOrderAmountEvent]{
		MsgBody: &EcOrderUpdateOrderAmountEvent{},
	}
}

func (s EcOrderUpdateOrderAmountListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcOrderUpdateOrderAmountEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcOrderUpdateOrderAmountEvent]))
}

type EcOrderUpdateOrderAmountEvent struct {
	OrderNo int64 `json:"orderNo,omitempty"`
}

// OnEcOrderUpdateOrderAmountEvent
// @id 1444
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=订单改价消息
func (s *Service) OnEcOrderUpdateOrderAmountEvent(handler msg.XinyunEventHandler[EcOrderUpdateOrderAmountEvent]) (service *Service) {
	var listener = &EcOrderUpdateOrderAmountListener{handler}
	s.Register(msg.WrapListener[EcOrderUpdateOrderAmountEvent](listener), "ec_order", "updateOrderAmount")
	return s
}
