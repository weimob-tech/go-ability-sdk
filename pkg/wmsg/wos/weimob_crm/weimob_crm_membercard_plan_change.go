package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmMembercardPlanChangeListener struct {
	handler msg.WosEventHandler[WeimobCrmMembercardPlanChangeEvent]
}

func (s WeimobCrmMembercardPlanChangeListener) NewMessage() msg.MsgRequest[WeimobCrmMembercardPlanChangeEvent] {
	return &msg.WosOpenMessage[WeimobCrmMembercardPlanChangeEvent]{
		MsgBody: &WeimobCrmMembercardPlanChangeEvent{},
	}
}

func (s WeimobCrmMembercardPlanChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmMembercardPlanChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmMembercardPlanChangeEvent]))
}

type WeimobCrmMembercardPlanChangeEvent struct {
	MembershipPlanId int64 `json:"membershipPlanId,omitempty"`
}

// OnWeimobCrmMembercardPlanChangeEvent
// @id 1281
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1281?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=会员卡方案变更消息)
func (s *Service) OnWeimobCrmMembercardPlanChangeEvent(handler msg.WosEventHandler[WeimobCrmMembercardPlanChangeEvent]) (service *Service) {
	var listener = &WeimobCrmMembercardPlanChangeListener{handler}
	s.Register(msg.WrapListener[WeimobCrmMembercardPlanChangeEvent](listener), "weimob_crm.membercard.plan", "change")
	return s
}
