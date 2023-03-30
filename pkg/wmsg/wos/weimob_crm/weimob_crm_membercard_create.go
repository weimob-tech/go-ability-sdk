package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmMembercardCreateListener struct {
	handler msg.WosEventHandler[WeimobCrmMembercardCreateEvent]
}

func (s WeimobCrmMembercardCreateListener) NewMessage() msg.MsgRequest[WeimobCrmMembercardCreateEvent] {
	return &msg.WosOpenMessage[WeimobCrmMembercardCreateEvent]{
		MsgBody: &WeimobCrmMembercardCreateEvent{},
	}
}

func (s WeimobCrmMembercardCreateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmMembercardCreateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmMembercardCreateEvent]))
}

type WeimobCrmMembercardCreateEvent struct {
	Wid               int64  `json:"wid,omitempty"`
	BosId             int64  `json:"bosId,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	Vid               int64  `json:"vid,omitempty"`
	BindType          int64  `json:"bindType,omitempty"`
	MembershipPlanId  int64  `json:"membershipPlanId,omitempty"`
	CustomCardNo      string `json:"customCardNo,omitempty"`
	Offline           int64  `json:"offline,omitempty"`
	GuideWid          int64  `json:"guideWid,omitempty"`
	SaasChannel       string `json:"saasChannel,omitempty"`
}

// OnWeimobCrmMembercardCreateEvent
// @id 1248
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1248?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=会员开卡事件)
func (s *Service) OnWeimobCrmMembercardCreateEvent(handler msg.WosEventHandler[WeimobCrmMembercardCreateEvent]) (service *Service) {
	var listener = &WeimobCrmMembercardCreateListener{handler}
	s.Register(msg.WrapListener[WeimobCrmMembercardCreateEvent](listener), "weimob_crm.membercard", "create")
	return s
}
