package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type MallOrderMallOrderMessageListener struct {
	handler msg.WosEventHandler[MallOrderMallOrderMessageEvent]
}

func (s MallOrderMallOrderMessageListener) NewMessage() msg.MsgRequest[MallOrderMallOrderMessageEvent] {
	return &msg.WosOpenMessage[MallOrderMallOrderMessageEvent]{
		MsgBody: &MallOrderMallOrderMessageEvent{},
	}
}

func (s MallOrderMallOrderMessageListener) OnMessage(ctx context.Context, message msg.MsgRequest[MallOrderMallOrderMessageEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[MallOrderMallOrderMessageEvent]))
}

type MallOrderMallOrderMessageEvent struct {
}

// OnMallOrderMallOrderMessageEvent
// @id 1186
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1186?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=产品融合订单消息)
func (s *Service) OnMallOrderMallOrderMessageEvent(handler msg.WosEventHandler[MallOrderMallOrderMessageEvent]) (service *Service) {
	var listener = &MallOrderMallOrderMessageListener{handler}
	s.Register(msg.WrapListener[MallOrderMallOrderMessageEvent](listener), "mall_order", "mall_order_message")
	return s
}
