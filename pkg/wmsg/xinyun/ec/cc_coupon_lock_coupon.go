package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type CcCouponLockCouponListener struct {
	handler msg.XinyunEventHandler[CcCouponLockCouponEvent]
}

func (s CcCouponLockCouponListener) NewMessage() msg.MsgRequest[CcCouponLockCouponEvent] {
	return &msg.XinyunOpenMessage[CcCouponLockCouponEvent]{
		MsgBody: &CcCouponLockCouponEvent{},
	}
}

func (s CcCouponLockCouponListener) OnMessage(ctx context.Context, message msg.MsgRequest[CcCouponLockCouponEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[CcCouponLockCouponEvent]))
}

type CcCouponLockCouponEvent struct {
	Pid            int64  `json:"pid,omitempty"`
	Wid            int64  `json:"wid,omitempty"`
	Code           string `json:"code,omitempty"`
	CardTemplateId int64  `json:"cardTemplateId,omitempty"`
	Channel        int64  `json:"channel,omitempty"`
}

// OnCcCouponLockCouponEvent
// @id 1690
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=券锁定消息
func (s *Service) OnCcCouponLockCouponEvent(handler msg.XinyunEventHandler[CcCouponLockCouponEvent]) (service *Service) {
	var listener = &CcCouponLockCouponListener{handler}
	s.Register(msg.WrapListener[CcCouponLockCouponEvent](listener), "cc_coupon", "lockCoupon")
	return s
}
