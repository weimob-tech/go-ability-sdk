package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeNicheDeleteListener struct {
	handler msg.XinyunEventHandler[XiaokeNicheDeleteEvent]
}

func (s XiaokeNicheDeleteListener) NewMessage() msg.MsgRequest[XiaokeNicheDeleteEvent] {
	return &msg.XinyunOpenMessage[XiaokeNicheDeleteEvent]{
		MsgBody: &XiaokeNicheDeleteEvent{},
	}
}

func (s XiaokeNicheDeleteListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeNicheDeleteEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeNicheDeleteEvent]))
}

type XiaokeNicheDeleteEvent struct {
	NicheKeys []string `json:"nicheKeys,omitempty"`
	BuildTime int      `json:"buildTime,omitempty"`
	Wid       int64    `json:"wid,omitempty"`
}

// OnXiaokeNicheDeleteEvent
// @id 1773
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=删除商机
func (s *Service) OnXiaokeNicheDeleteEvent(handler msg.XinyunEventHandler[XiaokeNicheDeleteEvent]) (service *Service) {
	var listener = &XiaokeNicheDeleteListener{handler}
	s.Register(msg.WrapListener[XiaokeNicheDeleteEvent](listener), "xiaoke_niche", "delete")
	return s
}
