package o2o

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type O2oStoreGoodsGoodsUpShelfListener struct {
	handler msg.XinyunEventHandler[O2oStoreGoodsGoodsUpShelfEvent]
}

func (s O2oStoreGoodsGoodsUpShelfListener) NewMessage() msg.MsgRequest[O2oStoreGoodsGoodsUpShelfEvent] {
	return &msg.XinyunOpenMessage[O2oStoreGoodsGoodsUpShelfEvent]{
		MsgBody: &O2oStoreGoodsGoodsUpShelfEvent{},
	}
}

func (s O2oStoreGoodsGoodsUpShelfListener) OnMessage(ctx context.Context, message msg.MsgRequest[O2oStoreGoodsGoodsUpShelfEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[O2oStoreGoodsGoodsUpShelfEvent]))
}

type O2oStoreGoodsGoodsUpShelfEvent struct {
	GoodsId      int64  `json:"goodsId,omitempty"`
	StoreId      int64  `json:"storeId,omitempty"`
	ThirdGoodsId string `json:"thirdGoodsId,omitempty"`
}

// OnO2oStoreGoodsGoodsUpShelfEvent
// @id 193
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=菜品上架
func (s *Service) OnO2oStoreGoodsGoodsUpShelfEvent(handler msg.XinyunEventHandler[O2oStoreGoodsGoodsUpShelfEvent]) (service *Service) {
	var listener = &O2oStoreGoodsGoodsUpShelfListener{handler}
	s.Register(msg.WrapListener[O2oStoreGoodsGoodsUpShelfEvent](listener), "o2o_store_goods", "goodsUpShelf")
	return s
}
