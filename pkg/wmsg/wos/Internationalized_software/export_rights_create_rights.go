package Internationalized_software

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type ExportRightsCreateRightsListener struct {
	handler msg.WosEventHandler[ExportRightsCreateRightsEvent]
}

func (s ExportRightsCreateRightsListener) NewMessage() msg.MsgRequest[ExportRightsCreateRightsEvent] {
	return &msg.WosOpenMessage[ExportRightsCreateRightsEvent]{
		MsgBody: &ExportRightsCreateRightsEvent{},
	}
}

func (s ExportRightsCreateRightsListener) OnMessage(ctx context.Context, message msg.MsgRequest[ExportRightsCreateRightsEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[ExportRightsCreateRightsEvent]))
}

type ExportRightsCreateRightsEvent struct {
	OrderNo     string `json:"orderNo,omitempty"`
	AftersaleNo string `json:"aftersaleNo,omitempty"`
}

// OnExportRightsCreateRightsEvent
// @id 1206
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1206?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商家新增售后)
func (s *Service) OnExportRightsCreateRightsEvent(handler msg.WosEventHandler[ExportRightsCreateRightsEvent]) (service *Service) {
	var listener = &ExportRightsCreateRightsListener{handler}
	s.Register(msg.WrapListener[ExportRightsCreateRightsEvent](listener), "export_rights", "createRights")
	return s
}
