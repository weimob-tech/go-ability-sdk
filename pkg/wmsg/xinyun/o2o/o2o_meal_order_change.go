package o2o

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type O2oMealOrderChangeListener struct {
	handler msg.XinyunEventHandler[O2oMealOrderChangeEvent]
}

func (s O2oMealOrderChangeListener) NewMessage() msg.MsgRequest[O2oMealOrderChangeEvent] {
	return &msg.XinyunOpenMessage[O2oMealOrderChangeEvent]{
		MsgBody: &O2oMealOrderChangeEvent{},
	}
}

func (s O2oMealOrderChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[O2oMealOrderChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[O2oMealOrderChangeEvent]))
}

type O2oMealOrderChangeEvent struct {
	StoreId        int64  `json:"storeId,omitempty"`
	OrderId        int64  `json:"orderId,omitempty"`
	OrderNo        string `json:"orderNo,omitempty"`
	OrderStatus    int    `json:"orderStatus,omitempty"`
	UpdateTime     int64  `json:"updateTime,omitempty"`
	StoreChannelId string `json:"storeChannelId,omitempty"`
}

// OnO2oMealOrderChangeEvent
// @id 1296
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=订单更新
func (s *Service) OnO2oMealOrderChangeEvent(handler msg.XinyunEventHandler[O2oMealOrderChangeEvent]) (service *Service) {
	var listener = &O2oMealOrderChangeListener{handler}
	s.Register(msg.WrapListener[O2oMealOrderChangeEvent](listener), "o2o_meal_order", "change")
	return s
}
