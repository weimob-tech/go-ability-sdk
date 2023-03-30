package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmCouponCreateBasicCouponListener struct {
	handler msg.WosEventHandler[WeimobCrmCouponCreateBasicCouponEvent]
}

func (s WeimobCrmCouponCreateBasicCouponListener) NewMessage() msg.MsgRequest[WeimobCrmCouponCreateBasicCouponEvent] {
	return &msg.WosOpenMessage[WeimobCrmCouponCreateBasicCouponEvent]{
		MsgBody: &WeimobCrmCouponCreateBasicCouponEvent{},
	}
}

func (s WeimobCrmCouponCreateBasicCouponListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmCouponCreateBasicCouponEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmCouponCreateBasicCouponEvent]))
}

type WeimobCrmCouponCreateBasicCouponEvent struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	Name             string `json:"name,omitempty"`
	SaasChannel      string `json:"saasChannel,omitempty"`
}

// OnWeimobCrmCouponCreateBasicCouponEvent
// @id 1361
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1361?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=新建优惠券基础模板消息)
func (s *Service) OnWeimobCrmCouponCreateBasicCouponEvent(handler msg.WosEventHandler[WeimobCrmCouponCreateBasicCouponEvent]) (service *Service) {
	var listener = &WeimobCrmCouponCreateBasicCouponListener{handler}
	s.Register(msg.WrapListener[WeimobCrmCouponCreateBasicCouponEvent](listener), "weimob_crm.coupon", "createBasicCoupon")
	return s
}
