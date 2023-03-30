package hotel

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type HOTELRefundAgreeRefundListener struct {
	handler msg.XinyunEventHandler[HOTELRefundAgreeRefundEvent]
}

func (s HOTELRefundAgreeRefundListener) NewMessage() msg.MsgRequest[HOTELRefundAgreeRefundEvent] {
	return &msg.XinyunOpenMessage[HOTELRefundAgreeRefundEvent]{
		MsgBody: &HOTELRefundAgreeRefundEvent{},
	}
}

func (s HOTELRefundAgreeRefundListener) OnMessage(ctx context.Context, message msg.MsgRequest[HOTELRefundAgreeRefundEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[HOTELRefundAgreeRefundEvent]))
}

type HOTELRefundAgreeRefundEvent struct {
	StoreId  int64  `json:"storeId,omitempty"`
	RefundNo int64  `json:"refundNo,omitempty"`
	OrderNo  string `json:"orderNo,omitempty"`
}

// OnHOTELRefundAgreeRefundEvent
// @id 2341
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=同意维权消息
func (s *Service) OnHOTELRefundAgreeRefundEvent(handler msg.XinyunEventHandler[HOTELRefundAgreeRefundEvent]) (service *Service) {
	var listener = &HOTELRefundAgreeRefundListener{handler}
	s.Register(msg.WrapListener[HOTELRefundAgreeRefundEvent](listener), "HOTEL_Refund", "agreeRefund")
	return s
}
