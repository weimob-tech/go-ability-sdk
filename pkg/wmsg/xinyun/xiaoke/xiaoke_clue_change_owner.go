package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeClueChangeOwnerListener struct {
	handler msg.XinyunEventHandler[XiaokeClueChangeOwnerEvent]
}

func (s XiaokeClueChangeOwnerListener) NewMessage() msg.MsgRequest[XiaokeClueChangeOwnerEvent] {
	return &msg.XinyunOpenMessage[XiaokeClueChangeOwnerEvent]{
		MsgBody: &XiaokeClueChangeOwnerEvent{},
	}
}

func (s XiaokeClueChangeOwnerListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeClueChangeOwnerEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeClueChangeOwnerEvent]))
}

type XiaokeClueChangeOwnerEvent struct {
	Keys       []string `json:"keys,omitempty"`
	NewOwnerId int64    `json:"newOwnerId,omitempty"`
	BuildTime  int64    `json:"buildTime,omitempty"`
	Wid        int64    `json:"wid,omitempty"`
}

// OnXiaokeClueChangeOwnerEvent
// @id 1756
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=线索变更所属人
func (s *Service) OnXiaokeClueChangeOwnerEvent(handler msg.XinyunEventHandler[XiaokeClueChangeOwnerEvent]) (service *Service) {
	var listener = &XiaokeClueChangeOwnerListener{handler}
	s.Register(msg.WrapListener[XiaokeClueChangeOwnerEvent](listener), "xiaoke_clue", "changeOwner")
	return s
}
