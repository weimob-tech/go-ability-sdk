package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmCouponUpdateCouponTemplateStatusListener struct {
	handler msg.WosEventHandler[WeimobCrmCouponUpdateCouponTemplateStatusEvent]
}

func (s WeimobCrmCouponUpdateCouponTemplateStatusListener) NewMessage() msg.MsgRequest[WeimobCrmCouponUpdateCouponTemplateStatusEvent] {
	return &msg.WosOpenMessage[WeimobCrmCouponUpdateCouponTemplateStatusEvent]{
		MsgBody: &WeimobCrmCouponUpdateCouponTemplateStatusEvent{},
	}
}

func (s WeimobCrmCouponUpdateCouponTemplateStatusListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmCouponUpdateCouponTemplateStatusEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmCouponUpdateCouponTemplateStatusEvent]))
}

type WeimobCrmCouponUpdateCouponTemplateStatusEvent struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	Status           int64  `json:"status,omitempty"`
	SaasChannel      string `json:"saasChannel,omitempty"`
}

// OnWeimobCrmCouponUpdateCouponTemplateStatusEvent
// @id 1343
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1343?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=优惠券状态变更消息)
func (s *Service) OnWeimobCrmCouponUpdateCouponTemplateStatusEvent(handler msg.WosEventHandler[WeimobCrmCouponUpdateCouponTemplateStatusEvent]) (service *Service) {
	var listener = &WeimobCrmCouponUpdateCouponTemplateStatusListener{handler}
	s.Register(msg.WrapListener[WeimobCrmCouponUpdateCouponTemplateStatusEvent](listener), "weimob_crm.coupon", "updateCouponTemplateStatus")
	return s
}
