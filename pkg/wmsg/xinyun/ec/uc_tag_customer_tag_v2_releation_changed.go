package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type UcTagCustomerTagV2ReleationChangedListener struct {
	handler msg.XinyunEventHandler[UcTagCustomerTagV2ReleationChangedEvent]
}

func (s UcTagCustomerTagV2ReleationChangedListener) NewMessage() msg.MsgRequest[UcTagCustomerTagV2ReleationChangedEvent] {
	return &msg.XinyunOpenMessage[UcTagCustomerTagV2ReleationChangedEvent]{
		MsgBody: &UcTagCustomerTagV2ReleationChangedEvent{},
	}
}

func (s UcTagCustomerTagV2ReleationChangedListener) OnMessage(ctx context.Context, message msg.MsgRequest[UcTagCustomerTagV2ReleationChangedEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[UcTagCustomerTagV2ReleationChangedEvent]))
}

type UcTagCustomerTagV2ReleationChangedEvent struct {
	Wids []int `json:"wids,omitempty"`
	Pid  int64 `json:"pid,omitempty"`
}

// OnUcTagCustomerTagV2ReleationChangedEvent
// @id 1617
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=人的标签信息变更
func (s *Service) OnUcTagCustomerTagV2ReleationChangedEvent(handler msg.XinyunEventHandler[UcTagCustomerTagV2ReleationChangedEvent]) (service *Service) {
	var listener = &UcTagCustomerTagV2ReleationChangedListener{handler}
	s.Register(msg.WrapListener[UcTagCustomerTagV2ReleationChangedEvent](listener), "uc_tag", "customerTagV2ReleationChanged")
	return s
}
