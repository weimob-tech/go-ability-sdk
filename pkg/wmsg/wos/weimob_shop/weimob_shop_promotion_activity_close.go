package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopPromotionActivityCloseListener struct {
	handler msg.WosEventHandler[WeimobShopPromotionActivityCloseEvent]
}

func (s WeimobShopPromotionActivityCloseListener) NewMessage() msg.MsgRequest[WeimobShopPromotionActivityCloseEvent] {
	return &msg.WosOpenMessage[WeimobShopPromotionActivityCloseEvent]{
		MsgBody: &WeimobShopPromotionActivityCloseEvent{},
	}
}

func (s WeimobShopPromotionActivityCloseListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopPromotionActivityCloseEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopPromotionActivityCloseEvent]))
}

type WeimobShopPromotionActivityCloseEvent struct {
	ActivityId   int64 `json:"activityId,omitempty"`
	ActivityType int64 `json:"activityType,omitempty"`
}

// OnWeimobShopPromotionActivityCloseEvent
// @id 1328
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1328?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=活动结束)
func (s *Service) OnWeimobShopPromotionActivityCloseEvent(handler msg.WosEventHandler[WeimobShopPromotionActivityCloseEvent]) (service *Service) {
	var listener = &WeimobShopPromotionActivityCloseListener{handler}
	s.Register(msg.WrapListener[WeimobShopPromotionActivityCloseEvent](listener), "weimob_shop.promotion", "activityClose")
	return s
}
