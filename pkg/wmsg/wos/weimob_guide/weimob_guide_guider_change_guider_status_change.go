package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobGuideGuiderChangeGuiderStatusChangeListener struct {
	handler msg.WosEventHandler[WeimobGuideGuiderChangeGuiderStatusChangeEvent]
}

func (s WeimobGuideGuiderChangeGuiderStatusChangeListener) NewMessage() msg.MsgRequest[WeimobGuideGuiderChangeGuiderStatusChangeEvent] {
	return &msg.WosOpenMessage[WeimobGuideGuiderChangeGuiderStatusChangeEvent]{
		MsgBody: &WeimobGuideGuiderChangeGuiderStatusChangeEvent{},
	}
}

func (s WeimobGuideGuiderChangeGuiderStatusChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobGuideGuiderChangeGuiderStatusChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobGuideGuiderChangeGuiderStatusChangeEvent]))
}

type WeimobGuideGuiderChangeGuiderStatusChangeEvent struct {
	BosId                   int64  `json:"bosId,omitempty"`
	GuiderId                string `json:"guiderId,omitempty"`
	IsUsed                  int64  `json:"isUsed,omitempty"`
	UpdateTime              string `json:"updateTime,omitempty"`
	GuiderWid               int64  `json:"guiderWid,omitempty"`
	GuiderVid               int64  `json:"guiderVid,omitempty"`
	GuiderVidType           int64  `json:"guiderVidType,omitempty"`
	SourceProductId         int64  `json:"sourceProductId,omitempty"`
	SourceProductInstanceId int64  `json:"sourceProductInstanceId,omitempty"`
}

// OnWeimobGuideGuiderChangeGuiderStatusChangeEvent
// @id 1309
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1309?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=导购启用状态变更)
func (s *Service) OnWeimobGuideGuiderChangeGuiderStatusChangeEvent(handler msg.WosEventHandler[WeimobGuideGuiderChangeGuiderStatusChangeEvent]) (service *Service) {
	var listener = &WeimobGuideGuiderChangeGuiderStatusChangeListener{handler}
	s.Register(msg.WrapListener[WeimobGuideGuiderChangeGuiderStatusChangeEvent](listener), "weimob_guide.guider.change", "guiderStatusChange")
	return s
}
