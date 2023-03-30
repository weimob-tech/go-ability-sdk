package hotel

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type HOTELBookGoodsOrderAddBookGoodsOrderListener struct {
	handler msg.XinyunEventHandler[HOTELBookGoodsOrderAddBookGoodsOrderEvent]
}

func (s HOTELBookGoodsOrderAddBookGoodsOrderListener) NewMessage() msg.MsgRequest[HOTELBookGoodsOrderAddBookGoodsOrderEvent] {
	return &msg.XinyunOpenMessage[HOTELBookGoodsOrderAddBookGoodsOrderEvent]{
		MsgBody: &HOTELBookGoodsOrderAddBookGoodsOrderEvent{},
	}
}

func (s HOTELBookGoodsOrderAddBookGoodsOrderListener) OnMessage(ctx context.Context, message msg.MsgRequest[HOTELBookGoodsOrderAddBookGoodsOrderEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[HOTELBookGoodsOrderAddBookGoodsOrderEvent]))
}

type HOTELBookGoodsOrderAddBookGoodsOrderEvent struct {
	OrderNo     string `json:"orderNo,omitempty"`
	GoodsName   string `json:"goodsName,omitempty"`
	Count       string `json:"count,omitempty"`
	Type        string `json:"type,omitempty"`
	GoodsCharge string `json:"goodsCharge,omitempty"`
	StoreId     string `json:"storeId,omitempty"`
}

// OnHOTELBookGoodsOrderAddBookGoodsOrderEvent
// @id 1439
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=预约商品订单下单成功消息
func (s *Service) OnHOTELBookGoodsOrderAddBookGoodsOrderEvent(handler msg.XinyunEventHandler[HOTELBookGoodsOrderAddBookGoodsOrderEvent]) (service *Service) {
	var listener = &HOTELBookGoodsOrderAddBookGoodsOrderListener{handler}
	s.Register(msg.WrapListener[HOTELBookGoodsOrderAddBookGoodsOrderEvent](listener), "HOTEL_BookGoodsOrder", "addBookGoodsOrder")
	return s
}
