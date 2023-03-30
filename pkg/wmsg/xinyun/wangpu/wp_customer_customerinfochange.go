package wangpu

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WpCustomerCustomerinfochangeListener struct {
	handler msg.XinyunEventHandler[WpCustomerCustomerinfochangeEvent]
}

func (s WpCustomerCustomerinfochangeListener) NewMessage() msg.MsgRequest[WpCustomerCustomerinfochangeEvent] {
	return &msg.XinyunOpenMessage[WpCustomerCustomerinfochangeEvent]{
		MsgBody: &WpCustomerCustomerinfochangeEvent{},
	}
}

func (s WpCustomerCustomerinfochangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[WpCustomerCustomerinfochangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[WpCustomerCustomerinfochangeEvent]))
}

type WpCustomerCustomerinfochangeEvent struct {
	CustomerNos []string `json:"customer_nos,omitempty"`
	ChangeTime  int64    `json:"change_time,omitempty"`
}

// OnWpCustomerCustomerinfochangeEvent
// @id 1268
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=客户信息变更
func (s *Service) OnWpCustomerCustomerinfochangeEvent(handler msg.XinyunEventHandler[WpCustomerCustomerinfochangeEvent]) (service *Service) {
	var listener = &WpCustomerCustomerinfochangeListener{handler}
	s.Register(msg.WrapListener[WpCustomerCustomerinfochangeEvent](listener), "wp_customer", "customerinfochange")
	return s
}
