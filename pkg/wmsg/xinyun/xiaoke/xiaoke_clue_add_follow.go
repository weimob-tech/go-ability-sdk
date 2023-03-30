package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeClueAddFollowListener struct {
	handler msg.XinyunEventHandler[XiaokeClueAddFollowEvent]
}

func (s XiaokeClueAddFollowListener) NewMessage() msg.MsgRequest[XiaokeClueAddFollowEvent] {
	return &msg.XinyunOpenMessage[XiaokeClueAddFollowEvent]{
		MsgBody: &XiaokeClueAddFollowEvent{},
	}
}

func (s XiaokeClueAddFollowListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeClueAddFollowEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeClueAddFollowEvent]))
}

type XiaokeClueAddFollowEvent struct {
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

// OnXiaokeClueAddFollowEvent
// @id 1758
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=新建跟进记录
func (s *Service) OnXiaokeClueAddFollowEvent(handler msg.XinyunEventHandler[XiaokeClueAddFollowEvent]) (service *Service) {
	var listener = &XiaokeClueAddFollowListener{handler}
	s.Register(msg.WrapListener[XiaokeClueAddFollowEvent](listener), "xiaoke_clue", "addFollow")
	return s
}
