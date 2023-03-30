package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeOrderDeleteListener struct {
	handler msg.XinyunEventHandler[XiaokeOrderDeleteEvent]
}

func (s XiaokeOrderDeleteListener) NewMessage() msg.MsgRequest[XiaokeOrderDeleteEvent] {
	return &msg.XinyunOpenMessage[XiaokeOrderDeleteEvent]{
		MsgBody: &XiaokeOrderDeleteEvent{},
	}
}

func (s XiaokeOrderDeleteListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeOrderDeleteEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeOrderDeleteEvent]))
}

type XiaokeOrderDeleteEvent struct {
	OrderNumberList   []int64 `json:"orderNumberList,omitempty"`
	IsDel             int     `json:"isDel,omitempty"`
	LastUpdateUserWid int64   `json:"lastUpdateUserWid,omitempty"`
	LastUpdateTime    int64   `json:"lastUpdateTime,omitempty"`
}

// OnXiaokeOrderDeleteEvent
// @id 2685
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=删除订单
func (s *Service) OnXiaokeOrderDeleteEvent(handler msg.XinyunEventHandler[XiaokeOrderDeleteEvent]) (service *Service) {
	var listener = &XiaokeOrderDeleteListener{handler}
	s.Register(msg.WrapListener[XiaokeOrderDeleteEvent](listener), "xiaoke_order", "delete")
	return s
}
