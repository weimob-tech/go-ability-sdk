package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopOrderUpdateListener struct {
	handler msg.WosEventHandler[WeimobShopOrderUpdateEvent]
}

func (s WeimobShopOrderUpdateListener) NewMessage() msg.MsgRequest[WeimobShopOrderUpdateEvent] {
	return &msg.WosOpenMessage[WeimobShopOrderUpdateEvent]{
		MsgBody: &WeimobShopOrderUpdateEvent{},
	}
}

func (s WeimobShopOrderUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopOrderUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopOrderUpdateEvent]))
}

type WeimobShopOrderUpdateEvent struct {
	BizSourceType     int64 `json:"bizSourceType,omitempty"`
	BosId             int64 `json:"bosId,omitempty"`
	ChannelType       int64 `json:"channelType,omitempty"`
	CreateTime        int64 `json:"createTime,omitempty"`
	DeliveryType      int64 `json:"deliveryType,omitempty"`
	MerchantId        int64 `json:"merchantId,omitempty"`
	OrderNo           int64 `json:"orderNo,omitempty"`
	OrderSource       int64 `json:"orderSource,omitempty"`
	OrderStatus       int64 `json:"orderStatus,omitempty"`
	OrderType         int64 `json:"orderType,omitempty"`
	ParentOrderNo     int64 `json:"parentOrderNo,omitempty"`
	PaymentStatus     int64 `json:"paymentStatus,omitempty"`
	PaymentType       int64 `json:"paymentType,omitempty"`
	ProductId         int64 `json:"productId,omitempty"`
	ProductInstanceId int64 `json:"productInstanceId,omitempty"`
	SaleChannelType   int64 `json:"saleChannelType,omitempty"`
	UpdateTime        int64 `json:"updateTime,omitempty"`
	UpdateType        int64 `json:"updateType,omitempty"`
	Vid               int64 `json:"vid,omitempty"`
	Wid               int64 `json:"wid,omitempty"`
}

// OnWeimobShopOrderUpdateEvent
// @id 1302
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1302?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=订单内容变更消息)
func (s *Service) OnWeimobShopOrderUpdateEvent(handler msg.WosEventHandler[WeimobShopOrderUpdateEvent]) (service *Service) {
	var listener = &WeimobShopOrderUpdateListener{handler}
	s.Register(msg.WrapListener[WeimobShopOrderUpdateEvent](listener), "weimob_shop.order", "update")
	return s
}
