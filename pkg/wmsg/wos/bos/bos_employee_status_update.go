package bos

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type BosEmployeeStatusUpdateListener struct {
	handler msg.WosEventHandler[BosEmployeeStatusUpdateEvent]
}

func (s BosEmployeeStatusUpdateListener) NewMessage() msg.MsgRequest[BosEmployeeStatusUpdateEvent] {
	return &msg.WosOpenMessage[BosEmployeeStatusUpdateEvent]{
		MsgBody: &BosEmployeeStatusUpdateEvent{},
	}
}

func (s BosEmployeeStatusUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[BosEmployeeStatusUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[BosEmployeeStatusUpdateEvent]))
}

type BosEmployeeStatusUpdateEvent struct {
	Wid    int64 `json:"wid,omitempty"`
	Status int64 `json:"status,omitempty"`
}

// OnBosEmployeeStatusUpdateEvent
// @id 1387
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1387?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=员工状态变更消息)
func (s *Service) OnBosEmployeeStatusUpdateEvent(handler msg.WosEventHandler[BosEmployeeStatusUpdateEvent]) (service *Service) {
	var listener = &BosEmployeeStatusUpdateListener{handler}
	s.Register(msg.WrapListener[BosEmployeeStatusUpdateEvent](listener), "bos.employee.status", "update")
	return s
}
