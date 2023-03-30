package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcPromotionPromotionChangeListener struct {
	handler msg.XinyunEventHandler[EcPromotionPromotionChangeEvent]
}

func (s EcPromotionPromotionChangeListener) NewMessage() msg.MsgRequest[EcPromotionPromotionChangeEvent] {
	return &msg.XinyunOpenMessage[EcPromotionPromotionChangeEvent]{
		MsgBody: &EcPromotionPromotionChangeEvent{},
	}
}

func (s EcPromotionPromotionChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcPromotionPromotionChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcPromotionPromotionChangeEvent]))
}

type EcPromotionPromotionChangeEvent struct {
	ActivityId   int64   `json:"activityId,omitempty"`
	ActivityType int     `json:"activityType,omitempty"`
	GoodsIdList  []int64 `json:"goodsIdList,omitempty"`
	DateTime     int     `json:"dateTime,omitempty"`
	MsgType      int     `json:"msgType,omitempty"`
	OperateType  int     `json:"operateType,omitempty"`
}

// OnEcPromotionPromotionChangeEvent
// @id 1786
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=营销活动变更消息
func (s *Service) OnEcPromotionPromotionChangeEvent(handler msg.XinyunEventHandler[EcPromotionPromotionChangeEvent]) (service *Service) {
	var listener = &EcPromotionPromotionChangeListener{handler}
	s.Register(msg.WrapListener[EcPromotionPromotionChangeEvent](listener), "ec_promotion", "promotionChange")
	return s
}
