package wangpu

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WPOrderStatusChangeListener struct {
	handler msg.XinyunEventHandler[WPOrderStatusChangeEvent]
}

func (s WPOrderStatusChangeListener) NewMessage() msg.MsgRequest[WPOrderStatusChangeEvent] {
	return &msg.XinyunOpenMessage[WPOrderStatusChangeEvent]{
		MsgBody: &WPOrderStatusChangeEvent{},
	}
}

func (s WPOrderStatusChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[WPOrderStatusChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[WPOrderStatusChangeEvent]))
}

type WPOrderStatusChangeEvent struct {
	OrderNo string `json:"order_no,omitempty"`
}

// OnWPOrderStatusChangeEvent
// @id 1294
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=订单状态变更
func (s *Service) OnWPOrderStatusChangeEvent(handler msg.XinyunEventHandler[WPOrderStatusChangeEvent]) (service *Service) {
	var listener = &WPOrderStatusChangeListener{handler}
	s.Register(msg.WrapListener[WPOrderStatusChangeEvent](listener), "WP_order", "statusChange")
	return s
}
