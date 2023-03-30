package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type CcCouponUpdateCouponTemplateListener struct {
	handler msg.XinyunEventHandler[CcCouponUpdateCouponTemplateEvent]
}

func (s CcCouponUpdateCouponTemplateListener) NewMessage() msg.MsgRequest[CcCouponUpdateCouponTemplateEvent] {
	return &msg.XinyunOpenMessage[CcCouponUpdateCouponTemplateEvent]{
		MsgBody: &CcCouponUpdateCouponTemplateEvent{},
	}
}

func (s CcCouponUpdateCouponTemplateListener) OnMessage(ctx context.Context, message msg.MsgRequest[CcCouponUpdateCouponTemplateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[CcCouponUpdateCouponTemplateEvent]))
}

type CcCouponUpdateCouponTemplateEvent struct {
	Pid            int64 `json:"pid,omitempty"`
	CardTemplateId int64 `json:"cardTemplateId,omitempty"`
	OperateUserId  int64 `json:"operateUserId,omitempty"`
	Channel        int64 `json:"channel,omitempty"`
}

// OnCcCouponUpdateCouponTemplateEvent
// @id 1686
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=修改券模板消息
func (s *Service) OnCcCouponUpdateCouponTemplateEvent(handler msg.XinyunEventHandler[CcCouponUpdateCouponTemplateEvent]) (service *Service) {
	var listener = &CcCouponUpdateCouponTemplateListener{handler}
	s.Register(msg.WrapListener[CcCouponUpdateCouponTemplateEvent](listener), "cc_coupon", "updateCouponTemplate")
	return s
}
