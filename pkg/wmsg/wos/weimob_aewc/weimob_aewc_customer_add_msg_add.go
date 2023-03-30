package weimob_aewc

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobAewcCustomerAddMsgAddListener struct {
	handler msg.WosEventHandler[WeimobAewcCustomerAddMsgAddEvent]
}

func (s WeimobAewcCustomerAddMsgAddListener) NewMessage() msg.MsgRequest[WeimobAewcCustomerAddMsgAddEvent] {
	return &msg.WosOpenMessage[WeimobAewcCustomerAddMsgAddEvent]{
		MsgBody: &WeimobAewcCustomerAddMsgAddEvent{},
	}
}

func (s WeimobAewcCustomerAddMsgAddListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobAewcCustomerAddMsgAddEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobAewcCustomerAddMsgAddEvent]))
}

type WeimobAewcCustomerAddMsgAddEvent struct {
	BosId               int64   `json:"bosId,omitempty"`
	Wid                 int64   `json:"wid,omitempty"`
	AddTime             int64   `json:"addTime,omitempty"`
	EmployeeId          string  `json:"employeeId,omitempty"`
	ExternalUserId      string  `json:"externalUserId,omitempty"`
	Unionid             string  `json:"unionid,omitempty"`
	TxUserId            string  `json:"txUserId,omitempty"`
	State               string  `json:"state,omitempty"`
	CustomerId          string  `json:"customerId,omitempty"`
	Name                string  `json:"name,omitempty"`
	CustomWelcomeCode   string  `json:"customWelcomeCode,omitempty"`
	WelcomeCode         string  `json:"welcomeCode,omitempty"`
	Corpid              string  `json:"corpid,omitempty"`
	AddWay              int64   `json:"addWay,omitempty"`
	HistoryCustomerSync int64   `json:"historyCustomerSync,omitempty"`
	RemarkMobiles       []int64 `json:"remarkMobiles,omitempty"`
	ExternalUserType    int64   `json:"externalUserType,omitempty"`
	UserStateChangeFlag bool    `json:"userStateChangeFlag,omitempty"`
}

// OnWeimobAewcCustomerAddMsgAddEvent
// @id 1404
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1404?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=企微助手二次分发添加好友消息)
func (s *Service) OnWeimobAewcCustomerAddMsgAddEvent(handler msg.WosEventHandler[WeimobAewcCustomerAddMsgAddEvent]) (service *Service) {
	var listener = &WeimobAewcCustomerAddMsgAddListener{handler}
	s.Register(msg.WrapListener[WeimobAewcCustomerAddMsgAddEvent](listener), "weimob_aewc.customer.add.msg", "add")
	return s
}
