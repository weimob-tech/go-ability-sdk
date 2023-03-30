package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopPromotionActivityDeleteListener struct {
	handler msg.WosEventHandler[WeimobShopPromotionActivityDeleteEvent]
}

func (s WeimobShopPromotionActivityDeleteListener) NewMessage() msg.MsgRequest[WeimobShopPromotionActivityDeleteEvent] {
	return &msg.WosOpenMessage[WeimobShopPromotionActivityDeleteEvent]{
		MsgBody: &WeimobShopPromotionActivityDeleteEvent{},
	}
}

func (s WeimobShopPromotionActivityDeleteListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopPromotionActivityDeleteEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopPromotionActivityDeleteEvent]))
}

type WeimobShopPromotionActivityDeleteEvent struct {
	ActivityId   int64 `json:"activityId,omitempty"`
	ActivityType int64 `json:"activityType,omitempty"`
}

// OnWeimobShopPromotionActivityDeleteEvent
// @id 1330
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1330?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=活动删除)
func (s *Service) OnWeimobShopPromotionActivityDeleteEvent(handler msg.WosEventHandler[WeimobShopPromotionActivityDeleteEvent]) (service *Service) {
	var listener = &WeimobShopPromotionActivityDeleteListener{handler}
	s.Register(msg.WrapListener[WeimobShopPromotionActivityDeleteEvent](listener), "weimob_shop.promotion", "activityDelete")
	return s
}
