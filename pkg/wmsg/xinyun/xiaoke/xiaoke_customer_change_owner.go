package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeCustomerChangeOwnerListener struct {
	handler msg.XinyunEventHandler[XiaokeCustomerChangeOwnerEvent]
}

func (s XiaokeCustomerChangeOwnerListener) NewMessage() msg.MsgRequest[XiaokeCustomerChangeOwnerEvent] {
	return &msg.XinyunOpenMessage[XiaokeCustomerChangeOwnerEvent]{
		MsgBody: &XiaokeCustomerChangeOwnerEvent{},
	}
}

func (s XiaokeCustomerChangeOwnerListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeCustomerChangeOwnerEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeCustomerChangeOwnerEvent]))
}

type XiaokeCustomerChangeOwnerEvent struct {
	Keys       []string `json:"keys,omitempty"`
	NewOwnerId int64    `json:"newOwnerId,omitempty"`
	BuildTime  int64    `json:"buildTime,omitempty"`
	Wid        int64    `json:"wid,omitempty"`
}

// OnXiaokeCustomerChangeOwnerEvent
// @id 1765
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=客户变更所属人
func (s *Service) OnXiaokeCustomerChangeOwnerEvent(handler msg.XinyunEventHandler[XiaokeCustomerChangeOwnerEvent]) (service *Service) {
	var listener = &XiaokeCustomerChangeOwnerListener{handler}
	s.Register(msg.WrapListener[XiaokeCustomerChangeOwnerEvent](listener), "xiaoke_customer", "changeOwner")
	return s
}
