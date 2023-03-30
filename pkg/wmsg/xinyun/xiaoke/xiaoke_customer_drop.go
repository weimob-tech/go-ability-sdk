package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeCustomerDropListener struct {
	handler msg.XinyunEventHandler[XiaokeCustomerDropEvent]
}

func (s XiaokeCustomerDropListener) NewMessage() msg.MsgRequest[XiaokeCustomerDropEvent] {
	return &msg.XinyunOpenMessage[XiaokeCustomerDropEvent]{
		MsgBody: &XiaokeCustomerDropEvent{},
	}
}

func (s XiaokeCustomerDropListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeCustomerDropEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeCustomerDropEvent]))
}

type XiaokeCustomerDropEvent struct {
	Keys      []string `json:"keys,omitempty"`
	Owner     int64    `json:"owner,omitempty"`
	BuildTime int64    `json:"buildTime,omitempty"`
	Wid       int64    `json:"wid,omitempty"`
}

// OnXiaokeCustomerDropEvent
// @id 1761
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=客户掉保公海
func (s *Service) OnXiaokeCustomerDropEvent(handler msg.XinyunEventHandler[XiaokeCustomerDropEvent]) (service *Service) {
	var listener = &XiaokeCustomerDropListener{handler}
	s.Register(msg.WrapListener[XiaokeCustomerDropEvent](listener), "xiaoke_customer", "drop")
	return s
}
