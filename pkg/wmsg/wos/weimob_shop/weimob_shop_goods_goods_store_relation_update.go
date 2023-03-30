package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobShopGoodsGoodsStoreRelationUpdateListener struct {
	handler msg.WosEventHandler[WeimobShopGoodsGoodsStoreRelationUpdateEvent]
}

func (s WeimobShopGoodsGoodsStoreRelationUpdateListener) NewMessage() msg.MsgRequest[WeimobShopGoodsGoodsStoreRelationUpdateEvent] {
	return &msg.WosOpenMessage[WeimobShopGoodsGoodsStoreRelationUpdateEvent]{
		MsgBody: &WeimobShopGoodsGoodsStoreRelationUpdateEvent{},
	}
}

func (s WeimobShopGoodsGoodsStoreRelationUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobShopGoodsGoodsStoreRelationUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobShopGoodsGoodsStoreRelationUpdateEvent]))
}

type WeimobShopGoodsGoodsStoreRelationUpdateEvent struct {
	GoodsIdList []WeimobShopGoodsGoodsStoreRelationUpdateGoodsIdList `json:"goodsIdList,omitempty"`
	VidList     []WeimobShopGoodsGoodsStoreRelationUpdateVidList     `json:"vidList,omitempty"`
	IsAssigned  bool                                                 `json:"isAssigned,omitempty"`
	IsOnline    bool                                                 `json:"isOnline,omitempty"`
}

type WeimobShopGoodsGoodsStoreRelationUpdateGoodsIdList struct {
}
type WeimobShopGoodsGoodsStoreRelationUpdateVidList struct {
}

// OnWeimobShopGoodsGoodsStoreRelationUpdateEvent
// @id 1231
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1231?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=门店商品关系变更)
func (s *Service) OnWeimobShopGoodsGoodsStoreRelationUpdateEvent(handler msg.WosEventHandler[WeimobShopGoodsGoodsStoreRelationUpdateEvent]) (service *Service) {
	var listener = &WeimobShopGoodsGoodsStoreRelationUpdateListener{handler}
	s.Register(msg.WrapListener[WeimobShopGoodsGoodsStoreRelationUpdateEvent](listener), "weimob_shop.goods", "goodsStoreRelationUpdate")
	return s
}
