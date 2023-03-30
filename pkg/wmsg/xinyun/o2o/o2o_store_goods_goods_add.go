package o2o

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type O2oStoreGoodsGoodsAddListener struct {
	handler msg.XinyunEventHandler[O2oStoreGoodsGoodsAddEvent]
}

func (s O2oStoreGoodsGoodsAddListener) NewMessage() msg.MsgRequest[O2oStoreGoodsGoodsAddEvent] {
	return &msg.XinyunOpenMessage[O2oStoreGoodsGoodsAddEvent]{
		MsgBody: &O2oStoreGoodsGoodsAddEvent{},
	}
}

func (s O2oStoreGoodsGoodsAddListener) OnMessage(ctx context.Context, message msg.MsgRequest[O2oStoreGoodsGoodsAddEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[O2oStoreGoodsGoodsAddEvent]))
}

type O2oStoreGoodsGoodsAddEvent struct {
	GoodsId      int64  `json:"goodsId,omitempty"`
	StoreId      int64  `json:"storeId,omitempty"`
	ThirdGoodsId string `json:"thirdGoodsId,omitempty"`
}

// OnO2oStoreGoodsGoodsAddEvent
// @id 189
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=菜品新增消息
func (s *Service) OnO2oStoreGoodsGoodsAddEvent(handler msg.XinyunEventHandler[O2oStoreGoodsGoodsAddEvent]) (service *Service) {
	var listener = &O2oStoreGoodsGoodsAddListener{handler}
	s.Register(msg.WrapListener[O2oStoreGoodsGoodsAddEvent](listener), "o2o_store_goods", "goodsAdd")
	return s
}
