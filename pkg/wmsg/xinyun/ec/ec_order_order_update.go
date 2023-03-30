package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcOrderOrderUpdateListener struct {
	handler msg.XinyunEventHandler[EcOrderOrderUpdateEvent]
}

func (s EcOrderOrderUpdateListener) NewMessage() msg.MsgRequest[EcOrderOrderUpdateEvent] {
	return &msg.XinyunOpenMessage[EcOrderOrderUpdateEvent]{
		MsgBody: &EcOrderOrderUpdateEvent{},
	}
}

func (s EcOrderOrderUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcOrderOrderUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcOrderOrderUpdateEvent]))
}

type EcOrderOrderUpdateEvent struct {
	OrderNo     int64  `json:"orderNo,omitempty"`
	OperateType int    `json:"operateType,omitempty"`
	FlagRank    int    `json:"flagRank,omitempty"`
	FlagContent string `json:"flagContent,omitempty"`
	OrderSource int    `json:"orderSource,omitempty"`
	ChannelType int    `json:"channelType,omitempty"`
	OrderType   int    `json:"orderType,omitempty"`
}

// OnEcOrderOrderUpdateEvent
// @id 1707
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=订单更新公共消息
func (s *Service) OnEcOrderOrderUpdateEvent(handler msg.XinyunEventHandler[EcOrderOrderUpdateEvent]) (service *Service) {
	var listener = &EcOrderOrderUpdateListener{handler}
	s.Register(msg.WrapListener[EcOrderOrderUpdateEvent](listener), "ec_order", "orderUpdate")
	return s
}
