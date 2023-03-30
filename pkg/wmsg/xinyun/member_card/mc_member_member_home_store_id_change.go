package member_card

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type McMemberMemberHomeStoreIdChangeListener struct {
	handler msg.XinyunEventHandler[McMemberMemberHomeStoreIdChangeEvent]
}

func (s McMemberMemberHomeStoreIdChangeListener) NewMessage() msg.MsgRequest[McMemberMemberHomeStoreIdChangeEvent] {
	return &msg.XinyunOpenMessage[McMemberMemberHomeStoreIdChangeEvent]{
		MsgBody: &McMemberMemberHomeStoreIdChangeEvent{},
	}
}

func (s McMemberMemberHomeStoreIdChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[McMemberMemberHomeStoreIdChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[McMemberMemberHomeStoreIdChangeEvent]))
}

type McMemberMemberHomeStoreIdChangeEvent struct {
	Wid            int64 `json:"wid,omitempty"`
	OldHomeStoreId int64 `json:"oldHomeStoreId,omitempty"`
	NewHomeStoreId int64 `json:"newHomeStoreId,omitempty"`
}

// OnMcMemberMemberHomeStoreIdChangeEvent
// @id 1418
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=归属门店变更消息
func (s *Service) OnMcMemberMemberHomeStoreIdChangeEvent(handler msg.XinyunEventHandler[McMemberMemberHomeStoreIdChangeEvent]) (service *Service) {
	var listener = &McMemberMemberHomeStoreIdChangeListener{handler}
	s.Register(msg.WrapListener[McMemberMemberHomeStoreIdChangeEvent](listener), "mc_member", "memberHomeStoreIdChange")
	return s
}
