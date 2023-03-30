package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobGuideGuiderChangeGuiderVidChangeListener struct {
	handler msg.WosEventHandler[WeimobGuideGuiderChangeGuiderVidChangeEvent]
}

func (s WeimobGuideGuiderChangeGuiderVidChangeListener) NewMessage() msg.MsgRequest[WeimobGuideGuiderChangeGuiderVidChangeEvent] {
	return &msg.WosOpenMessage[WeimobGuideGuiderChangeGuiderVidChangeEvent]{
		MsgBody: &WeimobGuideGuiderChangeGuiderVidChangeEvent{},
	}
}

func (s WeimobGuideGuiderChangeGuiderVidChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobGuideGuiderChangeGuiderVidChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobGuideGuiderChangeGuiderVidChangeEvent]))
}

type WeimobGuideGuiderChangeGuiderVidChangeEvent struct {
	BosId                   int64  `json:"bosId,omitempty"`
	GuiderId                string `json:"guiderId,omitempty"`
	UpdateTime              string `json:"updateTime,omitempty"`
	GuiderWid               int64  `json:"guiderWid,omitempty"`
	GuiderVid               int64  `json:"guiderVid,omitempty"`
	GuiderVidType           int64  `json:"guiderVidType,omitempty"`
	NewGuiderVid            int64  `json:"newGuiderVid,omitempty"`
	NewGuiderVidType        int64  `json:"newGuiderVidType,omitempty"`
	NewGuiderId             string `json:"newGuiderId,omitempty"`
	SourceProductId         int64  `json:"sourceProductId,omitempty"`
	SourceProductInstanceId int64  `json:"sourceProductInstanceId,omitempty"`
}

// OnWeimobGuideGuiderChangeGuiderVidChangeEvent
// @id 1313
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1313?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=导购的归属门店变更)
func (s *Service) OnWeimobGuideGuiderChangeGuiderVidChangeEvent(handler msg.WosEventHandler[WeimobGuideGuiderChangeGuiderVidChangeEvent]) (service *Service) {
	var listener = &WeimobGuideGuiderChangeGuiderVidChangeListener{handler}
	s.Register(msg.WrapListener[WeimobGuideGuiderChangeGuiderVidChangeEvent](listener), "weimob_guide.guider.change", "guiderVidChange")
	return s
}
