package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type CcCouponUpdateCouponTemplateStatusListener struct {
	handler msg.XinyunEventHandler[CcCouponUpdateCouponTemplateStatusEvent]
}

func (s CcCouponUpdateCouponTemplateStatusListener) NewMessage() msg.MsgRequest[CcCouponUpdateCouponTemplateStatusEvent] {
	return &msg.XinyunOpenMessage[CcCouponUpdateCouponTemplateStatusEvent]{
		MsgBody: &CcCouponUpdateCouponTemplateStatusEvent{},
	}
}

func (s CcCouponUpdateCouponTemplateStatusListener) OnMessage(ctx context.Context, message msg.MsgRequest[CcCouponUpdateCouponTemplateStatusEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[CcCouponUpdateCouponTemplateStatusEvent]))
}

type CcCouponUpdateCouponTemplateStatusEvent struct {
	Pid            int64 `json:"pid,omitempty"`
	CardTemplateId int64 `json:"cardTemplateId,omitempty"`
	Status         int   `json:"status,omitempty"`
	Channel        int64 `json:"channel,omitempty"`
	Timestamp      int64 `json:"timestamp,omitempty"`
}

// OnCcCouponUpdateCouponTemplateStatusEvent
// @id 1687
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=券模板状态变更消息
func (s *Service) OnCcCouponUpdateCouponTemplateStatusEvent(handler msg.XinyunEventHandler[CcCouponUpdateCouponTemplateStatusEvent]) (service *Service) {
	var listener = &CcCouponUpdateCouponTemplateStatusListener{handler}
	s.Register(msg.WrapListener[CcCouponUpdateCouponTemplateStatusEvent](listener), "cc_coupon", "updateCouponTemplateStatus")
	return s
}
