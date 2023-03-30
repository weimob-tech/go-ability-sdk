package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcRightsCreateRightsListener struct {
	handler msg.XinyunEventHandler[EcRightsCreateRightsEvent]
}

func (s EcRightsCreateRightsListener) NewMessage() msg.MsgRequest[EcRightsCreateRightsEvent] {
	return &msg.XinyunOpenMessage[EcRightsCreateRightsEvent]{
		MsgBody: &EcRightsCreateRightsEvent{},
	}
}

func (s EcRightsCreateRightsListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcRightsCreateRightsEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcRightsCreateRightsEvent]))
}

type EcRightsCreateRightsEvent struct {
	RightsId int64 `json:"rightsId,omitempty"`
	OrderNo  int64 `json:"orderNo,omitempty"`
	StoreId  int64 `json:"storeId,omitempty"`
}

// OnEcRightsCreateRightsEvent
// @id 478
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=新增售后消息
func (s *Service) OnEcRightsCreateRightsEvent(handler msg.XinyunEventHandler[EcRightsCreateRightsEvent]) (service *Service) {
	var listener = &EcRightsCreateRightsListener{handler}
	s.Register(msg.WrapListener[EcRightsCreateRightsEvent](listener), "ec_rights", "createRights")
	return s
}
