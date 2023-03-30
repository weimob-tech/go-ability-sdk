package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmCustomerUpdateListener struct {
	handler msg.WosEventHandler[WeimobCrmCustomerUpdateEvent]
}

func (s WeimobCrmCustomerUpdateListener) NewMessage() msg.MsgRequest[WeimobCrmCustomerUpdateEvent] {
	return &msg.WosOpenMessage[WeimobCrmCustomerUpdateEvent]{
		MsgBody: &WeimobCrmCustomerUpdateEvent{},
	}
}

func (s WeimobCrmCustomerUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmCustomerUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmCustomerUpdateEvent]))
}

type WeimobCrmCustomerUpdateEvent struct {
	Wid int64 `json:"wid,omitempty"`
}

// OnWeimobCrmCustomerUpdateEvent
// @id 1355
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1355?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=客户基础信息变更消息)
func (s *Service) OnWeimobCrmCustomerUpdateEvent(handler msg.WosEventHandler[WeimobCrmCustomerUpdateEvent]) (service *Service) {
	var listener = &WeimobCrmCustomerUpdateListener{handler}
	s.Register(msg.WrapListener[WeimobCrmCustomerUpdateEvent](listener), "weimob_crm.customer", "update")
	return s
}
