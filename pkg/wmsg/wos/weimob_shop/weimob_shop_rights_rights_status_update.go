package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopRightsRightsStatusUpdateListener struct {
	handler msg.WosEventHandler[WeimobShopRightsRightsStatusUpdateEvent]
}

func (s WeimobShopRightsRightsStatusUpdateListener) NewMessage() msg.MsgRequest[WeimobShopRightsRightsStatusUpdateEvent] {
	return &msg.WosOpenMessage[WeimobShopRightsRightsStatusUpdateEvent]{
		MsgBody: &WeimobShopRightsRightsStatusUpdateEvent{},
	}
}

func (s WeimobShopRightsRightsStatusUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopRightsRightsStatusUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopRightsRightsStatusUpdateEvent]))
}

type WeimobShopRightsRightsStatusUpdateEvent struct {
	RightsInfo    WeimobShopRightsRightsStatusUpdateRightsInfo `json:"rightsInfo,omitempty"`
	OrderInfo     WeimobShopRightsRightsStatusUpdateOrderInfo  `json:"orderInfo,omitempty"`
	RightsOrderNo int64                                        `json:"rightsOrderNo,omitempty"`
	OrderNo       int64                                        `json:"orderNo,omitempty"`
	StatusBefore  int64                                        `json:"statusBefore,omitempty"`
	Status        int64                                        `json:"status,omitempty"`
}

type WeimobShopRightsRightsStatusUpdateRightsInfo struct {
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
type WeimobShopRightsRightsStatusUpdateOrderInfo struct {
	ParentOrderNo int64 `json:"parentOrderNo,omitempty"`
	OrderType     int64 `json:"orderType,omitempty"`
	OrderSource   int64 `json:"orderSource,omitempty"`
	ChannelType   int64 `json:"channelType,omitempty"`
	BizSourceType int64 `json:"bizSourceType,omitempty"`
}

// OnWeimobShopRightsRightsStatusUpdateEvent
// @id 1218
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1218?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=售后状态变化消息)
func (s *Service) OnWeimobShopRightsRightsStatusUpdateEvent(handler msg.WosEventHandler[WeimobShopRightsRightsStatusUpdateEvent]) (service *Service) {
	var listener = &WeimobShopRightsRightsStatusUpdateListener{handler}
	s.Register(msg.WrapListener[WeimobShopRightsRightsStatusUpdateEvent](listener), "weimob_shop.rights", "rightsStatusUpdate")
	return s
}
