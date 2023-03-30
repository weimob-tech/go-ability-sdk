package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type MallRightsMallRightsMessageListener struct {
	handler msg.WosEventHandler[MallRightsMallRightsMessageEvent]
}

func (s MallRightsMallRightsMessageListener) NewMessage() msg.MsgRequest[MallRightsMallRightsMessageEvent] {
	return &msg.WosOpenMessage[MallRightsMallRightsMessageEvent]{
		MsgBody: &MallRightsMallRightsMessageEvent{},
	}
}

func (s MallRightsMallRightsMessageListener) OnMessage(ctx context.Context, message msg.MsgRequest[MallRightsMallRightsMessageEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[MallRightsMallRightsMessageEvent]))
}

type MallRightsMallRightsMessageEvent struct {
	Obj MallRightsMallRightsMessageObj `json:"obj,omitempty"`
}

type MallRightsMallRightsMessageObj struct {
}

// OnMallRightsMallRightsMessageEvent
// @id 1195
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1195?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商城售后消息)
func (s *Service) OnMallRightsMallRightsMessageEvent(handler msg.WosEventHandler[MallRightsMallRightsMessageEvent]) (service *Service) {
	var listener = &MallRightsMallRightsMessageListener{handler}
	s.Register(msg.WrapListener[MallRightsMallRightsMessageEvent](listener), "mall_rights", "mall_rights_message")
	return s
}
