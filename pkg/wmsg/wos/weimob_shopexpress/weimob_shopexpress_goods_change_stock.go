package weimob_shopexpress

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopexpressGoodsChangeStockListener struct {
	handler msg.WosEventHandler[WeimobShopexpressGoodsChangeStockEvent]
}

func (s WeimobShopexpressGoodsChangeStockListener) NewMessage() msg.MsgRequest[WeimobShopexpressGoodsChangeStockEvent] {
	return &msg.WosOpenMessage[WeimobShopexpressGoodsChangeStockEvent]{
		MsgBody: &WeimobShopexpressGoodsChangeStockEvent{},
	}
}

func (s WeimobShopexpressGoodsChangeStockListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopexpressGoodsChangeStockEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopexpressGoodsChangeStockEvent]))
}

type WeimobShopexpressGoodsChangeStockEvent struct {
	GoodsCode string `json:"goodsCode,omitempty"`
	SkuCode   string `json:"skuCode,omitempty"`
}

// OnWeimobShopexpressGoodsChangeStockEvent
// @id 1375
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1375?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商品库存更新消息)
func (s *Service) OnWeimobShopexpressGoodsChangeStockEvent(handler msg.WosEventHandler[WeimobShopexpressGoodsChangeStockEvent]) (service *Service) {
	var listener = &WeimobShopexpressGoodsChangeStockListener{handler}
	s.Register(msg.WrapListener[WeimobShopexpressGoodsChangeStockEvent](listener), "weimob_shopexpress.goods", "changeStock")
	return s
}
