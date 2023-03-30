package bos

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type BosUserBeginMergeListener struct {
	handler msg.WosEventHandler[BosUserBeginMergeEvent]
}

func (s BosUserBeginMergeListener) NewMessage() msg.MsgRequest[BosUserBeginMergeEvent] {
	return &msg.WosOpenMessage[BosUserBeginMergeEvent]{
		MsgBody: &BosUserBeginMergeEvent{},
	}
}

func (s BosUserBeginMergeListener) OnMessage(ctx context.Context, message msg.MsgRequest[BosUserBeginMergeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[BosUserBeginMergeEvent]))
}

type BosUserBeginMergeEvent struct {
	WidList []int64 `json:"widList,omitempty"`
	Time    string  `json:"time,omitempty"`
}

// OnBosUserBeginMergeEvent
// @id 1389
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1389?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=账号开始合并消息)
func (s *Service) OnBosUserBeginMergeEvent(handler msg.WosEventHandler[BosUserBeginMergeEvent]) (service *Service) {
	var listener = &BosUserBeginMergeListener{handler}
	s.Register(msg.WrapListener[BosUserBeginMergeEvent](listener), "bos.user", "beginMerge")
	return s
}
