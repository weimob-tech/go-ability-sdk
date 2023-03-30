package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeCustomerAddFollowListener struct {
	handler msg.XinyunEventHandler[XiaokeCustomerAddFollowEvent]
}

func (s XiaokeCustomerAddFollowListener) NewMessage() msg.MsgRequest[XiaokeCustomerAddFollowEvent] {
	return &msg.XinyunOpenMessage[XiaokeCustomerAddFollowEvent]{
		MsgBody: &XiaokeCustomerAddFollowEvent{},
	}
}

func (s XiaokeCustomerAddFollowListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeCustomerAddFollowEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeCustomerAddFollowEvent]))
}

type XiaokeCustomerAddFollowEvent struct {
	BusinessId      int64   `json:"businessId,omitempty"`
	Behavior        string  `json:"behavior,omitempty"`
	OpType          byte    `json:"opType,omitempty"`
	OpContentText   string  `json:"opContentText,omitempty"`
	OpContent       string  `json:"opContent,omitempty"`
	Remark          string  `json:"remark,omitempty"`
	OpTime          int     `json:"opTime,omitempty"`
	ImgAddress      string  `json:"imgAddress,omitempty"`
	CheckInAddress  string  `json:"checkInAddress,omitempty"`
	Longitude       float64 `json:"longitude,omitempty"`
	Latitude        float64 `json:"latitude,omitempty"`
	ClueDescription string  `json:"clueDescription,omitempty"`
	VoiceText       string  `json:"voiceText,omitempty"`
}

// OnXiaokeCustomerAddFollowEvent
// @id 1767
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=新建跟进记录
func (s *Service) OnXiaokeCustomerAddFollowEvent(handler msg.XinyunEventHandler[XiaokeCustomerAddFollowEvent]) (service *Service) {
	var listener = &XiaokeCustomerAddFollowListener{handler}
	s.Register(msg.WrapListener[XiaokeCustomerAddFollowEvent](listener), "xiaoke_customer", "addFollow")
	return s
}
