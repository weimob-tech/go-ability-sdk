package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcGuideUpdateGuiderListener struct {
	handler msg.XinyunEventHandler[EcGuideUpdateGuiderEvent]
}

func (s EcGuideUpdateGuiderListener) NewMessage() msg.MsgRequest[EcGuideUpdateGuiderEvent] {
	return &msg.XinyunOpenMessage[EcGuideUpdateGuiderEvent]{
		MsgBody: &EcGuideUpdateGuiderEvent{},
	}
}

func (s EcGuideUpdateGuiderListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcGuideUpdateGuiderEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcGuideUpdateGuiderEvent]))
}

type EcGuideUpdateGuiderEvent struct {
	StoreId        int64  `json:"storeId,omitempty"`
	GuiderWid      int64  `json:"guiderWid,omitempty"`
	JobNumber      string `json:"jobNumber,omitempty"`
	GuiderPhone    string `json:"guiderPhone,omitempty"`
	IsUsed         int    `json:"isUsed,omitempty"`
	IsExclusiveCus int    `json:"isExclusiveCus,omitempty"`
}

// OnEcGuideUpdateGuiderEvent
// @id 1271
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=更新导购消息
func (s *Service) OnEcGuideUpdateGuiderEvent(handler msg.XinyunEventHandler[EcGuideUpdateGuiderEvent]) (service *Service) {
	var listener = &EcGuideUpdateGuiderListener{handler}
	s.Register(msg.WrapListener[EcGuideUpdateGuiderEvent](listener), "ec_guide", "updateGuider")
	return s
}
