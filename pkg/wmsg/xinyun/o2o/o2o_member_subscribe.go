package o2o

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type O2oMemberSubscribeListener struct {
	handler msg.XinyunEventHandler[O2oMemberSubscribeEvent]
}

func (s O2oMemberSubscribeListener) NewMessage() msg.MsgRequest[O2oMemberSubscribeEvent] {
	return &msg.XinyunOpenMessage[O2oMemberSubscribeEvent]{
		MsgBody: &O2oMemberSubscribeEvent{},
	}
}

func (s O2oMemberSubscribeListener) OnMessage(ctx context.Context, message msg.MsgRequest[O2oMemberSubscribeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[O2oMemberSubscribeEvent]))
}

type O2oMemberSubscribeEvent struct {
	MerchantId int64  `json:"merchantId,omitempty"`
	OpenId     string `json:"openId,omitempty"`
	CreateTime int64  `json:"createTime,omitempty"`
}

// OnO2oMemberSubscribeEvent
// @id 468
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=用户关注
func (s *Service) OnO2oMemberSubscribeEvent(handler msg.XinyunEventHandler[O2oMemberSubscribeEvent]) (service *Service) {
	var listener = &O2oMemberSubscribeListener{handler}
	s.Register(msg.WrapListener[O2oMemberSubscribeEvent](listener), "o2o_member", "subscribe")
	return s
}
