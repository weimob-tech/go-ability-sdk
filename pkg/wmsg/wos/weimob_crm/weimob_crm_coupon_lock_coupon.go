package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmCouponLockCouponListener struct {
	handler msg.WosEventHandler[WeimobCrmCouponLockCouponEvent]
}

func (s WeimobCrmCouponLockCouponListener) NewMessage() msg.MsgRequest[WeimobCrmCouponLockCouponEvent] {
	return &msg.WosOpenMessage[WeimobCrmCouponLockCouponEvent]{
		MsgBody: &WeimobCrmCouponLockCouponEvent{},
	}
}

func (s WeimobCrmCouponLockCouponListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmCouponLockCouponEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmCouponLockCouponEvent]))
}

type WeimobCrmCouponLockCouponEvent struct {
	Wid         int64  `json:"wid,omitempty"`
	Code        string `json:"code,omitempty"`
	SaasChannel string `json:"saasChannel,omitempty"`
}

// OnWeimobCrmCouponLockCouponEvent
// @id 1346
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1346?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=券锁定消息)
func (s *Service) OnWeimobCrmCouponLockCouponEvent(handler msg.WosEventHandler[WeimobCrmCouponLockCouponEvent]) (service *Service) {
	var listener = &WeimobCrmCouponLockCouponListener{handler}
	s.Register(msg.WrapListener[WeimobCrmCouponLockCouponEvent](listener), "weimob_crm.coupon", "lockCoupon")
	return s
}
