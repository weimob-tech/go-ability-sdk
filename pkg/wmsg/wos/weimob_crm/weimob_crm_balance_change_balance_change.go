package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmBalanceChangeBalanceChangeListener struct {
	handler msg.WosEventHandler[WeimobCrmBalanceChangeBalanceChangeEvent]
}

func (s WeimobCrmBalanceChangeBalanceChangeListener) NewMessage() msg.MsgRequest[WeimobCrmBalanceChangeBalanceChangeEvent] {
	return &msg.WosOpenMessage[WeimobCrmBalanceChangeBalanceChangeEvent]{
		MsgBody: &WeimobCrmBalanceChangeBalanceChangeEvent{},
	}
}

func (s WeimobCrmBalanceChangeBalanceChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmBalanceChangeBalanceChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmBalanceChangeBalanceChangeEvent]))
}

type WeimobCrmBalanceChangeBalanceChangeEvent struct {
	ProductId         int64  `json:"productId,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	SourceProductId   int64  `json:"sourceProductId,omitempty"`
	Tcode             string `json:"tcode,omitempty"`
	BosId             int64  `json:"bosId,omitempty"`
	Vid               int64  `json:"vid,omitempty"`
	Wid               int64  `json:"wid,omitempty"`
	BalancePlanId     int64  `json:"balancePlanId,omitempty"`
	TransType         int64  `json:"transType,omitempty"`
	BizType           string `json:"bizType,omitempty"`
	BizUniqueId       int64  `json:"bizUniqueId,omitempty"`
	ChangAmount       int64  `json:"changAmount,omitempty"`
	TotalAmount       int64  `json:"totalAmount,omitempty"`
	ChangeTime        string `json:"changeTime,omitempty"`
}

// OnWeimobCrmBalanceChangeBalanceChangeEvent
// @id 1246
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1246?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=余额变动消息)
func (s *Service) OnWeimobCrmBalanceChangeBalanceChangeEvent(handler msg.WosEventHandler[WeimobCrmBalanceChangeBalanceChangeEvent]) (service *Service) {
	var listener = &WeimobCrmBalanceChangeBalanceChangeListener{handler}
	s.Register(msg.WrapListener[WeimobCrmBalanceChangeBalanceChangeEvent](listener), "weimob_crm.balance.change", "balanceChange")
	return s
}
