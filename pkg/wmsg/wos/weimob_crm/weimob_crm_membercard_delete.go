package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmMembercardDeleteListener struct {
	handler msg.WosEventHandler[WeimobCrmMembercardDeleteEvent]
}

func (s WeimobCrmMembercardDeleteListener) NewMessage() msg.MsgRequest[WeimobCrmMembercardDeleteEvent] {
	return &msg.WosOpenMessage[WeimobCrmMembercardDeleteEvent]{
		MsgBody: &WeimobCrmMembercardDeleteEvent{},
	}
}

func (s WeimobCrmMembercardDeleteListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmMembercardDeleteEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmMembercardDeleteEvent]))
}

type WeimobCrmMembercardDeleteEvent struct {
	membershipPlanId  int64  `json:" membershipPlanId,omitempty"`
	CustomCardNo      string `json:"customCardNo,omitempty"`
	Wid               int64  `json:"wid,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
}

// OnWeimobCrmMembercardDeleteEvent
// @id 1278
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1278?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除会员卡消息)
func (s *Service) OnWeimobCrmMembercardDeleteEvent(handler msg.WosEventHandler[WeimobCrmMembercardDeleteEvent]) (service *Service) {
	var listener = &WeimobCrmMembercardDeleteListener{handler}
	s.Register(msg.WrapListener[WeimobCrmMembercardDeleteEvent](listener), "weimob_crm.membercard", "delete")
	return s
}
