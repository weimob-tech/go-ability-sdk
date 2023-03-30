package member_card

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type McMemberImportCustomerListener struct {
	handler msg.XinyunEventHandler[McMemberImportCustomerEvent]
}

func (s McMemberImportCustomerListener) NewMessage() msg.MsgRequest[McMemberImportCustomerEvent] {
	return &msg.XinyunOpenMessage[McMemberImportCustomerEvent]{
		MsgBody: &McMemberImportCustomerEvent{},
	}
}

func (s McMemberImportCustomerListener) OnMessage(ctx context.Context, message msg.MsgRequest[McMemberImportCustomerEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[McMemberImportCustomerEvent]))
}

type McMemberImportCustomerEvent struct {
	WidList []int `json:"widList,omitempty"`
}

// OnMcMemberImportCustomerEvent
// @id 1447
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=导入客户消息
func (s *Service) OnMcMemberImportCustomerEvent(handler msg.XinyunEventHandler[McMemberImportCustomerEvent]) (service *Service) {
	var listener = &McMemberImportCustomerListener{handler}
	s.Register(msg.WrapListener[McMemberImportCustomerEvent](listener), "mc_member", "importCustomer")
	return s
}
