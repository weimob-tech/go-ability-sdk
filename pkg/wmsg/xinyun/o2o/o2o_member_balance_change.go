package o2o

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type O2oMemberBalanceChangeListener struct {
	handler msg.XinyunEventHandler[O2oMemberBalanceChangeEvent]
}

func (s O2oMemberBalanceChangeListener) NewMessage() msg.MsgRequest[O2oMemberBalanceChangeEvent] {
	return &msg.XinyunOpenMessage[O2oMemberBalanceChangeEvent]{
		MsgBody: &O2oMemberBalanceChangeEvent{},
	}
}

func (s O2oMemberBalanceChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[O2oMemberBalanceChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[O2oMemberBalanceChangeEvent]))
}

type O2oMemberBalanceChangeEvent struct {
	MerchantId    int64  `json:"merchantId,omitempty"`
	OpenId        string `json:"openId,omitempty"`
	Balance       int    `json:"balance,omitempty"`
	BalanceChange int    `json:"balanceChange,omitempty"`
	Comment       string `json:"comment,omitempty"`
	CreateTime    int64  `json:"createTime,omitempty"`
	Wid           string `json:"wid,omitempty"`
	CustomerId    int64  `json:"customerId,omitempty"`
	Code          string `json:"code,omitempty"`
}

// OnO2oMemberBalanceChangeEvent
// @id 472
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=余额变更
func (s *Service) OnO2oMemberBalanceChangeEvent(handler msg.XinyunEventHandler[O2oMemberBalanceChangeEvent]) (service *Service) {
	var listener = &O2oMemberBalanceChangeListener{handler}
	s.Register(msg.WrapListener[O2oMemberBalanceChangeEvent](listener), "o2o_member", "balanceChange")
	return s
}
