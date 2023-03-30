package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeClueDropListener struct {
	handler msg.XinyunEventHandler[XiaokeClueDropEvent]
}

func (s XiaokeClueDropListener) NewMessage() msg.MsgRequest[XiaokeClueDropEvent] {
	return &msg.XinyunOpenMessage[XiaokeClueDropEvent]{
		MsgBody: &XiaokeClueDropEvent{},
	}
}

func (s XiaokeClueDropListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeClueDropEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeClueDropEvent]))
}

type XiaokeClueDropEvent struct {
	Keys      []string `json:"keys,omitempty"`
	Owner     int64    `json:"owner,omitempty"`
	BuildTime int64    `json:"buildTime,omitempty"`
	Wid       int64    `json:"wid,omitempty"`
}

// OnXiaokeClueDropEvent
// @id 1751
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=线索掉保公海
func (s *Service) OnXiaokeClueDropEvent(handler msg.XinyunEventHandler[XiaokeClueDropEvent]) (service *Service) {
	var listener = &XiaokeClueDropListener{handler}
	s.Register(msg.WrapListener[XiaokeClueDropEvent](listener), "xiaoke_clue", "drop")
	return s
}
