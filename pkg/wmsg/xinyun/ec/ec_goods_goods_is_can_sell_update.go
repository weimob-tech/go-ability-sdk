package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcGoodsGoodsIsCanSellUpdateListener struct {
	handler msg.XinyunEventHandler[EcGoodsGoodsIsCanSellUpdateEvent]
}

func (s EcGoodsGoodsIsCanSellUpdateListener) NewMessage() msg.MsgRequest[EcGoodsGoodsIsCanSellUpdateEvent] {
	return &msg.XinyunOpenMessage[EcGoodsGoodsIsCanSellUpdateEvent]{
		MsgBody: &EcGoodsGoodsIsCanSellUpdateEvent{},
	}
}

func (s EcGoodsGoodsIsCanSellUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcGoodsGoodsIsCanSellUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcGoodsGoodsIsCanSellUpdateEvent]))
}

type EcGoodsGoodsIsCanSellUpdateEvent struct {
	GoodsIdList []int64 `json:"goodsIdList,omitempty"`
	IsCanSell   string  `json:"isCanSell,omitempty"`
}

// OnEcGoodsGoodsIsCanSellUpdateEvent
// @id 1897
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=商品禁售
func (s *Service) OnEcGoodsGoodsIsCanSellUpdateEvent(handler msg.XinyunEventHandler[EcGoodsGoodsIsCanSellUpdateEvent]) (service *Service) {
	var listener = &EcGoodsGoodsIsCanSellUpdateListener{handler}
	s.Register(msg.WrapListener[EcGoodsGoodsIsCanSellUpdateEvent](listener), "ec_goods", "goodsIsCanSellUpdate")
	return s
}
