package Internationalized_software

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type ExportOrderOrderStatusChangeListener struct {
	handler msg.WosEventHandler[ExportOrderOrderStatusChangeEvent]
}

func (s ExportOrderOrderStatusChangeListener) NewMessage() msg.MsgRequest[ExportOrderOrderStatusChangeEvent] {
	return &msg.WosOpenMessage[ExportOrderOrderStatusChangeEvent]{
		MsgBody: &ExportOrderOrderStatusChangeEvent{},
	}
}

func (s ExportOrderOrderStatusChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[ExportOrderOrderStatusChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[ExportOrderOrderStatusChangeEvent]))
}

type ExportOrderOrderStatusChangeEvent struct {
	OrderNo         string `json:"orderNo,omitempty"`
	OrderStatus     string `json:"orderStatus,omitempty"`
	OrderStatusText string `json:"orderStatusText,omitempty"`
}

// OnExportOrderOrderStatusChangeEvent
// @id 1205
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1205?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=订单状态发生变更)
func (s *Service) OnExportOrderOrderStatusChangeEvent(handler msg.WosEventHandler[ExportOrderOrderStatusChangeEvent]) (service *Service) {
	var listener = &ExportOrderOrderStatusChangeListener{handler}
	s.Register(msg.WrapListener[ExportOrderOrderStatusChangeEvent](listener), "export_order", "orderStatusChange")
	return s
}
