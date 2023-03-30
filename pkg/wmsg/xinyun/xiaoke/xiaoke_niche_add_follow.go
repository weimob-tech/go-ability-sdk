package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeNicheAddFollowListener struct {
	handler msg.XinyunEventHandler[XiaokeNicheAddFollowEvent]
}

func (s XiaokeNicheAddFollowListener) NewMessage() msg.MsgRequest[XiaokeNicheAddFollowEvent] {
	return &msg.XinyunOpenMessage[XiaokeNicheAddFollowEvent]{
		MsgBody: &XiaokeNicheAddFollowEvent{},
	}
}

func (s XiaokeNicheAddFollowListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeNicheAddFollowEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeNicheAddFollowEvent]))
}

type XiaokeNicheAddFollowEvent struct {
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

// OnXiaokeNicheAddFollowEvent
// @id 1774
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=新建跟进记录
func (s *Service) OnXiaokeNicheAddFollowEvent(handler msg.XinyunEventHandler[XiaokeNicheAddFollowEvent]) (service *Service) {
	var listener = &XiaokeNicheAddFollowListener{handler}
	s.Register(msg.WrapListener[XiaokeNicheAddFollowEvent](listener), "xiaoke_niche", "addFollow")
	return s
}
