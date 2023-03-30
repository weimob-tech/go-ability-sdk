package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmCouponExpireCouponListener struct {
	handler msg.WosEventHandler[WeimobCrmCouponExpireCouponEvent]
}

func (s WeimobCrmCouponExpireCouponListener) NewMessage() msg.MsgRequest[WeimobCrmCouponExpireCouponEvent] {
	return &msg.WosOpenMessage[WeimobCrmCouponExpireCouponEvent]{
		MsgBody: &WeimobCrmCouponExpireCouponEvent{},
	}
}

func (s WeimobCrmCouponExpireCouponListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmCouponExpireCouponEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmCouponExpireCouponEvent]))
}

type WeimobCrmCouponExpireCouponEvent struct {
	Code             string `json:"code,omitempty"`
	Wid              int64  `json:"wid,omitempty"`
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	SubScene         int64  `json:"subScene,omitempty"`
	SubSceneId       string `json:"subSceneId,omitempty"`
	SaasChannel      string `json:"saasChannel,omitempty"`
}

// OnWeimobCrmCouponExpireCouponEvent
// @id 1350
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1350?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=优惠券过期消息)
func (s *Service) OnWeimobCrmCouponExpireCouponEvent(handler msg.WosEventHandler[WeimobCrmCouponExpireCouponEvent]) (service *Service) {
	var listener = &WeimobCrmCouponExpireCouponListener{handler}
	s.Register(msg.WrapListener[WeimobCrmCouponExpireCouponEvent](listener), "weimob_crm.coupon", "expireCoupon")
	return s
}
