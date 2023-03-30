package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type CcCouponUpdateCouponStockListener struct {
	handler msg.XinyunEventHandler[CcCouponUpdateCouponStockEvent]
}

func (s CcCouponUpdateCouponStockListener) NewMessage() msg.MsgRequest[CcCouponUpdateCouponStockEvent] {
	return &msg.XinyunOpenMessage[CcCouponUpdateCouponStockEvent]{
		MsgBody: &CcCouponUpdateCouponStockEvent{},
	}
}

func (s CcCouponUpdateCouponStockListener) OnMessage(ctx context.Context, message msg.MsgRequest[CcCouponUpdateCouponStockEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[CcCouponUpdateCouponStockEvent]))
}

type CcCouponUpdateCouponStockEvent struct {
	Pid            int64 `json:"pid,omitempty"`
	CardTemplateId int64 `json:"cardTemplateId,omitempty"`
	Stock          int   `json:"stock,omitempty"`
	TotalStock     int   `json:"totalStock,omitempty"`
	IsIncr         bool  `json:"isIncr,omitempty"`
	Channel        int64 `json:"channel,omitempty"`
}

// OnCcCouponUpdateCouponStockEvent
// @id 1688
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=券模板库存更新消息
func (s *Service) OnCcCouponUpdateCouponStockEvent(handler msg.XinyunEventHandler[CcCouponUpdateCouponStockEvent]) (service *Service) {
	var listener = &CcCouponUpdateCouponStockListener{handler}
	s.Register(msg.WrapListener[CcCouponUpdateCouponStockEvent](listener), "cc_coupon", "updateCouponStock")
	return s
}
