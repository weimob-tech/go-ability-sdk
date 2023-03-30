package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobGuideGuiderRelationshipRelationshipChangeListener struct {
	handler msg.WosEventHandler[WeimobGuideGuiderRelationshipRelationshipChangeEvent]
}

func (s WeimobGuideGuiderRelationshipRelationshipChangeListener) NewMessage() msg.MsgRequest[WeimobGuideGuiderRelationshipRelationshipChangeEvent] {
	return &msg.WosOpenMessage[WeimobGuideGuiderRelationshipRelationshipChangeEvent]{
		MsgBody: &WeimobGuideGuiderRelationshipRelationshipChangeEvent{},
	}
}

func (s WeimobGuideGuiderRelationshipRelationshipChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobGuideGuiderRelationshipRelationshipChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobGuideGuiderRelationshipRelationshipChangeEvent]))
}

type WeimobGuideGuiderRelationshipRelationshipChangeEvent struct {
	BosId                   string `json:"bosId,omitempty"`
	BeforeVid               string `json:"beforeVid,omitempty"`
	AfterVid                string `json:"afterVid,omitempty"`
	CustomWid               string `json:"customWid,omitempty"`
	BeforeGuiderWid         string `json:"beforeGuiderWid,omitempty"`
	AfterGuiderWid          string `json:"afterGuiderWid,omitempty"`
	EventType               string `json:"eventType,omitempty"`
	SourceProductId         string `json:"sourceProductId,omitempty"`
	SourceProductInstanceId string `json:"sourceProductInstanceId,omitempty"`
	Time                    int64  `json:"time,omitempty"`
	BeforeGuiderId          string `json:"beforeGuiderId,omitempty"`
	AfterGuiderId           string `json:"afterGuiderId,omitempty"`
}

// OnWeimobGuideGuiderRelationshipRelationshipChangeEvent
// @id 1240
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1240?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=导购关系链变更消息)
func (s *Service) OnWeimobGuideGuiderRelationshipRelationshipChangeEvent(handler msg.WosEventHandler[WeimobGuideGuiderRelationshipRelationshipChangeEvent]) (service *Service) {
	var listener = &WeimobGuideGuiderRelationshipRelationshipChangeListener{handler}
	s.Register(msg.WrapListener[WeimobGuideGuiderRelationshipRelationshipChangeEvent](listener), "weimob_guide.guider.relationship", "relationshipChange")
	return s
}
