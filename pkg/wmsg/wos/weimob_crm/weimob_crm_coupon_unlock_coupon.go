package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmCouponUnlockCouponListener struct {
	handler msg.WosEventHandler[WeimobCrmCouponUnlockCouponEvent]
}

func (s WeimobCrmCouponUnlockCouponListener) NewMessage() msg.MsgRequest[WeimobCrmCouponUnlockCouponEvent] {
	return &msg.WosOpenMessage[WeimobCrmCouponUnlockCouponEvent]{
		MsgBody: &WeimobCrmCouponUnlockCouponEvent{},
	}
}

func (s WeimobCrmCouponUnlockCouponListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmCouponUnlockCouponEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmCouponUnlockCouponEvent]))
}

type WeimobCrmCouponUnlockCouponEvent struct {
	Code             string `json:"code,omitempty"`
	Wid              int64  `json:"wid,omitempty"`
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	SaasChannel      string `json:"saasChannel,omitempty"`
}

// OnWeimobCrmCouponUnlockCouponEvent
// @id 1347
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1347?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=券解锁消息)
func (s *Service) OnWeimobCrmCouponUnlockCouponEvent(handler msg.WosEventHandler[WeimobCrmCouponUnlockCouponEvent]) (service *Service) {
	var listener = &WeimobCrmCouponUnlockCouponListener{handler}
	s.Register(msg.WrapListener[WeimobCrmCouponUnlockCouponEvent](listener), "weimob_crm.coupon", "unlockCoupon")
	return s
}
