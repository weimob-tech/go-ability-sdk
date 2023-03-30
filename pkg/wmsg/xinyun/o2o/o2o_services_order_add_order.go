package o2o

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type O2oServicesOrderAddOrderListener struct {
	handler msg.XinyunEventHandler[O2oServicesOrderAddOrderEvent]
}

func (s O2oServicesOrderAddOrderListener) NewMessage() msg.MsgRequest[O2oServicesOrderAddOrderEvent] {
	return &msg.XinyunOpenMessage[O2oServicesOrderAddOrderEvent]{
		MsgBody: &O2oServicesOrderAddOrderEvent{},
	}
}

func (s O2oServicesOrderAddOrderListener) OnMessage(ctx context.Context, message msg.MsgRequest[O2oServicesOrderAddOrderEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[O2oServicesOrderAddOrderEvent]))
}

type O2oServicesOrderAddOrderEvent struct {
	OrderNo     string `json:"orderNo,omitempty"`
	OrderType   int    `json:"orderType,omitempty"`
	OrderTime   int64  `json:"orderTime,omitempty"`
	OrderStatus int    `json:"orderStatus,omitempty"`
	PayStatus   int    `json:"payStatus,omitempty"`
	StoreId     int64  `json:"storeId,omitempty"`
}

// OnO2oServicesOrderAddOrderEvent
// @id 1056
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=下单消息
func (s *Service) OnO2oServicesOrderAddOrderEvent(handler msg.XinyunEventHandler[O2oServicesOrderAddOrderEvent]) (service *Service) {
	var listener = &O2oServicesOrderAddOrderListener{handler}
	s.Register(msg.WrapListener[O2oServicesOrderAddOrderEvent](listener), "o2o_services_order", "addOrder")
	return s
}
