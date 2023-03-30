package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopGoodsUpdateListener struct {
	handler msg.WosEventHandler[WeimobShopGoodsUpdateEvent]
}

func (s WeimobShopGoodsUpdateListener) NewMessage() msg.MsgRequest[WeimobShopGoodsUpdateEvent] {
	return &msg.WosOpenMessage[WeimobShopGoodsUpdateEvent]{
		MsgBody: &WeimobShopGoodsUpdateEvent{},
	}
}

func (s WeimobShopGoodsUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopGoodsUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopGoodsUpdateEvent]))
}

type WeimobShopGoodsUpdateEvent struct {
	GoodsId int64 `json:"goodsId,omitempty"`
	Vid     int64 `json:"vid,omitempty"`
}

// OnWeimobShopGoodsUpdateEvent
// @id 1263
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1263?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=修改商品)
func (s *Service) OnWeimobShopGoodsUpdateEvent(handler msg.WosEventHandler[WeimobShopGoodsUpdateEvent]) (service *Service) {
	var listener = &WeimobShopGoodsUpdateListener{handler}
	s.Register(msg.WrapListener[WeimobShopGoodsUpdateEvent](listener), "weimob_shop.goods", "update")
	return s
}
