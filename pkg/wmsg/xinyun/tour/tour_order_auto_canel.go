package tour

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type TourOrderAutoCanelListener struct {
	handler msg.XinyunEventHandler[TourOrderAutoCanelEvent]
}

func (s TourOrderAutoCanelListener) NewMessage() msg.MsgRequest[TourOrderAutoCanelEvent] {
	return &msg.XinyunOpenMessage[TourOrderAutoCanelEvent]{
		MsgBody: &TourOrderAutoCanelEvent{},
	}
}

func (s TourOrderAutoCanelListener) OnMessage(ctx context.Context, message msg.MsgRequest[TourOrderAutoCanelEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[TourOrderAutoCanelEvent]))
}

type TourOrderAutoCanelEvent struct {
	OrderNo string `json:"orderNo,omitempty"`
}

// OnTourOrderAutoCanelEvent
// @id 1034
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=订单自动取消
func (s *Service) OnTourOrderAutoCanelEvent(handler msg.XinyunEventHandler[TourOrderAutoCanelEvent]) (service *Service) {
	var listener = &TourOrderAutoCanelListener{handler}
	s.Register(msg.WrapListener[TourOrderAutoCanelEvent](listener), "tour_order", "autoCanel")
	return s
}
