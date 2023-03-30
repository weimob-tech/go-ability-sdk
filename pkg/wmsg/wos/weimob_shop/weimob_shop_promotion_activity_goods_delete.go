package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopPromotionActivityGoodsDeleteListener struct {
	handler msg.WosEventHandler[WeimobShopPromotionActivityGoodsDeleteEvent]
}

func (s WeimobShopPromotionActivityGoodsDeleteListener) NewMessage() msg.MsgRequest[WeimobShopPromotionActivityGoodsDeleteEvent] {
	return &msg.WosOpenMessage[WeimobShopPromotionActivityGoodsDeleteEvent]{
		MsgBody: &WeimobShopPromotionActivityGoodsDeleteEvent{},
	}
}

func (s WeimobShopPromotionActivityGoodsDeleteListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopPromotionActivityGoodsDeleteEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopPromotionActivityGoodsDeleteEvent]))
}

type WeimobShopPromotionActivityGoodsDeleteEvent struct {
	ActivityId   int64   `json:"activityId,omitempty"`
	ActivityType int64   `json:"activityType,omitempty"`
	GoodsIdList  []int64 `json:"goodsIdList,omitempty"`
}

// OnWeimobShopPromotionActivityGoodsDeleteEvent
// @id 1333
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1333?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=活动商品删除)
func (s *Service) OnWeimobShopPromotionActivityGoodsDeleteEvent(handler msg.WosEventHandler[WeimobShopPromotionActivityGoodsDeleteEvent]) (service *Service) {
	var listener = &WeimobShopPromotionActivityGoodsDeleteListener{handler}
	s.Register(msg.WrapListener[WeimobShopPromotionActivityGoodsDeleteEvent](listener), "weimob_shop.promotion", "activityGoodsDelete")
	return s
}
