package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcOrderUpdateCycleDeliveryAddressListener struct {
	handler msg.XinyunEventHandler[EcOrderUpdateCycleDeliveryAddressEvent]
}

func (s EcOrderUpdateCycleDeliveryAddressListener) NewMessage() msg.MsgRequest[EcOrderUpdateCycleDeliveryAddressEvent] {
	return &msg.XinyunOpenMessage[EcOrderUpdateCycleDeliveryAddressEvent]{
		MsgBody: &EcOrderUpdateCycleDeliveryAddressEvent{},
	}
}

func (s EcOrderUpdateCycleDeliveryAddressListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcOrderUpdateCycleDeliveryAddressEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcOrderUpdateCycleDeliveryAddressEvent]))
}

type EcOrderUpdateCycleDeliveryAddressEvent struct {
	OrderNo           int64  `json:"orderNo,omitempty"`
	Pid               int64  `json:"pid,omitempty"`
	OrderStatus       int    `json:"orderStatus,omitempty"`
	EnableDelivery    string `json:"enableDelivery,omitempty"`
	ReceiverProvince  string `json:"receiverProvince,omitempty"`
	ReceiverCity      string `json:"receiverCity,omitempty"`
	ReceiverCounty    string `json:"receiverCounty,omitempty"`
	ReceiverArea      string `json:"receiverArea,omitempty"`
	ReceiverAddress   string `json:"receiverAddress,omitempty"`
	ReceiverLongitude string `json:"receiverLongitude,omitempty"`
	ReceiverLatitude  string `json:"receiverLatitude,omitempty"`
	ReceiverName      string `json:"receiverName,omitempty"`
	ReceiverMobile    string `json:"receiverMobile,omitempty"`
	ReceiverZip       string `json:"receiverZip,omitempty"`
	Datetime          string `json:"datetime,omitempty"`
}

// OnEcOrderUpdateCycleDeliveryAddressEvent
// @id 1610
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=周期购修改配送地址消息
func (s *Service) OnEcOrderUpdateCycleDeliveryAddressEvent(handler msg.XinyunEventHandler[EcOrderUpdateCycleDeliveryAddressEvent]) (service *Service) {
	var listener = &EcOrderUpdateCycleDeliveryAddressListener{handler}
	s.Register(msg.WrapListener[EcOrderUpdateCycleDeliveryAddressEvent](listener), "ec_order", "updateCycleDeliveryAddress")
	return s
}
