package kld

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type KLDOrderAddProductOrderListener struct {
	handler msg.XinyunEventHandler[KLDOrderAddProductOrderEvent]
}

func (s KLDOrderAddProductOrderListener) NewMessage() msg.MsgRequest[KLDOrderAddProductOrderEvent] {
	return &msg.XinyunOpenMessage[KLDOrderAddProductOrderEvent]{
		MsgBody: &KLDOrderAddProductOrderEvent{},
	}
}

func (s KLDOrderAddProductOrderListener) OnMessage(ctx context.Context, message msg.MsgRequest[KLDOrderAddProductOrderEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[KLDOrderAddProductOrderEvent]))
}

type KLDOrderAddProductOrderEvent struct {
	MegBody    KLDOrderAddProductOrderMeg_body `json:"meg_body,omitempty"`
	RealAmount float64                         `json:"realAmount,omitempty"`
	Amount     float64                         `json:"amount,omitempty"`
	OrderNo    string                          `json:"orderNo,omitempty"`
	OpenId     string                          `json:"openId,omitempty"`
	ClientId   string                          `json:"client_id,omitempty"`
	Storeid    string                          `json:"storeid,omitempty"`
}

type KLDOrderAddProductOrderMeg_body struct {
}

// OnKLDOrderAddProductOrderEvent
// @id 260
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=订单下单成功
func (s *Service) OnKLDOrderAddProductOrderEvent(handler msg.XinyunEventHandler[KLDOrderAddProductOrderEvent]) (service *Service) {
	var listener = &KLDOrderAddProductOrderListener{handler}
	s.Register(msg.WrapListener[KLDOrderAddProductOrderEvent](listener), "KLD_Order", "addProductOrder")
	return s
}
