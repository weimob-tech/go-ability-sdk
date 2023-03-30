package Internationalized_software

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type ExportRightsCompleteRightsListener struct {
	handler msg.WosEventHandler[ExportRightsCompleteRightsEvent]
}

func (s ExportRightsCompleteRightsListener) NewMessage() msg.MsgRequest[ExportRightsCompleteRightsEvent] {
	return &msg.WosOpenMessage[ExportRightsCompleteRightsEvent]{
		MsgBody: &ExportRightsCompleteRightsEvent{},
	}
}

func (s ExportRightsCompleteRightsListener) OnMessage(ctx context.Context, message msg.MsgRequest[ExportRightsCompleteRightsEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[ExportRightsCompleteRightsEvent]))
}

type ExportRightsCompleteRightsEvent struct {
	OrderNo     string `json:"orderNo,omitempty"`
	AftersaleNo string `json:"aftersaleNo,omitempty"`
}

// OnExportRightsCompleteRightsEvent
// @id 1207
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1207?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商家完成售后)
func (s *Service) OnExportRightsCompleteRightsEvent(handler msg.WosEventHandler[ExportRightsCompleteRightsEvent]) (service *Service) {
	var listener = &ExportRightsCompleteRightsListener{handler}
	s.Register(msg.WrapListener[ExportRightsCompleteRightsEvent](listener), "export_rights", "completeRights")
	return s
}
