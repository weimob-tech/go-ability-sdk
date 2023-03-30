package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobGuideGuiderCreationGuiderCreationListener struct {
	handler msg.WosEventHandler[WeimobGuideGuiderCreationGuiderCreationEvent]
}

func (s WeimobGuideGuiderCreationGuiderCreationListener) NewMessage() msg.MsgRequest[WeimobGuideGuiderCreationGuiderCreationEvent] {
	return &msg.WosOpenMessage[WeimobGuideGuiderCreationGuiderCreationEvent]{
		MsgBody: &WeimobGuideGuiderCreationGuiderCreationEvent{},
	}
}

func (s WeimobGuideGuiderCreationGuiderCreationListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobGuideGuiderCreationGuiderCreationEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobGuideGuiderCreationGuiderCreationEvent]))
}

type WeimobGuideGuiderCreationGuiderCreationEvent struct {
	BosId                   int64  `json:"bosId,omitempty"`
	GuiderVid               int64  `json:"guiderVid,omitempty"`
	GuiderWid               int64  `json:"guiderWid,omitempty"`
	SourceProductId         int64  `json:"sourceProductId,omitempty"`
	SourceProductInstanceId int64  `json:"sourceProductInstanceId,omitempty"`
	CreateTime              string `json:"createTime,omitempty"`
	GuiderName              string `json:"guiderName,omitempty"`
	GuiderId                string `json:"guiderId,omitempty"`
	GuiderPhone             string `json:"guiderPhone,omitempty"`
	GuiderVidType           int64  `json:"guiderVidType,omitempty"`
	IsUsed                  int64  `json:"isUsed,omitempty"`
	IsExclusiveCus          int64  `json:"isExclusiveCus,omitempty"`
}

// OnWeimobGuideGuiderCreationGuiderCreationEvent
// @id 1242
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1242?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=新增导购)
func (s *Service) OnWeimobGuideGuiderCreationGuiderCreationEvent(handler msg.WosEventHandler[WeimobGuideGuiderCreationGuiderCreationEvent]) (service *Service) {
	var listener = &WeimobGuideGuiderCreationGuiderCreationListener{handler}
	s.Register(msg.WrapListener[WeimobGuideGuiderCreationGuiderCreationEvent](listener), "weimob_guide.guider.creation", "guiderCreation")
	return s
}
