package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcGuideDeleteGuiderListener struct {
	handler msg.XinyunEventHandler[EcGuideDeleteGuiderEvent]
}

func (s EcGuideDeleteGuiderListener) NewMessage() msg.MsgRequest[EcGuideDeleteGuiderEvent] {
	return &msg.XinyunOpenMessage[EcGuideDeleteGuiderEvent]{
		MsgBody: &EcGuideDeleteGuiderEvent{},
	}
}

func (s EcGuideDeleteGuiderListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcGuideDeleteGuiderEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcGuideDeleteGuiderEvent]))
}

type EcGuideDeleteGuiderEvent struct {
	StoreId     int64  `json:"storeId,omitempty"`
	GuiderWid   int64  `json:"guiderWid,omitempty"`
	JobNumber   string `json:"jobNumber,omitempty"`
	GuiderPhone string `json:"guiderPhone,omitempty"`
}

// OnEcGuideDeleteGuiderEvent
// @id 1272
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=删除导购消息
func (s *Service) OnEcGuideDeleteGuiderEvent(handler msg.XinyunEventHandler[EcGuideDeleteGuiderEvent]) (service *Service) {
	var listener = &EcGuideDeleteGuiderListener{handler}
	s.Register(msg.WrapListener[EcGuideDeleteGuiderEvent](listener), "ec_guide", "deleteGuider")
	return s
}
