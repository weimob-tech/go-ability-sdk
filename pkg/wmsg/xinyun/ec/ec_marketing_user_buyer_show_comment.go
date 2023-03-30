package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcMarketingUserBuyerShowCommentListener struct {
	handler msg.XinyunEventHandler[EcMarketingUserBuyerShowCommentEvent]
}

func (s EcMarketingUserBuyerShowCommentListener) NewMessage() msg.MsgRequest[EcMarketingUserBuyerShowCommentEvent] {
	return &msg.XinyunOpenMessage[EcMarketingUserBuyerShowCommentEvent]{
		MsgBody: &EcMarketingUserBuyerShowCommentEvent{},
	}
}

func (s EcMarketingUserBuyerShowCommentListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcMarketingUserBuyerShowCommentEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcMarketingUserBuyerShowCommentEvent]))
}

type EcMarketingUserBuyerShowCommentEvent struct {
	Wid   int64 `json:"wid,omitempty"`
	BizId int64 `json:"bizId,omitempty"`
}

// OnEcMarketingUserBuyerShowCommentEvent
// @id 1902
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=买家秀评论
func (s *Service) OnEcMarketingUserBuyerShowCommentEvent(handler msg.XinyunEventHandler[EcMarketingUserBuyerShowCommentEvent]) (service *Service) {
	var listener = &EcMarketingUserBuyerShowCommentListener{handler}
	s.Register(msg.WrapListener[EcMarketingUserBuyerShowCommentEvent](listener), "ec_marketing", "userBuyerShowComment")
	return s
}
