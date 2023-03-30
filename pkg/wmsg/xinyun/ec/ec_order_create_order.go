package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcOrderCreateOrderListener struct {
	handler msg.XinyunEventHandler[EcOrderCreateOrderEvent]
}

func (s EcOrderCreateOrderListener) NewMessage() msg.MsgRequest[EcOrderCreateOrderEvent] {
	return &msg.XinyunOpenMessage[EcOrderCreateOrderEvent]{
		MsgBody: &EcOrderCreateOrderEvent{},
	}
}

func (s EcOrderCreateOrderListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcOrderCreateOrderEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcOrderCreateOrderEvent]))
}

type EcOrderCreateOrderEvent struct {
	OrderNo     int64 `json:"orderNo,omitempty"`
	CreateTime  int   `json:"createTime,omitempty"`
	ChannelType int   `json:"channelType,omitempty"`
}

// OnEcOrderCreateOrderEvent
// @id 371
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=订单新增
func (s *Service) OnEcOrderCreateOrderEvent(handler msg.XinyunEventHandler[EcOrderCreateOrderEvent]) (service *Service) {
	var listener = &EcOrderCreateOrderListener{handler}
	s.Register(msg.WrapListener[EcOrderCreateOrderEvent](listener), "ec_order", "createOrder")
	return s
}
