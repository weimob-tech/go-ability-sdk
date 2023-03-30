package tour

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type TourOrderPaymentSuccessListener struct {
	handler msg.XinyunEventHandler[TourOrderPaymentSuccessEvent]
}

func (s TourOrderPaymentSuccessListener) NewMessage() msg.MsgRequest[TourOrderPaymentSuccessEvent] {
	return &msg.XinyunOpenMessage[TourOrderPaymentSuccessEvent]{
		MsgBody: &TourOrderPaymentSuccessEvent{},
	}
}

func (s TourOrderPaymentSuccessListener) OnMessage(ctx context.Context, message msg.MsgRequest[TourOrderPaymentSuccessEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[TourOrderPaymentSuccessEvent]))
}

type TourOrderPaymentSuccessEvent struct {
	OrderNo string `json:"orderNo,omitempty"`
}

// OnTourOrderPaymentSuccessEvent
// @id 1036
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=订单支付成功
func (s *Service) OnTourOrderPaymentSuccessEvent(handler msg.XinyunEventHandler[TourOrderPaymentSuccessEvent]) (service *Service) {
	var listener = &TourOrderPaymentSuccessListener{handler}
	s.Register(msg.WrapListener[TourOrderPaymentSuccessEvent](listener), "tour_order", "paymentSuccess")
	return s
}
