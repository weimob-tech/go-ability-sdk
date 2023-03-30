package marketing

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type InterScrapecardPlayCardListener struct {
	handler msg.XinyunEventHandler[InterScrapecardPlayCardEvent]
}

func (s InterScrapecardPlayCardListener) NewMessage() msg.MsgRequest[InterScrapecardPlayCardEvent] {
	return &msg.XinyunOpenMessage[InterScrapecardPlayCardEvent]{
		MsgBody: &InterScrapecardPlayCardEvent{},
	}
}

func (s InterScrapecardPlayCardListener) OnMessage(ctx context.Context, message msg.MsgRequest[InterScrapecardPlayCardEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[InterScrapecardPlayCardEvent]))
}

type InterScrapecardPlayCardEvent struct {
	Unionid  string `json:"unionid,omitempty"`
	Wid      int64  `json:"wid,omitempty"`
	Openid   string `json:"openid,omitempty"`
	Playtime int64  `json:"playtime,omitempty"`
}

// OnInterScrapecardPlayCardEvent
// @id 1482
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=粉丝成功刮卡行为消息
func (s *Service) OnInterScrapecardPlayCardEvent(handler msg.XinyunEventHandler[InterScrapecardPlayCardEvent]) (service *Service) {
	var listener = &InterScrapecardPlayCardListener{handler}
	s.Register(msg.WrapListener[InterScrapecardPlayCardEvent](listener), "inter_scrapecard", "playCard")
	return s
}
