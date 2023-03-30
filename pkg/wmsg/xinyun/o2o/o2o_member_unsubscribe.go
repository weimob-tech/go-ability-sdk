package o2o

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type O2oMemberUnsubscribeListener struct {
	handler msg.XinyunEventHandler[O2oMemberUnsubscribeEvent]
}

func (s O2oMemberUnsubscribeListener) NewMessage() msg.MsgRequest[O2oMemberUnsubscribeEvent] {
	return &msg.XinyunOpenMessage[O2oMemberUnsubscribeEvent]{
		MsgBody: &O2oMemberUnsubscribeEvent{},
	}
}

func (s O2oMemberUnsubscribeListener) OnMessage(ctx context.Context, message msg.MsgRequest[O2oMemberUnsubscribeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[O2oMemberUnsubscribeEvent]))
}

type O2oMemberUnsubscribeEvent struct {
	MerchantId int64  `json:"merchantId,omitempty"`
	OpenId     string `json:"openId,omitempty"`
	CreateTime int64  `json:"createTime,omitempty"`
}

// OnO2oMemberUnsubscribeEvent
// @id 469
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=用户取消关注
func (s *Service) OnO2oMemberUnsubscribeEvent(handler msg.XinyunEventHandler[O2oMemberUnsubscribeEvent]) (service *Service) {
	var listener = &O2oMemberUnsubscribeListener{handler}
	s.Register(msg.WrapListener[O2oMemberUnsubscribeEvent](listener), "o2o_member", "unsubscribe")
	return s
}
