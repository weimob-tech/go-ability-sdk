package o2o

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type O2oMemberPointChangeListener struct {
	handler msg.XinyunEventHandler[O2oMemberPointChangeEvent]
}

func (s O2oMemberPointChangeListener) NewMessage() msg.MsgRequest[O2oMemberPointChangeEvent] {
	return &msg.XinyunOpenMessage[O2oMemberPointChangeEvent]{
		MsgBody: &O2oMemberPointChangeEvent{},
	}
}

func (s O2oMemberPointChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[O2oMemberPointChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[O2oMemberPointChangeEvent]))
}

type O2oMemberPointChangeEvent struct {
	Point       int    `json:"point,omitempty"`
	PointChange int    `json:"pointChange,omitempty"`
	Comment     string `json:"comment,omitempty"`
	CreateTime  int64  `json:"createTime,omitempty"`
	MerchantId  int64  `json:"merchantId,omitempty"`
	OpenId      string `json:"openId,omitempty"`
	Wid         string `json:"wid,omitempty"`
	CustomerId  int64  `json:"customerId,omitempty"`
	Code        string `json:"code,omitempty"`
}

// OnO2oMemberPointChangeEvent
// @id 471
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=积分变更
func (s *Service) OnO2oMemberPointChangeEvent(handler msg.XinyunEventHandler[O2oMemberPointChangeEvent]) (service *Service) {
	var listener = &O2oMemberPointChangeListener{handler}
	s.Register(msg.WrapListener[O2oMemberPointChangeEvent](listener), "o2o_member", "pointChange")
	return s
}
