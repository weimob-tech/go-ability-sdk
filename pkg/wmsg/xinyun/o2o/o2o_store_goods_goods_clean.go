package o2o

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type O2oStoreGoodsGoodsCleanListener struct {
	handler msg.XinyunEventHandler[O2oStoreGoodsGoodsCleanEvent]
}

func (s O2oStoreGoodsGoodsCleanListener) NewMessage() msg.MsgRequest[O2oStoreGoodsGoodsCleanEvent] {
	return &msg.XinyunOpenMessage[O2oStoreGoodsGoodsCleanEvent]{
		MsgBody: &O2oStoreGoodsGoodsCleanEvent{},
	}
}

func (s O2oStoreGoodsGoodsCleanListener) OnMessage(ctx context.Context, message msg.MsgRequest[O2oStoreGoodsGoodsCleanEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[O2oStoreGoodsGoodsCleanEvent]))
}

type O2oStoreGoodsGoodsCleanEvent struct {
	GoodsId      int64  `json:"goodsId,omitempty"`
	StoreId      int64  `json:"storeId,omitempty"`
	ThirdGoodsId string `json:"thirdGoodsId,omitempty"`
}

// OnO2oStoreGoodsGoodsCleanEvent
// @id 192
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=菜品估清
func (s *Service) OnO2oStoreGoodsGoodsCleanEvent(handler msg.XinyunEventHandler[O2oStoreGoodsGoodsCleanEvent]) (service *Service) {
	var listener = &O2oStoreGoodsGoodsCleanListener{handler}
	s.Register(msg.WrapListener[O2oStoreGoodsGoodsCleanEvent](listener), "o2o_store_goods", "goodsClean")
	return s
}
