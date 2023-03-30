package bos

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type BosUserPhoneRebindEventListener struct {
	handler msg.WosEventHandler[BosUserPhoneRebindEventEvent]
}

func (s BosUserPhoneRebindEventListener) NewMessage() msg.MsgRequest[BosUserPhoneRebindEventEvent] {
	return &msg.WosOpenMessage[BosUserPhoneRebindEventEvent]{
		MsgBody: &BosUserPhoneRebindEventEvent{},
	}
}

func (s BosUserPhoneRebindEventListener) OnMessage(ctx context.Context, message msg.MsgRequest[BosUserPhoneRebindEventEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[BosUserPhoneRebindEventEvent]))
}

type BosUserPhoneRebindEventEvent struct {
	Wid      int64  `json:"wid,omitempty"`
	BosId    int64  `json:"bosId,omitempty"`
	OldPhone string `json:"oldPhone,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

// OnBosUserPhoneRebindEventEvent
// @id 1215
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1215?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=用户手机号换绑消息)
func (s *Service) OnBosUserPhoneRebindEventEvent(handler msg.WosEventHandler[BosUserPhoneRebindEventEvent]) (service *Service) {
	var listener = &BosUserPhoneRebindEventListener{handler}
	s.Register(msg.WrapListener[BosUserPhoneRebindEventEvent](listener), "bos.user", "phoneRebindEvent")
	return s
}
