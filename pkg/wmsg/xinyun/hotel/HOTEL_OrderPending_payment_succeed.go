package hotel

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type HOTELOrderPendingPaymentSucceedListener struct {
	handler msg.XinyunEventHandler[HOTELOrderPendingPaymentSucceedEvent]
}

func (s HOTELOrderPendingPaymentSucceedListener) NewMessage() msg.MsgRequest[HOTELOrderPendingPaymentSucceedEvent] {
	return &msg.XinyunOpenMessage[HOTELOrderPendingPaymentSucceedEvent]{
		MsgBody: &HOTELOrderPendingPaymentSucceedEvent{},
	}
}

func (s HOTELOrderPendingPaymentSucceedListener) OnMessage(ctx context.Context, message msg.MsgRequest[HOTELOrderPendingPaymentSucceedEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[HOTELOrderPendingPaymentSucceedEvent]))
}

type HOTELOrderPendingPaymentSucceedEvent struct {
	OrderNo       string  `json:"orderNo,omitempty"`
	PaymentAmount float64 `json:"paymentAmount,omitempty"`
}

// OnHOTELOrderPendingPaymentSucceedEvent
// @id 1293
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=待付订单支付成功通知
func (s *Service) OnHOTELOrderPendingPaymentSucceedEvent(handler msg.XinyunEventHandler[HOTELOrderPendingPaymentSucceedEvent]) (service *Service) {
	var listener = &HOTELOrderPendingPaymentSucceedListener{handler}
	s.Register(msg.WrapListener[HOTELOrderPendingPaymentSucceedEvent](listener), "HOTEL_OrderPending", "paymentSucceed")
	return s
}
