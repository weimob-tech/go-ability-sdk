package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopGoodsOnlineStatusUpdateListener struct {
	handler msg.WosEventHandler[WeimobShopGoodsOnlineStatusUpdateEvent]
}

func (s WeimobShopGoodsOnlineStatusUpdateListener) NewMessage() msg.MsgRequest[WeimobShopGoodsOnlineStatusUpdateEvent] {
	return &msg.WosOpenMessage[WeimobShopGoodsOnlineStatusUpdateEvent]{
		MsgBody: &WeimobShopGoodsOnlineStatusUpdateEvent{},
	}
}

func (s WeimobShopGoodsOnlineStatusUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopGoodsOnlineStatusUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopGoodsOnlineStatusUpdateEvent]))
}

type WeimobShopGoodsOnlineStatusUpdateEvent struct {
	IsOnline    bool    `json:"isOnline,omitempty"`
	GoodsIdList []int64 `json:"goodsIdList,omitempty"`
	Vid         int64   `json:"vid,omitempty"`
}

// OnWeimobShopGoodsOnlineStatusUpdateEvent
// @id 1261
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1261?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=修改商品上下架状态)
func (s *Service) OnWeimobShopGoodsOnlineStatusUpdateEvent(handler msg.WosEventHandler[WeimobShopGoodsOnlineStatusUpdateEvent]) (service *Service) {
	var listener = &WeimobShopGoodsOnlineStatusUpdateListener{handler}
	s.Register(msg.WrapListener[WeimobShopGoodsOnlineStatusUpdateEvent](listener), "weimob_shop.goods", "onlineStatusUpdate")
	return s
}
