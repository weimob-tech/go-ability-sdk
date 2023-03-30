package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcRightsCancelRightsListener struct {
	handler msg.XinyunEventHandler[EcRightsCancelRightsEvent]
}

func (s EcRightsCancelRightsListener) NewMessage() msg.MsgRequest[EcRightsCancelRightsEvent] {
	return &msg.XinyunOpenMessage[EcRightsCancelRightsEvent]{
		MsgBody: &EcRightsCancelRightsEvent{},
	}
}

func (s EcRightsCancelRightsListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcRightsCancelRightsEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcRightsCancelRightsEvent]))
}

type EcRightsCancelRightsEvent struct {
	RightsId    int64 `json:"rightsId,omitempty"`
	OrderNo     int64 `json:"orderNo,omitempty"`
	StoreId     int64 `json:"storeId,omitempty"`
	OrderSource int   `json:"orderSource,omitempty"`
	ChannelType int   `json:"channelType,omitempty"`
	OrderType   int   `json:"orderType,omitempty"`
}

// OnEcRightsCancelRightsEvent
// @id 480
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=取消售后消息
func (s *Service) OnEcRightsCancelRightsEvent(handler msg.XinyunEventHandler[EcRightsCancelRightsEvent]) (service *Service) {
	var listener = &EcRightsCancelRightsListener{handler}
	s.Register(msg.WrapListener[EcRightsCancelRightsEvent](listener), "ec_rights", "cancelRights")
	return s
}
