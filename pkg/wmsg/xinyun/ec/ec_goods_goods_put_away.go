package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcGoodsGoodsPutAwayListener struct {
	handler msg.XinyunEventHandler[EcGoodsGoodsPutAwayEvent]
}

func (s EcGoodsGoodsPutAwayListener) NewMessage() msg.MsgRequest[EcGoodsGoodsPutAwayEvent] {
	return &msg.XinyunOpenMessage[EcGoodsGoodsPutAwayEvent]{
		MsgBody: &EcGoodsGoodsPutAwayEvent{},
	}
}

func (s EcGoodsGoodsPutAwayListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcGoodsGoodsPutAwayEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcGoodsGoodsPutAwayEvent]))
}

type EcGoodsGoodsPutAwayEvent struct {
	GoodsIdList []int64 `json:"goodsIdList,omitempty"`
	PutAway     string  `json:"putAway,omitempty"`
	StoreId     string  `json:"storeId,omitempty"`
}

// OnEcGoodsGoodsPutAwayEvent
// @id 1896
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=商品上下架
func (s *Service) OnEcGoodsGoodsPutAwayEvent(handler msg.XinyunEventHandler[EcGoodsGoodsPutAwayEvent]) (service *Service) {
	var listener = &EcGoodsGoodsPutAwayListener{handler}
	s.Register(msg.WrapListener[EcGoodsGoodsPutAwayEvent](listener), "ec_goods", "goodsPutAway")
	return s
}
