package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopRightsStatusUpdateListener struct {
	handler msg.WosEventHandler[WeimobShopRightsStatusUpdateEvent]
}

func (s WeimobShopRightsStatusUpdateListener) NewMessage() msg.MsgRequest[WeimobShopRightsStatusUpdateEvent] {
	return &msg.WosOpenMessage[WeimobShopRightsStatusUpdateEvent]{
		MsgBody: &WeimobShopRightsStatusUpdateEvent{},
	}
}

func (s WeimobShopRightsStatusUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopRightsStatusUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopRightsStatusUpdateEvent]))
}

type WeimobShopRightsStatusUpdateEvent struct {
	RightsInfo   WeimobShopRightsStatusUpdateRightsInfo `json:"rightsInfo,omitempty"`
	OrderInfo    WeimobShopRightsStatusUpdateOrderInfo  `json:"orderInfo,omitempty"`
	RightsId     int64                                  `json:"rightsId,omitempty"`
	OrderNo      int64                                  `json:"orderNo,omitempty"`
	StatusBefore int64                                  `json:"statusBefore,omitempty"`
	Status       int64                                  `json:"status,omitempty"`
}

type WeimobShopRightsStatusUpdateRightsInfo struct {
	Vid             int64 `json:"vid,omitempty"`
	ProcessVid      int64 `json:"processVid,omitempty"`
	RightsType      int64 `json:"rightsType,omitempty"`
	RightsStatus    int64 `json:"rightsStatus,omitempty"`
	RightsCauseType int64 `json:"rightsCauseType,omitempty"`
	RightsSource    int64 `json:"rightsSource,omitempty"`
	RefundType      int64 `json:"refundType,omitempty"`
}
type WeimobShopRightsStatusUpdateOrderInfo struct {
	ParentOrderNo int64 `json:"parentOrderNo,omitempty"`
	OrderType     int64 `json:"orderType,omitempty"`
	OrderSource   int64 `json:"orderSource,omitempty"`
	ChannelType   int64 `json:"channelType,omitempty"`
	BizSourceType int64 `json:"bizSourceType,omitempty"`
}

// OnWeimobShopRightsStatusUpdateEvent
// @id 1304
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1304?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=售后状态变更消息)
func (s *Service) OnWeimobShopRightsStatusUpdateEvent(handler msg.WosEventHandler[WeimobShopRightsStatusUpdateEvent]) (service *Service) {
	var listener = &WeimobShopRightsStatusUpdateListener{handler}
	s.Register(msg.WrapListener[WeimobShopRightsStatusUpdateEvent](listener), "weimob_shop.rights", "statusUpdate")
	return s
}
