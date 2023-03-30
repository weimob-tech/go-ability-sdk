package tour

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type TourMemberActiveLocalMemberListener struct {
	handler msg.XinyunEventHandler[TourMemberActiveLocalMemberEvent]
}

func (s TourMemberActiveLocalMemberListener) NewMessage() msg.MsgRequest[TourMemberActiveLocalMemberEvent] {
	return &msg.XinyunOpenMessage[TourMemberActiveLocalMemberEvent]{
		MsgBody: &TourMemberActiveLocalMemberEvent{},
	}
}

func (s TourMemberActiveLocalMemberListener) OnMessage(ctx context.Context, message msg.MsgRequest[TourMemberActiveLocalMemberEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[TourMemberActiveLocalMemberEvent]))
}

type TourMemberActiveLocalMemberEvent struct {
	Type  int    `json:"type,omitempty"`
	Code  string `json:"code,omitempty"`
	Phone string `json:"phone,omitempty"`
	Wid   string `json:"wid,omitempty"`
}

// OnTourMemberActiveLocalMemberEvent
// @id 1039
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=会员开卡通知
func (s *Service) OnTourMemberActiveLocalMemberEvent(handler msg.XinyunEventHandler[TourMemberActiveLocalMemberEvent]) (service *Service) {
	var listener = &TourMemberActiveLocalMemberListener{handler}
	s.Register(msg.WrapListener[TourMemberActiveLocalMemberEvent](listener), "tour_member", "activeLocalMember")
	return s
}
