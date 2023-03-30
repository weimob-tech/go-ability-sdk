package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeOrderNullityListener struct {
	handler msg.XinyunEventHandler[XiaokeOrderNullityEvent]
}

func (s XiaokeOrderNullityListener) NewMessage() msg.MsgRequest[XiaokeOrderNullityEvent] {
	return &msg.XinyunOpenMessage[XiaokeOrderNullityEvent]{
		MsgBody: &XiaokeOrderNullityEvent{},
	}
}

func (s XiaokeOrderNullityListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeOrderNullityEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeOrderNullityEvent]))
}

type XiaokeOrderNullityEvent struct {
	OrderNumberList        []int64 `json:"orderNumberList,omitempty"`
	CauseOfAbandonment     string  `json:"causeOfAbandonment,omitempty"`
	CauseOfAbandonmentCode string  `json:"causeOfAbandonmentCode,omitempty"`
	LastUpdateUserWid      int64   `json:"lastUpdateUserWid,omitempty"`
	LastUpdateTime         int64   `json:"lastUpdateTime,omitempty"`
}

// OnXiaokeOrderNullityEvent
// @id 2686
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=作废订单
func (s *Service) OnXiaokeOrderNullityEvent(handler msg.XinyunEventHandler[XiaokeOrderNullityEvent]) (service *Service) {
	var listener = &XiaokeOrderNullityListener{handler}
	s.Register(msg.WrapListener[XiaokeOrderNullityEvent](listener), "xiaoke_order", "nullity")
	return s
}
