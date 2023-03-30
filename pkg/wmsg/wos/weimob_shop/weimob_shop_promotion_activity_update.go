package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopPromotionActivityUpdateListener struct {
	handler msg.WosEventHandler[WeimobShopPromotionActivityUpdateEvent]
}

func (s WeimobShopPromotionActivityUpdateListener) NewMessage() msg.MsgRequest[WeimobShopPromotionActivityUpdateEvent] {
	return &msg.WosOpenMessage[WeimobShopPromotionActivityUpdateEvent]{
		MsgBody: &WeimobShopPromotionActivityUpdateEvent{},
	}
}

func (s WeimobShopPromotionActivityUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopPromotionActivityUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopPromotionActivityUpdateEvent]))
}

type WeimobShopPromotionActivityUpdateEvent struct {
	ActivityId   int64 `json:"activityId,omitempty"`
	ActivityType int64 `json:"activityType,omitempty"`
}

// OnWeimobShopPromotionActivityUpdateEvent
// @id 1329
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1329?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=活动修改)
func (s *Service) OnWeimobShopPromotionActivityUpdateEvent(handler msg.WosEventHandler[WeimobShopPromotionActivityUpdateEvent]) (service *Service) {
	var listener = &WeimobShopPromotionActivityUpdateListener{handler}
	s.Register(msg.WrapListener[WeimobShopPromotionActivityUpdateEvent](listener), "weimob_shop.promotion", "activityUpdate")
	return s
}
