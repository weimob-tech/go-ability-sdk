package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcGuideAddGuiderListener struct {
	handler msg.XinyunEventHandler[EcGuideAddGuiderEvent]
}

func (s EcGuideAddGuiderListener) NewMessage() msg.MsgRequest[EcGuideAddGuiderEvent] {
	return &msg.XinyunOpenMessage[EcGuideAddGuiderEvent]{
		MsgBody: &EcGuideAddGuiderEvent{},
	}
}

func (s EcGuideAddGuiderListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcGuideAddGuiderEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcGuideAddGuiderEvent]))
}

type EcGuideAddGuiderEvent struct {
	StoreId        int64  `json:"storeId,omitempty"`
	GuiderWid      int64  `json:"guiderWid,omitempty"`
	JobNumber      string `json:"jobNumber,omitempty"`
	GuiderPhone    string `json:"guiderPhone,omitempty"`
	IsUsed         int    `json:"isUsed,omitempty"`
	IsExclusiveCus int    `json:"isExclusiveCus,omitempty"`
}

// OnEcGuideAddGuiderEvent
// @id 1270
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=新增导购消息
func (s *Service) OnEcGuideAddGuiderEvent(handler msg.XinyunEventHandler[EcGuideAddGuiderEvent]) (service *Service) {
	var listener = &EcGuideAddGuiderListener{handler}
	s.Register(msg.WrapListener[EcGuideAddGuiderEvent](listener), "ec_guide", "addGuider")
	return s
}
