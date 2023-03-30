package coupon

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type CouponCodeUseCouponCodeListener struct {
	handler msg.XinyunEventHandler[CouponCodeUseCouponCodeEvent]
}

func (s CouponCodeUseCouponCodeListener) NewMessage() msg.MsgRequest[CouponCodeUseCouponCodeEvent] {
	return &msg.XinyunOpenMessage[CouponCodeUseCouponCodeEvent]{
		MsgBody: &CouponCodeUseCouponCodeEvent{},
	}
}

func (s CouponCodeUseCouponCodeListener) OnMessage(ctx context.Context, message msg.MsgRequest[CouponCodeUseCouponCodeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[CouponCodeUseCouponCodeEvent]))
}

type CouponCodeUseCouponCodeEvent struct {
	Pid  int64  `json:"pid,omitempty"`
	Wid  int64  `json:"wid,omitempty"`
	Code string `json:"code,omitempty"`
}

// OnCouponCodeUseCouponCodeEvent
// @id 1166
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=优惠码核销通知
func (s *Service) OnCouponCodeUseCouponCodeEvent(handler msg.XinyunEventHandler[CouponCodeUseCouponCodeEvent]) (service *Service) {
	var listener = &CouponCodeUseCouponCodeListener{handler}
	s.Register(msg.WrapListener[CouponCodeUseCouponCodeEvent](listener), "coupon_code", "useCouponCode")
	return s
}
