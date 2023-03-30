package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeOrderCreateApprovalListener struct {
	handler msg.XinyunEventHandler[XiaokeOrderCreateApprovalEvent]
}

func (s XiaokeOrderCreateApprovalListener) NewMessage() msg.MsgRequest[XiaokeOrderCreateApprovalEvent] {
	return &msg.XinyunOpenMessage[XiaokeOrderCreateApprovalEvent]{
		MsgBody: &XiaokeOrderCreateApprovalEvent{},
	}
}

func (s XiaokeOrderCreateApprovalListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeOrderCreateApprovalEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeOrderCreateApprovalEvent]))
}

type XiaokeOrderCreateApprovalEvent struct {
	OrderNumber    string `json:"orderNumber,omitempty"`
	WorkFlowStatus int64  `json:"workFlowStatus,omitempty"`
	BuildTime      int64  `json:"buildTime,omitempty"`
	Wid            int64  `json:"wid,omitempty"`
}

// OnXiaokeOrderCreateApprovalEvent
// @id 3065
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=创建审批流
func (s *Service) OnXiaokeOrderCreateApprovalEvent(handler msg.XinyunEventHandler[XiaokeOrderCreateApprovalEvent]) (service *Service) {
	var listener = &XiaokeOrderCreateApprovalListener{handler}
	s.Register(msg.WrapListener[XiaokeOrderCreateApprovalEvent](listener), "xiaoke_order", "createApproval")
	return s
}
