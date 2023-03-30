package kld

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type KLDMemberCardOpenMemberCardListener struct {
	handler msg.XinyunEventHandler[KLDMemberCardOpenMemberCardEvent]
}

func (s KLDMemberCardOpenMemberCardListener) NewMessage() msg.MsgRequest[KLDMemberCardOpenMemberCardEvent] {
	return &msg.XinyunOpenMessage[KLDMemberCardOpenMemberCardEvent]{
		MsgBody: &KLDMemberCardOpenMemberCardEvent{},
	}
}

func (s KLDMemberCardOpenMemberCardListener) OnMessage(ctx context.Context, message msg.MsgRequest[KLDMemberCardOpenMemberCardEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[KLDMemberCardOpenMemberCardEvent]))
}

type KLDMemberCardOpenMemberCardEvent struct {
	MegBody      KLDMemberCardOpenMemberCardMeg_body `json:"meg_body,omitempty"`
	ClientId     string                              `json:"client_id,omitempty"`
	MemberCardNo string                              `json:"MemberCardNo,omitempty"`
	Phone        string                              `json:"Phone,omitempty"`
	Name         string                              `json:"Name,omitempty"`
}

type KLDMemberCardOpenMemberCardMeg_body struct {
}

// OnKLDMemberCardOpenMemberCardEvent
// @id 322
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=会员卡开卡订阅消息
func (s *Service) OnKLDMemberCardOpenMemberCardEvent(handler msg.XinyunEventHandler[KLDMemberCardOpenMemberCardEvent]) (service *Service) {
	var listener = &KLDMemberCardOpenMemberCardListener{handler}
	s.Register(msg.WrapListener[KLDMemberCardOpenMemberCardEvent](listener), "KLD_MemberCard", "openMemberCard")
	return s
}
