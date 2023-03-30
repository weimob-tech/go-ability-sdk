package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmCustomerChangeListener struct {
	handler msg.WosEventHandler[WeimobCrmCustomerChangeEvent]
}

func (s WeimobCrmCustomerChangeListener) NewMessage() msg.MsgRequest[WeimobCrmCustomerChangeEvent] {
	return &msg.WosOpenMessage[WeimobCrmCustomerChangeEvent]{
		MsgBody: &WeimobCrmCustomerChangeEvent{},
	}
}

func (s WeimobCrmCustomerChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmCustomerChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmCustomerChangeEvent]))
}

type WeimobCrmCustomerChangeEvent struct {
	Wid int64 `json:"wid,omitempty"`
}

// OnWeimobCrmCustomerChangeEvent
// @id 1311
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1311?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=客户基础信息变更消息)
func (s *Service) OnWeimobCrmCustomerChangeEvent(handler msg.WosEventHandler[WeimobCrmCustomerChangeEvent]) (service *Service) {
	var listener = &WeimobCrmCustomerChangeListener{handler}
	s.Register(msg.WrapListener[WeimobCrmCustomerChangeEvent](listener), "weimob_crm.customer", "change")
	return s
}
