package weimob_shopexpress

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopexpressOrderChangeStatusListener struct {
	handler msg.WosEventHandler[WeimobShopexpressOrderChangeStatusEvent]
}

func (s WeimobShopexpressOrderChangeStatusListener) NewMessage() msg.MsgRequest[WeimobShopexpressOrderChangeStatusEvent] {
	return &msg.WosOpenMessage[WeimobShopexpressOrderChangeStatusEvent]{
		MsgBody: &WeimobShopexpressOrderChangeStatusEvent{},
	}
}

func (s WeimobShopexpressOrderChangeStatusListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopexpressOrderChangeStatusEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopexpressOrderChangeStatusEvent]))
}

type WeimobShopexpressOrderChangeStatusEvent struct {
	OrderNo         string `json:"orderNo,omitempty"`
	OrderStatus     string `json:"orderStatus,omitempty"`
	OrderStatusText string `json:"orderStatusText,omitempty"`
}

// OnWeimobShopexpressOrderChangeStatusEvent
// @id 1339
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1339?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=订单状态发生变更)
func (s *Service) OnWeimobShopexpressOrderChangeStatusEvent(handler msg.WosEventHandler[WeimobShopexpressOrderChangeStatusEvent]) (service *Service) {
	var listener = &WeimobShopexpressOrderChangeStatusListener{handler}
	s.Register(msg.WrapListener[WeimobShopexpressOrderChangeStatusEvent](listener), "weimob_shopexpress.order", "changeStatus")
	return s
}
