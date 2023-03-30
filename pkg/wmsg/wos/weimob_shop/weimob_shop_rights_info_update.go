package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopRightsInfoUpdateListener struct {
	handler msg.WosEventHandler[WeimobShopRightsInfoUpdateEvent]
}

func (s WeimobShopRightsInfoUpdateListener) NewMessage() msg.MsgRequest[WeimobShopRightsInfoUpdateEvent] {
	return &msg.WosOpenMessage[WeimobShopRightsInfoUpdateEvent]{
		MsgBody: &WeimobShopRightsInfoUpdateEvent{},
	}
}

func (s WeimobShopRightsInfoUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopRightsInfoUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopRightsInfoUpdateEvent]))
}

type WeimobShopRightsInfoUpdateEvent struct {
	FieldContents WeimobShopRightsInfoUpdateFieldContents `json:"fieldContents,omitempty"`
	RightsInfo    WeimobShopRightsInfoUpdateRightsInfo    `json:"rightsInfo,omitempty"`
	OrderInfo     WeimobShopRightsInfoUpdateOrderInfo     `json:"orderInfo,omitempty"`
	OperateType   int64                                   `json:"operateType,omitempty"`
	RightsId      int64                                   `json:"rightsId,omitempty"`
	OrderNo       int64                                   `json:"orderNo,omitempty"`
	Fields        []int64                                 `json:"fields,omitempty"`
}

type WeimobShopRightsInfoUpdateFieldContents struct {
	FlagRank    string `json:"flagRank,omitempty"`
	FlagContent string `json:"flagContent,omitempty"`
}
type WeimobShopRightsInfoUpdateRightsInfo struct {
	Vid             int64 `json:"vid,omitempty"`
	ProcessVid      int64 `json:"processVid,omitempty"`
	RightsType      int64 `json:"rightsType,omitempty"`
	RightsStatus    int64 `json:"rightsStatus,omitempty"`
	RightsCauseType int64 `json:"rightsCauseType,omitempty"`
	RightsSource    int64 `json:"rightsSource,omitempty"`
	RefundType      int64 `json:"refundType,omitempty"`
}
type WeimobShopRightsInfoUpdateOrderInfo struct {
	ParentOrderNo int64 `json:"parentOrderNo,omitempty"`
	OrderType     int64 `json:"orderType,omitempty"`
	OrderSource   int64 `json:"orderSource,omitempty"`
	ChannelType   int64 `json:"channelType,omitempty"`
	BizSourceType int64 `json:"bizSourceType,omitempty"`
}

// OnWeimobShopRightsInfoUpdateEvent
// @id 1305
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1305?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=售后内容变更消息)
func (s *Service) OnWeimobShopRightsInfoUpdateEvent(handler msg.WosEventHandler[WeimobShopRightsInfoUpdateEvent]) (service *Service) {
	var listener = &WeimobShopRightsInfoUpdateListener{handler}
	s.Register(msg.WrapListener[WeimobShopRightsInfoUpdateEvent](listener), "weimob_shop.rights", "infoUpdate")
	return s
}
