package niepan

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type NiepanChannelRightsCreateRightsListener struct {
	handler msg.WosEventHandler[NiepanChannelRightsCreateRightsEvent]
}

func (s NiepanChannelRightsCreateRightsListener) NewMessage() msg.MsgRequest[NiepanChannelRightsCreateRightsEvent] {
	return &msg.WosOpenMessage[NiepanChannelRightsCreateRightsEvent]{
		MsgBody: &NiepanChannelRightsCreateRightsEvent{},
	}
}

func (s NiepanChannelRightsCreateRightsListener) OnMessage(ctx context.Context, message msg.MsgRequest[NiepanChannelRightsCreateRightsEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[NiepanChannelRightsCreateRightsEvent]))
}

type NiepanChannelRightsCreateRightsEvent struct {
	AuxiliaryInfo      []NiepanChannelRightsCreateRightsAuxiliaryInfo `json:"auxiliaryInfo,omitempty"`
	MenuResInfo        NiepanChannelRightsCreateRightsMenuResInfo     `json:"menuResInfo,omitempty"`
	StartTime          NiepanChannelRightsCreateRightsStartTime       `json:"startTime,omitempty"`
	MsgType            string                                         `json:"msgType,omitempty"`
	ProductId          int64                                          `json:"productId,omitempty"`
	ProductInstanceId  int64                                          `json:"productInstanceId,omitempty"`
	RecordId           int64                                          `json:"recordId,omitempty"`
	MerchantId         int64                                          `json:"merchantId,omitempty"`
	BosId              int64                                          `json:"bosId,omitempty"`
	ProductVersionName string                                         `json:"productVersionName,omitempty"`
	EndTime            int64                                          `json:"endTime,omitempty"`
	ProductVersionId   int64                                          `json:"productVersionId,omitempty"`
}

type NiepanChannelRightsCreateRightsAuxiliaryInfo struct {
	Code  string  `json:"code,omitempty"`
	Num   int64   `json:"num,omitempty"`
	Type  int64   `json:"type,omitempty"`
	Name  string  `json:"name,omitempty"`
	Value []int64 `json:"value,omitempty"`
}
type NiepanChannelRightsCreateRightsMenuResInfo struct {
}
type NiepanChannelRightsCreateRightsStartTime struct {
}

// OnNiepanChannelRightsCreateRightsEvent
// @id 1334
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1334?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=权益创建消息推送微盟云)
func (s *Service) OnNiepanChannelRightsCreateRightsEvent(handler msg.WosEventHandler[NiepanChannelRightsCreateRightsEvent]) (service *Service) {
	var listener = &NiepanChannelRightsCreateRightsListener{handler}
	s.Register(msg.WrapListener[NiepanChannelRightsCreateRightsEvent](listener), "niepan.channel.rights", "createRights")
	return s
}
