package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeProductDeleteListener struct {
	handler msg.XinyunEventHandler[XiaokeProductDeleteEvent]
}

func (s XiaokeProductDeleteListener) NewMessage() msg.MsgRequest[XiaokeProductDeleteEvent] {
	return &msg.XinyunOpenMessage[XiaokeProductDeleteEvent]{
		MsgBody: &XiaokeProductDeleteEvent{},
	}
}

func (s XiaokeProductDeleteListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeProductDeleteEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeProductDeleteEvent]))
}

type XiaokeProductDeleteEvent struct {
	ProductUniqueRequestList []map[string]any `json:"productUniqueRequestList,omitempty"`
	BuildTime                int64            `json:"buildTime,omitempty"`
	Wid                      int64            `json:"wid,omitempty"`
}

// OnXiaokeProductDeleteEvent
// @id 2690
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=删除产品
func (s *Service) OnXiaokeProductDeleteEvent(handler msg.XinyunEventHandler[XiaokeProductDeleteEvent]) (service *Service) {
	var listener = &XiaokeProductDeleteListener{handler}
	s.Register(msg.WrapListener[XiaokeProductDeleteEvent](listener), "xiaoke_product", "delete")
	return s
}
