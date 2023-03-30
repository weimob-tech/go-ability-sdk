package wangpu

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WPReturnOrderAddListener struct {
	handler msg.XinyunEventHandler[WPReturnOrderAddEvent]
}

func (s WPReturnOrderAddListener) NewMessage() msg.MsgRequest[WPReturnOrderAddEvent] {
	return &msg.XinyunOpenMessage[WPReturnOrderAddEvent]{
		MsgBody: &WPReturnOrderAddEvent{},
	}
}

func (s WPReturnOrderAddListener) OnMessage(ctx context.Context, message msg.MsgRequest[WPReturnOrderAddEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[WPReturnOrderAddEvent]))
}

type WPReturnOrderAddEvent struct {
	OrderNo       string `json:"order_no,omitempty"`
	CreateTime    string `json:"create_time,omitempty"`
	ReturnOrderNo string `json:"return_order_no,omitempty"`
}

// OnWPReturnOrderAddEvent
// @id 1266
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=维权单新增
func (s *Service) OnWPReturnOrderAddEvent(handler msg.XinyunEventHandler[WPReturnOrderAddEvent]) (service *Service) {
	var listener = &WPReturnOrderAddListener{handler}
	s.Register(msg.WrapListener[WPReturnOrderAddEvent](listener), "WP_returnOrder", "add")
	return s
}
