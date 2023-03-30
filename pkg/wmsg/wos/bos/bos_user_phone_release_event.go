package bos

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type BosUserPhoneReleaseEventListener struct {
	handler msg.WosEventHandler[BosUserPhoneReleaseEventEvent]
}

func (s BosUserPhoneReleaseEventListener) NewMessage() msg.MsgRequest[BosUserPhoneReleaseEventEvent] {
	return &msg.WosOpenMessage[BosUserPhoneReleaseEventEvent]{
		MsgBody: &BosUserPhoneReleaseEventEvent{},
	}
}

func (s BosUserPhoneReleaseEventListener) OnMessage(ctx context.Context, message msg.MsgRequest[BosUserPhoneReleaseEventEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[BosUserPhoneReleaseEventEvent]))
}

type BosUserPhoneReleaseEventEvent struct {
	Wid   int64  `json:"wid,omitempty"`
	Phone string `json:"phone,omitempty"`
	BosId int64  `json:"bosId,omitempty"`
	Time  int64  `json:"time,omitempty"`
}

// OnBosUserPhoneReleaseEventEvent
// @id 1365
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1365?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=手机号解绑消息)
func (s *Service) OnBosUserPhoneReleaseEventEvent(handler msg.WosEventHandler[BosUserPhoneReleaseEventEvent]) (service *Service) {
	var listener = &BosUserPhoneReleaseEventListener{handler}
	s.Register(msg.WrapListener[BosUserPhoneReleaseEventEvent](listener), "bos.user", "phoneReleaseEvent")
	return s
}
