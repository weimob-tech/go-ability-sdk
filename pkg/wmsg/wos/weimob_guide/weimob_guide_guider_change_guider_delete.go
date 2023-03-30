package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobGuideGuiderChangeGuiderDeleteListener struct {
	handler msg.WosEventHandler[WeimobGuideGuiderChangeGuiderDeleteEvent]
}

func (s WeimobGuideGuiderChangeGuiderDeleteListener) NewMessage() msg.MsgRequest[WeimobGuideGuiderChangeGuiderDeleteEvent] {
	return &msg.WosOpenMessage[WeimobGuideGuiderChangeGuiderDeleteEvent]{
		MsgBody: &WeimobGuideGuiderChangeGuiderDeleteEvent{},
	}
}

func (s WeimobGuideGuiderChangeGuiderDeleteListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobGuideGuiderChangeGuiderDeleteEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobGuideGuiderChangeGuiderDeleteEvent]))
}

type WeimobGuideGuiderChangeGuiderDeleteEvent struct {
	BosId                   int64  `json:"bosId,omitempty"`
	GuiderId                string `json:"guiderId,omitempty"`
	UpdateTime              string `json:"updateTime,omitempty"`
	GuiderWid               int64  `json:"guiderWid,omitempty"`
	GuiderVid               int64  `json:"guiderVid,omitempty"`
	GuiderVidType           int64  `json:"guiderVidType,omitempty"`
	SourceProductId         int64  `json:"sourceProductId,omitempty"`
	SourceProductInstanceId int64  `json:"sourceProductInstanceId,omitempty"`
}

// OnWeimobGuideGuiderChangeGuiderDeleteEvent
// @id 1307
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1307?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除导购)
func (s *Service) OnWeimobGuideGuiderChangeGuiderDeleteEvent(handler msg.WosEventHandler[WeimobGuideGuiderChangeGuiderDeleteEvent]) (service *Service) {
	var listener = &WeimobGuideGuiderChangeGuiderDeleteListener{handler}
	s.Register(msg.WrapListener[WeimobGuideGuiderChangeGuiderDeleteEvent](listener), "weimob_guide.guider.change", "guiderDelete")
	return s
}
