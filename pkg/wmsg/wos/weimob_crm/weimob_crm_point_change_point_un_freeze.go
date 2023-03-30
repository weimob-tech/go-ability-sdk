package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmPointChangePointUnFreezeListener struct {
	handler msg.WosEventHandler[WeimobCrmPointChangePointUnFreezeEvent]
}

func (s WeimobCrmPointChangePointUnFreezeListener) NewMessage() msg.MsgRequest[WeimobCrmPointChangePointUnFreezeEvent] {
	return &msg.WosOpenMessage[WeimobCrmPointChangePointUnFreezeEvent]{
		MsgBody: &WeimobCrmPointChangePointUnFreezeEvent{},
	}
}

func (s WeimobCrmPointChangePointUnFreezeListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmPointChangePointUnFreezeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmPointChangePointUnFreezeEvent]))
}

type WeimobCrmPointChangePointUnFreezeEvent struct {
	BosId             int64  `json:"bosId,omitempty"`
	Wid               int64  `json:"wid,omitempty"`
	Vid               int64  `json:"vid,omitempty"`
	ProductId         int64  `json:"productId,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	PointPlanId       int64  `json:"pointPlanId,omitempty"`
	SourceProductId   int64  `json:"sourceProductId,omitempty"`
	TransTime         string `json:"transTime,omitempty"`
	TransNo           int64  `json:"transNo,omitempty"`
	RequestId         string `json:"requestId,omitempty"`
	RequestType       string `json:"requestType,omitempty"`
	Point             string `json:"point,omitempty"`
}

// OnWeimobCrmPointChangePointUnFreezeEvent
// @id 1414
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1414?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=积分解冻消息)
func (s *Service) OnWeimobCrmPointChangePointUnFreezeEvent(handler msg.WosEventHandler[WeimobCrmPointChangePointUnFreezeEvent]) (service *Service) {
	var listener = &WeimobCrmPointChangePointUnFreezeListener{handler}
	s.Register(msg.WrapListener[WeimobCrmPointChangePointUnFreezeEvent](listener), "weimob_crm.point.change", "pointUnFreeze")
	return s
}
