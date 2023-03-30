package weimob_shopexpress

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopexpressRightsCreateListener struct {
	handler msg.WosEventHandler[WeimobShopexpressRightsCreateEvent]
}

func (s WeimobShopexpressRightsCreateListener) NewMessage() msg.MsgRequest[WeimobShopexpressRightsCreateEvent] {
	return &msg.WosOpenMessage[WeimobShopexpressRightsCreateEvent]{
		MsgBody: &WeimobShopexpressRightsCreateEvent{},
	}
}

func (s WeimobShopexpressRightsCreateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopexpressRightsCreateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopexpressRightsCreateEvent]))
}

type WeimobShopexpressRightsCreateEvent struct {
	OrderNo     string `json:"orderNo,omitempty"`
	AftersaleNo string `json:"aftersaleNo,omitempty"`
}

// OnWeimobShopexpressRightsCreateEvent
// @id 1338
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1338?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商家新增售后)
func (s *Service) OnWeimobShopexpressRightsCreateEvent(handler msg.WosEventHandler[WeimobShopexpressRightsCreateEvent]) (service *Service) {
	var listener = &WeimobShopexpressRightsCreateListener{handler}
	s.Register(msg.WrapListener[WeimobShopexpressRightsCreateEvent](listener), "weimob_shopexpress.rights", "create")
	return s
}
