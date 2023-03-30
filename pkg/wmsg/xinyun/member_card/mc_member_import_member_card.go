package member_card

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type McMemberImportMemberCardListener struct {
	handler msg.XinyunEventHandler[McMemberImportMemberCardEvent]
}

func (s McMemberImportMemberCardListener) NewMessage() msg.MsgRequest[McMemberImportMemberCardEvent] {
	return &msg.XinyunOpenMessage[McMemberImportMemberCardEvent]{
		MsgBody: &McMemberImportMemberCardEvent{},
	}
}

func (s McMemberImportMemberCardListener) OnMessage(ctx context.Context, message msg.MsgRequest[McMemberImportMemberCardEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[McMemberImportMemberCardEvent]))
}

type McMemberImportMemberCardEvent struct {
	WidList              []int `json:"widList,omitempty"`
	MemberCardTemplateId int64 `json:"memberCardTemplateId,omitempty"`
}

// OnMcMemberImportMemberCardEvent
// @id 1448
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=导入会员消息
func (s *Service) OnMcMemberImportMemberCardEvent(handler msg.XinyunEventHandler[McMemberImportMemberCardEvent]) (service *Service) {
	var listener = &McMemberImportMemberCardListener{handler}
	s.Register(msg.WrapListener[McMemberImportMemberCardEvent](listener), "mc_member", "importMemberCard")
	return s
}
