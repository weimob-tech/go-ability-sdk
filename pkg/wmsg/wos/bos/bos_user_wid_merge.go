package bos

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type BosUserWidMergeListener struct {
	handler msg.WosEventHandler[BosUserWidMergeEvent]
}

func (s BosUserWidMergeListener) NewMessage() msg.MsgRequest[BosUserWidMergeEvent] {
	return &msg.WosOpenMessage[BosUserWidMergeEvent]{
		MsgBody: &BosUserWidMergeEvent{},
	}
}

func (s BosUserWidMergeListener) OnMessage(ctx context.Context, message msg.MsgRequest[BosUserWidMergeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[BosUserWidMergeEvent]))
}

type BosUserWidMergeEvent struct {
	Wid   int64 `json:"wid,omitempty"`
	BosId int64 `json:"bosId,omitempty"`
}

// OnBosUserWidMergeEvent
// @id 1352
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1352?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=账号合并完成消息)
func (s *Service) OnBosUserWidMergeEvent(handler msg.WosEventHandler[BosUserWidMergeEvent]) (service *Service) {
	var listener = &BosUserWidMergeListener{handler}
	s.Register(msg.WrapListener[BosUserWidMergeEvent](listener), "bos.user", "widMerge")
	return s
}
