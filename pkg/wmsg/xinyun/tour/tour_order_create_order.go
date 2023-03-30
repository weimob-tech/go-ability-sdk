package tour

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type TourOrderCreateOrderListener struct {
	handler msg.XinyunEventHandler[TourOrderCreateOrderEvent]
}

func (s TourOrderCreateOrderListener) NewMessage() msg.MsgRequest[TourOrderCreateOrderEvent] {
	return &msg.XinyunOpenMessage[TourOrderCreateOrderEvent]{
		MsgBody: &TourOrderCreateOrderEvent{},
	}
}

func (s TourOrderCreateOrderListener) OnMessage(ctx context.Context, message msg.MsgRequest[TourOrderCreateOrderEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[TourOrderCreateOrderEvent]))
}

type TourOrderCreateOrderEvent struct {
	OrderNo string `json:"orderNo,omitempty"`
}

// OnTourOrderCreateOrderEvent
// @id 1037
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=新增订单
func (s *Service) OnTourOrderCreateOrderEvent(handler msg.XinyunEventHandler[TourOrderCreateOrderEvent]) (service *Service) {
	var listener = &TourOrderCreateOrderListener{handler}
	s.Register(msg.WrapListener[TourOrderCreateOrderEvent](listener), "tour_order", "createOrder")
	return s
}
