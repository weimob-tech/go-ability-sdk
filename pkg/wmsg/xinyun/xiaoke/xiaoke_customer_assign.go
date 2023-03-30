package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeCustomerAssignListener struct {
	handler msg.XinyunEventHandler[XiaokeCustomerAssignEvent]
}

func (s XiaokeCustomerAssignListener) NewMessage() msg.MsgRequest[XiaokeCustomerAssignEvent] {
	return &msg.XinyunOpenMessage[XiaokeCustomerAssignEvent]{
		MsgBody: &XiaokeCustomerAssignEvent{},
	}
}

func (s XiaokeCustomerAssignListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeCustomerAssignEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeCustomerAssignEvent]))
}

type XiaokeCustomerAssignEvent struct {
	NicheKeys    []string `json:"nicheKeys,omitempty"`
	ContactKeys  []string `json:"contactKeys,omitempty"`
	CustomerKeys []string `json:"customerKeys,omitempty"`
	Owner        int64    `json:"owner,omitempty"`
	BuildTime    int64    `json:"buildTime,omitempty"`
	Wid          int64    `json:"wid,omitempty"`
}

// OnXiaokeCustomerAssignEvent
// @id 1769
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=客户公海数据分配
func (s *Service) OnXiaokeCustomerAssignEvent(handler msg.XinyunEventHandler[XiaokeCustomerAssignEvent]) (service *Service) {
	var listener = &XiaokeCustomerAssignListener{handler}
	s.Register(msg.WrapListener[XiaokeCustomerAssignEvent](listener), "xiaoke_customer", "assign")
	return s
}
