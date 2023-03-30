package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcRightsReturnRightsGoodsListener struct {
	handler msg.XinyunEventHandler[EcRightsReturnRightsGoodsEvent]
}

func (s EcRightsReturnRightsGoodsListener) NewMessage() msg.MsgRequest[EcRightsReturnRightsGoodsEvent] {
	return &msg.XinyunOpenMessage[EcRightsReturnRightsGoodsEvent]{
		MsgBody: &EcRightsReturnRightsGoodsEvent{},
	}
}

func (s EcRightsReturnRightsGoodsListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcRightsReturnRightsGoodsEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcRightsReturnRightsGoodsEvent]))
}

type EcRightsReturnRightsGoodsEvent struct {
	RightsId    int64 `json:"rightsId,omitempty"`
	OrderNo     int64 `json:"orderNo,omitempty"`
	StoreId     int64 `json:"storeId,omitempty"`
	OrderSource int   `json:"orderSource,omitempty"`
	ChannelType int   `json:"channelType,omitempty"`
	OrderType   int   `json:"orderType,omitempty"`
}

// OnEcRightsReturnRightsGoodsEvent
// @id 1252
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=售后退货消息
func (s *Service) OnEcRightsReturnRightsGoodsEvent(handler msg.XinyunEventHandler[EcRightsReturnRightsGoodsEvent]) (service *Service) {
	var listener = &EcRightsReturnRightsGoodsListener{handler}
	s.Register(msg.WrapListener[EcRightsReturnRightsGoodsEvent](listener), "ec_rights", "returnRightsGoods")
	return s
}
