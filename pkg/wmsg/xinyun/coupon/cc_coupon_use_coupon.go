package coupon

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type CcCouponUseCouponListener struct {
	handler msg.XinyunEventHandler[CcCouponUseCouponEvent]
}

func (s CcCouponUseCouponListener) NewMessage() msg.MsgRequest[CcCouponUseCouponEvent] {
	return &msg.XinyunOpenMessage[CcCouponUseCouponEvent]{
		MsgBody: &CcCouponUseCouponEvent{},
	}
}

func (s CcCouponUseCouponListener) OnMessage(ctx context.Context, message msg.MsgRequest[CcCouponUseCouponEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[CcCouponUseCouponEvent]))
}

type CcCouponUseCouponEvent struct {
	Pid              int64  `json:"pid,omitempty"`
	Wid              int64  `json:"wid,omitempty"`
	Code             string `json:"code,omitempty"`
	CardTemplateId   string `json:"cardTemplateId,omitempty"`
	UseAdaptObjectId int64  `json:"useAdaptObjectId,omitempty"`
	Channel          int64  `json:"channel,omitempty"`
	TransactionId    int64  `json:"transactionId,omitempty"`
}

// OnCcCouponUseCouponEvent
// @id 1164
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=优惠券核销消息
func (s *Service) OnCcCouponUseCouponEvent(handler msg.XinyunEventHandler[CcCouponUseCouponEvent]) (service *Service) {
	var listener = &CcCouponUseCouponListener{handler}
	s.Register(msg.WrapListener[CcCouponUseCouponEvent](listener), "cc_coupon", "useCoupon")
	return s
}
