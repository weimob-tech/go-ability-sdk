package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeOrderChangeOwnerListener struct {
	handler msg.XinyunEventHandler[XiaokeOrderChangeOwnerEvent]
}

func (s XiaokeOrderChangeOwnerListener) NewMessage() msg.MsgRequest[XiaokeOrderChangeOwnerEvent] {
	return &msg.XinyunOpenMessage[XiaokeOrderChangeOwnerEvent]{
		MsgBody: &XiaokeOrderChangeOwnerEvent{},
	}
}

func (s XiaokeOrderChangeOwnerListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeOrderChangeOwnerEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeOrderChangeOwnerEvent]))
}

type XiaokeOrderChangeOwnerEvent struct {
	Keys       []string `json:"keys,omitempty"`
	NewOwnerId int64    `json:"newOwnerId,omitempty"`
	UserWid    int64    `json:"userWid,omitempty"`
}

// OnXiaokeOrderChangeOwnerEvent
// @id 2687
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=订单变更所属人
func (s *Service) OnXiaokeOrderChangeOwnerEvent(handler msg.XinyunEventHandler[XiaokeOrderChangeOwnerEvent]) (service *Service) {
	var listener = &XiaokeOrderChangeOwnerListener{handler}
	s.Register(msg.WrapListener[XiaokeOrderChangeOwnerEvent](listener), "xiaoke_order", "changeOwner")
	return s
}
