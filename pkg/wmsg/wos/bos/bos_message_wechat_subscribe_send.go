package bos

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type BosMessageWechatSubscribeSendListener struct {
	handler msg.WosEventHandler[BosMessageWechatSubscribeSendEvent]
}

func (s BosMessageWechatSubscribeSendListener) NewMessage() msg.MsgRequest[BosMessageWechatSubscribeSendEvent] {
	return &msg.WosOpenMessage[BosMessageWechatSubscribeSendEvent]{
		MsgBody: &BosMessageWechatSubscribeSendEvent{},
	}
}

func (s BosMessageWechatSubscribeSendListener) OnMessage(ctx context.Context, message msg.MsgRequest[BosMessageWechatSubscribeSendEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[BosMessageWechatSubscribeSendEvent]))
}

type BosMessageWechatSubscribeSendEvent struct {
	RequestId string `json:"requestId,omitempty"`
	Wid       int64  `json:"wid,omitempty"`
	ErrorCode string `json:"errorCode,omitempty"`
	ErrorMsg  string `json:"errorMsg,omitempty"`
}

// OnBosMessageWechatSubscribeSendEvent
// @id 1424
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1424?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=微信消息发送结果通知)
func (s *Service) OnBosMessageWechatSubscribeSendEvent(handler msg.WosEventHandler[BosMessageWechatSubscribeSendEvent]) (service *Service) {
	var listener = &BosMessageWechatSubscribeSendListener{handler}
	s.Register(msg.WrapListener[BosMessageWechatSubscribeSendEvent](listener), "bos.message.wechat.subscribe", "send")
	return s
}
