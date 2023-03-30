package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcRightsCompleteRightsListener struct {
	handler msg.XinyunEventHandler[EcRightsCompleteRightsEvent]
}

func (s EcRightsCompleteRightsListener) NewMessage() msg.MsgRequest[EcRightsCompleteRightsEvent] {
	return &msg.XinyunOpenMessage[EcRightsCompleteRightsEvent]{
		MsgBody: &EcRightsCompleteRightsEvent{},
	}
}

func (s EcRightsCompleteRightsListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcRightsCompleteRightsEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcRightsCompleteRightsEvent]))
}

type EcRightsCompleteRightsEvent struct {
	RightsId    int64 `json:"rightsId,omitempty"`
	OrderNo     int64 `json:"orderNo,omitempty"`
	StoreId     int64 `json:"storeId,omitempty"`
	OrderSource int   `json:"orderSource,omitempty"`
	ChannelType int   `json:"channelType,omitempty"`
	OrderType   int   `json:"orderType,omitempty"`
}

// OnEcRightsCompleteRightsEvent
// @id 479
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=完成售后消息
func (s *Service) OnEcRightsCompleteRightsEvent(handler msg.XinyunEventHandler[EcRightsCompleteRightsEvent]) (service *Service) {
	var listener = &EcRightsCompleteRightsListener{handler}
	s.Register(msg.WrapListener[EcRightsCompleteRightsEvent](listener), "ec_rights", "completeRights")
	return s
}
