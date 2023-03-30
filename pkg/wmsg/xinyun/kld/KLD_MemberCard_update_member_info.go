package kld

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type KLDMemberCardUpdateMemberInfoListener struct {
	handler msg.XinyunEventHandler[KLDMemberCardUpdateMemberInfoEvent]
}

func (s KLDMemberCardUpdateMemberInfoListener) NewMessage() msg.MsgRequest[KLDMemberCardUpdateMemberInfoEvent] {
	return &msg.XinyunOpenMessage[KLDMemberCardUpdateMemberInfoEvent]{
		MsgBody: &KLDMemberCardUpdateMemberInfoEvent{},
	}
}

func (s KLDMemberCardUpdateMemberInfoListener) OnMessage(ctx context.Context, message msg.MsgRequest[KLDMemberCardUpdateMemberInfoEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[KLDMemberCardUpdateMemberInfoEvent]))
}

type KLDMemberCardUpdateMemberInfoEvent struct {
	MegBody      KLDMemberCardUpdateMemberInfoMeg_body `json:"meg_body,omitempty"`
	ClientId     string                                `json:"client_id,omitempty"`
	MemberCardNo string                                `json:"MemberCardNo,omitempty"`
}

type KLDMemberCardUpdateMemberInfoMeg_body struct {
}

// OnKLDMemberCardUpdateMemberInfoEvent
// @id 324
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=会员基本信息变更
func (s *Service) OnKLDMemberCardUpdateMemberInfoEvent(handler msg.XinyunEventHandler[KLDMemberCardUpdateMemberInfoEvent]) (service *Service) {
	var listener = &KLDMemberCardUpdateMemberInfoListener{handler}
	s.Register(msg.WrapListener[KLDMemberCardUpdateMemberInfoEvent](listener), "KLD_MemberCard", "updateMemberInfo")
	return s
}
