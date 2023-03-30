package uc

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type UcUserBeginMergeListener struct {
	handler msg.XinyunEventHandler[UcUserBeginMergeEvent]
}

func (s UcUserBeginMergeListener) NewMessage() msg.MsgRequest[UcUserBeginMergeEvent] {
	return &msg.XinyunOpenMessage[UcUserBeginMergeEvent]{
		MsgBody: &UcUserBeginMergeEvent{},
	}
}

func (s UcUserBeginMergeListener) OnMessage(ctx context.Context, message msg.MsgRequest[UcUserBeginMergeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[UcUserBeginMergeEvent]))
}

type UcUserBeginMergeEvent struct {
	WidList []int64 `json:"widList,omitempty"`
	Time    int64   `json:"time,omitempty"`
}

// OnUcUserBeginMergeEvent
// @id 3695
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=账号开始合并消息
func (s *Service) OnUcUserBeginMergeEvent(handler msg.XinyunEventHandler[UcUserBeginMergeEvent]) (service *Service) {
	var listener = &UcUserBeginMergeListener{handler}
	s.Register(msg.WrapListener[UcUserBeginMergeEvent](listener), "uc_user", "begin_merge")
	return s
}
