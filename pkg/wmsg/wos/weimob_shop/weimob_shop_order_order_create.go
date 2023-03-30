package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopOrderOrderCreateListener struct {
	handler msg.WosEventHandler[WeimobShopOrderOrderCreateEvent]
}

func (s WeimobShopOrderOrderCreateListener) NewMessage() msg.MsgRequest[WeimobShopOrderOrderCreateEvent] {
	return &msg.WosOpenMessage[WeimobShopOrderOrderCreateEvent]{
		MsgBody: &WeimobShopOrderOrderCreateEvent{},
	}
}

func (s WeimobShopOrderOrderCreateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopOrderOrderCreateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopOrderOrderCreateEvent]))
}

type WeimobShopOrderOrderCreateEvent struct {
}

// OnWeimobShopOrderOrderCreateEvent
// @id 1217
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1217?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=订单新增消息)
func (s *Service) OnWeimobShopOrderOrderCreateEvent(handler msg.WosEventHandler[WeimobShopOrderOrderCreateEvent]) (service *Service) {
	var listener = &WeimobShopOrderOrderCreateListener{handler}
	s.Register(msg.WrapListener[WeimobShopOrderOrderCreateEvent](listener), "weimob_shop.order", "orderCreate")
	return s
}
