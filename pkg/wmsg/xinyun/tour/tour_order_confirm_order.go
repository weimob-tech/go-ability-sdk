package tour

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type TourOrderConfirmOrderListener struct {
	handler msg.XinyunEventHandler[TourOrderConfirmOrderEvent]
}

func (s TourOrderConfirmOrderListener) NewMessage() msg.MsgRequest[TourOrderConfirmOrderEvent] {
	return &msg.XinyunOpenMessage[TourOrderConfirmOrderEvent]{
		MsgBody: &TourOrderConfirmOrderEvent{},
	}
}

func (s TourOrderConfirmOrderListener) OnMessage(ctx context.Context, message msg.MsgRequest[TourOrderConfirmOrderEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[TourOrderConfirmOrderEvent]))
}

type TourOrderConfirmOrderEvent struct {
	OrderNo string `json:"orderNo,omitempty"`
}

// OnTourOrderConfirmOrderEvent
// @id 1035
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=订单确认
func (s *Service) OnTourOrderConfirmOrderEvent(handler msg.XinyunEventHandler[TourOrderConfirmOrderEvent]) (service *Service) {
	var listener = &TourOrderConfirmOrderListener{handler}
	s.Register(msg.WrapListener[TourOrderConfirmOrderEvent](listener), "tour_order", "confirmOrder")
	return s
}
