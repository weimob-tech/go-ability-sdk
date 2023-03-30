package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeClueEditListener struct {
	handler msg.XinyunEventHandler[XiaokeClueEditEvent]
}

func (s XiaokeClueEditListener) NewMessage() msg.MsgRequest[XiaokeClueEditEvent] {
	return &msg.XinyunOpenMessage[XiaokeClueEditEvent]{
		MsgBody: &XiaokeClueEditEvent{},
	}
}

func (s XiaokeClueEditListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeClueEditEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeClueEditEvent]))
}

type XiaokeClueEditEvent struct {
	Keys       []string `json:"keys,omitempty"`
	NewOwnerId int64    `json:"newOwnerId,omitempty"`
	BuildTime  int64    `json:"buildTime,omitempty"`
	Wid        int64    `json:"wid,omitempty"`
}

// OnXiaokeClueEditEvent
// @id 1757
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=编辑线索
func (s *Service) OnXiaokeClueEditEvent(handler msg.XinyunEventHandler[XiaokeClueEditEvent]) (service *Service) {
	var listener = &XiaokeClueEditListener{handler}
	s.Register(msg.WrapListener[XiaokeClueEditEvent](listener), "xiaoke_clue", "edit")
	return s
}
