package member_card

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type McMemberMemberGrowthChangeListener struct {
	handler msg.XinyunEventHandler[McMemberMemberGrowthChangeEvent]
}

func (s McMemberMemberGrowthChangeListener) NewMessage() msg.MsgRequest[McMemberMemberGrowthChangeEvent] {
	return &msg.XinyunOpenMessage[McMemberMemberGrowthChangeEvent]{
		MsgBody: &McMemberMemberGrowthChangeEvent{},
	}
}

func (s McMemberMemberGrowthChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[McMemberMemberGrowthChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[McMemberMemberGrowthChangeEvent]))
}

type McMemberMemberGrowthChangeEvent struct {
	Pid          int64  `json:"pid,omitempty"`
	Wid          int64  `json:"wid,omitempty"`
	Source       int64  `json:"source,omitempty"`
	MemberCode   string `json:"memberCode,omitempty"`
	ChangeTime   int64  `json:"changeTime,omitempty"`
	ChangeReason string `json:"changeReason,omitempty"`
	ChangeGrowth string `json:"changeGrowth,omitempty"`
	Type         int64  `json:"type,omitempty"`
	ChangeType   int64  `json:"changeType,omitempty"`
}

// OnMcMemberMemberGrowthChangeEvent
// @id 626
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=会员成长值变更
func (s *Service) OnMcMemberMemberGrowthChangeEvent(handler msg.XinyunEventHandler[McMemberMemberGrowthChangeEvent]) (service *Service) {
	var listener = &McMemberMemberGrowthChangeListener{handler}
	s.Register(msg.WrapListener[McMemberMemberGrowthChangeEvent](listener), "mc_member", "memberGrowthChange")
	return s
}
