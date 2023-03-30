package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcStockAddInventoryOrderListener struct {
	handler msg.XinyunEventHandler[EcStockAddInventoryOrderEvent]
}

func (s EcStockAddInventoryOrderListener) NewMessage() msg.MsgRequest[EcStockAddInventoryOrderEvent] {
	return &msg.XinyunOpenMessage[EcStockAddInventoryOrderEvent]{
		MsgBody: &EcStockAddInventoryOrderEvent{},
	}
}

func (s EcStockAddInventoryOrderListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcStockAddInventoryOrderEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcStockAddInventoryOrderEvent]))
}

type EcStockAddInventoryOrderEvent struct {
	OrderId       int64  `json:"orderId,omitempty"`
	ReferId       int64  `json:"referId,omitempty"`
	InventoryType int    `json:"inventoryType,omitempty"`
	InventoryName string `json:"inventoryName,omitempty"`
}

// OnEcStockAddInventoryOrderEvent
// @id 591
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=新增出入库单据消息
func (s *Service) OnEcStockAddInventoryOrderEvent(handler msg.XinyunEventHandler[EcStockAddInventoryOrderEvent]) (service *Service) {
	var listener = &EcStockAddInventoryOrderListener{handler}
	s.Register(msg.WrapListener[EcStockAddInventoryOrderEvent](listener), "ec_stock", "add_inventory_order")
	return s
}
