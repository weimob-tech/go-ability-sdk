package o2o

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type O2oServicesGoodsVerifyListener struct {
	handler msg.XinyunEventHandler[O2oServicesGoodsVerifyEvent]
}

func (s O2oServicesGoodsVerifyListener) NewMessage() msg.MsgRequest[O2oServicesGoodsVerifyEvent] {
	return &msg.XinyunOpenMessage[O2oServicesGoodsVerifyEvent]{
		MsgBody: &O2oServicesGoodsVerifyEvent{},
	}
}

func (s O2oServicesGoodsVerifyListener) OnMessage(ctx context.Context, message msg.MsgRequest[O2oServicesGoodsVerifyEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[O2oServicesGoodsVerifyEvent]))
}

type O2oServicesGoodsVerifyEvent struct {
	CardNo     string `json:"cardNo,omitempty"`
	VerifierId int64  `json:"verifierId,omitempty"`
	VerifyTime int64  `json:"verifyTime,omitempty"`
	SerialNo   string `json:"serialNo,omitempty"`
	StoreId    int64  `json:"storeId,omitempty"`
}

// OnO2oServicesGoodsVerifyEvent
// @id 1053
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=商品核销
func (s *Service) OnO2oServicesGoodsVerifyEvent(handler msg.XinyunEventHandler[O2oServicesGoodsVerifyEvent]) (service *Service) {
	var listener = &O2oServicesGoodsVerifyListener{handler}
	s.Register(msg.WrapListener[O2oServicesGoodsVerifyEvent](listener), "o2o_services_goods", "verify")
	return s
}
