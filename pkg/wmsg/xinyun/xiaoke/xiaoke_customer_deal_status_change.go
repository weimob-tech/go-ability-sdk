package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeCustomerDealStatusChangeListener struct {
	handler msg.XinyunEventHandler[XiaokeCustomerDealStatusChangeEvent]
}

func (s XiaokeCustomerDealStatusChangeListener) NewMessage() msg.MsgRequest[XiaokeCustomerDealStatusChangeEvent] {
	return &msg.XinyunOpenMessage[XiaokeCustomerDealStatusChangeEvent]{
		MsgBody: &XiaokeCustomerDealStatusChangeEvent{},
	}
}

func (s XiaokeCustomerDealStatusChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeCustomerDealStatusChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeCustomerDealStatusChangeEvent]))
}

type XiaokeCustomerDealStatusChangeEvent struct {
	CustomerKey string `json:"customerKey,omitempty"`
	DealStatus  int    `json:"dealStatus,omitempty"`
	BuildTime   int    `json:"buildTime,omitempty"`
	Wid         int64  `json:"wid,omitempty"`
}

// OnXiaokeCustomerDealStatusChangeEvent
// @id 1764
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=客户成交状态变更
func (s *Service) OnXiaokeCustomerDealStatusChangeEvent(handler msg.XinyunEventHandler[XiaokeCustomerDealStatusChangeEvent]) (service *Service) {
	var listener = &XiaokeCustomerDealStatusChangeListener{handler}
	s.Register(msg.WrapListener[XiaokeCustomerDealStatusChangeEvent](listener), "xiaoke_customer", "dealStatusChange")
	return s
}
