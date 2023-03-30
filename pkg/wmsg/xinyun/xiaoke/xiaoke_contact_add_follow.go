package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeContactAddFollowListener struct {
	handler msg.XinyunEventHandler[XiaokeContactAddFollowEvent]
}

func (s XiaokeContactAddFollowListener) NewMessage() msg.MsgRequest[XiaokeContactAddFollowEvent] {
	return &msg.XinyunOpenMessage[XiaokeContactAddFollowEvent]{
		MsgBody: &XiaokeContactAddFollowEvent{},
	}
}

func (s XiaokeContactAddFollowListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeContactAddFollowEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeContactAddFollowEvent]))
}

type XiaokeContactAddFollowEvent struct {
	BusinessId      int64   `json:"businessId,omitempty"`
	Behavior        string  `json:"behavior,omitempty"`
	OpType          byte    `json:"opType,omitempty"`
	OpContentText   string  `json:"opContentText,omitempty"`
	OpContent       string  `json:"opContent,omitempty"`
	Remark          string  `json:"remark,omitempty"`
	OpTime          int     `json:"opTime,omitempty"`
	ImgAddress      string  `json:"imgAddress,omitempty"`
	VoiceAddress    string  `json:"voiceAddress,omitempty"`
	CheckInAddress  string  `json:"checkInAddress,omitempty"`
	Longitude       float64 `json:"longitude,omitempty"`
	Latitude        float64 `json:"latitude,omitempty"`
	FileAddress     string  `json:"fileAddress,omitempty"`
	ClueDescription string  `json:"clueDescription,omitempty"`
	VoiceText       string  `json:"voiceText,omitempty"`
}

// OnXiaokeContactAddFollowEvent
// @id 1856
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=新建联系人跟进记录
func (s *Service) OnXiaokeContactAddFollowEvent(handler msg.XinyunEventHandler[XiaokeContactAddFollowEvent]) (service *Service) {
	var listener = &XiaokeContactAddFollowListener{handler}
	s.Register(msg.WrapListener[XiaokeContactAddFollowEvent](listener), "xiaoke_contact", "addFollow")
	return s
}
