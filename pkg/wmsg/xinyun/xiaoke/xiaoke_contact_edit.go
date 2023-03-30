package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeContactEditListener struct {
	handler msg.XinyunEventHandler[XiaokeContactEditEvent]
}

func (s XiaokeContactEditListener) NewMessage() msg.MsgRequest[XiaokeContactEditEvent] {
	return &msg.XinyunOpenMessage[XiaokeContactEditEvent]{
		MsgBody: &XiaokeContactEditEvent{},
	}
}

func (s XiaokeContactEditListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeContactEditEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeContactEditEvent]))
}

type XiaokeContactEditEvent struct {
	FieldKeyValue XiaokeContactEditFieldKeyValue `json:"fieldKeyValue,omitempty"`
	Id            string                         `json:"id,omitempty"`
}

type XiaokeContactEditFieldKeyValue struct {
	SContactName string `json:"s_contact_name,omitempty"`
	STelephone   string `json:"s_telephone,omitempty"`
	SMobilePhone string `json:"s_mobile_phone,omitempty"`
	SCustomerKey string `json:"s_customer_key,omitempty"`
	SGender      string `json:"s_gender,omitempty"`
	SPost        string `json:"s_post,omitempty"`
	SBirthday    int    `json:"s_birthday,omitempty"`
	SQq          string `json:"s_qq,omitempty"`
}

// OnXiaokeContactEditEvent
// @id 1883
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=编辑联系人
func (s *Service) OnXiaokeContactEditEvent(handler msg.XinyunEventHandler[XiaokeContactEditEvent]) (service *Service) {
	var listener = &XiaokeContactEditListener{handler}
	s.Register(msg.WrapListener[XiaokeContactEditEvent](listener), "xiaoke_contact", "edit")
	return s
}
