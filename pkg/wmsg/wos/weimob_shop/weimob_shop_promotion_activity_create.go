package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopPromotionActivityCreateListener struct {
	handler msg.WosEventHandler[WeimobShopPromotionActivityCreateEvent]
}

func (s WeimobShopPromotionActivityCreateListener) NewMessage() msg.MsgRequest[WeimobShopPromotionActivityCreateEvent] {
	return &msg.WosOpenMessage[WeimobShopPromotionActivityCreateEvent]{
		MsgBody: &WeimobShopPromotionActivityCreateEvent{},
	}
}

func (s WeimobShopPromotionActivityCreateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopPromotionActivityCreateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopPromotionActivityCreateEvent]))
}

type WeimobShopPromotionActivityCreateEvent struct {
	ActivityId   int64 `json:"activityId,omitempty"`
	ActivityType int64 `json:"activityType,omitempty"`
}

// OnWeimobShopPromotionActivityCreateEvent
// @id 1326
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1326?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=活动创建)
func (s *Service) OnWeimobShopPromotionActivityCreateEvent(handler msg.WosEventHandler[WeimobShopPromotionActivityCreateEvent]) (service *Service) {
	var listener = &WeimobShopPromotionActivityCreateListener{handler}
	s.Register(msg.WrapListener[WeimobShopPromotionActivityCreateEvent](listener), "weimob_shop.promotion", "activityCreate")
	return s
}
