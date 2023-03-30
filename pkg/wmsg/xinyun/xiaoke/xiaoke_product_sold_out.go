package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeProductSoldOutListener struct {
	handler msg.XinyunEventHandler[XiaokeProductSoldOutEvent]
}

func (s XiaokeProductSoldOutListener) NewMessage() msg.MsgRequest[XiaokeProductSoldOutEvent] {
	return &msg.XinyunOpenMessage[XiaokeProductSoldOutEvent]{
		MsgBody: &XiaokeProductSoldOutEvent{},
	}
}

func (s XiaokeProductSoldOutListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeProductSoldOutEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeProductSoldOutEvent]))
}

type XiaokeProductSoldOutEvent struct {
	ProductUniqueRequestList []map[string]any `json:"productUniqueRequestList,omitempty"`
	BuildTime                int64            `json:"buildTime,omitempty"`
	Wid                      int64            `json:"wid,omitempty"`
}

// OnXiaokeProductSoldOutEvent
// @id 2692
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=下架产品
func (s *Service) OnXiaokeProductSoldOutEvent(handler msg.XinyunEventHandler[XiaokeProductSoldOutEvent]) (service *Service) {
	var listener = &XiaokeProductSoldOutListener{handler}
	s.Register(msg.WrapListener[XiaokeProductSoldOutEvent](listener), "xiaoke_product", "soldOut")
	return s
}
