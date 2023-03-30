package member_card

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type McMemberMemberBaseInfoChangeListener struct {
	handler msg.XinyunEventHandler[McMemberMemberBaseInfoChangeEvent]
}

func (s McMemberMemberBaseInfoChangeListener) NewMessage() msg.MsgRequest[McMemberMemberBaseInfoChangeEvent] {
	return &msg.XinyunOpenMessage[McMemberMemberBaseInfoChangeEvent]{
		MsgBody: &McMemberMemberBaseInfoChangeEvent{},
	}
}

func (s McMemberMemberBaseInfoChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[McMemberMemberBaseInfoChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[McMemberMemberBaseInfoChangeEvent]))
}

type McMemberMemberBaseInfoChangeEvent struct {
	Pid        int64  `json:"pid,omitempty"`
	Wid        int64  `json:"wid,omitempty"`
	Source     int    `json:"source,omitempty"`
	MemberCode string `json:"memberCode,omitempty"`
	Type       int    `json:"type,omitempty"`
}

// OnMcMemberMemberBaseInfoChangeEvent
// @id 625
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=会员基本信息变更
func (s *Service) OnMcMemberMemberBaseInfoChangeEvent(handler msg.XinyunEventHandler[McMemberMemberBaseInfoChangeEvent]) (service *Service) {
	var listener = &McMemberMemberBaseInfoChangeListener{handler}
	s.Register(msg.WrapListener[McMemberMemberBaseInfoChangeEvent](listener), "mc_member", "memberBaseInfoChange")
	return s
}
