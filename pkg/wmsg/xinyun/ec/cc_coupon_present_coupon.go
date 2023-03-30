package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type CcCouponPresentCouponListener struct {
	handler msg.XinyunEventHandler[CcCouponPresentCouponEvent]
}

func (s CcCouponPresentCouponListener) NewMessage() msg.MsgRequest[CcCouponPresentCouponEvent] {
	return &msg.XinyunOpenMessage[CcCouponPresentCouponEvent]{
		MsgBody: &CcCouponPresentCouponEvent{},
	}
}

func (s CcCouponPresentCouponListener) OnMessage(ctx context.Context, message msg.MsgRequest[CcCouponPresentCouponEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[CcCouponPresentCouponEvent]))
}

type CcCouponPresentCouponEvent struct {
	Pid        int64  `json:"pid,omitempty"`
	FromWid    int64  `json:"fromWid,omitempty"`
	ToWid      int64  `json:"toWid,omitempty"`
	Code       string `json:"code,omitempty"`
	StatusType int    `json:"statusType,omitempty"`
	Channel    int64  `json:"channel,omitempty"`
}

// OnCcCouponPresentCouponEvent
// @id 1693
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=券转赠状态消息
func (s *Service) OnCcCouponPresentCouponEvent(handler msg.XinyunEventHandler[CcCouponPresentCouponEvent]) (service *Service) {
	var listener = &CcCouponPresentCouponListener{handler}
	s.Register(msg.WrapListener[CcCouponPresentCouponEvent](listener), "cc_coupon", "presentCoupon")
	return s
}
