package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopGoodsDeleteListener struct {
	handler msg.WosEventHandler[WeimobShopGoodsDeleteEvent]
}

func (s WeimobShopGoodsDeleteListener) NewMessage() msg.MsgRequest[WeimobShopGoodsDeleteEvent] {
	return &msg.WosOpenMessage[WeimobShopGoodsDeleteEvent]{
		MsgBody: &WeimobShopGoodsDeleteEvent{},
	}
}

func (s WeimobShopGoodsDeleteListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopGoodsDeleteEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopGoodsDeleteEvent]))
}

type WeimobShopGoodsDeleteEvent struct {
	GoodsIdList []int64 `json:"goodsIdList,omitempty"`
	Vid         int64   `json:"vid,omitempty"`
}

// OnWeimobShopGoodsDeleteEvent
// @id 1264
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1264?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除商品)
func (s *Service) OnWeimobShopGoodsDeleteEvent(handler msg.WosEventHandler[WeimobShopGoodsDeleteEvent]) (service *Service) {
	var listener = &WeimobShopGoodsDeleteListener{handler}
	s.Register(msg.WrapListener[WeimobShopGoodsDeleteEvent](listener), "weimob_shop.goods", "delete")
	return s
}
