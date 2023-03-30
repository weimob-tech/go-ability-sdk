package tour

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type TourOrderUserCanelListener struct {
	handler msg.XinyunEventHandler[TourOrderUserCanelEvent]
}

func (s TourOrderUserCanelListener) NewMessage() msg.MsgRequest[TourOrderUserCanelEvent] {
	return &msg.XinyunOpenMessage[TourOrderUserCanelEvent]{
		MsgBody: &TourOrderUserCanelEvent{},
	}
}

func (s TourOrderUserCanelListener) OnMessage(ctx context.Context, message msg.MsgRequest[TourOrderUserCanelEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[TourOrderUserCanelEvent]))
}

type TourOrderUserCanelEvent struct {
	OrderNo string `json:"orderNo,omitempty"`
}

// OnTourOrderUserCanelEvent
// @id 1033
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=买家取消订单
func (s *Service) OnTourOrderUserCanelEvent(handler msg.XinyunEventHandler[TourOrderUserCanelEvent]) (service *Service) {
	var listener = &TourOrderUserCanelListener{handler}
	s.Register(msg.WrapListener[TourOrderUserCanelEvent](listener), "tour_order", "userCanel")
	return s
}
