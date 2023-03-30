package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmCustomerStatusChangeListener struct {
	handler msg.WosEventHandler[WeimobCrmCustomerStatusChangeEvent]
}

func (s WeimobCrmCustomerStatusChangeListener) NewMessage() msg.MsgRequest[WeimobCrmCustomerStatusChangeEvent] {
	return &msg.WosOpenMessage[WeimobCrmCustomerStatusChangeEvent]{
		MsgBody: &WeimobCrmCustomerStatusChangeEvent{},
	}
}

func (s WeimobCrmCustomerStatusChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmCustomerStatusChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmCustomerStatusChangeEvent]))
}

type WeimobCrmCustomerStatusChangeEvent struct {
	ChangeType int64 `json:"changeType,omitempty"`
	Wid        int64 `json:"wid,omitempty"`
}

// OnWeimobCrmCustomerStatusChangeEvent
// @id 1310
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1310?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=冻结/解冻客户消息)
func (s *Service) OnWeimobCrmCustomerStatusChangeEvent(handler msg.WosEventHandler[WeimobCrmCustomerStatusChangeEvent]) (service *Service) {
	var listener = &WeimobCrmCustomerStatusChangeListener{handler}
	s.Register(msg.WrapListener[WeimobCrmCustomerStatusChangeEvent](listener), "weimob_crm.customer.status", "change")
	return s
}
