package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeCustomerDeleteListener struct {
	handler msg.XinyunEventHandler[XiaokeCustomerDeleteEvent]
}

func (s XiaokeCustomerDeleteListener) NewMessage() msg.MsgRequest[XiaokeCustomerDeleteEvent] {
	return &msg.XinyunOpenMessage[XiaokeCustomerDeleteEvent]{
		MsgBody: &XiaokeCustomerDeleteEvent{},
	}
}

func (s XiaokeCustomerDeleteListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeCustomerDeleteEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeCustomerDeleteEvent]))
}

type XiaokeCustomerDeleteEvent struct {
	NicheKeys    []string         `json:"nicheKeys,omitempty"`
	CustomerKeys []map[string]any `json:"customerKeys,omitempty"`
	ContactKeys  []string         `json:"contactKeys,omitempty"`
	BuildTime    int              `json:"buildTime,omitempty"`
	Wid          int64            `json:"wid,omitempty"`
}

// OnXiaokeCustomerDeleteEvent
// @id 1763
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=删除客户
func (s *Service) OnXiaokeCustomerDeleteEvent(handler msg.XinyunEventHandler[XiaokeCustomerDeleteEvent]) (service *Service) {
	var listener = &XiaokeCustomerDeleteListener{handler}
	s.Register(msg.WrapListener[XiaokeCustomerDeleteEvent](listener), "xiaoke_customer", "delete")
	return s
}
