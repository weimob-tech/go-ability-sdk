package kld

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type KLDMemberCardUpdateMemberCardAccountInfoListener struct {
	handler msg.XinyunEventHandler[KLDMemberCardUpdateMemberCardAccountInfoEvent]
}

func (s KLDMemberCardUpdateMemberCardAccountInfoListener) NewMessage() msg.MsgRequest[KLDMemberCardUpdateMemberCardAccountInfoEvent] {
	return &msg.XinyunOpenMessage[KLDMemberCardUpdateMemberCardAccountInfoEvent]{
		MsgBody: &KLDMemberCardUpdateMemberCardAccountInfoEvent{},
	}
}

func (s KLDMemberCardUpdateMemberCardAccountInfoListener) OnMessage(ctx context.Context, message msg.MsgRequest[KLDMemberCardUpdateMemberCardAccountInfoEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[KLDMemberCardUpdateMemberCardAccountInfoEvent]))
}

type KLDMemberCardUpdateMemberCardAccountInfoEvent struct {
	MegBody      KLDMemberCardUpdateMemberCardAccountInfoMeg_body `json:"meg_body,omitempty"`
	ClientId     string                                           `json:"client_id,omitempty"`
	MemberCardNo string                                           `json:"MemberCardNo,omitempty"`
	Message      string                                           `json:"Message,omitempty"`
}

type KLDMemberCardUpdateMemberCardAccountInfoMeg_body struct {
}

// OnKLDMemberCardUpdateMemberCardAccountInfoEvent
// @id 323
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=会员卡账户变更消息
func (s *Service) OnKLDMemberCardUpdateMemberCardAccountInfoEvent(handler msg.XinyunEventHandler[KLDMemberCardUpdateMemberCardAccountInfoEvent]) (service *Service) {
	var listener = &KLDMemberCardUpdateMemberCardAccountInfoListener{handler}
	s.Register(msg.WrapListener[KLDMemberCardUpdateMemberCardAccountInfoEvent](listener), "KLD_MemberCard", "updateMemberCardAccountInfo")
	return s
}
