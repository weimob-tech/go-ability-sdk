package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmCustomerStatusUpdateListener struct {
	handler msg.WosEventHandler[WeimobCrmCustomerStatusUpdateEvent]
}

func (s WeimobCrmCustomerStatusUpdateListener) NewMessage() msg.MsgRequest[WeimobCrmCustomerStatusUpdateEvent] {
	return &msg.WosOpenMessage[WeimobCrmCustomerStatusUpdateEvent]{
		MsgBody: &WeimobCrmCustomerStatusUpdateEvent{},
	}
}

func (s WeimobCrmCustomerStatusUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmCustomerStatusUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmCustomerStatusUpdateEvent]))
}

type WeimobCrmCustomerStatusUpdateEvent struct {
	ChangeType int64 `json:"changeType,omitempty"`
	Wid        int64 `json:"wid,omitempty"`
}

// OnWeimobCrmCustomerStatusUpdateEvent
// @id 1356
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1356?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=冻结/解冻客户消息)
func (s *Service) OnWeimobCrmCustomerStatusUpdateEvent(handler msg.WosEventHandler[WeimobCrmCustomerStatusUpdateEvent]) (service *Service) {
	var listener = &WeimobCrmCustomerStatusUpdateListener{handler}
	s.Register(msg.WrapListener[WeimobCrmCustomerStatusUpdateEvent](listener), "weimob_crm.customer", "statusUpdate")
	return s
}
