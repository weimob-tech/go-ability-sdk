package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmPointChangePointCleanListener struct {
	handler msg.WosEventHandler[WeimobCrmPointChangePointCleanEvent]
}

func (s WeimobCrmPointChangePointCleanListener) NewMessage() msg.MsgRequest[WeimobCrmPointChangePointCleanEvent] {
	return &msg.WosOpenMessage[WeimobCrmPointChangePointCleanEvent]{
		MsgBody: &WeimobCrmPointChangePointCleanEvent{},
	}
}

func (s WeimobCrmPointChangePointCleanListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmPointChangePointCleanEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmPointChangePointCleanEvent]))
}

type WeimobCrmPointChangePointCleanEvent struct {
	BizType           string `json:"bizType,omitempty"`
	BizUniqueId       int64  `json:"bizUniqueId,omitempty"`
	BosId             int64  `json:"bosId,omitempty"`
	ChangPoint        int64  `json:"changPoint,omitempty"`
	PointPlanId       int64  `json:"pointPlanId,omitempty"`
	ProductId         int64  `json:"productId,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	SourceProductId   int64  `json:"sourceProductId,omitempty"`
	TotalPoint        int64  `json:"totalPoint,omitempty"`
	TransType         int64  `json:"transType,omitempty"`
	Vid               int64  `json:"vid,omitempty"`
	Wid               int64  `json:"wid,omitempty"`
}

// OnWeimobCrmPointChangePointCleanEvent
// @id 1412
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1412?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=积分清零消息)
func (s *Service) OnWeimobCrmPointChangePointCleanEvent(handler msg.WosEventHandler[WeimobCrmPointChangePointCleanEvent]) (service *Service) {
	var listener = &WeimobCrmPointChangePointCleanListener{handler}
	s.Register(msg.WrapListener[WeimobCrmPointChangePointCleanEvent](listener), "weimob_crm.point.change", "pointClean")
	return s
}
