package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcMarketingUserPraiseListener struct {
	handler msg.XinyunEventHandler[EcMarketingUserPraiseEvent]
}

func (s EcMarketingUserPraiseListener) NewMessage() msg.MsgRequest[EcMarketingUserPraiseEvent] {
	return &msg.XinyunOpenMessage[EcMarketingUserPraiseEvent]{
		MsgBody: &EcMarketingUserPraiseEvent{},
	}
}

func (s EcMarketingUserPraiseListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcMarketingUserPraiseEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcMarketingUserPraiseEvent]))
}

type EcMarketingUserPraiseEvent struct {
	Wid   int64 `json:"wid,omitempty"`
	BizId int64 `json:"bizId,omitempty"`
}

// OnEcMarketingUserPraiseEvent
// @id 1901
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=买家秀点赞
func (s *Service) OnEcMarketingUserPraiseEvent(handler msg.XinyunEventHandler[EcMarketingUserPraiseEvent]) (service *Service) {
	var listener = &EcMarketingUserPraiseListener{handler}
	s.Register(msg.WrapListener[EcMarketingUserPraiseEvent](listener), "ec_marketing", "userPraise")
	return s
}
