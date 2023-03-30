package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopOrderOrderStatusUpdateListener struct {
	handler msg.WosEventHandler[WeimobShopOrderOrderStatusUpdateEvent]
}

func (s WeimobShopOrderOrderStatusUpdateListener) NewMessage() msg.MsgRequest[WeimobShopOrderOrderStatusUpdateEvent] {
	return &msg.WosOpenMessage[WeimobShopOrderOrderStatusUpdateEvent]{
		MsgBody: &WeimobShopOrderOrderStatusUpdateEvent{},
	}
}

func (s WeimobShopOrderOrderStatusUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopOrderOrderStatusUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopOrderOrderStatusUpdateEvent]))
}

type WeimobShopOrderOrderStatusUpdateEvent struct {
}

// OnWeimobShopOrderOrderStatusUpdateEvent
// @id 1216
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1216?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=订单状态变更消息)
func (s *Service) OnWeimobShopOrderOrderStatusUpdateEvent(handler msg.WosEventHandler[WeimobShopOrderOrderStatusUpdateEvent]) (service *Service) {
	var listener = &WeimobShopOrderOrderStatusUpdateListener{handler}
	s.Register(msg.WrapListener[WeimobShopOrderOrderStatusUpdateEvent](listener), "weimob_shop.order", "orderStatusUpdate")
	return s
}
