package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmPointChangePointFreezeListener struct {
	handler msg.WosEventHandler[WeimobCrmPointChangePointFreezeEvent]
}

func (s WeimobCrmPointChangePointFreezeListener) NewMessage() msg.MsgRequest[WeimobCrmPointChangePointFreezeEvent] {
	return &msg.WosOpenMessage[WeimobCrmPointChangePointFreezeEvent]{
		MsgBody: &WeimobCrmPointChangePointFreezeEvent{},
	}
}

func (s WeimobCrmPointChangePointFreezeListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmPointChangePointFreezeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmPointChangePointFreezeEvent]))
}

type WeimobCrmPointChangePointFreezeEvent struct {
	BosId             int64  `json:"bosId,omitempty"`
	Wid               int64  `json:"wid,omitempty"`
	Vid               int64  `json:"vid,omitempty"`
	Point             string `json:"point,omitempty"`
	ProductId         int64  `json:"productId,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	PointPlanId       int64  `json:"pointPlanId,omitempty"`
	SourceProductId   int64  `json:"sourceProductId,omitempty"`
	TransTime         string `json:"transTime,omitempty"`
	TransNo           int64  `json:"transNo,omitempty"`
	ChangeType        string `json:"changeType,omitempty"`
	ChannelType       string `json:"channelType,omitempty"`
	OccurVid          int64  `json:"occurVid,omitempty"`
	Remark            string `json:"remark,omitempty"`
	OperatorWid       int64  `json:"operatorWid,omitempty"`
	RequestId         string `json:"requestId,omitempty"`
	RequestType       string `json:"requestType,omitempty"`
}

// OnWeimobCrmPointChangePointFreezeEvent
// @id 1415
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1415?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=积分冻结消息)
func (s *Service) OnWeimobCrmPointChangePointFreezeEvent(handler msg.WosEventHandler[WeimobCrmPointChangePointFreezeEvent]) (service *Service) {
	var listener = &WeimobCrmPointChangePointFreezeListener{handler}
	s.Register(msg.WrapListener[WeimobCrmPointChangePointFreezeEvent](listener), "weimob_crm.point.change", "pointFreeze")
	return s
}
