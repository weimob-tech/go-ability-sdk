package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcMarketingUserFeatureArticleCommentListener struct {
	handler msg.XinyunEventHandler[EcMarketingUserFeatureArticleCommentEvent]
}

func (s EcMarketingUserFeatureArticleCommentListener) NewMessage() msg.MsgRequest[EcMarketingUserFeatureArticleCommentEvent] {
	return &msg.XinyunOpenMessage[EcMarketingUserFeatureArticleCommentEvent]{
		MsgBody: &EcMarketingUserFeatureArticleCommentEvent{},
	}
}

func (s EcMarketingUserFeatureArticleCommentListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcMarketingUserFeatureArticleCommentEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcMarketingUserFeatureArticleCommentEvent]))
}

type EcMarketingUserFeatureArticleCommentEvent struct {
	Wid   int64 `json:"wid,omitempty"`
	BizId int64 `json:"bizId,omitempty"`
}

// OnEcMarketingUserFeatureArticleCommentEvent
// @id 1903
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=专题文章评论
func (s *Service) OnEcMarketingUserFeatureArticleCommentEvent(handler msg.XinyunEventHandler[EcMarketingUserFeatureArticleCommentEvent]) (service *Service) {
	var listener = &EcMarketingUserFeatureArticleCommentListener{handler}
	s.Register(msg.WrapListener[EcMarketingUserFeatureArticleCommentEvent](listener), "ec_marketing", "userFeatureArticleComment")
	return s
}
