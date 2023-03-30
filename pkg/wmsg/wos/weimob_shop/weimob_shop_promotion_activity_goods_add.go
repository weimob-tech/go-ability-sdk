package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopPromotionActivityGoodsAddListener struct {
	handler msg.WosEventHandler[WeimobShopPromotionActivityGoodsAddEvent]
}

func (s WeimobShopPromotionActivityGoodsAddListener) NewMessage() msg.MsgRequest[WeimobShopPromotionActivityGoodsAddEvent] {
	return &msg.WosOpenMessage[WeimobShopPromotionActivityGoodsAddEvent]{
		MsgBody: &WeimobShopPromotionActivityGoodsAddEvent{},
	}
}

func (s WeimobShopPromotionActivityGoodsAddListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopPromotionActivityGoodsAddEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopPromotionActivityGoodsAddEvent]))
}

type WeimobShopPromotionActivityGoodsAddEvent struct {
	ActivityId   int64   `json:"activityId,omitempty"`
	ActivityType int64   `json:"activityType,omitempty"`
	GoodsIdList  []int64 `json:"goodsIdList,omitempty"`
}

// OnWeimobShopPromotionActivityGoodsAddEvent
// @id 1331
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1331?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=活动商品创建)
func (s *Service) OnWeimobShopPromotionActivityGoodsAddEvent(handler msg.WosEventHandler[WeimobShopPromotionActivityGoodsAddEvent]) (service *Service) {
	var listener = &WeimobShopPromotionActivityGoodsAddListener{handler}
	s.Register(msg.WrapListener[WeimobShopPromotionActivityGoodsAddEvent](listener), "weimob_shop.promotion", "activityGoodsAdd")
	return s
}
