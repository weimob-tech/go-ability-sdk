package weimob_shopexpress

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopexpressGoodsChangeShellListener struct {
	handler msg.WosEventHandler[WeimobShopexpressGoodsChangeShellEvent]
}

func (s WeimobShopexpressGoodsChangeShellListener) NewMessage() msg.MsgRequest[WeimobShopexpressGoodsChangeShellEvent] {
	return &msg.WosOpenMessage[WeimobShopexpressGoodsChangeShellEvent]{
		MsgBody: &WeimobShopexpressGoodsChangeShellEvent{},
	}
}

func (s WeimobShopexpressGoodsChangeShellListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopexpressGoodsChangeShellEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopexpressGoodsChangeShellEvent]))
}

type WeimobShopexpressGoodsChangeShellEvent struct {
	GoodsCode   string `json:"goodsCode,omitempty"`
	GoodsStatus int64  `json:"goodsStatus,omitempty"`
}

// OnWeimobShopexpressGoodsChangeShellEvent
// @id 1335
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1335?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商品上下架变更)
func (s *Service) OnWeimobShopexpressGoodsChangeShellEvent(handler msg.WosEventHandler[WeimobShopexpressGoodsChangeShellEvent]) (service *Service) {
	var listener = &WeimobShopexpressGoodsChangeShellListener{handler}
	s.Register(msg.WrapListener[WeimobShopexpressGoodsChangeShellEvent](listener), "weimob_shopexpress.goods", "changeShell")
	return s
}
