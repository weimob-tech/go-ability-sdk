package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeUserDeleteListener struct {
	handler msg.XinyunEventHandler[XiaokeUserDeleteEvent]
}

func (s XiaokeUserDeleteListener) NewMessage() msg.MsgRequest[XiaokeUserDeleteEvent] {
	return &msg.XinyunOpenMessage[XiaokeUserDeleteEvent]{
		MsgBody: &XiaokeUserDeleteEvent{},
	}
}

func (s XiaokeUserDeleteListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeUserDeleteEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeUserDeleteEvent]))
}

type XiaokeUserDeleteEvent struct {
	WidList     []int  `json:"widList,omitempty"`
	PassUserWid string `json:"passUserWid,omitempty"`
}

// OnXiaokeUserDeleteEvent
// @id 1750
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=删除员工
func (s *Service) OnXiaokeUserDeleteEvent(handler msg.XinyunEventHandler[XiaokeUserDeleteEvent]) (service *Service) {
	var listener = &XiaokeUserDeleteListener{handler}
	s.Register(msg.WrapListener[XiaokeUserDeleteEvent](listener), "xiaoke_user", "delete")
	return s
}
