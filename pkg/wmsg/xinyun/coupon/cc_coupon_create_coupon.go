package coupon

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type CcCouponCreateCouponListener struct {
	handler msg.XinyunEventHandler[CcCouponCreateCouponEvent]
}

func (s CcCouponCreateCouponListener) NewMessage() msg.MsgRequest[CcCouponCreateCouponEvent] {
	return &msg.XinyunOpenMessage[CcCouponCreateCouponEvent]{
		MsgBody: &CcCouponCreateCouponEvent{},
	}
}

func (s CcCouponCreateCouponListener) OnMessage(ctx context.Context, message msg.MsgRequest[CcCouponCreateCouponEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[CcCouponCreateCouponEvent]))
}

type CcCouponCreateCouponEvent struct {
	Pid              int64  `json:"pid,omitempty"`
	CardTemplateId   int64  `json:"cardTemplateId,omitempty"`
	CardTemplateName string `json:"cardTemplateName,omitempty"`
	Type             string `json:"type,omitempty"`
}

// OnCcCouponCreateCouponEvent
// @id 1165
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=新建优惠券模板消息
func (s *Service) OnCcCouponCreateCouponEvent(handler msg.XinyunEventHandler[CcCouponCreateCouponEvent]) (service *Service) {
	var listener = &CcCouponCreateCouponListener{handler}
	s.Register(msg.WrapListener[CcCouponCreateCouponEvent](listener), "cc_coupon", "createCoupon")
	return s
}
