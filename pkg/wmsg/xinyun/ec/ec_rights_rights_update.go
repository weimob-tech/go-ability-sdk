package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcRightsRightsUpdateListener struct {
	handler msg.XinyunEventHandler[EcRightsRightsUpdateEvent]
}

func (s EcRightsRightsUpdateListener) NewMessage() msg.MsgRequest[EcRightsRightsUpdateEvent] {
	return &msg.XinyunOpenMessage[EcRightsRightsUpdateEvent]{
		MsgBody: &EcRightsRightsUpdateEvent{},
	}
}

func (s EcRightsRightsUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcRightsRightsUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcRightsRightsUpdateEvent]))
}

type EcRightsRightsUpdateEvent struct {
	OrderNo     int64  `json:"orderNo,omitempty"`
	RightsId    int64  `json:"rightsId,omitempty"`
	OperateType int    `json:"operateType,omitempty"`
	FlagRank    int    `json:"flagRank,omitempty"`
	FlagContent string `json:"flagContent,omitempty"`
	ChannelType int    `json:"channelType,omitempty"`
	OrderSource int    `json:"orderSource,omitempty"`
	OrderType   int    `json:"orderType,omitempty"`
}

// OnEcRightsRightsUpdateEvent
// @id 1708
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=售后更新公共消息
func (s *Service) OnEcRightsRightsUpdateEvent(handler msg.XinyunEventHandler[EcRightsRightsUpdateEvent]) (service *Service) {
	var listener = &EcRightsRightsUpdateListener{handler}
	s.Register(msg.WrapListener[EcRightsRightsUpdateEvent](listener), "ec_rights", "rightsUpdate")
	return s
}
