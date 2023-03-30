package coupon

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type CcCouponUpdateBasicTemplateListener struct {
	handler msg.XinyunEventHandler[CcCouponUpdateBasicTemplateEvent]
}

func (s CcCouponUpdateBasicTemplateListener) NewMessage() msg.MsgRequest[CcCouponUpdateBasicTemplateEvent] {
	return &msg.XinyunOpenMessage[CcCouponUpdateBasicTemplateEvent]{
		MsgBody: &CcCouponUpdateBasicTemplateEvent{},
	}
}

func (s CcCouponUpdateBasicTemplateListener) OnMessage(ctx context.Context, message msg.MsgRequest[CcCouponUpdateBasicTemplateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[CcCouponUpdateBasicTemplateEvent]))
}

type CcCouponUpdateBasicTemplateEvent struct {
	Params CcCouponUpdateBasicTemplateParams `json:"params,omitempty"`
}

type CcCouponUpdateBasicTemplateParams struct {
}

// OnCcCouponUpdateBasicTemplateEvent
// @id 3958
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=
func (s *Service) OnCcCouponUpdateBasicTemplateEvent(handler msg.XinyunEventHandler[CcCouponUpdateBasicTemplateEvent]) (service *Service) {
	var listener = &CcCouponUpdateBasicTemplateListener{handler}
	s.Register(msg.WrapListener[CcCouponUpdateBasicTemplateEvent](listener), "cc_coupon", "updateBasicTemplate")
	return s
}
