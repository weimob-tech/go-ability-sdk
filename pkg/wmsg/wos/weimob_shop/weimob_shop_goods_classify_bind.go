package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopGoodsClassifyBindListener struct {
	handler msg.WosEventHandler[WeimobShopGoodsClassifyBindEvent]
}

func (s WeimobShopGoodsClassifyBindListener) NewMessage() msg.MsgRequest[WeimobShopGoodsClassifyBindEvent] {
	return &msg.WosOpenMessage[WeimobShopGoodsClassifyBindEvent]{
		MsgBody: &WeimobShopGoodsClassifyBindEvent{},
	}
}

func (s WeimobShopGoodsClassifyBindListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopGoodsClassifyBindEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopGoodsClassifyBindEvent]))
}

type WeimobShopGoodsClassifyBindEvent struct {
	GoodsId int64 `json:"goodsId,omitempty"`
	Vid     int64 `json:"vid,omitempty"`
}

// OnWeimobShopGoodsClassifyBindEvent
// @id 1399
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1399?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商品分组关联关系变更)
func (s *Service) OnWeimobShopGoodsClassifyBindEvent(handler msg.WosEventHandler[WeimobShopGoodsClassifyBindEvent]) (service *Service) {
	var listener = &WeimobShopGoodsClassifyBindListener{handler}
	s.Register(msg.WrapListener[WeimobShopGoodsClassifyBindEvent](listener), "weimob_shop.goods.classify", "bind")
	return s
}
