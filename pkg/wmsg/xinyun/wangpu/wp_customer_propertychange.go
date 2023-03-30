package wangpu

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WpCustomerPropertychangeListener struct {
	handler msg.XinyunEventHandler[WpCustomerPropertychangeEvent]
}

func (s WpCustomerPropertychangeListener) NewMessage() msg.MsgRequest[WpCustomerPropertychangeEvent] {
	return &msg.XinyunOpenMessage[WpCustomerPropertychangeEvent]{
		MsgBody: &WpCustomerPropertychangeEvent{},
	}
}

func (s WpCustomerPropertychangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[WpCustomerPropertychangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[WpCustomerPropertychangeEvent]))
}

type WpCustomerPropertychangeEvent struct {
	CustomerNos []string `json:"customer_nos,omitempty"`
	ChangeTime  int64    `json:"change_time,omitempty"`
}

// OnWpCustomerPropertychangeEvent
// @id 176
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=客户资产变更
func (s *Service) OnWpCustomerPropertychangeEvent(handler msg.XinyunEventHandler[WpCustomerPropertychangeEvent]) (service *Service) {
	var listener = &WpCustomerPropertychangeListener{handler}
	s.Register(msg.WrapListener[WpCustomerPropertychangeEvent](listener), "wp_customer", "propertychange")
	return s
}
