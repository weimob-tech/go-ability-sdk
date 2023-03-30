package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type CcCouponUpdateCouponOwnerListener struct {
	handler msg.XinyunEventHandler[CcCouponUpdateCouponOwnerEvent]
}

func (s CcCouponUpdateCouponOwnerListener) NewMessage() msg.MsgRequest[CcCouponUpdateCouponOwnerEvent] {
	return &msg.XinyunOpenMessage[CcCouponUpdateCouponOwnerEvent]{
		MsgBody: &CcCouponUpdateCouponOwnerEvent{},
	}
}

func (s CcCouponUpdateCouponOwnerListener) OnMessage(ctx context.Context, message msg.MsgRequest[CcCouponUpdateCouponOwnerEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[CcCouponUpdateCouponOwnerEvent]))
}

type CcCouponUpdateCouponOwnerEvent struct {
	FromWid int64  `json:"fromWid,omitempty"`
	ToWid   int64  `json:"toWid,omitempty"`
	Pid     int64  `json:"pid,omitempty"`
	CardIds string `json:"cardIds,omitempty"`
}

// OnCcCouponUpdateCouponOwnerEvent
// @id 1689
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=身份合并券归属变更消息
func (s *Service) OnCcCouponUpdateCouponOwnerEvent(handler msg.XinyunEventHandler[CcCouponUpdateCouponOwnerEvent]) (service *Service) {
	var listener = &CcCouponUpdateCouponOwnerListener{handler}
	s.Register(msg.WrapListener[CcCouponUpdateCouponOwnerEvent](listener), "cc_coupon", "updateCouponOwner")
	return s
}
