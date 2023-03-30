package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmCouponPresentCouponListener struct {
	handler msg.WosEventHandler[WeimobCrmCouponPresentCouponEvent]
}

func (s WeimobCrmCouponPresentCouponListener) NewMessage() msg.MsgRequest[WeimobCrmCouponPresentCouponEvent] {
	return &msg.WosOpenMessage[WeimobCrmCouponPresentCouponEvent]{
		MsgBody: &WeimobCrmCouponPresentCouponEvent{},
	}
}

func (s WeimobCrmCouponPresentCouponListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmCouponPresentCouponEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmCouponPresentCouponEvent]))
}

type WeimobCrmCouponPresentCouponEvent struct {
	Code             string `json:"code,omitempty"`
	FromWid          int64  `json:"fromWid,omitempty"`
	ToWid            int64  `json:"toWid,omitempty"`
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	StatusType       int64  `json:"statusType,omitempty"`
	SaasChannel      string `json:"saasChannel,omitempty"`
}

// OnWeimobCrmCouponPresentCouponEvent
// @id 1349
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1349?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=券转赠状态消息)
func (s *Service) OnWeimobCrmCouponPresentCouponEvent(handler msg.WosEventHandler[WeimobCrmCouponPresentCouponEvent]) (service *Service) {
	var listener = &WeimobCrmCouponPresentCouponListener{handler}
	s.Register(msg.WrapListener[WeimobCrmCouponPresentCouponEvent](listener), "weimob_crm.coupon", "presentCoupon")
	return s
}
