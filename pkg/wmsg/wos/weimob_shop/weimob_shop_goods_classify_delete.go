package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopGoodsClassifyDeleteListener struct {
	handler msg.WosEventHandler[WeimobShopGoodsClassifyDeleteEvent]
}

func (s WeimobShopGoodsClassifyDeleteListener) NewMessage() msg.MsgRequest[WeimobShopGoodsClassifyDeleteEvent] {
	return &msg.WosOpenMessage[WeimobShopGoodsClassifyDeleteEvent]{
		MsgBody: &WeimobShopGoodsClassifyDeleteEvent{},
	}
}

func (s WeimobShopGoodsClassifyDeleteListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopGoodsClassifyDeleteEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopGoodsClassifyDeleteEvent]))
}

type WeimobShopGoodsClassifyDeleteEvent struct {
	ClassifyIdList []int64 `json:"classifyIdList,omitempty"`
	Vid            int64   `json:"vid,omitempty"`
}

// OnWeimobShopGoodsClassifyDeleteEvent
// @id 1400
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1400?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除商品分组)
func (s *Service) OnWeimobShopGoodsClassifyDeleteEvent(handler msg.WosEventHandler[WeimobShopGoodsClassifyDeleteEvent]) (service *Service) {
	var listener = &WeimobShopGoodsClassifyDeleteListener{handler}
	s.Register(msg.WrapListener[WeimobShopGoodsClassifyDeleteEvent](listener), "weimob_shop.goods.classify", "delete")
	return s
}
