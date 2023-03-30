package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcRightsAgreeReturnRightsGoodsListener struct {
	handler msg.XinyunEventHandler[EcRightsAgreeReturnRightsGoodsEvent]
}

func (s EcRightsAgreeReturnRightsGoodsListener) NewMessage() msg.MsgRequest[EcRightsAgreeReturnRightsGoodsEvent] {
	return &msg.XinyunOpenMessage[EcRightsAgreeReturnRightsGoodsEvent]{
		MsgBody: &EcRightsAgreeReturnRightsGoodsEvent{},
	}
}

func (s EcRightsAgreeReturnRightsGoodsListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcRightsAgreeReturnRightsGoodsEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcRightsAgreeReturnRightsGoodsEvent]))
}

type EcRightsAgreeReturnRightsGoodsEvent struct {
	RightsId    string `json:"rightsId,omitempty"`
	OrderNo     string `json:"orderNo,omitempty"`
	StoreId     string `json:"storeId,omitempty"`
	OrderSource int    `json:"orderSource,omitempty"`
	ChannelType int    `json:"channelType,omitempty"`
	OrderType   int    `json:"orderType,omitempty"`
}

// OnEcRightsAgreeReturnRightsGoodsEvent
// @id 1608
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=售后同意退货消息
func (s *Service) OnEcRightsAgreeReturnRightsGoodsEvent(handler msg.XinyunEventHandler[EcRightsAgreeReturnRightsGoodsEvent]) (service *Service) {
	var listener = &EcRightsAgreeReturnRightsGoodsListener{handler}
	s.Register(msg.WrapListener[EcRightsAgreeReturnRightsGoodsEvent](listener), "ec_rights", "agreeReturnRightsGoods")
	return s
}
