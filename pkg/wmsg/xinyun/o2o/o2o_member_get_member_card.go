package o2o

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type O2oMemberGetMemberCardListener struct {
	handler msg.XinyunEventHandler[O2oMemberGetMemberCardEvent]
}

func (s O2oMemberGetMemberCardListener) NewMessage() msg.MsgRequest[O2oMemberGetMemberCardEvent] {
	return &msg.XinyunOpenMessage[O2oMemberGetMemberCardEvent]{
		MsgBody: &O2oMemberGetMemberCardEvent{},
	}
}

func (s O2oMemberGetMemberCardListener) OnMessage(ctx context.Context, message msg.MsgRequest[O2oMemberGetMemberCardEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[O2oMemberGetMemberCardEvent]))
}

type O2oMemberGetMemberCardEvent struct {
	MerchantId int64  `json:"merchantId,omitempty"`
	OpenId     string `json:"openId,omitempty"`
	Code       string `json:"code,omitempty"`
	CreateTime int64  `json:"createTime,omitempty"`
}

// OnO2oMemberGetMemberCardEvent
// @id 470
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=领取会员卡
func (s *Service) OnO2oMemberGetMemberCardEvent(handler msg.XinyunEventHandler[O2oMemberGetMemberCardEvent]) (service *Service) {
	var listener = &O2oMemberGetMemberCardListener{handler}
	s.Register(msg.WrapListener[O2oMemberGetMemberCardEvent](listener), "o2o_member", "getMemberCard")
	return s
}
