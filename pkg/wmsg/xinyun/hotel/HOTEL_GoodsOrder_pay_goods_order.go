package hotel

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type HOTELGoodsOrderPayGoodsOrderListener struct {
	handler msg.XinyunEventHandler[HOTELGoodsOrderPayGoodsOrderEvent]
}

func (s HOTELGoodsOrderPayGoodsOrderListener) NewMessage() msg.MsgRequest[HOTELGoodsOrderPayGoodsOrderEvent] {
	return &msg.XinyunOpenMessage[HOTELGoodsOrderPayGoodsOrderEvent]{
		MsgBody: &HOTELGoodsOrderPayGoodsOrderEvent{},
	}
}

func (s HOTELGoodsOrderPayGoodsOrderListener) OnMessage(ctx context.Context, message msg.MsgRequest[HOTELGoodsOrderPayGoodsOrderEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[HOTELGoodsOrderPayGoodsOrderEvent]))
}

type HOTELGoodsOrderPayGoodsOrderEvent struct {
	GoodsId    string `json:"goodsId,omitempty"`
	GoodsName  string `json:"goodsName,omitempty"`
	Type       string `json:"type,omitempty"`
	Count      string `json:"count,omitempty"`
	RealAmount string `json:"realAmount,omitempty"`
	OrderNo    string `json:"orderNo,omitempty"`
}

// OnHOTELGoodsOrderPayGoodsOrderEvent
// @id 488
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=商品订单支付成功提醒
func (s *Service) OnHOTELGoodsOrderPayGoodsOrderEvent(handler msg.XinyunEventHandler[HOTELGoodsOrderPayGoodsOrderEvent]) (service *Service) {
	var listener = &HOTELGoodsOrderPayGoodsOrderListener{handler}
	s.Register(msg.WrapListener[HOTELGoodsOrderPayGoodsOrderEvent](listener), "HOTEL_GoodsOrder", "payGoodsOrder")
	return s
}
