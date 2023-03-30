package weimob_aewc

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobAewcCustomerGroupAddMemberListener struct {
	handler msg.WosEventHandler[WeimobAewcCustomerGroupAddMemberEvent]
}

func (s WeimobAewcCustomerGroupAddMemberListener) NewMessage() msg.MsgRequest[WeimobAewcCustomerGroupAddMemberEvent] {
	return &msg.WosOpenMessage[WeimobAewcCustomerGroupAddMemberEvent]{
		MsgBody: &WeimobAewcCustomerGroupAddMemberEvent{},
	}
}

func (s WeimobAewcCustomerGroupAddMemberListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobAewcCustomerGroupAddMemberEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobAewcCustomerGroupAddMemberEvent]))
}

type WeimobAewcCustomerGroupAddMemberEvent struct {
	Wid              int64  `json:"wid,omitempty"`
	GroupId          string `json:"groupId,omitempty"`
	MemberId         string `json:"memberId,omitempty"`
	Type             int64  `json:"type,omitempty"`
	JoinTime         string `json:"joinTime,omitempty"`
	JoinScene        int64  `json:"joinScene,omitempty"`
	Nickname         string `json:"nickname,omitempty"`
	GroupNickname    string `json:"groupNickname,omitempty"`
	InvitorOrgUserid string `json:"invitorOrgUserid,omitempty"`
}

// OnWeimobAewcCustomerGroupAddMemberEvent
// @id 1422
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1422?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=客户入群消息二次分发)
func (s *Service) OnWeimobAewcCustomerGroupAddMemberEvent(handler msg.WosEventHandler[WeimobAewcCustomerGroupAddMemberEvent]) (service *Service) {
	var listener = &WeimobAewcCustomerGroupAddMemberListener{handler}
	s.Register(msg.WrapListener[WeimobAewcCustomerGroupAddMemberEvent](listener), "weimob_aewc.customer.group", "addMember")
	return s
}
