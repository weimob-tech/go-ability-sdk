package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCrmPointChangePointChangeListener struct {
	handler msg.WosEventHandler[WeimobCrmPointChangePointChangeEvent]
}

func (s WeimobCrmPointChangePointChangeListener) NewMessage() msg.MsgRequest[WeimobCrmPointChangePointChangeEvent] {
	return &msg.WosOpenMessage[WeimobCrmPointChangePointChangeEvent]{
		MsgBody: &WeimobCrmPointChangePointChangeEvent{},
	}
}

func (s WeimobCrmPointChangePointChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCrmPointChangePointChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCrmPointChangePointChangeEvent]))
}

type WeimobCrmPointChangePointChangeEvent struct {
	ProductId         int64  `json:"productId,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	SourceProductId   int64  `json:"sourceProductId,omitempty"`
	Tcode             string `json:"tcode,omitempty"`
	BosId             int64  `json:"bosId,omitempty"`
	Vid               int64  `json:"vid,omitempty"`
	Wid               int64  `json:"wid,omitempty"`
	PointPlanId       int64  `json:"pointPlanId,omitempty"`
	TransType         int64  `json:"transType,omitempty"`
	BizType           string `json:"bizType,omitempty"`
	BizUniqueId       int64  `json:"bizUniqueId,omitempty"`
	ChangPoint        string `json:"changPoint,omitempty"`
	TotalPoint        string `json:"totalPoint,omitempty"`
	ChangeTime        int64  `json:"changeTime,omitempty"`
	EffectiveTime     int64  `json:"effectiveTime,omitempty"`
	ExpireTime        int64  `json:"expireTime,omitempty"`
}

// OnWeimobCrmPointChangePointChangeEvent
// @id 1245
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1245?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=积分变动消息)
func (s *Service) OnWeimobCrmPointChangePointChangeEvent(handler msg.WosEventHandler[WeimobCrmPointChangePointChangeEvent]) (service *Service) {
	var listener = &WeimobCrmPointChangePointChangeListener{handler}
	s.Register(msg.WrapListener[WeimobCrmPointChangePointChangeEvent](listener), "weimob_crm.point.change", "pointChange")
	return s
}
