package wangpu

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WPReturnOrderStatusChangeListener struct {
	handler msg.XinyunEventHandler[WPReturnOrderStatusChangeEvent]
}

func (s WPReturnOrderStatusChangeListener) NewMessage() msg.MsgRequest[WPReturnOrderStatusChangeEvent] {
	return &msg.XinyunOpenMessage[WPReturnOrderStatusChangeEvent]{
		MsgBody: &WPReturnOrderStatusChangeEvent{},
	}
}

func (s WPReturnOrderStatusChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[WPReturnOrderStatusChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[WPReturnOrderStatusChangeEvent]))
}

type WPReturnOrderStatusChangeEvent struct {
	OrderNo       string `json:"order_no,omitempty"`
	ReturnOrderNo string `json:"return_order_no,omitempty"`
}

// OnWPReturnOrderStatusChangeEvent
// @id 1267
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=维权单状态变更
func (s *Service) OnWPReturnOrderStatusChangeEvent(handler msg.XinyunEventHandler[WPReturnOrderStatusChangeEvent]) (service *Service) {
	var listener = &WPReturnOrderStatusChangeListener{handler}
	s.Register(msg.WrapListener[WPReturnOrderStatusChangeEvent](listener), "WP_returnOrder", "statusChange")
	return s
}
