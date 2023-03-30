package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeOrderChangeApprovalStatusListener struct {
	handler msg.XinyunEventHandler[XiaokeOrderChangeApprovalStatusEvent]
}

func (s XiaokeOrderChangeApprovalStatusListener) NewMessage() msg.MsgRequest[XiaokeOrderChangeApprovalStatusEvent] {
	return &msg.XinyunOpenMessage[XiaokeOrderChangeApprovalStatusEvent]{
		MsgBody: &XiaokeOrderChangeApprovalStatusEvent{},
	}
}

func (s XiaokeOrderChangeApprovalStatusListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeOrderChangeApprovalStatusEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeOrderChangeApprovalStatusEvent]))
}

type XiaokeOrderChangeApprovalStatusEvent struct {
	OrderNumber    string `json:"orderNumber,omitempty"`
	WorkFlowStatus int64  `json:"workFlowStatus,omitempty"`
	BuildTime      int64  `json:"buildTime,omitempty"`
	Wid            int64  `json:"wid,omitempty"`
}

// OnXiaokeOrderChangeApprovalStatusEvent
// @id 3255
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=审批流状态变更
func (s *Service) OnXiaokeOrderChangeApprovalStatusEvent(handler msg.XinyunEventHandler[XiaokeOrderChangeApprovalStatusEvent]) (service *Service) {
	var listener = &XiaokeOrderChangeApprovalStatusListener{handler}
	s.Register(msg.WrapListener[XiaokeOrderChangeApprovalStatusEvent](listener), "xiaoke_order", "changeApprovalStatus")
	return s
}
