package member_card

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type McMemberMemberRankChangeListener struct {
	handler msg.XinyunEventHandler[McMemberMemberRankChangeEvent]
}

func (s McMemberMemberRankChangeListener) NewMessage() msg.MsgRequest[McMemberMemberRankChangeEvent] {
	return &msg.XinyunOpenMessage[McMemberMemberRankChangeEvent]{
		MsgBody: &McMemberMemberRankChangeEvent{},
	}
}

func (s McMemberMemberRankChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[McMemberMemberRankChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[McMemberMemberRankChangeEvent]))
}

type McMemberMemberRankChangeEvent struct {
	Pid          int64  `json:"pid,omitempty"`
	Wid          int64  `json:"wid,omitempty"`
	Source       int    `json:"source,omitempty"`
	MemberCode   string `json:"memberCode,omitempty"`
	OldRankId    int64  `json:"oldRankId,omitempty"`
	NewRankId    int64  `json:"newRankId,omitempty"`
	ChangeReason string `json:"changeReason,omitempty"`
	ChangeTime   int64  `json:"changeTime,omitempty"`
	Type         int    `json:"type,omitempty"`
}

// OnMcMemberMemberRankChangeEvent
// @id 624
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=会员等级变更
func (s *Service) OnMcMemberMemberRankChangeEvent(handler msg.XinyunEventHandler[McMemberMemberRankChangeEvent]) (service *Service) {
	var listener = &McMemberMemberRankChangeListener{handler}
	s.Register(msg.WrapListener[McMemberMemberRankChangeEvent](listener), "mc_member", "memberRankChange")
	return s
}
