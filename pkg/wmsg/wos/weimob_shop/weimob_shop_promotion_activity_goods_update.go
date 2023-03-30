package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopPromotionActivityGoodsUpdateListener struct {
	handler msg.WosEventHandler[WeimobShopPromotionActivityGoodsUpdateEvent]
}

func (s WeimobShopPromotionActivityGoodsUpdateListener) NewMessage() msg.MsgRequest[WeimobShopPromotionActivityGoodsUpdateEvent] {
	return &msg.WosOpenMessage[WeimobShopPromotionActivityGoodsUpdateEvent]{
		MsgBody: &WeimobShopPromotionActivityGoodsUpdateEvent{},
	}
}

func (s WeimobShopPromotionActivityGoodsUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopPromotionActivityGoodsUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopPromotionActivityGoodsUpdateEvent]))
}

type WeimobShopPromotionActivityGoodsUpdateEvent struct {
	ActivityId   int64   `json:"activityId,omitempty"`
	ActivityType int64   `json:"activityType,omitempty"`
	GoodsIdList  []int64 `json:"goodsIdList,omitempty"`
}

// OnWeimobShopPromotionActivityGoodsUpdateEvent
// @id 1332
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1332?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=活动商品修改)
func (s *Service) OnWeimobShopPromotionActivityGoodsUpdateEvent(handler msg.WosEventHandler[WeimobShopPromotionActivityGoodsUpdateEvent]) (service *Service) {
	var listener = &WeimobShopPromotionActivityGoodsUpdateListener{handler}
	s.Register(msg.WrapListener[WeimobShopPromotionActivityGoodsUpdateEvent](listener), "weimob_shop.promotion", "activityGoodsUpdate")
	return s
}
