package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeCustomerApplyListener struct {
	handler msg.XinyunEventHandler[XiaokeCustomerApplyEvent]
}

func (s XiaokeCustomerApplyListener) NewMessage() msg.MsgRequest[XiaokeCustomerApplyEvent] {
	return &msg.XinyunOpenMessage[XiaokeCustomerApplyEvent]{
		MsgBody: &XiaokeCustomerApplyEvent{},
	}
}

func (s XiaokeCustomerApplyListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeCustomerApplyEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeCustomerApplyEvent]))
}

type XiaokeCustomerApplyEvent struct {
	NicheKeys    []string `json:"nicheKeys,omitempty"`
	ContactKeys  []string `json:"contactKeys,omitempty"`
	CustomerKeys []string `json:"customerKeys,omitempty"`
	Owner        int64    `json:"owner,omitempty"`
	BuildTime    int64    `json:"buildTime,omitempty"`
	Wid          int64    `json:"wid,omitempty"`
}

// OnXiaokeCustomerApplyEvent
// @id 1768
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=客户公海数据领取
func (s *Service) OnXiaokeCustomerApplyEvent(handler msg.XinyunEventHandler[XiaokeCustomerApplyEvent]) (service *Service) {
	var listener = &XiaokeCustomerApplyListener{handler}
	s.Register(msg.WrapListener[XiaokeCustomerApplyEvent](listener), "xiaoke_customer", "apply")
	return s
}
