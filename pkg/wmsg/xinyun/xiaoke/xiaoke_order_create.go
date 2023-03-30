package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeOrderCreateListener struct {
	handler msg.XinyunEventHandler[XiaokeOrderCreateEvent]
}

func (s XiaokeOrderCreateListener) NewMessage() msg.MsgRequest[XiaokeOrderCreateEvent] {
	return &msg.XinyunOpenMessage[XiaokeOrderCreateEvent]{
		MsgBody: &XiaokeOrderCreateEvent{},
	}
}

func (s XiaokeOrderCreateListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeOrderCreateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeOrderCreateEvent]))
}

type XiaokeOrderCreateEvent struct {
	List      []map[string]any `json:"list,omitempty"`
	Products  []map[string]any `json:"products,omitempty"`
	BuildTime int64            `json:"buildTime,omitempty"`
	Wid       int64            `json:"wid,omitempty"`
}

// OnXiaokeOrderCreateEvent
// @id 2683
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=创建订单
func (s *Service) OnXiaokeOrderCreateEvent(handler msg.XinyunEventHandler[XiaokeOrderCreateEvent]) (service *Service) {
	var listener = &XiaokeOrderCreateListener{handler}
	s.Register(msg.WrapListener[XiaokeOrderCreateEvent](listener), "xiaoke_order", "create")
	return s
}
