package bos

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type BosOrganizationUpdateListener struct {
	handler msg.WosEventHandler[BosOrganizationUpdateEvent]
}

func (s BosOrganizationUpdateListener) NewMessage() msg.MsgRequest[BosOrganizationUpdateEvent] {
	return &msg.WosOpenMessage[BosOrganizationUpdateEvent]{
		MsgBody: &BosOrganizationUpdateEvent{},
	}
}

func (s BosOrganizationUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[BosOrganizationUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[BosOrganizationUpdateEvent]))
}

type BosOrganizationUpdateEvent struct {
	Vid        int64 `json:"vid,omitempty"`
	UpdateType int64 `json:"updateType,omitempty"`
}

// OnBosOrganizationUpdateEvent
// @id 1243
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1243?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新组织基础信息)
func (s *Service) OnBosOrganizationUpdateEvent(handler msg.WosEventHandler[BosOrganizationUpdateEvent]) (service *Service) {
	var listener = &BosOrganizationUpdateListener{handler}
	s.Register(msg.WrapListener[BosOrganizationUpdateEvent](listener), "bos.organization", "update")
	return s
}
