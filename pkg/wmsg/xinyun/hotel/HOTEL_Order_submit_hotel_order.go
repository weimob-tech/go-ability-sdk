package hotel

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type HOTELOrderSubmitHotelOrderListener struct {
	handler msg.XinyunEventHandler[HOTELOrderSubmitHotelOrderEvent]
}

func (s HOTELOrderSubmitHotelOrderListener) NewMessage() msg.MsgRequest[HOTELOrderSubmitHotelOrderEvent] {
	return &msg.XinyunOpenMessage[HOTELOrderSubmitHotelOrderEvent]{
		MsgBody: &HOTELOrderSubmitHotelOrderEvent{},
	}
}

func (s HOTELOrderSubmitHotelOrderListener) OnMessage(ctx context.Context, message msg.MsgRequest[HOTELOrderSubmitHotelOrderEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[HOTELOrderSubmitHotelOrderEvent]))
}

type HOTELOrderSubmitHotelOrderEvent struct {
	RoomTypeName string `json:"roomTypeName,omitempty"`
	CheckInTime  string `json:"checkInTime,omitempty"`
	CheckOutTime string `json:"checkOutTime,omitempty"`
	RoomCharge   string `json:"roomCharge,omitempty"`
	OrderNo      string `json:"orderNo,omitempty"`
}

// OnHOTELOrderSubmitHotelOrderEvent
// @id 482
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=订房提交成功消息
func (s *Service) OnHOTELOrderSubmitHotelOrderEvent(handler msg.XinyunEventHandler[HOTELOrderSubmitHotelOrderEvent]) (service *Service) {
	var listener = &HOTELOrderSubmitHotelOrderListener{handler}
	s.Register(msg.WrapListener[HOTELOrderSubmitHotelOrderEvent](listener), "HOTEL_Order", "submitHotelOrder")
	return s
}
