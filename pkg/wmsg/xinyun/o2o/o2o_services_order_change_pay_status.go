package o2o

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type O2oServicesOrderChangePayStatusListener struct {
	handler msg.XinyunEventHandler[O2oServicesOrderChangePayStatusEvent]
}

func (s O2oServicesOrderChangePayStatusListener) NewMessage() msg.MsgRequest[O2oServicesOrderChangePayStatusEvent] {
	return &msg.XinyunOpenMessage[O2oServicesOrderChangePayStatusEvent]{
		MsgBody: &O2oServicesOrderChangePayStatusEvent{},
	}
}

func (s O2oServicesOrderChangePayStatusListener) OnMessage(ctx context.Context, message msg.MsgRequest[O2oServicesOrderChangePayStatusEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[O2oServicesOrderChangePayStatusEvent]))
}

type O2oServicesOrderChangePayStatusEvent struct {
	OrderNo     string `json:"orderNo,omitempty"`
	OrderType   int    `json:"orderType,omitempty"`
	OrderTime   int64  `json:"orderTime,omitempty"`
	OrderStatus int    `json:"orderStatus,omitempty"`
	PayStatus   int    `json:"payStatus,omitempty"`
	StoreId     int64  `json:"storeId,omitempty"`
}

// OnO2oServicesOrderChangePayStatusEvent
// @id 1061
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=支付状态变更
func (s *Service) OnO2oServicesOrderChangePayStatusEvent(handler msg.XinyunEventHandler[O2oServicesOrderChangePayStatusEvent]) (service *Service) {
	var listener = &O2oServicesOrderChangePayStatusListener{handler}
	s.Register(msg.WrapListener[O2oServicesOrderChangePayStatusEvent](listener), "o2o_services_order", "changePayStatus")
	return s
}
