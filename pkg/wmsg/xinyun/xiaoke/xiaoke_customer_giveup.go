package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeCustomerGiveupListener struct {
	handler msg.XinyunEventHandler[XiaokeCustomerGiveupEvent]
}

func (s XiaokeCustomerGiveupListener) NewMessage() msg.MsgRequest[XiaokeCustomerGiveupEvent] {
	return &msg.XinyunOpenMessage[XiaokeCustomerGiveupEvent]{
		MsgBody: &XiaokeCustomerGiveupEvent{},
	}
}

func (s XiaokeCustomerGiveupListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeCustomerGiveupEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeCustomerGiveupEvent]))
}

type XiaokeCustomerGiveupEvent struct {
	Keys      []string `json:"keys,omitempty"`
	Owner     int64    `json:"owner,omitempty"`
	Reason    string   `json:"reason,omitempty"`
	BuildTime int64    `json:"buildTime,omitempty"`
	Wid       int64    `json:"wid,omitempty"`
}

// OnXiaokeCustomerGiveupEvent
// @id 1762
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=客户放弃公海
func (s *Service) OnXiaokeCustomerGiveupEvent(handler msg.XinyunEventHandler[XiaokeCustomerGiveupEvent]) (service *Service) {
	var listener = &XiaokeCustomerGiveupListener{handler}
	s.Register(msg.WrapListener[XiaokeCustomerGiveupEvent](listener), "xiaoke_customer", "giveup")
	return s
}
