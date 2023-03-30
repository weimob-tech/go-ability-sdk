package member_card

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type McMemberOpenMemberCardListener struct {
	handler msg.XinyunEventHandler[McMemberOpenMemberCardEvent]
}

func (s McMemberOpenMemberCardListener) NewMessage() msg.MsgRequest[McMemberOpenMemberCardEvent] {
	return &msg.XinyunOpenMessage[McMemberOpenMemberCardEvent]{
		MsgBody: &McMemberOpenMemberCardEvent{},
	}
}

func (s McMemberOpenMemberCardListener) OnMessage(ctx context.Context, message msg.MsgRequest[McMemberOpenMemberCardEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[McMemberOpenMemberCardEvent]))
}

type McMemberOpenMemberCardEvent struct {
	Pid        int64  `json:"pid,omitempty"`
	Wid        int64  `json:"wid,omitempty"`
	MemberCode string `json:"memberCode,omitempty"`
	Type       int    `json:"type,omitempty"`
	StoreId    int64  `json:"storeId,omitempty"`
}

// OnMcMemberOpenMemberCardEvent
// @id 622
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=会员开卡消息
func (s *Service) OnMcMemberOpenMemberCardEvent(handler msg.XinyunEventHandler[McMemberOpenMemberCardEvent]) (service *Service) {
	var listener = &McMemberOpenMemberCardListener{handler}
	s.Register(msg.WrapListener[McMemberOpenMemberCardEvent](listener), "mc_member", "openMemberCard")
	return s
}
