package member_card

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type McCardTemplatePointOrLevelRuleChangeListener struct {
	handler msg.XinyunEventHandler[McCardTemplatePointOrLevelRuleChangeEvent]
}

func (s McCardTemplatePointOrLevelRuleChangeListener) NewMessage() msg.MsgRequest[McCardTemplatePointOrLevelRuleChangeEvent] {
	return &msg.XinyunOpenMessage[McCardTemplatePointOrLevelRuleChangeEvent]{
		MsgBody: &McCardTemplatePointOrLevelRuleChangeEvent{},
	}
}

func (s McCardTemplatePointOrLevelRuleChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[McCardTemplatePointOrLevelRuleChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[McCardTemplatePointOrLevelRuleChangeEvent]))
}

type McCardTemplatePointOrLevelRuleChangeEvent struct {
	Pid        int64 `json:"pid,omitempty"`
	Source     int   `json:"source,omitempty"`
	ChangeType int   `json:"changeType,omitempty"`
}

// OnMcCardTemplatePointOrLevelRuleChangeEvent
// @id 643
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=积分等级配置变更消息
func (s *Service) OnMcCardTemplatePointOrLevelRuleChangeEvent(handler msg.XinyunEventHandler[McCardTemplatePointOrLevelRuleChangeEvent]) (service *Service) {
	var listener = &McCardTemplatePointOrLevelRuleChangeListener{handler}
	s.Register(msg.WrapListener[McCardTemplatePointOrLevelRuleChangeEvent](listener), "mc_card_template", "pointOrLevelRuleChange")
	return s
}
