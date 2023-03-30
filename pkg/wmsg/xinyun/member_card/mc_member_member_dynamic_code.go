package member_card

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type McMemberMemberDynamicCodeListener struct {
	handler msg.XinyunEventHandler[McMemberMemberDynamicCodeEvent]
}

func (s McMemberMemberDynamicCodeListener) NewMessage() msg.MsgRequest[McMemberMemberDynamicCodeEvent] {
	return &msg.XinyunOpenMessage[McMemberMemberDynamicCodeEvent]{
		MsgBody: &McMemberMemberDynamicCodeEvent{},
	}
}

func (s McMemberMemberDynamicCodeListener) OnMessage(ctx context.Context, message msg.MsgRequest[McMemberMemberDynamicCodeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[McMemberMemberDynamicCodeEvent]))
}

type McMemberMemberDynamicCodeEvent struct {
	Pid                  int64  `json:"pid,omitempty"`
	Wid                  int64  `json:"wid,omitempty"`
	MemberCardTemplateId int64  `json:"memberCardTemplateId,omitempty"`
	Phone                string `json:"phone,omitempty"`
	DynamicCode          string `json:"dynamicCode,omitempty"`
	ExpireTime           int    `json:"expireTime,omitempty"`
}

// OnMcMemberMemberDynamicCodeEvent
// @id 1555
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=会员动态码
func (s *Service) OnMcMemberMemberDynamicCodeEvent(handler msg.XinyunEventHandler[McMemberMemberDynamicCodeEvent]) (service *Service) {
	var listener = &McMemberMemberDynamicCodeListener{handler}
	s.Register(msg.WrapListener[McMemberMemberDynamicCodeEvent](listener), "mc_member", "memberDynamicCode")
	return s
}
