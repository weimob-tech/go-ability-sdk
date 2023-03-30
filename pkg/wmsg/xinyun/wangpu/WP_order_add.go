package wangpu

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WPOrderAddListener struct {
	handler msg.XinyunEventHandler[WPOrderAddEvent]
}

func (s WPOrderAddListener) NewMessage() msg.MsgRequest[WPOrderAddEvent] {
	return &msg.XinyunOpenMessage[WPOrderAddEvent]{
		MsgBody: &WPOrderAddEvent{},
	}
}

func (s WPOrderAddListener) OnMessage(ctx context.Context, message msg.MsgRequest[WPOrderAddEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[WPOrderAddEvent]))
}

type WPOrderAddEvent struct {
	OrderNo     string `json:"order_no,omitempty"`
	CreateTime  string `json:"create_time,omitempty"`
	OrderSource string `json:"order_source,omitempty"`
}

// OnWPOrderAddEvent
// @id 1295
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=订单新增
func (s *Service) OnWPOrderAddEvent(handler msg.XinyunEventHandler[WPOrderAddEvent]) (service *Service) {
	var listener = &WPOrderAddListener{handler}
	s.Register(msg.WrapListener[WPOrderAddEvent](listener), "WP_order", "add")
	return s
}
