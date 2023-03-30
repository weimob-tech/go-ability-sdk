package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcRightsAgreeRightsListener struct {
	handler msg.XinyunEventHandler[EcRightsAgreeRightsEvent]
}

func (s EcRightsAgreeRightsListener) NewMessage() msg.MsgRequest[EcRightsAgreeRightsEvent] {
	return &msg.XinyunOpenMessage[EcRightsAgreeRightsEvent]{
		MsgBody: &EcRightsAgreeRightsEvent{},
	}
}

func (s EcRightsAgreeRightsListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcRightsAgreeRightsEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcRightsAgreeRightsEvent]))
}

type EcRightsAgreeRightsEvent struct {
	RightsId    int64 `json:"rightsId,omitempty"`
	OrderNo     int64 `json:"orderNo,omitempty"`
	StoreId     int64 `json:"storeId,omitempty"`
	OrderSource int   `json:"orderSource,omitempty"`
	ChannelType int   `json:"channelType,omitempty"`
	OrderType   int   `json:"orderType,omitempty"`
}

// OnEcRightsAgreeRightsEvent
// @id 1175
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=同意售后消息
func (s *Service) OnEcRightsAgreeRightsEvent(handler msg.XinyunEventHandler[EcRightsAgreeRightsEvent]) (service *Service) {
	var listener = &EcRightsAgreeRightsListener{handler}
	s.Register(msg.WrapListener[EcRightsAgreeRightsEvent](listener), "ec_rights", "agreeRights")
	return s
}
