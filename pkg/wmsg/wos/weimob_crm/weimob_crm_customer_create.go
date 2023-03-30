package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmCustomerCreateListener struct {
	handler msg.WosEventHandler[WeimobCrmCustomerCreateEvent]
}

func (s WeimobCrmCustomerCreateListener) NewMessage() msg.MsgRequest[WeimobCrmCustomerCreateEvent] {
	return &msg.WosOpenMessage[WeimobCrmCustomerCreateEvent]{
		MsgBody: &WeimobCrmCustomerCreateEvent{},
	}
}

func (s WeimobCrmCustomerCreateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmCustomerCreateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmCustomerCreateEvent]))
}

type WeimobCrmCustomerCreateEvent struct {
	MembershipCreateSceneType int64   `json:"membershipCreateSceneType,omitempty"`
	MembershipPlanId          int64   `json:"membershipPlanId,omitempty"`
	MembershipType            int64   `json:"membershipType,omitempty"`
	WidList                   []int64 `json:"widList,omitempty"`
}

// OnWeimobCrmCustomerCreateEvent
// @id 1280
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1280?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=成为客户消息)
func (s *Service) OnWeimobCrmCustomerCreateEvent(handler msg.WosEventHandler[WeimobCrmCustomerCreateEvent]) (service *Service) {
	var listener = &WeimobCrmCustomerCreateListener{handler}
	s.Register(msg.WrapListener[WeimobCrmCustomerCreateEvent](listener), "weimob_crm.customer", "create")
	return s
}
