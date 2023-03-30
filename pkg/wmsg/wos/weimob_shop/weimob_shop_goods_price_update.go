package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopGoodsPriceUpdateListener struct {
	handler msg.WosEventHandler[WeimobShopGoodsPriceUpdateEvent]
}

func (s WeimobShopGoodsPriceUpdateListener) NewMessage() msg.MsgRequest[WeimobShopGoodsPriceUpdateEvent] {
	return &msg.WosOpenMessage[WeimobShopGoodsPriceUpdateEvent]{
		MsgBody: &WeimobShopGoodsPriceUpdateEvent{},
	}
}

func (s WeimobShopGoodsPriceUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopGoodsPriceUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopGoodsPriceUpdateEvent]))
}

type WeimobShopGoodsPriceUpdateEvent struct {
	GoodsId int64 `json:"goodsId,omitempty"`
	SkuId   int64 `json:"skuId,omitempty"`
	Vid     int64 `json:"vid,omitempty"`
}

// OnWeimobShopGoodsPriceUpdateEvent
// @id 1265
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1265?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商品销售价变更)
func (s *Service) OnWeimobShopGoodsPriceUpdateEvent(handler msg.WosEventHandler[WeimobShopGoodsPriceUpdateEvent]) (service *Service) {
	var listener = &WeimobShopGoodsPriceUpdateListener{handler}
	s.Register(msg.WrapListener[WeimobShopGoodsPriceUpdateEvent](listener), "weimob_shop.goods", "priceUpdate")
	return s
}
