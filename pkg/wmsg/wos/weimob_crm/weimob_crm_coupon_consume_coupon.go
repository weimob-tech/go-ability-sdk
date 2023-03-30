package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmCouponConsumeCouponListener struct {
	handler msg.WosEventHandler[WeimobCrmCouponConsumeCouponEvent]
}

func (s WeimobCrmCouponConsumeCouponListener) NewMessage() msg.MsgRequest[WeimobCrmCouponConsumeCouponEvent] {
	return &msg.WosOpenMessage[WeimobCrmCouponConsumeCouponEvent]{
		MsgBody: &WeimobCrmCouponConsumeCouponEvent{},
	}
}

func (s WeimobCrmCouponConsumeCouponListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmCouponConsumeCouponEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmCouponConsumeCouponEvent]))
}

type WeimobCrmCouponConsumeCouponEvent struct {
	Wid              int64  `json:"wid,omitempty"`
	Code             string `json:"code,omitempty"`
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	Vid              int64  `json:"vid,omitempty"`
	UseScene         int64  `json:"useScene,omitempty"`
	OrderNo          int64  `json:"orderNo,omitempty"`
	UseTime          int64  `json:"useTime,omitempty"`
	SWid             int64  `json:"sWid,omitempty"`
	SubScene         int64  `json:"subScene,omitempty"`
	SubSceneId       string `json:"subSceneId,omitempty"`
	SaasChannel      string `json:"saasChannel,omitempty"`
}

// OnWeimobCrmCouponConsumeCouponEvent
// @id 1341
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1341?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=核销优惠券消息)
func (s *Service) OnWeimobCrmCouponConsumeCouponEvent(handler msg.WosEventHandler[WeimobCrmCouponConsumeCouponEvent]) (service *Service) {
	var listener = &WeimobCrmCouponConsumeCouponListener{handler}
	s.Register(msg.WrapListener[WeimobCrmCouponConsumeCouponEvent](listener), "weimob_crm.coupon", "consumeCoupon")
	return s
}
