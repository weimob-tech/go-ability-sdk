package weimob_shopexpress

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopexpressRightsCompleteListener struct {
	handler msg.WosEventHandler[WeimobShopexpressRightsCompleteEvent]
}

func (s WeimobShopexpressRightsCompleteListener) NewMessage() msg.MsgRequest[WeimobShopexpressRightsCompleteEvent] {
	return &msg.WosOpenMessage[WeimobShopexpressRightsCompleteEvent]{
		MsgBody: &WeimobShopexpressRightsCompleteEvent{},
	}
}

func (s WeimobShopexpressRightsCompleteListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopexpressRightsCompleteEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopexpressRightsCompleteEvent]))
}

type WeimobShopexpressRightsCompleteEvent struct {
	OrderNo     string `json:"orderNo,omitempty"`
	AftersaleNo string `json:"aftersaleNo,omitempty"`
}

// OnWeimobShopexpressRightsCompleteEvent
// @id 1337
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1337?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商家完成售后)
func (s *Service) OnWeimobShopexpressRightsCompleteEvent(handler msg.WosEventHandler[WeimobShopexpressRightsCompleteEvent]) (service *Service) {
	var listener = &WeimobShopexpressRightsCompleteListener{handler}
	s.Register(msg.WrapListener[WeimobShopexpressRightsCompleteEvent](listener), "weimob_shopexpress.rights", "complete")
	return s
}
