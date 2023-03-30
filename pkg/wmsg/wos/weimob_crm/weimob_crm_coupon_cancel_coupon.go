package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmCouponCancelCouponListener struct {
	handler msg.WosEventHandler[WeimobCrmCouponCancelCouponEvent]
}

func (s WeimobCrmCouponCancelCouponListener) NewMessage() msg.MsgRequest[WeimobCrmCouponCancelCouponEvent] {
	return &msg.WosOpenMessage[WeimobCrmCouponCancelCouponEvent]{
		MsgBody: &WeimobCrmCouponCancelCouponEvent{},
	}
}

func (s WeimobCrmCouponCancelCouponListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmCouponCancelCouponEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmCouponCancelCouponEvent]))
}

type WeimobCrmCouponCancelCouponEvent struct {
	Code             string `json:"code,omitempty"`
	Wid              int64  `json:"wid,omitempty"`
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	Type             int64  `json:"type,omitempty"`
	SubScene         int64  `json:"subScene,omitempty"`
	SubSceneId       string `json:"subSceneId,omitempty"`
	SaasChannel      string `json:"saasChannel,omitempty"`
}

// OnWeimobCrmCouponCancelCouponEvent
// @id 1348
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1348?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=券作废消息)
func (s *Service) OnWeimobCrmCouponCancelCouponEvent(handler msg.WosEventHandler[WeimobCrmCouponCancelCouponEvent]) (service *Service) {
	var listener = &WeimobCrmCouponCancelCouponListener{handler}
	s.Register(msg.WrapListener[WeimobCrmCouponCancelCouponEvent](listener), "weimob_crm.coupon", "cancelCoupon")
	return s
}
