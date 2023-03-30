package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopGoodsClassifyCreateListener struct {
	handler msg.WosEventHandler[WeimobShopGoodsClassifyCreateEvent]
}

func (s WeimobShopGoodsClassifyCreateListener) NewMessage() msg.MsgRequest[WeimobShopGoodsClassifyCreateEvent] {
	return &msg.WosOpenMessage[WeimobShopGoodsClassifyCreateEvent]{
		MsgBody: &WeimobShopGoodsClassifyCreateEvent{},
	}
}

func (s WeimobShopGoodsClassifyCreateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopGoodsClassifyCreateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopGoodsClassifyCreateEvent]))
}

type WeimobShopGoodsClassifyCreateEvent struct {
	ClassifyId int64 `json:"classifyId,omitempty"`
	Vid        int64 `json:"vid,omitempty"`
}

// OnWeimobShopGoodsClassifyCreateEvent
// @id 1402
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1402?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=新增商品分组)
func (s *Service) OnWeimobShopGoodsClassifyCreateEvent(handler msg.WosEventHandler[WeimobShopGoodsClassifyCreateEvent]) (service *Service) {
	var listener = &WeimobShopGoodsClassifyCreateListener{handler}
	s.Register(msg.WrapListener[WeimobShopGoodsClassifyCreateEvent](listener), "weimob_shop.goods.classify", "create")
	return s
}
