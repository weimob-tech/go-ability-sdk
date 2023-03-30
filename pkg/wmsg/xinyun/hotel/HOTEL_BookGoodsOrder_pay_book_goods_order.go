package hotel

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type HOTELBookGoodsOrderPayBookGoodsOrderListener struct {
	handler msg.XinyunEventHandler[HOTELBookGoodsOrderPayBookGoodsOrderEvent]
}

func (s HOTELBookGoodsOrderPayBookGoodsOrderListener) NewMessage() msg.MsgRequest[HOTELBookGoodsOrderPayBookGoodsOrderEvent] {
	return &msg.XinyunOpenMessage[HOTELBookGoodsOrderPayBookGoodsOrderEvent]{
		MsgBody: &HOTELBookGoodsOrderPayBookGoodsOrderEvent{},
	}
}

func (s HOTELBookGoodsOrderPayBookGoodsOrderListener) OnMessage(ctx context.Context, message msg.MsgRequest[HOTELBookGoodsOrderPayBookGoodsOrderEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[HOTELBookGoodsOrderPayBookGoodsOrderEvent]))
}

type HOTELBookGoodsOrderPayBookGoodsOrderEvent struct {
	OrderNo    string `json:"orderNo,omitempty"`
	GoodsName  string `json:"goodsName,omitempty"`
	Count      string `json:"count,omitempty"`
	RealAmount string `json:"realAmount,omitempty"`
	StoreId    string `json:"storeId,omitempty"`
}

// OnHOTELBookGoodsOrderPayBookGoodsOrderEvent
// @id 1440
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=预约商品订单支付成功消息
func (s *Service) OnHOTELBookGoodsOrderPayBookGoodsOrderEvent(handler msg.XinyunEventHandler[HOTELBookGoodsOrderPayBookGoodsOrderEvent]) (service *Service) {
	var listener = &HOTELBookGoodsOrderPayBookGoodsOrderListener{handler}
	s.Register(msg.WrapListener[HOTELBookGoodsOrderPayBookGoodsOrderEvent](listener), "HOTEL_BookGoodsOrder", "payBookGoodsOrder")
	return s
}
