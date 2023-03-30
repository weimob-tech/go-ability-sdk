package bos

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type BosUserPhoneBindEventListener struct {
	handler msg.WosEventHandler[BosUserPhoneBindEventEvent]
}

func (s BosUserPhoneBindEventListener) NewMessage() msg.MsgRequest[BosUserPhoneBindEventEvent] {
	return &msg.WosOpenMessage[BosUserPhoneBindEventEvent]{
		MsgBody: &BosUserPhoneBindEventEvent{},
	}
}

func (s BosUserPhoneBindEventListener) OnMessage(ctx context.Context, message msg.MsgRequest[BosUserPhoneBindEventEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[BosUserPhoneBindEventEvent]))
}

type BosUserPhoneBindEventEvent struct {
	Wid   int64  `json:"wid,omitempty"`
	Phone string `json:"phone,omitempty"`
	BosId int64  `json:"bosId,omitempty"`
}

// OnBosUserPhoneBindEventEvent
// @id 1214
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1214?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=用户手机号绑定消息)
func (s *Service) OnBosUserPhoneBindEventEvent(handler msg.WosEventHandler[BosUserPhoneBindEventEvent]) (service *Service) {
	var listener = &BosUserPhoneBindEventListener{handler}
	s.Register(msg.WrapListener[BosUserPhoneBindEventEvent](listener), "bos.user", "phoneBindEvent")
	return s
}
