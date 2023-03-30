package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcStoreStoreChangeListener struct {
	handler msg.XinyunEventHandler[EcStoreStoreChangeEvent]
}

func (s EcStoreStoreChangeListener) NewMessage() msg.MsgRequest[EcStoreStoreChangeEvent] {
	return &msg.XinyunOpenMessage[EcStoreStoreChangeEvent]{
		MsgBody: &EcStoreStoreChangeEvent{},
	}
}

func (s EcStoreStoreChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcStoreStoreChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcStoreStoreChangeEvent]))
}

type EcStoreStoreChangeEvent struct {
	StoreId    int64 `json:"storeId,omitempty"`
	ChangeType int   `json:"changeType,omitempty"`
}

// OnEcStoreStoreChangeEvent
// @id 1244
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=门店信息变更消息
func (s *Service) OnEcStoreStoreChangeEvent(handler msg.XinyunEventHandler[EcStoreStoreChangeEvent]) (service *Service) {
	var listener = &EcStoreStoreChangeListener{handler}
	s.Register(msg.WrapListener[EcStoreStoreChangeEvent](listener), "ec_store", "storeChange")
	return s
}
