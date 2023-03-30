package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcOrderOrderStatusChangeListener struct {
	handler msg.XinyunEventHandler[EcOrderOrderStatusChangeEvent]
}

func (s EcOrderOrderStatusChangeListener) NewMessage() msg.MsgRequest[EcOrderOrderStatusChangeEvent] {
	return &msg.XinyunOpenMessage[EcOrderOrderStatusChangeEvent]{
		MsgBody: &EcOrderOrderStatusChangeEvent{},
	}
}

func (s EcOrderOrderStatusChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcOrderOrderStatusChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcOrderOrderStatusChangeEvent]))
}

type EcOrderOrderStatusChangeEvent struct {
	OrderNo        int64 `json:"orderNo,omitempty"`
	OrderStatus    int   `json:"orderStatus,omitempty"`
	EnableDelivery int   `json:"enableDelivery,omitempty"`
	Datetime       int   `json:"datetime,omitempty"`
	StoreId        int64 `json:"storeId,omitempty"`
	OrderSource    int   `json:"orderSource,omitempty"`
	ChannelType    int   `json:"channelType,omitempty"`
	OrderType      int   `json:"orderType,omitempty"`
}

// OnEcOrderOrderStatusChangeEvent
// @id 372
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=订单状态变更消息
func (s *Service) OnEcOrderOrderStatusChangeEvent(handler msg.XinyunEventHandler[EcOrderOrderStatusChangeEvent]) (service *Service) {
	var listener = &EcOrderOrderStatusChangeListener{handler}
	s.Register(msg.WrapListener[EcOrderOrderStatusChangeEvent](listener), "ec_order", "orderStatusChange")
	return s
}
