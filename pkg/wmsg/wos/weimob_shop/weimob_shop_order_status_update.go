package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopOrderStatusUpdateListener struct {
	handler msg.WosEventHandler[WeimobShopOrderStatusUpdateEvent]
}

func (s WeimobShopOrderStatusUpdateListener) NewMessage() msg.MsgRequest[WeimobShopOrderStatusUpdateEvent] {
	return &msg.WosOpenMessage[WeimobShopOrderStatusUpdateEvent]{
		MsgBody: &WeimobShopOrderStatusUpdateEvent{},
	}
}

func (s WeimobShopOrderStatusUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopOrderStatusUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopOrderStatusUpdateEvent]))
}

type WeimobShopOrderStatusUpdateEvent struct {
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
	Vid               int64 `json:"vid,omitempty"`
	Wid               int64 `json:"wid,omitempty"`
}

// OnWeimobShopOrderStatusUpdateEvent
// @id 1300
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1300?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=订单状态变更消息)
func (s *Service) OnWeimobShopOrderStatusUpdateEvent(handler msg.WosEventHandler[WeimobShopOrderStatusUpdateEvent]) (service *Service) {
	var listener = &WeimobShopOrderStatusUpdateListener{handler}
	s.Register(msg.WrapListener[WeimobShopOrderStatusUpdateEvent](listener), "weimob_shop.order", "statusUpdate")
	return s
}
