package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmCouponCreateCouponTemplateListener struct {
	handler msg.WosEventHandler[WeimobCrmCouponCreateCouponTemplateEvent]
}

func (s WeimobCrmCouponCreateCouponTemplateListener) NewMessage() msg.MsgRequest[WeimobCrmCouponCreateCouponTemplateEvent] {
	return &msg.WosOpenMessage[WeimobCrmCouponCreateCouponTemplateEvent]{
		MsgBody: &WeimobCrmCouponCreateCouponTemplateEvent{},
	}
}

func (s WeimobCrmCouponCreateCouponTemplateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmCouponCreateCouponTemplateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmCouponCreateCouponTemplateEvent]))
}

type WeimobCrmCouponCreateCouponTemplateEvent struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	Name             string `json:"name,omitempty"`
	CouponType       int64  `json:"couponType,omitempty"`
	SaasChannel      string `json:"saasChannel,omitempty"`
}

// OnWeimobCrmCouponCreateCouponTemplateEvent
// @id 1363
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1363?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=新建优惠券消息)
func (s *Service) OnWeimobCrmCouponCreateCouponTemplateEvent(handler msg.WosEventHandler[WeimobCrmCouponCreateCouponTemplateEvent]) (service *Service) {
	var listener = &WeimobCrmCouponCreateCouponTemplateListener{handler}
	s.Register(msg.WrapListener[WeimobCrmCouponCreateCouponTemplateEvent](listener), "weimob_crm.coupon", "createCouponTemplate")
	return s
}
