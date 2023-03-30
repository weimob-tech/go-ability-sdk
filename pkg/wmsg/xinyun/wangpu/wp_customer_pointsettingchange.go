package wangpu

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WpCustomerPointsettingchangeListener struct {
	handler msg.XinyunEventHandler[WpCustomerPointsettingchangeEvent]
}

func (s WpCustomerPointsettingchangeListener) NewMessage() msg.MsgRequest[WpCustomerPointsettingchangeEvent] {
	return &msg.XinyunOpenMessage[WpCustomerPointsettingchangeEvent]{
		MsgBody: &WpCustomerPointsettingchangeEvent{},
	}
}

func (s WpCustomerPointsettingchangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[WpCustomerPointsettingchangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[WpCustomerPointsettingchangeEvent]))
}

type WpCustomerPointsettingchangeEvent struct {
	ChangeTime int64 `json:"change_time,omitempty"`
}

// OnWpCustomerPointsettingchangeEvent
// @id 178
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=积分配置变更
func (s *Service) OnWpCustomerPointsettingchangeEvent(handler msg.XinyunEventHandler[WpCustomerPointsettingchangeEvent]) (service *Service) {
	var listener = &WpCustomerPointsettingchangeListener{handler}
	s.Register(msg.WrapListener[WpCustomerPointsettingchangeEvent](listener), "wp_customer", "pointsettingchange")
	return s
}
