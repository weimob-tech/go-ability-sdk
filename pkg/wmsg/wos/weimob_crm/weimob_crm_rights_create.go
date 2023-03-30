package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmRightsCreateListener struct {
	handler msg.WosEventHandler[WeimobCrmRightsCreateEvent]
}

func (s WeimobCrmRightsCreateListener) NewMessage() msg.MsgRequest[WeimobCrmRightsCreateEvent] {
	return &msg.WosOpenMessage[WeimobCrmRightsCreateEvent]{
		MsgBody: &WeimobCrmRightsCreateEvent{},
	}
}

func (s WeimobCrmRightsCreateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmRightsCreateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmRightsCreateEvent]))
}

type WeimobCrmRightsCreateEvent struct {
	RightsInfo  WeimobCrmRightsCreateRightsInfo `json:"rightsInfo,omitempty"`
	RightsId    int64                           `json:"rightsId,omitempty"`
	OrderNo     int64                           `json:"orderNo,omitempty"`
	OutRightsNo string                          `json:"outRightsNo,omitempty"`
	OutOrderNo  string                          `json:"outOrderNo,omitempty"`
	Wid         int64                           `json:"wid,omitempty"`
}

type WeimobCrmRightsCreateRightsInfo struct {
	Vid          int64 `json:"vid,omitempty"`
	RightsStatus int64 `json:"rightsStatus,omitempty"`
}

// OnWeimobCrmRightsCreateEvent
// @id 1419
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1419?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=售后单创建消息)
func (s *Service) OnWeimobCrmRightsCreateEvent(handler msg.WosEventHandler[WeimobCrmRightsCreateEvent]) (service *Service) {
	var listener = &WeimobCrmRightsCreateListener{handler}
	s.Register(msg.WrapListener[WeimobCrmRightsCreateEvent](listener), "weimob_crm.rights", "create")
	return s
}
