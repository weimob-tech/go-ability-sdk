package kld

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type KLDCashAddCashOrderListener struct {
	handler msg.XinyunEventHandler[KLDCashAddCashOrderEvent]
}

func (s KLDCashAddCashOrderListener) NewMessage() msg.MsgRequest[KLDCashAddCashOrderEvent] {
	return &msg.XinyunOpenMessage[KLDCashAddCashOrderEvent]{
		MsgBody: &KLDCashAddCashOrderEvent{},
	}
}

func (s KLDCashAddCashOrderListener) OnMessage(ctx context.Context, message msg.MsgRequest[KLDCashAddCashOrderEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[KLDCashAddCashOrderEvent]))
}

type KLDCashAddCashOrderEvent struct {
	MegBody    KLDCashAddCashOrderMeg_body `json:"meg_body,omitempty"`
	ClientId   string                      `json:"client_id,omitempty"`
	RealAmount float64                     `json:"realAmount,omitempty"`
	Amount     float64                     `json:"amount,omitempty"`
	OrderNo    string                      `json:"orderNo,omitempty"`
	OpenId     string                      `json:"openId,omitempty"`
}

type KLDCashAddCashOrderMeg_body struct {
}

// OnKLDCashAddCashOrderEvent
// @id 245
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=新增支付成功收银单
func (s *Service) OnKLDCashAddCashOrderEvent(handler msg.XinyunEventHandler[KLDCashAddCashOrderEvent]) (service *Service) {
	var listener = &KLDCashAddCashOrderListener{handler}
	s.Register(msg.WrapListener[KLDCashAddCashOrderEvent](listener), "KLD_Cash", "addCashOrder")
	return s
}
