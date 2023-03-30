package o2o

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type O2oOrderMenuAddListener struct {
	handler msg.XinyunEventHandler[O2oOrderMenuAddEvent]
}

func (s O2oOrderMenuAddListener) NewMessage() msg.MsgRequest[O2oOrderMenuAddEvent] {
	return &msg.XinyunOpenMessage[O2oOrderMenuAddEvent]{
		MsgBody: &O2oOrderMenuAddEvent{},
	}
}

func (s O2oOrderMenuAddListener) OnMessage(ctx context.Context, message msg.MsgRequest[O2oOrderMenuAddEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[O2oOrderMenuAddEvent]))
}

type O2oOrderMenuAddEvent struct {
	MenuItems      []map[string]any `json:"menuItems,omitempty"`
	StoreId        int64            `json:"storeId,omitempty"`
	OrderId        int64            `json:"orderId,omitempty"`
	OrderNo        string           `json:"orderNo,omitempty"`
	OrderStatus    int              `json:"orderStatus,omitempty"`
	UpdateTime     int64            `json:"updateTime,omitempty"`
	SerialNo       int              `json:"serialNo,omitempty"`
	MenuStatus     int              `json:"menuStatus,omitempty"`
	StoreChannelId string           `json:"storeChannelId,omitempty"`
}

// OnO2oOrderMenuAddEvent
// @id 1298
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=加菜创建
func (s *Service) OnO2oOrderMenuAddEvent(handler msg.XinyunEventHandler[O2oOrderMenuAddEvent]) (service *Service) {
	var listener = &O2oOrderMenuAddListener{handler}
	s.Register(msg.WrapListener[O2oOrderMenuAddEvent](listener), "o2o_order_menu", "add")
	return s
}
