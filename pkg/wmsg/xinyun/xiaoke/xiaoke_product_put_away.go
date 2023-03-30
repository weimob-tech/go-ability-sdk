package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeProductPutAwayListener struct {
	handler msg.XinyunEventHandler[XiaokeProductPutAwayEvent]
}

func (s XiaokeProductPutAwayListener) NewMessage() msg.MsgRequest[XiaokeProductPutAwayEvent] {
	return &msg.XinyunOpenMessage[XiaokeProductPutAwayEvent]{
		MsgBody: &XiaokeProductPutAwayEvent{},
	}
}

func (s XiaokeProductPutAwayListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeProductPutAwayEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeProductPutAwayEvent]))
}

type XiaokeProductPutAwayEvent struct {
	ProductUniqueRequestList []map[string]any `json:"productUniqueRequestList,omitempty"`
	BuildTime                int64            `json:"buildTime,omitempty"`
	Wid                      int64            `json:"wid,omitempty"`
}

// OnXiaokeProductPutAwayEvent
// @id 2691
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=上架产品
func (s *Service) OnXiaokeProductPutAwayEvent(handler msg.XinyunEventHandler[XiaokeProductPutAwayEvent]) (service *Service) {
	var listener = &XiaokeProductPutAwayListener{handler}
	s.Register(msg.WrapListener[XiaokeProductPutAwayEvent](listener), "xiaoke_product", "putAway")
	return s
}
