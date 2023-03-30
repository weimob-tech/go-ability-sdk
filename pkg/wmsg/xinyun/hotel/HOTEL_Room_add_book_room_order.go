package hotel

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type HOTELRoomAddBookRoomOrderListener struct {
	handler msg.XinyunEventHandler[HOTELRoomAddBookRoomOrderEvent]
}

func (s HOTELRoomAddBookRoomOrderListener) NewMessage() msg.MsgRequest[HOTELRoomAddBookRoomOrderEvent] {
	return &msg.XinyunOpenMessage[HOTELRoomAddBookRoomOrderEvent]{
		MsgBody: &HOTELRoomAddBookRoomOrderEvent{},
	}
}

func (s HOTELRoomAddBookRoomOrderListener) OnMessage(ctx context.Context, message msg.MsgRequest[HOTELRoomAddBookRoomOrderEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[HOTELRoomAddBookRoomOrderEvent]))
}

type HOTELRoomAddBookRoomOrderEvent struct {
	OrderNo string `json:"orderNo,omitempty"`
	Pid     string `json:"pid,omitempty"`
	StoreId string `json:"storeId,omitempty"`
}

// OnHOTELRoomAddBookRoomOrderEvent
// @id 534
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=办理预约提醒
func (s *Service) OnHOTELRoomAddBookRoomOrderEvent(handler msg.XinyunEventHandler[HOTELRoomAddBookRoomOrderEvent]) (service *Service) {
	var listener = &HOTELRoomAddBookRoomOrderListener{handler}
	s.Register(msg.WrapListener[HOTELRoomAddBookRoomOrderEvent](listener), "HOTEL_Room", "addBookRoomOrder")
	return s
}
