package o2o

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type O2oBaseOrderPayChangeListener struct {
	handler msg.XinyunEventHandler[O2oBaseOrderPayChangeEvent]
}

func (s O2oBaseOrderPayChangeListener) NewMessage() msg.MsgRequest[O2oBaseOrderPayChangeEvent] {
	return &msg.XinyunOpenMessage[O2oBaseOrderPayChangeEvent]{
		MsgBody: &O2oBaseOrderPayChangeEvent{},
	}
}

func (s O2oBaseOrderPayChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[O2oBaseOrderPayChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[O2oBaseOrderPayChangeEvent]))
}

type O2oBaseOrderPayChangeEvent struct {
	Params O2oBaseOrderPayChangeParams `json:"params,omitempty"`
}

type O2oBaseOrderPayChangeParams struct {
}

// OnO2oBaseOrderPayChangeEvent
// @id 1428
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=
func (s *Service) OnO2oBaseOrderPayChangeEvent(handler msg.XinyunEventHandler[O2oBaseOrderPayChangeEvent]) (service *Service) {
	var listener = &O2oBaseOrderPayChangeListener{handler}
	s.Register(msg.WrapListener[O2oBaseOrderPayChangeEvent](listener), "o2o_base_order", "payChange")
	return s
}
