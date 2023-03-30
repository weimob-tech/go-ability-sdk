package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopGoodsClassifyUpdateListener struct {
	handler msg.WosEventHandler[WeimobShopGoodsClassifyUpdateEvent]
}

func (s WeimobShopGoodsClassifyUpdateListener) NewMessage() msg.MsgRequest[WeimobShopGoodsClassifyUpdateEvent] {
	return &msg.WosOpenMessage[WeimobShopGoodsClassifyUpdateEvent]{
		MsgBody: &WeimobShopGoodsClassifyUpdateEvent{},
	}
}

func (s WeimobShopGoodsClassifyUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopGoodsClassifyUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopGoodsClassifyUpdateEvent]))
}

type WeimobShopGoodsClassifyUpdateEvent struct {
	ClassifyIdList []int64 `json:"classifyIdList,omitempty"`
	Vid            int64   `json:"vid,omitempty"`
}

// OnWeimobShopGoodsClassifyUpdateEvent
// @id 1401
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1401?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新商品分组)
func (s *Service) OnWeimobShopGoodsClassifyUpdateEvent(handler msg.WosEventHandler[WeimobShopGoodsClassifyUpdateEvent]) (service *Service) {
	var listener = &WeimobShopGoodsClassifyUpdateListener{handler}
	s.Register(msg.WrapListener[WeimobShopGoodsClassifyUpdateEvent](listener), "weimob_shop.goods.classify", "update")
	return s
}
