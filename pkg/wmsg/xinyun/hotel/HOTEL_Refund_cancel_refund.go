package hotel

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type HOTELRefundCancelRefundListener struct {
	handler msg.XinyunEventHandler[HOTELRefundCancelRefundEvent]
}

func (s HOTELRefundCancelRefundListener) NewMessage() msg.MsgRequest[HOTELRefundCancelRefundEvent] {
	return &msg.XinyunOpenMessage[HOTELRefundCancelRefundEvent]{
		MsgBody: &HOTELRefundCancelRefundEvent{},
	}
}

func (s HOTELRefundCancelRefundListener) OnMessage(ctx context.Context, message msg.MsgRequest[HOTELRefundCancelRefundEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[HOTELRefundCancelRefundEvent]))
}

type HOTELRefundCancelRefundEvent struct {
	OrderNo  string `json:"orderNo,omitempty"`
	RefundNo string `json:"refundNo,omitempty"`
	StoreId  string `json:"storeId,omitempty"`
}

// OnHOTELRefundCancelRefundEvent
// @id 2342
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=取消维权消息
func (s *Service) OnHOTELRefundCancelRefundEvent(handler msg.XinyunEventHandler[HOTELRefundCancelRefundEvent]) (service *Service) {
	var listener = &HOTELRefundCancelRefundListener{handler}
	s.Register(msg.WrapListener[HOTELRefundCancelRefundEvent](listener), "HOTEL_Refund", "cancelRefund")
	return s
}
