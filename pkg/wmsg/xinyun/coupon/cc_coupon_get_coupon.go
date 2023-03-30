package coupon

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type CcCouponGetCouponListener struct {
	handler msg.XinyunEventHandler[CcCouponGetCouponEvent]
}

func (s CcCouponGetCouponListener) NewMessage() msg.MsgRequest[CcCouponGetCouponEvent] {
	return &msg.XinyunOpenMessage[CcCouponGetCouponEvent]{
		MsgBody: &CcCouponGetCouponEvent{},
	}
}

func (s CcCouponGetCouponListener) OnMessage(ctx context.Context, message msg.MsgRequest[CcCouponGetCouponEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[CcCouponGetCouponEvent]))
}

type CcCouponGetCouponEvent struct {
	Pid                  int64  `json:"pid,omitempty"`
	Wid                  int64  `json:"wid,omitempty"`
	Code                 string `json:"code,omitempty"`
	Type                 int    `json:"type,omitempty"`
	CardTemplateId       int64  `json:"cardTemplateId,omitempty"`
	Channel              int    `json:"channel,omitempty"`
	IsGiveByFriend       int64  `json:"isGiveByFriend,omitempty"`
	IsExternalGeneration int    `json:"isExternalGeneration,omitempty"`
	AcquireStoreId       int64  `json:"acquireStoreId,omitempty"`
	ExpDate              int64  `json:"expDate,omitempty"`
}

// OnCcCouponGetCouponEvent
// @id 1163
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=领取优惠券消息
func (s *Service) OnCcCouponGetCouponEvent(handler msg.XinyunEventHandler[CcCouponGetCouponEvent]) (service *Service) {
	var listener = &CcCouponGetCouponListener{handler}
	s.Register(msg.WrapListener[CcCouponGetCouponEvent](listener), "cc_coupon", "getCoupon")
	return s
}
