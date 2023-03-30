package member

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcLoginAgreeProtocolListener struct {
	handler msg.XinyunEventHandler[EcLoginAgreeProtocolEvent]
}

func (s EcLoginAgreeProtocolListener) NewMessage() msg.MsgRequest[EcLoginAgreeProtocolEvent] {
	return &msg.XinyunOpenMessage[EcLoginAgreeProtocolEvent]{
		MsgBody: &EcLoginAgreeProtocolEvent{},
	}
}

func (s EcLoginAgreeProtocolListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcLoginAgreeProtocolEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcLoginAgreeProtocolEvent]))
}

type EcLoginAgreeProtocolEvent struct {
	Pid               int64    `json:"pid,omitempty"`
	StoreId           int64    `json:"storeId,omitempty"`
	AgreementType     int64    `json:"agreementType,omitempty"`
	ProtocolFileNames []string `json:"protocolFileNames,omitempty"`
	Signature         string   `json:"signature,omitempty"`
	Wid               int64    `json:"wid,omitempty"`
}

// OnEcLoginAgreeProtocolEvent
// @id 1910
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=登录设置-同意协议消息
func (s *Service) OnEcLoginAgreeProtocolEvent(handler msg.XinyunEventHandler[EcLoginAgreeProtocolEvent]) (service *Service) {
	var listener = &EcLoginAgreeProtocolListener{handler}
	s.Register(msg.WrapListener[EcLoginAgreeProtocolEvent](listener), "ec_login", "agreeProtocol")
	return s
}
