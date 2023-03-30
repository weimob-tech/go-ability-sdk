package member_card

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type McMemberBecomeCustomerListener struct {
	handler msg.XinyunEventHandler[McMemberBecomeCustomerEvent]
}

func (s McMemberBecomeCustomerListener) NewMessage() msg.MsgRequest[McMemberBecomeCustomerEvent] {
	return &msg.XinyunOpenMessage[McMemberBecomeCustomerEvent]{
		MsgBody: &McMemberBecomeCustomerEvent{},
	}
}

func (s McMemberBecomeCustomerListener) OnMessage(ctx context.Context, message msg.MsgRequest[McMemberBecomeCustomerEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[McMemberBecomeCustomerEvent]))
}

type McMemberBecomeCustomerEvent struct {
	Wid int64 `json:"wid,omitempty"`
}

// OnMcMemberBecomeCustomerEvent
// @id 1446
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=成为客户消息
func (s *Service) OnMcMemberBecomeCustomerEvent(handler msg.XinyunEventHandler[McMemberBecomeCustomerEvent]) (service *Service) {
	var listener = &McMemberBecomeCustomerListener{handler}
	s.Register(msg.WrapListener[McMemberBecomeCustomerEvent](listener), "mc_member", "becomeCustomer")
	return s
}
