package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmCouponUpdateBasicCouponTemplateListener struct {
	handler msg.WosEventHandler[WeimobCrmCouponUpdateBasicCouponTemplateEvent]
}

func (s WeimobCrmCouponUpdateBasicCouponTemplateListener) NewMessage() msg.MsgRequest[WeimobCrmCouponUpdateBasicCouponTemplateEvent] {
	return &msg.WosOpenMessage[WeimobCrmCouponUpdateBasicCouponTemplateEvent]{
		MsgBody: &WeimobCrmCouponUpdateBasicCouponTemplateEvent{},
	}
}

func (s WeimobCrmCouponUpdateBasicCouponTemplateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmCouponUpdateBasicCouponTemplateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmCouponUpdateBasicCouponTemplateEvent]))
}

type WeimobCrmCouponUpdateBasicCouponTemplateEvent struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	Name             string `json:"name,omitempty"`
	SaasChannel      string `json:"saasChannel,omitempty"`
}

// OnWeimobCrmCouponUpdateBasicCouponTemplateEvent
// @id 1360
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1360?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=修改优惠券基础模板消息)
func (s *Service) OnWeimobCrmCouponUpdateBasicCouponTemplateEvent(handler msg.WosEventHandler[WeimobCrmCouponUpdateBasicCouponTemplateEvent]) (service *Service) {
	var listener = &WeimobCrmCouponUpdateBasicCouponTemplateListener{handler}
	s.Register(msg.WrapListener[WeimobCrmCouponUpdateBasicCouponTemplateEvent](listener), "weimob_crm.coupon", "updateBasicCouponTemplate")
	return s
}
