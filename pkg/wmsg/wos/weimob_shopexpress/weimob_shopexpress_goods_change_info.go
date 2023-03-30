package weimob_shopexpress

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopexpressGoodsChangeInfoListener struct {
	handler msg.WosEventHandler[WeimobShopexpressGoodsChangeInfoEvent]
}

func (s WeimobShopexpressGoodsChangeInfoListener) NewMessage() msg.MsgRequest[WeimobShopexpressGoodsChangeInfoEvent] {
	return &msg.WosOpenMessage[WeimobShopexpressGoodsChangeInfoEvent]{
		MsgBody: &WeimobShopexpressGoodsChangeInfoEvent{},
	}
}

func (s WeimobShopexpressGoodsChangeInfoListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopexpressGoodsChangeInfoEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopexpressGoodsChangeInfoEvent]))
}

type WeimobShopexpressGoodsChangeInfoEvent struct {
	GoodsCode string `json:"goodsCode,omitempty"`
}

// OnWeimobShopexpressGoodsChangeInfoEvent
// @id 1336
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1336?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商品信息变更)
func (s *Service) OnWeimobShopexpressGoodsChangeInfoEvent(handler msg.WosEventHandler[WeimobShopexpressGoodsChangeInfoEvent]) (service *Service) {
	var listener = &WeimobShopexpressGoodsChangeInfoListener{handler}
	s.Register(msg.WrapListener[WeimobShopexpressGoodsChangeInfoEvent](listener), "weimob_shopexpress.goods", "changeInfo")
	return s
}
