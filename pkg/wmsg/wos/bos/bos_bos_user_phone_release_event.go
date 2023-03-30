package bos

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type BosBosUserPhoneReleaseEventListener struct {
	handler msg.WosEventHandler[BosBosUserPhoneReleaseEventEvent]
}

func (s BosBosUserPhoneReleaseEventListener) NewMessage() msg.MsgRequest[BosBosUserPhoneReleaseEventEvent] {
	return &msg.WosOpenMessage[BosBosUserPhoneReleaseEventEvent]{
		MsgBody: &BosBosUserPhoneReleaseEventEvent{},
	}
}

func (s BosBosUserPhoneReleaseEventListener) OnMessage(ctx context.Context, message msg.MsgRequest[BosBosUserPhoneReleaseEventEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[BosBosUserPhoneReleaseEventEvent]))
}

type BosBosUserPhoneReleaseEventEvent struct {
	Wid   int64 `json:"wid,omitempty"`
	Phone int64 `json:"phone,omitempty"`
	BosId int64 `json:"bosId,omitempty"`
	Time  int64 `json:"time,omitempty"`
}

// OnBosBosUserPhoneReleaseEventEvent
// @id 1364
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1364?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=用户手机号解绑消息)
func (s *Service) OnBosBosUserPhoneReleaseEventEvent(handler msg.WosEventHandler[BosBosUserPhoneReleaseEventEvent]) (service *Service) {
	var listener = &BosBosUserPhoneReleaseEventListener{handler}
	s.Register(msg.WrapListener[BosBosUserPhoneReleaseEventEvent](listener), "bos.bos.user", "phoneReleaseEvent")
	return s
}
