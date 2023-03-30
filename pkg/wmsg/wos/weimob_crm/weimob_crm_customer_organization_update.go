package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmCustomerOrganizationUpdateListener struct {
	handler msg.WosEventHandler[WeimobCrmCustomerOrganizationUpdateEvent]
}

func (s WeimobCrmCustomerOrganizationUpdateListener) NewMessage() msg.MsgRequest[WeimobCrmCustomerOrganizationUpdateEvent] {
	return &msg.WosOpenMessage[WeimobCrmCustomerOrganizationUpdateEvent]{
		MsgBody: &WeimobCrmCustomerOrganizationUpdateEvent{},
	}
}

func (s WeimobCrmCustomerOrganizationUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmCustomerOrganizationUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmCustomerOrganizationUpdateEvent]))
}

type WeimobCrmCustomerOrganizationUpdateEvent struct {
	Wid          int64 `json:"wid,omitempty"`
	NewBelongVid int64 `json:"newBelongVid,omitempty"`
	OldBelongVid int64 `json:"oldBelongVid,omitempty"`
}

// OnWeimobCrmCustomerOrganizationUpdateEvent
// @id 1357
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1357?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=归属门店变更消息)
func (s *Service) OnWeimobCrmCustomerOrganizationUpdateEvent(handler msg.WosEventHandler[WeimobCrmCustomerOrganizationUpdateEvent]) (service *Service) {
	var listener = &WeimobCrmCustomerOrganizationUpdateListener{handler}
	s.Register(msg.WrapListener[WeimobCrmCustomerOrganizationUpdateEvent](listener), "weimob_crm.customer", "organizationUpdate")
	return s
}
