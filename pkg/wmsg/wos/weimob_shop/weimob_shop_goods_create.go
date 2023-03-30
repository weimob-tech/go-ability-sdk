package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopGoodsCreateListener struct {
	handler msg.WosEventHandler[WeimobShopGoodsCreateEvent]
}

func (s WeimobShopGoodsCreateListener) NewMessage() msg.MsgRequest[WeimobShopGoodsCreateEvent] {
	return &msg.WosOpenMessage[WeimobShopGoodsCreateEvent]{
		MsgBody: &WeimobShopGoodsCreateEvent{},
	}
}

func (s WeimobShopGoodsCreateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopGoodsCreateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopGoodsCreateEvent]))
}

type WeimobShopGoodsCreateEvent struct {
	Vid     string `json:"vid,omitempty"`
	GoodsId string `json:"goodsId,omitempty"`
}

// OnWeimobShopGoodsCreateEvent
// @id 1325
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1325?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=创建商品消息)
func (s *Service) OnWeimobShopGoodsCreateEvent(handler msg.WosEventHandler[WeimobShopGoodsCreateEvent]) (service *Service) {
	var listener = &WeimobShopGoodsCreateListener{handler}
	s.Register(msg.WrapListener[WeimobShopGoodsCreateEvent](listener), "weimob_shop.goods", "create")
	return s
}
