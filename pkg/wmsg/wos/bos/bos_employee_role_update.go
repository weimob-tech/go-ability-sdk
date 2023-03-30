package bos

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type BosEmployeeRoleUpdateListener struct {
	handler msg.WosEventHandler[BosEmployeeRoleUpdateEvent]
}

func (s BosEmployeeRoleUpdateListener) NewMessage() msg.MsgRequest[BosEmployeeRoleUpdateEvent] {
	return &msg.WosOpenMessage[BosEmployeeRoleUpdateEvent]{
		MsgBody: &BosEmployeeRoleUpdateEvent{},
	}
}

func (s BosEmployeeRoleUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[BosEmployeeRoleUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[BosEmployeeRoleUpdateEvent]))
}

type BosEmployeeRoleUpdateEvent struct {
	OperateRelations []BosEmployeeRoleUpdateOperateRelations `json:"operateRelations,omitempty"`
	Wid              int64                                   `json:"wid,omitempty"`
}

type BosEmployeeRoleUpdateOperateRelations struct {
	VidRelationDetailList []BosEmployeeRoleUpdateVidRelationDetailList `json:"vidRelationDetailList,omitempty"`
	RoleId                int64                                        `json:"roleId,omitempty"`
	RoleType              int64                                        `json:"roleType,omitempty"`
	ExtendAttr            string                                       `json:"extendAttr,omitempty"`
}
type BosEmployeeRoleUpdateVidRelationDetailList struct {
	VidList     []int64 `json:"vidList,omitempty"`
	FinalStatus int64   `json:"finalStatus,omitempty"`
}

// OnBosEmployeeRoleUpdateEvent
// @id 1388
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1388?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=员工组织角色关系变更消息)
func (s *Service) OnBosEmployeeRoleUpdateEvent(handler msg.WosEventHandler[BosEmployeeRoleUpdateEvent]) (service *Service) {
	var listener = &BosEmployeeRoleUpdateListener{handler}
	s.Register(msg.WrapListener[BosEmployeeRoleUpdateEvent](listener), "bos.employee.role", "update")
	return s
}
