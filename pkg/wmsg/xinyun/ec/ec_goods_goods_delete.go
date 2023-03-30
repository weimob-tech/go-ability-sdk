package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcGoodsGoodsDeleteListener struct {
	handler msg.XinyunEventHandler[EcGoodsGoodsDeleteEvent]
}

func (s EcGoodsGoodsDeleteListener) NewMessage() msg.MsgRequest[EcGoodsGoodsDeleteEvent] {
	return &msg.XinyunOpenMessage[EcGoodsGoodsDeleteEvent]{
		MsgBody: &EcGoodsGoodsDeleteEvent{},
	}
}

func (s EcGoodsGoodsDeleteListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcGoodsGoodsDeleteEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcGoodsGoodsDeleteEvent]))
}

type EcGoodsGoodsDeleteEvent struct {
	GoodsIdList string `json:"goodsIdList,omitempty"`
}

// OnEcGoodsGoodsDeleteEvent
// @id 1273
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=删除商品消息
func (s *Service) OnEcGoodsGoodsDeleteEvent(handler msg.XinyunEventHandler[EcGoodsGoodsDeleteEvent]) (service *Service) {
	var listener = &EcGoodsGoodsDeleteListener{handler}
	s.Register(msg.WrapListener[EcGoodsGoodsDeleteEvent](listener), "ec_goods", "goodsDelete")
	return s
}
