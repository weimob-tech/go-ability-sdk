package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcPricePriceUpdateListener struct {
	handler msg.XinyunEventHandler[EcPricePriceUpdateEvent]
}

func (s EcPricePriceUpdateListener) NewMessage() msg.MsgRequest[EcPricePriceUpdateEvent] {
	return &msg.XinyunOpenMessage[EcPricePriceUpdateEvent]{
		MsgBody: &EcPricePriceUpdateEvent{},
	}
}

func (s EcPricePriceUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcPricePriceUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcPricePriceUpdateEvent]))
}

type EcPricePriceUpdateEvent struct {
	GoodsIdList int64 `json:"goodsIdList,omitempty"`
	StoreId     int64 `json:"storeId,omitempty"`
}

// OnEcPricePriceUpdateEvent
// @id 548
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=商品销售价变更
func (s *Service) OnEcPricePriceUpdateEvent(handler msg.XinyunEventHandler[EcPricePriceUpdateEvent]) (service *Service) {
	var listener = &EcPricePriceUpdateListener{handler}
	s.Register(msg.WrapListener[EcPricePriceUpdateEvent](listener), "ec_price", "price_update")
	return s
}
