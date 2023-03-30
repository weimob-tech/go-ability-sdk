package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeClueGiveupListener struct {
	handler msg.XinyunEventHandler[XiaokeClueGiveupEvent]
}

func (s XiaokeClueGiveupListener) NewMessage() msg.MsgRequest[XiaokeClueGiveupEvent] {
	return &msg.XinyunOpenMessage[XiaokeClueGiveupEvent]{
		MsgBody: &XiaokeClueGiveupEvent{},
	}
}

func (s XiaokeClueGiveupListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeClueGiveupEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeClueGiveupEvent]))
}

type XiaokeClueGiveupEvent struct {
	Keys      []string `json:"keys,omitempty"`
	Owner     int64    `json:"owner,omitempty"`
	Reason    string   `json:"reason,omitempty"`
	BuildTime int64    `json:"buildTime,omitempty"`
	Wid       int64    `json:"wid,omitempty"`
}

// OnXiaokeClueGiveupEvent
// @id 1752
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=线索放弃公海
func (s *Service) OnXiaokeClueGiveupEvent(handler msg.XinyunEventHandler[XiaokeClueGiveupEvent]) (service *Service) {
	var listener = &XiaokeClueGiveupListener{handler}
	s.Register(msg.WrapListener[XiaokeClueGiveupEvent](listener), "xiaoke_clue", "giveup")
	return s
}
