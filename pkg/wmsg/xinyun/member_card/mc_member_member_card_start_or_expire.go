package member_card

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type McMemberMemberCardStartOrExpireListener struct {
	handler msg.XinyunEventHandler[McMemberMemberCardStartOrExpireEvent]
}

func (s McMemberMemberCardStartOrExpireListener) NewMessage() msg.MsgRequest[McMemberMemberCardStartOrExpireEvent] {
	return &msg.XinyunOpenMessage[McMemberMemberCardStartOrExpireEvent]{
		MsgBody: &McMemberMemberCardStartOrExpireEvent{},
	}
}

func (s McMemberMemberCardStartOrExpireListener) OnMessage(ctx context.Context, message msg.MsgRequest[McMemberMemberCardStartOrExpireEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[McMemberMemberCardStartOrExpireEvent]))
}

type McMemberMemberCardStartOrExpireEvent struct {
	MemberCardTemplateId int64 `json:"memberCardTemplateId,omitempty"`
	Wid                  int64 `json:"wid,omitempty"`
	Type                 int   `json:"type,omitempty"`
}

// OnMcMemberMemberCardStartOrExpireEvent
// @id 1473
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=会员卡状态变更消息（生效/失效）
func (s *Service) OnMcMemberMemberCardStartOrExpireEvent(handler msg.XinyunEventHandler[McMemberMemberCardStartOrExpireEvent]) (service *Service) {
	var listener = &McMemberMemberCardStartOrExpireListener{handler}
	s.Register(msg.WrapListener[McMemberMemberCardStartOrExpireEvent](listener), "mc_member", "memberCardStartOrExpire")
	return s
}
