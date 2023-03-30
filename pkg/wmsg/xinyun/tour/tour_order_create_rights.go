package tour

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type TourOrderCreateRightsListener struct {
	handler msg.XinyunEventHandler[TourOrderCreateRightsEvent]
}

func (s TourOrderCreateRightsListener) NewMessage() msg.MsgRequest[TourOrderCreateRightsEvent] {
	return &msg.XinyunOpenMessage[TourOrderCreateRightsEvent]{
		MsgBody: &TourOrderCreateRightsEvent{},
	}
}

func (s TourOrderCreateRightsListener) OnMessage(ctx context.Context, message msg.MsgRequest[TourOrderCreateRightsEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[TourOrderCreateRightsEvent]))
}

type TourOrderCreateRightsEvent struct {
	OrderNo string `json:"orderNo,omitempty"`
	RightNo string `json:"rightNo,omitempty"`
}

// OnTourOrderCreateRightsEvent
// @id 1038
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=新增维权
func (s *Service) OnTourOrderCreateRightsEvent(handler msg.XinyunEventHandler[TourOrderCreateRightsEvent]) (service *Service) {
	var listener = &TourOrderCreateRightsListener{handler}
	s.Register(msg.WrapListener[TourOrderCreateRightsEvent](listener), "tour_order", "createRights")
	return s
}
