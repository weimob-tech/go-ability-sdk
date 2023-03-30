package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmMembercardLevelChangeListener struct {
	handler msg.WosEventHandler[WeimobCrmMembercardLevelChangeEvent]
}

func (s WeimobCrmMembercardLevelChangeListener) NewMessage() msg.MsgRequest[WeimobCrmMembercardLevelChangeEvent] {
	return &msg.WosOpenMessage[WeimobCrmMembercardLevelChangeEvent]{
		MsgBody: &WeimobCrmMembercardLevelChangeEvent{},
	}
}

func (s WeimobCrmMembercardLevelChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmMembercardLevelChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmMembercardLevelChangeEvent]))
}

type WeimobCrmMembercardLevelChangeEvent struct {
	BosId             int64  `json:"bosId	,omitempty"`
	Reason            string `json:"reason,omitempty"`
	ChangeTime        int64  `json:"changeTime,omitempty"`
	NewLevelId        int64  `json:"newLevelId,omitempty"`
	OldLevelId        int64  `json:"oldLevelId,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	Vid               int64  `json:"vid,omitempty"`
	Wid               int64  `json:"wid,omitempty"`
}

// OnWeimobCrmMembercardLevelChangeEvent
// @id 1314
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1314?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=会员等级变更消息)
func (s *Service) OnWeimobCrmMembercardLevelChangeEvent(handler msg.WosEventHandler[WeimobCrmMembercardLevelChangeEvent]) (service *Service) {
	var listener = &WeimobCrmMembercardLevelChangeListener{handler}
	s.Register(msg.WrapListener[WeimobCrmMembercardLevelChangeEvent](listener), "weimob_crm.membercard.level", "change")
	return s
}
