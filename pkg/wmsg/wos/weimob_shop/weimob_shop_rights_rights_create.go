package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopRightsRightsCreateListener struct {
	handler msg.WosEventHandler[WeimobShopRightsRightsCreateEvent]
}

func (s WeimobShopRightsRightsCreateListener) NewMessage() msg.MsgRequest[WeimobShopRightsRightsCreateEvent] {
	return &msg.WosOpenMessage[WeimobShopRightsRightsCreateEvent]{
		MsgBody: &WeimobShopRightsRightsCreateEvent{},
	}
}

func (s WeimobShopRightsRightsCreateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopRightsRightsCreateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopRightsRightsCreateEvent]))
}

type WeimobShopRightsRightsCreateEvent struct {
	RightsInfo    WeimobShopRightsRightsCreateRightsInfo `json:"rightsInfo,omitempty"`
	OrderInfo     WeimobShopRightsRightsCreateOrderInfo  `json:"orderInfo,omitempty"`
	RightsOrderNo int64                                  `json:"rightsOrderNo,omitempty"`
	OrderNo       int64                                  `json:"orderNo,omitempty"`
}

type WeimobShopRightsRightsCreateRightsInfo struct {
	BosId             int64 `json:"bosId,omitempty"`
	MerchantId        int64 `json:"merchantId,omitempty"`
	ProductId         int64 `json:"productId,omitempty"`
	ProductInstanceId int64 `json:"productInstanceId,omitempty"`
	Vid               int64 `json:"vid,omitempty"`
	ProcessVid        int64 `json:"processVid,omitempty"`
	RightsType        int64 `json:"rightsType,omitempty"`
	RightsStatus      int64 `json:"rightsStatus,omitempty"`
	RightsCauseType   int64 `json:"rightsCauseType,omitempty"`
	RightsSource      int64 `json:"rightsSource,omitempty"`
	RefundType        int64 `json:"refundType,omitempty"`
}
type WeimobShopRightsRightsCreateOrderInfo struct {
	ParentOrderNo int64 `json:"parentOrderNo,omitempty"`
	OrderType     int64 `json:"orderType,omitempty"`
	OrderSource   int64 `json:"orderSource,omitempty"`
	ChannelType   int64 `json:"channelType,omitempty"`
	BizSourceType int64 `json:"bizSourceType,omitempty"`
}

// OnWeimobShopRightsRightsCreateEvent
// @id 1219
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1219?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=售后创建消息)
func (s *Service) OnWeimobShopRightsRightsCreateEvent(handler msg.WosEventHandler[WeimobShopRightsRightsCreateEvent]) (service *Service) {
	var listener = &WeimobShopRightsRightsCreateListener{handler}
	s.Register(msg.WrapListener[WeimobShopRightsRightsCreateEvent](listener), "weimob_shop.rights", "rightsCreate")
	return s
}
