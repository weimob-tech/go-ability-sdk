package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopOrderCreateListener struct {
	handler msg.WosEventHandler[WeimobShopOrderCreateEvent]
}

func (s WeimobShopOrderCreateListener) NewMessage() msg.MsgRequest[WeimobShopOrderCreateEvent] {
	return &msg.WosOpenMessage[WeimobShopOrderCreateEvent]{
		MsgBody: &WeimobShopOrderCreateEvent{},
	}
}

func (s WeimobShopOrderCreateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopOrderCreateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopOrderCreateEvent]))
}

type WeimobShopOrderCreateEvent struct {
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

// OnWeimobShopOrderCreateEvent
// @id 1301
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1301?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=订单新增消息)
func (s *Service) OnWeimobShopOrderCreateEvent(handler msg.WosEventHandler[WeimobShopOrderCreateEvent]) (service *Service) {
	var listener = &WeimobShopOrderCreateListener{handler}
	s.Register(msg.WrapListener[WeimobShopOrderCreateEvent](listener), "weimob_shop.order", "create")
	return s
}
