package weimob_shopexpress

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopexpressGoodsDeleteListener struct {
	handler msg.WosEventHandler[WeimobShopexpressGoodsDeleteEvent]
}

func (s WeimobShopexpressGoodsDeleteListener) NewMessage() msg.MsgRequest[WeimobShopexpressGoodsDeleteEvent] {
	return &msg.WosOpenMessage[WeimobShopexpressGoodsDeleteEvent]{
		MsgBody: &WeimobShopexpressGoodsDeleteEvent{},
	}
}

func (s WeimobShopexpressGoodsDeleteListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopexpressGoodsDeleteEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopexpressGoodsDeleteEvent]))
}

type WeimobShopexpressGoodsDeleteEvent struct {
	GoodsCodeList []int64 `json:"goodsCodeList,omitempty"`
}

// OnWeimobShopexpressGoodsDeleteEvent
// @id 1376
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1376?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商品删除消息)
func (s *Service) OnWeimobShopexpressGoodsDeleteEvent(handler msg.WosEventHandler[WeimobShopexpressGoodsDeleteEvent]) (service *Service) {
	var listener = &WeimobShopexpressGoodsDeleteListener{handler}
	s.Register(msg.WrapListener[WeimobShopexpressGoodsDeleteEvent](listener), "weimob_shopexpress.goods", "delete")
	return s
}
