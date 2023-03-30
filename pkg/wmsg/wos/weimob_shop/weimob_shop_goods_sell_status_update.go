package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopGoodsSellStatusUpdateListener struct {
	handler msg.WosEventHandler[WeimobShopGoodsSellStatusUpdateEvent]
}

func (s WeimobShopGoodsSellStatusUpdateListener) NewMessage() msg.MsgRequest[WeimobShopGoodsSellStatusUpdateEvent] {
	return &msg.WosOpenMessage[WeimobShopGoodsSellStatusUpdateEvent]{
		MsgBody: &WeimobShopGoodsSellStatusUpdateEvent{},
	}
}

func (s WeimobShopGoodsSellStatusUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopGoodsSellStatusUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopGoodsSellStatusUpdateEvent]))
}

type WeimobShopGoodsSellStatusUpdateEvent struct {
	GoodsIdList []WeimobShopGoodsSellStatusUpdateGoodsIdList `json:"goodsIdList,omitempty"`
	IsCanShell  bool                                         `json:"isCanShell,omitempty"`
}

type WeimobShopGoodsSellStatusUpdateGoodsIdList struct {
}

// OnWeimobShopGoodsSellStatusUpdateEvent
// @id 1262
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1262?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=修改商品可售状态)
func (s *Service) OnWeimobShopGoodsSellStatusUpdateEvent(handler msg.WosEventHandler[WeimobShopGoodsSellStatusUpdateEvent]) (service *Service) {
	var listener = &WeimobShopGoodsSellStatusUpdateListener{handler}
	s.Register(msg.WrapListener[WeimobShopGoodsSellStatusUpdateEvent](listener), "weimob_shop.goods", "sellStatusUpdate")
	return s
}
