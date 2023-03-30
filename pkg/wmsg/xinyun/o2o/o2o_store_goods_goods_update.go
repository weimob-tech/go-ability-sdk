package o2o

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type O2oStoreGoodsGoodsUpdateListener struct {
	handler msg.XinyunEventHandler[O2oStoreGoodsGoodsUpdateEvent]
}

func (s O2oStoreGoodsGoodsUpdateListener) NewMessage() msg.MsgRequest[O2oStoreGoodsGoodsUpdateEvent] {
	return &msg.XinyunOpenMessage[O2oStoreGoodsGoodsUpdateEvent]{
		MsgBody: &O2oStoreGoodsGoodsUpdateEvent{},
	}
}

func (s O2oStoreGoodsGoodsUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[O2oStoreGoodsGoodsUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[O2oStoreGoodsGoodsUpdateEvent]))
}

type O2oStoreGoodsGoodsUpdateEvent struct {
	GoodsId      string `json:"goodsId,omitempty"`
	StoreId      int64  `json:"storeId,omitempty"`
	ThirdGoodsId string `json:"thirdGoodsId,omitempty"`
}

// OnO2oStoreGoodsGoodsUpdateEvent
// @id 190
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=菜品更新
func (s *Service) OnO2oStoreGoodsGoodsUpdateEvent(handler msg.XinyunEventHandler[O2oStoreGoodsGoodsUpdateEvent]) (service *Service) {
	var listener = &O2oStoreGoodsGoodsUpdateListener{handler}
	s.Register(msg.WrapListener[O2oStoreGoodsGoodsUpdateEvent](listener), "o2o_store_goods", "goodsUpdate")
	return s
}
