package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmCouponUpdateCouponTemplateListener struct {
	handler msg.WosEventHandler[WeimobCrmCouponUpdateCouponTemplateEvent]
}

func (s WeimobCrmCouponUpdateCouponTemplateListener) NewMessage() msg.MsgRequest[WeimobCrmCouponUpdateCouponTemplateEvent] {
	return &msg.WosOpenMessage[WeimobCrmCouponUpdateCouponTemplateEvent]{
		MsgBody: &WeimobCrmCouponUpdateCouponTemplateEvent{},
	}
}

func (s WeimobCrmCouponUpdateCouponTemplateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmCouponUpdateCouponTemplateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmCouponUpdateCouponTemplateEvent]))
}

type WeimobCrmCouponUpdateCouponTemplateEvent struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	Name             string `json:"name,omitempty"`
	SaasChannel      string `json:"saasChannel,omitempty"`
}

// OnWeimobCrmCouponUpdateCouponTemplateEvent
// @id 1342
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1342?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=修改优惠券消息)
func (s *Service) OnWeimobCrmCouponUpdateCouponTemplateEvent(handler msg.WosEventHandler[WeimobCrmCouponUpdateCouponTemplateEvent]) (service *Service) {
	var listener = &WeimobCrmCouponUpdateCouponTemplateListener{handler}
	s.Register(msg.WrapListener[WeimobCrmCouponUpdateCouponTemplateEvent](listener), "weimob_crm.coupon", "updateCouponTemplate")
	return s
}
