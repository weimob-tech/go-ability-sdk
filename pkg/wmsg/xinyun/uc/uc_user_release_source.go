package uc

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type UcUserReleaseSourceListener struct {
	handler msg.XinyunEventHandler[UcUserReleaseSourceEvent]
}

func (s UcUserReleaseSourceListener) NewMessage() msg.MsgRequest[UcUserReleaseSourceEvent] {
	return &msg.XinyunOpenMessage[UcUserReleaseSourceEvent]{
		MsgBody: &UcUserReleaseSourceEvent{},
	}
}

func (s UcUserReleaseSourceListener) OnMessage(ctx context.Context, message msg.MsgRequest[UcUserReleaseSourceEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[UcUserReleaseSourceEvent]))
}

type UcUserReleaseSourceEvent struct {
	Wid          int64  `json:"wid,omitempty"`
	Source       int64  `json:"source,omitempty"`
	SourceAppid  string `json:"sourceAppid,omitempty"`
	SourceOpenid string `json:"sourceOpenid,omitempty"`
}

// OnUcUserReleaseSourceEvent
// @id 1486
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=wid渠道释放消息
func (s *Service) OnUcUserReleaseSourceEvent(handler msg.XinyunEventHandler[UcUserReleaseSourceEvent]) (service *Service) {
	var listener = &UcUserReleaseSourceListener{handler}
	s.Register(msg.WrapListener[UcUserReleaseSourceEvent](listener), "uc_user", "release_source")
	return s
}
