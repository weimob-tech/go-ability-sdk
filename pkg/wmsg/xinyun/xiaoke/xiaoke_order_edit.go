package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeOrderEditListener struct {
	handler msg.XinyunEventHandler[XiaokeOrderEditEvent]
}

func (s XiaokeOrderEditListener) NewMessage() msg.MsgRequest[XiaokeOrderEditEvent] {
	return &msg.XinyunOpenMessage[XiaokeOrderEditEvent]{
		MsgBody: &XiaokeOrderEditEvent{},
	}
}

func (s XiaokeOrderEditListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeOrderEditEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeOrderEditEvent]))
}

type XiaokeOrderEditEvent struct {
	List      []map[string]any `json:"list,omitempty"`
	Products  []map[string]any `json:"products,omitempty"`
	BuildTime int64            `json:"buildTime,omitempty"`
	Wid       int64            `json:"wid,omitempty"`
}

// OnXiaokeOrderEditEvent
// @id 2684
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=编辑订单
func (s *Service) OnXiaokeOrderEditEvent(handler msg.XinyunEventHandler[XiaokeOrderEditEvent]) (service *Service) {
	var listener = &XiaokeOrderEditListener{handler}
	s.Register(msg.WrapListener[XiaokeOrderEditEvent](listener), "xiaoke_order", "edit")
	return s
}
