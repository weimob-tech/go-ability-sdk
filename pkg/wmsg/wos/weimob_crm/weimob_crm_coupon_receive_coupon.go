package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmCouponReceiveCouponListener struct {
	handler msg.WosEventHandler[WeimobCrmCouponReceiveCouponEvent]
}

func (s WeimobCrmCouponReceiveCouponListener) NewMessage() msg.MsgRequest[WeimobCrmCouponReceiveCouponEvent] {
	return &msg.WosOpenMessage[WeimobCrmCouponReceiveCouponEvent]{
		MsgBody: &WeimobCrmCouponReceiveCouponEvent{},
	}
}

func (s WeimobCrmCouponReceiveCouponListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmCouponReceiveCouponEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmCouponReceiveCouponEvent]))
}

type WeimobCrmCouponReceiveCouponEvent struct {
	Wid              int64  `json:"wid,omitempty"`
	Code             string `json:"code,omitempty"`
	CouponType       int64  `json:"couponType,omitempty"`
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	Vid              int64  `json:"vid,omitempty"`
	ExpireTime       int64  `json:"expireTime,omitempty"`
	SubScene         int64  `json:"subScene,omitempty"`
	SubSceneId       string `json:"subSceneId,omitempty"`
	SWid             int64  `json:"sWid,omitempty"`
	ShareGuiderWid   int64  `json:"shareGuiderWid,omitempty"`
	SaasChannel      string `json:"saasChannel,omitempty"`
}

// OnWeimobCrmCouponReceiveCouponEvent
// @id 1259
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1259?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=领券消息)
func (s *Service) OnWeimobCrmCouponReceiveCouponEvent(handler msg.WosEventHandler[WeimobCrmCouponReceiveCouponEvent]) (service *Service) {
	var listener = &WeimobCrmCouponReceiveCouponListener{handler}
	s.Register(msg.WrapListener[WeimobCrmCouponReceiveCouponEvent](listener), "weimob_crm.coupon", "receiveCoupon")
	return s
}
