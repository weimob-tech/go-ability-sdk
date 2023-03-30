package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmOrderCreateListener struct {
	handler msg.WosEventHandler[WeimobCrmOrderCreateEvent]
}

func (s WeimobCrmOrderCreateListener) NewMessage() msg.MsgRequest[WeimobCrmOrderCreateEvent] {
	return &msg.WosOpenMessage[WeimobCrmOrderCreateEvent]{
		MsgBody: &WeimobCrmOrderCreateEvent{},
	}
}

func (s WeimobCrmOrderCreateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmOrderCreateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmOrderCreateEvent]))
}

type WeimobCrmOrderCreateEvent struct {
	OrderNo     int64  `json:"orderNo,omitempty"`
	OutOrderNo  string `json:"outOrderNo,omitempty"`
	OrderStatus int64  `json:"orderStatus,omitempty"`
	Vid         int64  `json:"vid,omitempty"`
	Wid         int64  `json:"wid,omitempty"`
}

// OnWeimobCrmOrderCreateEvent
// @id 1418
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1418?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=订单创建消息)
func (s *Service) OnWeimobCrmOrderCreateEvent(handler msg.WosEventHandler[WeimobCrmOrderCreateEvent]) (service *Service) {
	var listener = &WeimobCrmOrderCreateListener{handler}
	s.Register(msg.WrapListener[WeimobCrmOrderCreateEvent](listener), "weimob_crm.order", "create")
	return s
}
