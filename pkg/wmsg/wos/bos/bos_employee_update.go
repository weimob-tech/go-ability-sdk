package bos

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type BosEmployeeUpdateListener struct {
	handler msg.WosEventHandler[BosEmployeeUpdateEvent]
}

func (s BosEmployeeUpdateListener) NewMessage() msg.MsgRequest[BosEmployeeUpdateEvent] {
	return &msg.WosOpenMessage[BosEmployeeUpdateEvent]{
		MsgBody: &BosEmployeeUpdateEvent{},
	}
}

func (s BosEmployeeUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[BosEmployeeUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[BosEmployeeUpdateEvent]))
}

type BosEmployeeUpdateEvent struct {
	Wid         int64 `json:"wid,omitempty"`
	OperateType int64 `json:"operateType,omitempty"`
}

// OnBosEmployeeUpdateEvent
// @id 1386
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1386?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=员工基本信息变更消息)
func (s *Service) OnBosEmployeeUpdateEvent(handler msg.WosEventHandler[BosEmployeeUpdateEvent]) (service *Service) {
	var listener = &BosEmployeeUpdateListener{handler}
	s.Register(msg.WrapListener[BosEmployeeUpdateEvent](listener), "bos.employee", "update")
	return s
}
