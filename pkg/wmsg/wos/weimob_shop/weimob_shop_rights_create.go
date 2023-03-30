package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopRightsCreateListener struct {
	handler msg.WosEventHandler[WeimobShopRightsCreateEvent]
}

func (s WeimobShopRightsCreateListener) NewMessage() msg.MsgRequest[WeimobShopRightsCreateEvent] {
	return &msg.WosOpenMessage[WeimobShopRightsCreateEvent]{
		MsgBody: &WeimobShopRightsCreateEvent{},
	}
}

func (s WeimobShopRightsCreateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopRightsCreateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopRightsCreateEvent]))
}

type WeimobShopRightsCreateEvent struct {
	RightsInfo WeimobShopRightsCreateRightsInfo `json:"rightsInfo,omitempty"`
	OrderInfo  WeimobShopRightsCreateOrderInfo  `json:"orderInfo,omitempty"`
	RightsId   int64                            `json:"rightsId,omitempty"`
	OrderNo    int64                            `json:"orderNo,omitempty"`
}

type WeimobShopRightsCreateRightsInfo struct {
	Vid             int64 `json:"vid,omitempty"`
	ProcessVid      int64 `json:"processVid,omitempty"`
	RightsType      int64 `json:"rightsType,omitempty"`
	RightsStatus    int64 `json:"rightsStatus,omitempty"`
	RightsCauseType int64 `json:"rightsCauseType,omitempty"`
	RightsSource    int64 `json:"rightsSource,omitempty"`
	RefundType      int64 `json:"refundType,omitempty"`
}
type WeimobShopRightsCreateOrderInfo struct {
	ParentOrderNo int64 `json:"parentOrderNo,omitempty"`
	OrderType     int64 `json:"orderType,omitempty"`
	OrderSource   int64 `json:"orderSource,omitempty"`
	ChannelType   int64 `json:"channelType,omitempty"`
	BizSourceType int64 `json:"bizSourceType,omitempty"`
}

// OnWeimobShopRightsCreateEvent
// @id 1303
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1303?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=售后创建消息)
func (s *Service) OnWeimobShopRightsCreateEvent(handler msg.WosEventHandler[WeimobShopRightsCreateEvent]) (service *Service) {
	var listener = &WeimobShopRightsCreateListener{handler}
	s.Register(msg.WrapListener[WeimobShopRightsCreateEvent](listener), "weimob_shop.rights", "create")
	return s
}
