package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobGuideGuiderChangeGuiderUpdateListener struct {
	handler msg.WosEventHandler[WeimobGuideGuiderChangeGuiderUpdateEvent]
}

func (s WeimobGuideGuiderChangeGuiderUpdateListener) NewMessage() msg.MsgRequest[WeimobGuideGuiderChangeGuiderUpdateEvent] {
	return &msg.WosOpenMessage[WeimobGuideGuiderChangeGuiderUpdateEvent]{
		MsgBody: &WeimobGuideGuiderChangeGuiderUpdateEvent{},
	}
}

func (s WeimobGuideGuiderChangeGuiderUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobGuideGuiderChangeGuiderUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobGuideGuiderChangeGuiderUpdateEvent]))
}

type WeimobGuideGuiderChangeGuiderUpdateEvent struct {
	BosId                   string `json:"bosId,omitempty"`
	GuiderName              string `json:"guiderName,omitempty"`
	GuiderId                string `json:"guiderId,omitempty"`
	UpdateTime              string `json:"updateTime,omitempty"`
	GuiderWid               int64  `json:"guiderWid,omitempty"`
	GuiderVid               int64  `json:"guiderVid,omitempty"`
	GuiderVidType           int64  `json:"guiderVidType,omitempty"`
	JobNumber               string `json:"jobNumber,omitempty"`
	IsExclusiveCus          int64  `json:"isExclusiveCus,omitempty"`
	GuiderPhone             string `json:"guiderPhone,omitempty"`
	SourceProductId         int64  `json:"sourceProductId,omitempty"`
	SourceProductInstanceId int64  `json:"sourceProductInstanceId,omitempty"`
}

// OnWeimobGuideGuiderChangeGuiderUpdateEvent
// @id 1308
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1308?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=导购信息更新)
func (s *Service) OnWeimobGuideGuiderChangeGuiderUpdateEvent(handler msg.WosEventHandler[WeimobGuideGuiderChangeGuiderUpdateEvent]) (service *Service) {
	var listener = &WeimobGuideGuiderChangeGuiderUpdateListener{handler}
	s.Register(msg.WrapListener[WeimobGuideGuiderChangeGuiderUpdateEvent](listener), "weimob_guide.guider.change", "guiderUpdate")
	return s
}
