package hotel

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type HOTELGoodsOrderGoodsConsumedListener struct {
	handler msg.XinyunEventHandler[HOTELGoodsOrderGoodsConsumedEvent]
}

func (s HOTELGoodsOrderGoodsConsumedListener) NewMessage() msg.MsgRequest[HOTELGoodsOrderGoodsConsumedEvent] {
	return &msg.XinyunOpenMessage[HOTELGoodsOrderGoodsConsumedEvent]{
		MsgBody: &HOTELGoodsOrderGoodsConsumedEvent{},
	}
}

func (s HOTELGoodsOrderGoodsConsumedListener) OnMessage(ctx context.Context, message msg.MsgRequest[HOTELGoodsOrderGoodsConsumedEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[HOTELGoodsOrderGoodsConsumedEvent]))
}

type HOTELGoodsOrderGoodsConsumedEvent struct {
	SnNo     string `json:"snNo,omitempty"`
	StoreId  string `json:"storeId,omitempty"`
	Operator string `json:"operator,omitempty"`
}

// OnHOTELGoodsOrderGoodsConsumedEvent
// @id 490
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=商品核销提醒
func (s *Service) OnHOTELGoodsOrderGoodsConsumedEvent(handler msg.XinyunEventHandler[HOTELGoodsOrderGoodsConsumedEvent]) (service *Service) {
	var listener = &HOTELGoodsOrderGoodsConsumedListener{handler}
	s.Register(msg.WrapListener[HOTELGoodsOrderGoodsConsumedEvent](listener), "HOTEL_GoodsOrder", "goodsConsumed")
	return s
}
