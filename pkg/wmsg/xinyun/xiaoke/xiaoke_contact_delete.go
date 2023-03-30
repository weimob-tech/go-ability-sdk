package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeContactDeleteListener struct {
	handler msg.XinyunEventHandler[XiaokeContactDeleteEvent]
}

func (s XiaokeContactDeleteListener) NewMessage() msg.MsgRequest[XiaokeContactDeleteEvent] {
	return &msg.XinyunOpenMessage[XiaokeContactDeleteEvent]{
		MsgBody: &XiaokeContactDeleteEvent{},
	}
}

func (s XiaokeContactDeleteListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeContactDeleteEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeContactDeleteEvent]))
}

type XiaokeContactDeleteEvent struct {
	ContactKeys []string `json:"contactKeys,omitempty"`
	BuildTime   int      `json:"buildTime,omitempty"`
	Wid         int64    `json:"wid,omitempty"`
}

// OnXiaokeContactDeleteEvent
// @id 1776
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=删除联系人
func (s *Service) OnXiaokeContactDeleteEvent(handler msg.XinyunEventHandler[XiaokeContactDeleteEvent]) (service *Service) {
	var listener = &XiaokeContactDeleteListener{handler}
	s.Register(msg.WrapListener[XiaokeContactDeleteEvent](listener), "xiaoke_contact", "delete")
	return s
}
