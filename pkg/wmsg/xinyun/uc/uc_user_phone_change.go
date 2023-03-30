package uc

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type UcUserPhoneChangeListener struct {
	handler msg.XinyunEventHandler[UcUserPhoneChangeEvent]
}

func (s UcUserPhoneChangeListener) NewMessage() msg.MsgRequest[UcUserPhoneChangeEvent] {
	return &msg.XinyunOpenMessage[UcUserPhoneChangeEvent]{
		MsgBody: &UcUserPhoneChangeEvent{},
	}
}

func (s UcUserPhoneChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[UcUserPhoneChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[UcUserPhoneChangeEvent]))
}

type UcUserPhoneChangeEvent struct {
	Wid          int64    `json:"wid,omitempty"`
	OldPhoneList []string `json:"oldPhoneList,omitempty"`
	NewPhone     string   `json:"newPhone,omitempty"`
}

// OnUcUserPhoneChangeEvent
// @id 1443
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=修改手机号
func (s *Service) OnUcUserPhoneChangeEvent(handler msg.XinyunEventHandler[UcUserPhoneChangeEvent]) (service *Service) {
	var listener = &UcUserPhoneChangeListener{handler}
	s.Register(msg.WrapListener[UcUserPhoneChangeEvent](listener), "uc_user", "phone_change")
	return s
}
