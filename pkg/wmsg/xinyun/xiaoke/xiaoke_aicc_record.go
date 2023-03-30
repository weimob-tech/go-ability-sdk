package xiaoke

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type XiaokeAiccRecordListener struct {
	handler msg.XinyunEventHandler[XiaokeAiccRecordEvent]
}

func (s XiaokeAiccRecordListener) NewMessage() msg.MsgRequest[XiaokeAiccRecordEvent] {
	return &msg.XinyunOpenMessage[XiaokeAiccRecordEvent]{
		MsgBody: &XiaokeAiccRecordEvent{},
	}
}

func (s XiaokeAiccRecordListener) OnMessage(ctx context.Context, message msg.MsgRequest[XiaokeAiccRecordEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[XiaokeAiccRecordEvent]))
}

type XiaokeAiccRecordEvent struct {
	NumberData      XiaokeAiccRecordNumberData `json:"numberData,omitempty"`
	AAudioUrl       string                     `json:"aAudioUrl,omitempty"`
	ATaskName       string                     `json:"aTaskName,omitempty"`
	AUserId         string                     `json:"aUserId,omitempty"`
	Bill            int64                      `json:"bill,omitempty"`
	CallDate        int64                      `json:"callDate,omitempty"`
	CallId          string                     `json:"callId,omitempty"`
	LevelId         int64                      `json:"levelId,omitempty"`
	LevelName       string                     `json:"levelName,omitempty"`
	RecordId        string                     `json:"recordId,omitempty"`
	Status          int64                      `json:"status,omitempty"`
	TaskId          int64                      `json:"taskId,omitempty"`
	TurnCount       int64                      `json:"turnCount,omitempty"`
	AutoAddWechat   int64                      `json:"autoAddWechat,omitempty"`
	AddWechatStatus int64                      `json:"addWechatStatus,omitempty"`
	BosId           int64                      `json:"bosId,omitempty"`
}

type XiaokeAiccRecordNumberData struct {
	HandUpDate int64  `json:"handUpDate,omitempty"`
	Number     string `json:"number,omitempty"`
}

// OnXiaokeAiccRecordEvent
// @id 3735
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=通话记录回调
func (s *Service) OnXiaokeAiccRecordEvent(handler msg.XinyunEventHandler[XiaokeAiccRecordEvent]) (service *Service) {
	var listener = &XiaokeAiccRecordListener{handler}
	s.Register(msg.WrapListener[XiaokeAiccRecordEvent](listener), "xiaoke_aicc", "record")
	return s
}
