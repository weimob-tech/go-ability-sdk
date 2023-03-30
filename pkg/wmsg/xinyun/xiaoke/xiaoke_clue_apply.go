package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeClueApplyListener struct {
	handler msg.XinyunEventHandler[XiaokeClueApplyEvent]
}

func (s XiaokeClueApplyListener) NewMessage() msg.MsgRequest[XiaokeClueApplyEvent] {
	return &msg.XinyunOpenMessage[XiaokeClueApplyEvent]{
		MsgBody: &XiaokeClueApplyEvent{},
	}
}

func (s XiaokeClueApplyListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeClueApplyEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeClueApplyEvent]))
}

type XiaokeClueApplyEvent struct {
	MsgSignature string `json:"msgSignature,omitempty"`
}

// OnXiaokeClueApplyEvent
// @id 1759
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=线索公海数据领取
func (s *Service) OnXiaokeClueApplyEvent(handler msg.XinyunEventHandler[XiaokeClueApplyEvent]) (service *Service) {
	var listener = &XiaokeClueApplyListener{handler}
	s.Register(msg.WrapListener[XiaokeClueApplyEvent](listener), "xiaoke_clue", "apply")
	return s
}
