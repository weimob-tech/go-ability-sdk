package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeClueAssignListener struct {
	handler msg.XinyunEventHandler[XiaokeClueAssignEvent]
}

func (s XiaokeClueAssignListener) NewMessage() msg.MsgRequest[XiaokeClueAssignEvent] {
	return &msg.XinyunOpenMessage[XiaokeClueAssignEvent]{
		MsgBody: &XiaokeClueAssignEvent{},
	}
}

func (s XiaokeClueAssignListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeClueAssignEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeClueAssignEvent]))
}

type XiaokeClueAssignEvent struct {
	ClueKeys  []string `json:"clueKeys,omitempty"`
	Owner     int64    `json:"owner,omitempty"`
	BuildTime int64    `json:"buildTime,omitempty"`
	Wid       int64    `json:"wid,omitempty"`
}

// OnXiaokeClueAssignEvent
// @id 1760
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=线索公海数据分配
func (s *Service) OnXiaokeClueAssignEvent(handler msg.XinyunEventHandler[XiaokeClueAssignEvent]) (service *Service) {
	var listener = &XiaokeClueAssignListener{handler}
	s.Register(msg.WrapListener[XiaokeClueAssignEvent](listener), "xiaoke_clue", "assign")
	return s
}
