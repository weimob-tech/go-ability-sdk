package weimob_notes

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobNotesContentConsumerActionCommentListener struct {
	handler msg.WosEventHandler[WeimobNotesContentConsumerActionCommentEvent]
}

func (s WeimobNotesContentConsumerActionCommentListener) NewMessage() msg.MsgRequest[WeimobNotesContentConsumerActionCommentEvent] {
	return &msg.WosOpenMessage[WeimobNotesContentConsumerActionCommentEvent]{
		MsgBody: &WeimobNotesContentConsumerActionCommentEvent{},
	}
}

func (s WeimobNotesContentConsumerActionCommentListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobNotesContentConsumerActionCommentEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobNotesContentConsumerActionCommentEvent]))
}

type WeimobNotesContentConsumerActionCommentEvent struct {
	Wid       int64 `json:"wid,omitempty"`
	ContentId int64 `json:"contentId,omitempty"`
}

// OnWeimobNotesContentConsumerActionCommentEvent
// @id 1287
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1287?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=用户评论内容消息)
func (s *Service) OnWeimobNotesContentConsumerActionCommentEvent(handler msg.WosEventHandler[WeimobNotesContentConsumerActionCommentEvent]) (service *Service) {
	var listener = &WeimobNotesContentConsumerActionCommentListener{handler}
	s.Register(msg.WrapListener[WeimobNotesContentConsumerActionCommentEvent](listener), "weimob_notes.content.consumer.action", "comment")
	return s
}
