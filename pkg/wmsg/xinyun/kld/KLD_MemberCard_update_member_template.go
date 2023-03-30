package kld

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type KLDMemberCardUpdateMemberTemplateListener struct {
	handler msg.XinyunEventHandler[KLDMemberCardUpdateMemberTemplateEvent]
}

func (s KLDMemberCardUpdateMemberTemplateListener) NewMessage() msg.MsgRequest[KLDMemberCardUpdateMemberTemplateEvent] {
	return &msg.XinyunOpenMessage[KLDMemberCardUpdateMemberTemplateEvent]{
		MsgBody: &KLDMemberCardUpdateMemberTemplateEvent{},
	}
}

func (s KLDMemberCardUpdateMemberTemplateListener) OnMessage(ctx context.Context, message msg.MsgRequest[KLDMemberCardUpdateMemberTemplateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[KLDMemberCardUpdateMemberTemplateEvent]))
}

type KLDMemberCardUpdateMemberTemplateEvent struct {
	MegBody  KLDMemberCardUpdateMemberTemplateMeg_body `json:"meg_body,omitempty"`
	ClientId string                                    `json:"client_id,omitempty"`
	Message  string                                    `json:"message,omitempty"`
}

type KLDMemberCardUpdateMemberTemplateMeg_body struct {
}

// OnKLDMemberCardUpdateMemberTemplateEvent
// @id 325
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=会员卡配置变更
func (s *Service) OnKLDMemberCardUpdateMemberTemplateEvent(handler msg.XinyunEventHandler[KLDMemberCardUpdateMemberTemplateEvent]) (service *Service) {
	var listener = &KLDMemberCardUpdateMemberTemplateListener{handler}
	s.Register(msg.WrapListener[KLDMemberCardUpdateMemberTemplateEvent](listener), "KLD_MemberCard", "updateMemberTemplate")
	return s
}
