package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmMembercardGrowingcreditChangeListener struct {
	handler msg.WosEventHandler[WeimobCrmMembercardGrowingcreditChangeEvent]
}

func (s WeimobCrmMembercardGrowingcreditChangeListener) NewMessage() msg.MsgRequest[WeimobCrmMembercardGrowingcreditChangeEvent] {
	return &msg.WosOpenMessage[WeimobCrmMembercardGrowingcreditChangeEvent]{
		MsgBody: &WeimobCrmMembercardGrowingcreditChangeEvent{},
	}
}

func (s WeimobCrmMembercardGrowingcreditChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmMembercardGrowingcreditChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmMembercardGrowingcreditChangeEvent]))
}

type WeimobCrmMembercardGrowingcreditChangeEvent struct {
	ChangeReason     string `json:"changeReason,omitempty"`
	ChangeTime       int64  `json:"changeTime,omitempty"`
	Wid              int64  `json:"wid,omitempty"`
	ChangeGrowth     int64  `json:"changeGrowth,omitempty"`
	MembershipPlanId int64  `json:"membershipPlanId,omitempty"`
}

// OnWeimobCrmMembercardGrowingcreditChangeEvent
// @id 1282
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1282?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=会员成长值变更消息)
func (s *Service) OnWeimobCrmMembercardGrowingcreditChangeEvent(handler msg.WosEventHandler[WeimobCrmMembercardGrowingcreditChangeEvent]) (service *Service) {
	var listener = &WeimobCrmMembercardGrowingcreditChangeListener{handler}
	s.Register(msg.WrapListener[WeimobCrmMembercardGrowingcreditChangeEvent](listener), "weimob_crm.membercard.growingcredit", "change")
	return s
}
