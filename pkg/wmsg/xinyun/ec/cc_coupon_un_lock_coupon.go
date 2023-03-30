package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type CcCouponUnLockCouponListener struct {
	handler msg.XinyunEventHandler[CcCouponUnLockCouponEvent]
}

func (s CcCouponUnLockCouponListener) NewMessage() msg.MsgRequest[CcCouponUnLockCouponEvent] {
	return &msg.XinyunOpenMessage[CcCouponUnLockCouponEvent]{
		MsgBody: &CcCouponUnLockCouponEvent{},
	}
}

func (s CcCouponUnLockCouponListener) OnMessage(ctx context.Context, message msg.MsgRequest[CcCouponUnLockCouponEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[CcCouponUnLockCouponEvent]))
}

type CcCouponUnLockCouponEvent struct {
	Pid            int64  `json:"pid,omitempty"`
	Wid            int64  `json:"wid,omitempty"`
	Code           string `json:"code,omitempty"`
	CardTemplateId int64  `json:"cardTemplateId,omitempty"`
	Channel        int64  `json:"channel,omitempty"`
}

// OnCcCouponUnLockCouponEvent
// @id 1691
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=券解锁消息
func (s *Service) OnCcCouponUnLockCouponEvent(handler msg.XinyunEventHandler[CcCouponUnLockCouponEvent]) (service *Service) {
	var listener = &CcCouponUnLockCouponListener{handler}
	s.Register(msg.WrapListener[CcCouponUnLockCouponEvent](listener), "cc_coupon", "unLockCoupon")
	return s
}
