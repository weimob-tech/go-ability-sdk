package o2o

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type O2oStoreGoodsGoodsDeleteListener struct {
	handler msg.XinyunEventHandler[O2oStoreGoodsGoodsDeleteEvent]
}

func (s O2oStoreGoodsGoodsDeleteListener) NewMessage() msg.MsgRequest[O2oStoreGoodsGoodsDeleteEvent] {
	return &msg.XinyunOpenMessage[O2oStoreGoodsGoodsDeleteEvent]{
		MsgBody: &O2oStoreGoodsGoodsDeleteEvent{},
	}
}

func (s O2oStoreGoodsGoodsDeleteListener) OnMessage(ctx context.Context, message msg.MsgRequest[O2oStoreGoodsGoodsDeleteEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[O2oStoreGoodsGoodsDeleteEvent]))
}

type O2oStoreGoodsGoodsDeleteEvent struct {
	GoodsId      int64  `json:"goodsId,omitempty"`
	StoreId      int64  `json:"storeId,omitempty"`
	ThirdGoodsId string `json:"thirdGoodsId,omitempty"`
}

// OnO2oStoreGoodsGoodsDeleteEvent
// @id 191
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=菜品删除
func (s *Service) OnO2oStoreGoodsGoodsDeleteEvent(handler msg.XinyunEventHandler[O2oStoreGoodsGoodsDeleteEvent]) (service *Service) {
	var listener = &O2oStoreGoodsGoodsDeleteListener{handler}
	s.Register(msg.WrapListener[O2oStoreGoodsGoodsDeleteEvent](listener), "o2o_store_goods", "goodsDelete")
	return s
}
