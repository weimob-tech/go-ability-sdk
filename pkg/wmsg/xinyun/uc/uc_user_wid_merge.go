package uc

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type UcUserWidMergeListener struct {
	handler msg.XinyunEventHandler[UcUserWidMergeEvent]
}

func (s UcUserWidMergeListener) NewMessage() msg.MsgRequest[UcUserWidMergeEvent] {
	return &msg.XinyunOpenMessage[UcUserWidMergeEvent]{
		MsgBody: &UcUserWidMergeEvent{},
	}
}

func (s UcUserWidMergeListener) OnMessage(ctx context.Context, message msg.MsgRequest[UcUserWidMergeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[UcUserWidMergeEvent]))
}

type UcUserWidMergeEvent struct {
	Pid int64 `json:"pid,omitempty"`
	Wid int64 `json:"wid,omitempty"`
}

// OnUcUserWidMergeEvent
// @id 1155
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=wid合并通知消息
func (s *Service) OnUcUserWidMergeEvent(handler msg.XinyunEventHandler[UcUserWidMergeEvent]) (service *Service) {
	var listener = &UcUserWidMergeListener{handler}
	s.Register(msg.WrapListener[UcUserWidMergeEvent](listener), "uc_user", "wid_merge")
	return s
}
