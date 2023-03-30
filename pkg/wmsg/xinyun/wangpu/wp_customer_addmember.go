package wangpu

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WpCustomerAddmemberListener struct {
	handler msg.XinyunEventHandler[WpCustomerAddmemberEvent]
}

func (s WpCustomerAddmemberListener) NewMessage() msg.MsgRequest[WpCustomerAddmemberEvent] {
	return &msg.XinyunOpenMessage[WpCustomerAddmemberEvent]{
		MsgBody: &WpCustomerAddmemberEvent{},
	}
}

func (s WpCustomerAddmemberListener) OnMessage(ctx context.Context, message msg.MsgRequest[WpCustomerAddmemberEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[WpCustomerAddmemberEvent]))
}

type WpCustomerAddmemberEvent struct {
	ChangeTime int64 `json:"change_time,omitempty"`
}

// OnWpCustomerAddmemberEvent
// @id 174
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=会员卡配置变更
func (s *Service) OnWpCustomerAddmemberEvent(handler msg.XinyunEventHandler[WpCustomerAddmemberEvent]) (service *Service) {
	var listener = &WpCustomerAddmemberListener{handler}
	s.Register(msg.WrapListener[WpCustomerAddmemberEvent](listener), "wp_customer", "addmember")
	return s
}
