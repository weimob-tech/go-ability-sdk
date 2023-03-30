package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type CcCouponExpireCouponListener struct {
	handler msg.XinyunEventHandler[CcCouponExpireCouponEvent]
}

func (s CcCouponExpireCouponListener) NewMessage() msg.MsgRequest[CcCouponExpireCouponEvent] {
	return &msg.XinyunOpenMessage[CcCouponExpireCouponEvent]{
		MsgBody: &CcCouponExpireCouponEvent{},
	}
}

func (s CcCouponExpireCouponListener) OnMessage(ctx context.Context, message msg.MsgRequest[CcCouponExpireCouponEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[CcCouponExpireCouponEvent]))
}

type CcCouponExpireCouponEvent struct {
	Wid  int64  `json:"wid,omitempty"`
	Code string `json:"code,omitempty"`
}

// OnCcCouponExpireCouponEvent
// @id 1705
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=优惠券过期消息
func (s *Service) OnCcCouponExpireCouponEvent(handler msg.XinyunEventHandler[CcCouponExpireCouponEvent]) (service *Service) {
	var listener = &CcCouponExpireCouponListener{handler}
	s.Register(msg.WrapListener[CcCouponExpireCouponEvent](listener), "cc_coupon", "expireCoupon")
	return s
}
