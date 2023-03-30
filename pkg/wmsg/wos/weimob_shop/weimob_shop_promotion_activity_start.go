package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopPromotionActivityStartListener struct {
	handler msg.WosEventHandler[WeimobShopPromotionActivityStartEvent]
}

func (s WeimobShopPromotionActivityStartListener) NewMessage() msg.MsgRequest[WeimobShopPromotionActivityStartEvent] {
	return &msg.WosOpenMessage[WeimobShopPromotionActivityStartEvent]{
		MsgBody: &WeimobShopPromotionActivityStartEvent{},
	}
}

func (s WeimobShopPromotionActivityStartListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopPromotionActivityStartEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopPromotionActivityStartEvent]))
}

type WeimobShopPromotionActivityStartEvent struct {
	ActivityId   int64 `json:"activityId,omitempty"`
	ActivityType int64 `json:"activityType,omitempty"`
}

// OnWeimobShopPromotionActivityStartEvent
// @id 1327
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1327?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=活动开始)
func (s *Service) OnWeimobShopPromotionActivityStartEvent(handler msg.WosEventHandler[WeimobShopPromotionActivityStartEvent]) (service *Service) {
	var listener = &WeimobShopPromotionActivityStartListener{handler}
	s.Register(msg.WrapListener[WeimobShopPromotionActivityStartEvent](listener), "weimob_shop.promotion", "activityStart")
	return s
}
