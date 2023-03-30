package member_card

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type McCardTemplateMemberCardTemplateChangeListener struct {
	handler msg.XinyunEventHandler[McCardTemplateMemberCardTemplateChangeEvent]
}

func (s McCardTemplateMemberCardTemplateChangeListener) NewMessage() msg.MsgRequest[McCardTemplateMemberCardTemplateChangeEvent] {
	return &msg.XinyunOpenMessage[McCardTemplateMemberCardTemplateChangeEvent]{
		MsgBody: &McCardTemplateMemberCardTemplateChangeEvent{},
	}
}

func (s McCardTemplateMemberCardTemplateChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[McCardTemplateMemberCardTemplateChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[McCardTemplateMemberCardTemplateChangeEvent]))
}

type McCardTemplateMemberCardTemplateChangeEvent struct {
	Pid                  int64 `json:"pid,omitempty"`
	Source               int   `json:"source,omitempty"`
	CardTemplateId       int64 `json:"cardTemplateId,omitempty"`
	MemberCardTemplateId int64 `json:"memberCardTemplateId,omitempty"`
}

// OnMcCardTemplateMemberCardTemplateChangeEvent
// @id 629
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=会员卡模板信息变更
func (s *Service) OnMcCardTemplateMemberCardTemplateChangeEvent(handler msg.XinyunEventHandler[McCardTemplateMemberCardTemplateChangeEvent]) (service *Service) {
	var listener = &McCardTemplateMemberCardTemplateChangeListener{handler}
	s.Register(msg.WrapListener[McCardTemplateMemberCardTemplateChangeEvent](listener), "mc_card_template", "memberCardTemplateChange")
	return s
}
