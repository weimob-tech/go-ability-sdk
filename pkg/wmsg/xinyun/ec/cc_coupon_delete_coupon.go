package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type CcCouponDeleteCouponListener struct {
	handler msg.XinyunEventHandler[CcCouponDeleteCouponEvent]
}

func (s CcCouponDeleteCouponListener) NewMessage() msg.MsgRequest[CcCouponDeleteCouponEvent] {
	return &msg.XinyunOpenMessage[CcCouponDeleteCouponEvent]{
		MsgBody: &CcCouponDeleteCouponEvent{},
	}
}

func (s CcCouponDeleteCouponListener) OnMessage(ctx context.Context, message msg.MsgRequest[CcCouponDeleteCouponEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[CcCouponDeleteCouponEvent]))
}

type CcCouponDeleteCouponEvent struct {
	Pid            int64  `json:"pid,omitempty"`
	CardTemplateId int64  `json:"cardTemplateId,omitempty"`
	Wid            int64  `json:"wid,omitempty"`
	Code           string `json:"code,omitempty"`
	Channel        int64  `json:"channel,omitempty"`
}

// OnCcCouponDeleteCouponEvent
// @id 1692
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=券作废消息
func (s *Service) OnCcCouponDeleteCouponEvent(handler msg.XinyunEventHandler[CcCouponDeleteCouponEvent]) (service *Service) {
	var listener = &CcCouponDeleteCouponListener{handler}
	s.Register(msg.WrapListener[CcCouponDeleteCouponEvent](listener), "cc_coupon", "deleteCoupon")
	return s
}
