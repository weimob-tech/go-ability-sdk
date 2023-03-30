package member_card

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type McMemberDeleteMemberCardListener struct {
	handler msg.XinyunEventHandler[McMemberDeleteMemberCardEvent]
}

func (s McMemberDeleteMemberCardListener) NewMessage() msg.MsgRequest[McMemberDeleteMemberCardEvent] {
	return &msg.XinyunOpenMessage[McMemberDeleteMemberCardEvent]{
		MsgBody: &McMemberDeleteMemberCardEvent{},
	}
}

func (s McMemberDeleteMemberCardListener) OnMessage(ctx context.Context, message msg.MsgRequest[McMemberDeleteMemberCardEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[McMemberDeleteMemberCardEvent]))
}

type McMemberDeleteMemberCardEvent struct {
	Pid          int64  `json:"pid,omitempty"`
	Wid          int64  `json:"wid,omitempty"`
	McTemplateId int64  `json:"mcTemplateId,omitempty"`
	Mid          int64  `json:"mid,omitempty"`
	Code         string `json:"code,omitempty"`
}

// OnMcMemberDeleteMemberCardEvent
// @id 1703
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=删除会员卡消息
func (s *Service) OnMcMemberDeleteMemberCardEvent(handler msg.XinyunEventHandler[McMemberDeleteMemberCardEvent]) (service *Service) {
	var listener = &McMemberDeleteMemberCardListener{handler}
	s.Register(msg.WrapListener[McMemberDeleteMemberCardEvent](listener), "mc_member", "deleteMemberCard")
	return s
}
