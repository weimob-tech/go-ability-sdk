package hotel

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type HOTELOrderConfirmHotelOrderListener struct {
	handler msg.XinyunEventHandler[HOTELOrderConfirmHotelOrderEvent]
}

func (s HOTELOrderConfirmHotelOrderListener) NewMessage() msg.MsgRequest[HOTELOrderConfirmHotelOrderEvent] {
	return &msg.XinyunOpenMessage[HOTELOrderConfirmHotelOrderEvent]{
		MsgBody: &HOTELOrderConfirmHotelOrderEvent{},
	}
}

func (s HOTELOrderConfirmHotelOrderListener) OnMessage(ctx context.Context, message msg.MsgRequest[HOTELOrderConfirmHotelOrderEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[HOTELOrderConfirmHotelOrderEvent]))
}

type HOTELOrderConfirmHotelOrderEvent struct {
	RoomTypeName string `json:"roomTypeName,omitempty"`
	RoomNumber   string `json:"roomNumber,omitempty"`
	OrderNo      string `json:"orderNo,omitempty"`
}

// OnHOTELOrderConfirmHotelOrderEvent
// @id 483
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=订房确认提醒
func (s *Service) OnHOTELOrderConfirmHotelOrderEvent(handler msg.XinyunEventHandler[HOTELOrderConfirmHotelOrderEvent]) (service *Service) {
	var listener = &HOTELOrderConfirmHotelOrderListener{handler}
	s.Register(msg.WrapListener[HOTELOrderConfirmHotelOrderEvent](listener), "HOTEL_Order", "confirmHotelOrder")
	return s
}
