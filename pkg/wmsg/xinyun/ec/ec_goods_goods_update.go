package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcGoodsGoodsUpdateListener struct {
	handler msg.XinyunEventHandler[EcGoodsGoodsUpdateEvent]
}

func (s EcGoodsGoodsUpdateListener) NewMessage() msg.MsgRequest[EcGoodsGoodsUpdateEvent] {
	return &msg.XinyunOpenMessage[EcGoodsGoodsUpdateEvent]{
		MsgBody: &EcGoodsGoodsUpdateEvent{},
	}
}

func (s EcGoodsGoodsUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcGoodsGoodsUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcGoodsGoodsUpdateEvent]))
}

type EcGoodsGoodsUpdateEvent struct {
	GoodsIdList string `json:"goodsIdList,omitempty"`
}

// OnEcGoodsGoodsUpdateEvent
// @id 1324
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=修改商品消息
func (s *Service) OnEcGoodsGoodsUpdateEvent(handler msg.XinyunEventHandler[EcGoodsGoodsUpdateEvent]) (service *Service) {
	var listener = &EcGoodsGoodsUpdateListener{handler}
	s.Register(msg.WrapListener[EcGoodsGoodsUpdateEvent](listener), "ec_goods", "goodsUpdate")
	return s
}
