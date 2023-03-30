package hotel

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type HOTELRoomCheckInHotelListener struct {
	handler msg.XinyunEventHandler[HOTELRoomCheckInHotelEvent]
}

func (s HOTELRoomCheckInHotelListener) NewMessage() msg.MsgRequest[HOTELRoomCheckInHotelEvent] {
	return &msg.XinyunOpenMessage[HOTELRoomCheckInHotelEvent]{
		MsgBody: &HOTELRoomCheckInHotelEvent{},
	}
}

func (s HOTELRoomCheckInHotelListener) OnMessage(ctx context.Context, message msg.MsgRequest[HOTELRoomCheckInHotelEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[HOTELRoomCheckInHotelEvent]))
}

type HOTELRoomCheckInHotelEvent struct {
	OrderNo string `json:"orderNo,omitempty"`
	Pid     string `json:"pid,omitempty"`
	StoreId string `json:"storeId,omitempty"`
}

// OnHOTELRoomCheckInHotelEvent
// @id 486
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=酒店入住提醒
func (s *Service) OnHOTELRoomCheckInHotelEvent(handler msg.XinyunEventHandler[HOTELRoomCheckInHotelEvent]) (service *Service) {
	var listener = &HOTELRoomCheckInHotelListener{handler}
	s.Register(msg.WrapListener[HOTELRoomCheckInHotelEvent](listener), "HOTEL_Room", "checkInHotel")
	return s
}
