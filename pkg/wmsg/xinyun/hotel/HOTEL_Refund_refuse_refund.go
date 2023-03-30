package hotel

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type HOTELRefundRefuseRefundListener struct {
	handler msg.XinyunEventHandler[HOTELRefundRefuseRefundEvent]
}

func (s HOTELRefundRefuseRefundListener) NewMessage() msg.MsgRequest[HOTELRefundRefuseRefundEvent] {
	return &msg.XinyunOpenMessage[HOTELRefundRefuseRefundEvent]{
		MsgBody: &HOTELRefundRefuseRefundEvent{},
	}
}

func (s HOTELRefundRefuseRefundListener) OnMessage(ctx context.Context, message msg.MsgRequest[HOTELRefundRefuseRefundEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[HOTELRefundRefuseRefundEvent]))
}

type HOTELRefundRefuseRefundEvent struct {
	OrderNo  string `json:"orderNo,omitempty"`
	RefundNo string `json:"refundNo,omitempty"`
	StoreId  string `json:"storeId,omitempty"`
}

// OnHOTELRefundRefuseRefundEvent
// @id 2340
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=拒绝维权消息
func (s *Service) OnHOTELRefundRefuseRefundEvent(handler msg.XinyunEventHandler[HOTELRefundRefuseRefundEvent]) (service *Service) {
	var listener = &HOTELRefundRefuseRefundListener{handler}
	s.Register(msg.WrapListener[HOTELRefundRefuseRefundEvent](listener), "HOTEL_Refund", "refuseRefund")
	return s
}
