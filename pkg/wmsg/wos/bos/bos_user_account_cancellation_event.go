package bos

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type BosUserAccountCancellationEventListener struct {
	handler msg.WosEventHandler[BosUserAccountCancellationEventEvent]
}

func (s BosUserAccountCancellationEventListener) NewMessage() msg.MsgRequest[BosUserAccountCancellationEventEvent] {
	return &msg.WosOpenMessage[BosUserAccountCancellationEventEvent]{
		MsgBody: &BosUserAccountCancellationEventEvent{},
	}
}

func (s BosUserAccountCancellationEventListener) OnMessage(ctx context.Context, message msg.MsgRequest[BosUserAccountCancellationEventEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[BosUserAccountCancellationEventEvent]))
}

type BosUserAccountCancellationEventEvent struct {
	WidSourceList []BosUserAccountCancellationEventWidSourceList `json:"widSourceList,omitempty"`
	Wid           int64                                          `json:"wid,omitempty"`
	BosId         int64                                          `json:"bosId,omitempty"`
	LogoutTime    int64                                          `json:"logoutTime,omitempty"`
	Time          int64                                          `json:"time,omitempty"`
}

type BosUserAccountCancellationEventWidSourceList struct {
	Source      int64  `json:"source,omitempty"`
	AppId       string `json:"appId,omitempty"`
	OriginalId  string `json:"originalId,omitempty"`
	GmtCreate   int64  `json:"gmtCreate,omitempty"`
	GmtModified int64  `json:"gmtModified,omitempty"`
}

// OnBosUserAccountCancellationEventEvent
// @id 1366
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1366?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=用户注销消息)
func (s *Service) OnBosUserAccountCancellationEventEvent(handler msg.WosEventHandler[BosUserAccountCancellationEventEvent]) (service *Service) {
	var listener = &BosUserAccountCancellationEventListener{handler}
	s.Register(msg.WrapListener[BosUserAccountCancellationEventEvent](listener), "bos.user", "accountCancellationEvent")
	return s
}
