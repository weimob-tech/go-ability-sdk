package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcGoodsAssignStatusChangedListener struct {
	handler msg.XinyunEventHandler[EcGoodsAssignStatusChangedEvent]
}

func (s EcGoodsAssignStatusChangedListener) NewMessage() msg.MsgRequest[EcGoodsAssignStatusChangedEvent] {
	return &msg.XinyunOpenMessage[EcGoodsAssignStatusChangedEvent]{
		MsgBody: &EcGoodsAssignStatusChangedEvent{},
	}
}

func (s EcGoodsAssignStatusChangedListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcGoodsAssignStatusChangedEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcGoodsAssignStatusChangedEvent]))
}

type EcGoodsAssignStatusChangedEvent struct {
	GoodsIdList []int64 `json:"goodsIdList,omitempty"`
	StoreIdList []int   `json:"storeIdList,omitempty"`
	AssignType  int     `json:"assignType,omitempty"`
}

// OnEcGoodsAssignStatusChangedEvent
// @id 1677
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=分配商品逆向消息
func (s *Service) OnEcGoodsAssignStatusChangedEvent(handler msg.XinyunEventHandler[EcGoodsAssignStatusChangedEvent]) (service *Service) {
	var listener = &EcGoodsAssignStatusChangedListener{handler}
	s.Register(msg.WrapListener[EcGoodsAssignStatusChangedEvent](listener), "ec_goods", "assign_status_changed")
	return s
}
