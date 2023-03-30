package o2o

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type O2oServicesOrderChangeOrderStatusListener struct {
	handler msg.XinyunEventHandler[O2oServicesOrderChangeOrderStatusEvent]
}

func (s O2oServicesOrderChangeOrderStatusListener) NewMessage() msg.MsgRequest[O2oServicesOrderChangeOrderStatusEvent] {
	return &msg.XinyunOpenMessage[O2oServicesOrderChangeOrderStatusEvent]{
		MsgBody: &O2oServicesOrderChangeOrderStatusEvent{},
	}
}

func (s O2oServicesOrderChangeOrderStatusListener) OnMessage(ctx context.Context, message msg.MsgRequest[O2oServicesOrderChangeOrderStatusEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[O2oServicesOrderChangeOrderStatusEvent]))
}

type O2oServicesOrderChangeOrderStatusEvent struct {
	OrderNo     string `json:"orderNo,omitempty"`
	OrderType   int    `json:"orderType,omitempty"`
	OrderTime   int64  `json:"orderTime,omitempty"`
	OrderStatus int    `json:"orderStatus,omitempty"`
	PayStatus   int    `json:"payStatus,omitempty"`
	StoreId     int64  `json:"storeId,omitempty"`
}

// OnO2oServicesOrderChangeOrderStatusEvent
// @id 1059
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=订单状态变更
func (s *Service) OnO2oServicesOrderChangeOrderStatusEvent(handler msg.XinyunEventHandler[O2oServicesOrderChangeOrderStatusEvent]) (service *Service) {
	var listener = &O2oServicesOrderChangeOrderStatusListener{handler}
	s.Register(msg.WrapListener[O2oServicesOrderChangeOrderStatusEvent](listener), "o2o_services_order", "changeOrderStatus")
	return s
}
