package hotel

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type HOTELOrderHotelOrderFailListener struct {
	handler msg.XinyunEventHandler[HOTELOrderHotelOrderFailEvent]
}

func (s HOTELOrderHotelOrderFailListener) NewMessage() msg.MsgRequest[HOTELOrderHotelOrderFailEvent] {
	return &msg.XinyunOpenMessage[HOTELOrderHotelOrderFailEvent]{
		MsgBody: &HOTELOrderHotelOrderFailEvent{},
	}
}

func (s HOTELOrderHotelOrderFailListener) OnMessage(ctx context.Context, message msg.MsgRequest[HOTELOrderHotelOrderFailEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[HOTELOrderHotelOrderFailEvent]))
}

type HOTELOrderHotelOrderFailEvent struct {
	RoomTypeName string `json:"roomTypeName,omitempty"`
	RoomNumber   string `json:"roomNumber,omitempty"`
	OrderNo      string `json:"orderNo,omitempty"`
}

// OnHOTELOrderHotelOrderFailEvent
// @id 484
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=订房失败提醒
func (s *Service) OnHOTELOrderHotelOrderFailEvent(handler msg.XinyunEventHandler[HOTELOrderHotelOrderFailEvent]) (service *Service) {
	var listener = &HOTELOrderHotelOrderFailListener{handler}
	s.Register(msg.WrapListener[HOTELOrderHotelOrderFailEvent](listener), "HOTEL_Order", "hotelOrderFail")
	return s
}
