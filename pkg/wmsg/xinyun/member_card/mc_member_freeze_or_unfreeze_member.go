package member_card

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type McMemberFreezeOrUnfreezeMemberListener struct {
	handler msg.XinyunEventHandler[McMemberFreezeOrUnfreezeMemberEvent]
}

func (s McMemberFreezeOrUnfreezeMemberListener) NewMessage() msg.MsgRequest[McMemberFreezeOrUnfreezeMemberEvent] {
	return &msg.XinyunOpenMessage[McMemberFreezeOrUnfreezeMemberEvent]{
		MsgBody: &McMemberFreezeOrUnfreezeMemberEvent{},
	}
}

func (s McMemberFreezeOrUnfreezeMemberListener) OnMessage(ctx context.Context, message msg.MsgRequest[McMemberFreezeOrUnfreezeMemberEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[McMemberFreezeOrUnfreezeMemberEvent]))
}

type McMemberFreezeOrUnfreezeMemberEvent struct {
	Wid        int64  `json:"wid,omitempty"`
	MemberCode string `json:"memberCode,omitempty"`
	ChangeType int    `json:"changeType,omitempty"`
	Type       int    `json:"type,omitempty"`
}

// OnMcMemberFreezeOrUnfreezeMemberEvent
// @id 623
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=冻结/解冻会员
func (s *Service) OnMcMemberFreezeOrUnfreezeMemberEvent(handler msg.XinyunEventHandler[McMemberFreezeOrUnfreezeMemberEvent]) (service *Service) {
	var listener = &McMemberFreezeOrUnfreezeMemberListener{handler}
	s.Register(msg.WrapListener[McMemberFreezeOrUnfreezeMemberEvent](listener), "mc_member", "freezeOrUnfreezeMember")
	return s
}
