package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeClueDeleteListener struct {
	handler msg.XinyunEventHandler[XiaokeClueDeleteEvent]
}

func (s XiaokeClueDeleteListener) NewMessage() msg.MsgRequest[XiaokeClueDeleteEvent] {
	return &msg.XinyunOpenMessage[XiaokeClueDeleteEvent]{
		MsgBody: &XiaokeClueDeleteEvent{},
	}
}

func (s XiaokeClueDeleteListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeClueDeleteEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeClueDeleteEvent]))
}

type XiaokeClueDeleteEvent struct {
	ClueKeys  []string `json:"clueKeys,omitempty"`
	BuildTime int      `json:"buildTime,omitempty"`
	Wid       int64    `json:"wid,omitempty"`
}

// OnXiaokeClueDeleteEvent
// @id 1754
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=删除线索
func (s *Service) OnXiaokeClueDeleteEvent(handler msg.XinyunEventHandler[XiaokeClueDeleteEvent]) (service *Service) {
	var listener = &XiaokeClueDeleteListener{handler}
	s.Register(msg.WrapListener[XiaokeClueDeleteEvent](listener), "xiaoke_clue", "delete")
	return s
}
