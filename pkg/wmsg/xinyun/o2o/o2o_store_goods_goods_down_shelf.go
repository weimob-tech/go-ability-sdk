package o2o

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type O2oStoreGoodsGoodsDownShelfListener struct {
	handler msg.XinyunEventHandler[O2oStoreGoodsGoodsDownShelfEvent]
}

func (s O2oStoreGoodsGoodsDownShelfListener) NewMessage() msg.MsgRequest[O2oStoreGoodsGoodsDownShelfEvent] {
	return &msg.XinyunOpenMessage[O2oStoreGoodsGoodsDownShelfEvent]{
		MsgBody: &O2oStoreGoodsGoodsDownShelfEvent{},
	}
}

func (s O2oStoreGoodsGoodsDownShelfListener) OnMessage(ctx context.Context, message msg.MsgRequest[O2oStoreGoodsGoodsDownShelfEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[O2oStoreGoodsGoodsDownShelfEvent]))
}

type O2oStoreGoodsGoodsDownShelfEvent struct {
	GoodsId      int64  `json:"goodsId,omitempty"`
	StoreId      int64  `json:"storeId,omitempty"`
	ThirdGoodsId string `json:"thirdGoodsId,omitempty"`
}

// OnO2oStoreGoodsGoodsDownShelfEvent
// @id 194
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=菜品下架
func (s *Service) OnO2oStoreGoodsGoodsDownShelfEvent(handler msg.XinyunEventHandler[O2oStoreGoodsGoodsDownShelfEvent]) (service *Service) {
	var listener = &O2oStoreGoodsGoodsDownShelfListener{handler}
	s.Register(msg.WrapListener[O2oStoreGoodsGoodsDownShelfEvent](listener), "o2o_store_goods", "goodsDownShelf")
	return s
}
