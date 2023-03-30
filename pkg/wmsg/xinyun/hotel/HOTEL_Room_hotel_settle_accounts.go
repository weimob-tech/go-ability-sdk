package hotel

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type HOTELRoomHotelSettleAccountsListener struct {
	handler msg.XinyunEventHandler[HOTELRoomHotelSettleAccountsEvent]
}

func (s HOTELRoomHotelSettleAccountsListener) NewMessage() msg.MsgRequest[HOTELRoomHotelSettleAccountsEvent] {
	return &msg.XinyunOpenMessage[HOTELRoomHotelSettleAccountsEvent]{
		MsgBody: &HOTELRoomHotelSettleAccountsEvent{},
	}
}

func (s HOTELRoomHotelSettleAccountsListener) OnMessage(ctx context.Context, message msg.MsgRequest[HOTELRoomHotelSettleAccountsEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[HOTELRoomHotelSettleAccountsEvent]))
}

type HOTELRoomHotelSettleAccountsEvent struct {
	OrderNo string `json:"orderNo,omitempty"`
	Pid     string `json:"pid,omitempty"`
	StoreId string `json:"storeId,omitempty"`
}

// OnHOTELRoomHotelSettleAccountsEvent
// @id 1778
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=酒店住宿结算
func (s *Service) OnHOTELRoomHotelSettleAccountsEvent(handler msg.XinyunEventHandler[HOTELRoomHotelSettleAccountsEvent]) (service *Service) {
	var listener = &HOTELRoomHotelSettleAccountsListener{handler}
	s.Register(msg.WrapListener[HOTELRoomHotelSettleAccountsEvent](listener), "HOTEL_Room", "hotelSettleAccounts")
	return s
}
