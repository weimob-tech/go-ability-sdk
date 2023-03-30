package wangpu

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WPGoodsPriceChangeListener struct {
	handler msg.XinyunEventHandler[WPGoodsPriceChangeEvent]
}

func (s WPGoodsPriceChangeListener) NewMessage() msg.MsgRequest[WPGoodsPriceChangeEvent] {
	return &msg.XinyunOpenMessage[WPGoodsPriceChangeEvent]{
		MsgBody: &WPGoodsPriceChangeEvent{},
	}
}

func (s WPGoodsPriceChangeListener) OnMessage(ctx context.Context, message msg.MsgRequest[WPGoodsPriceChangeEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[WPGoodsPriceChangeEvent]))
}

type WPGoodsPriceChangeEvent struct {
	GoodsId    string `json:"goods_id,omitempty"`
	UpdateTime int64  `json:"update_time,omitempty"`
}

// OnWPGoodsPriceChangeEvent
// @id 268
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=商品信息变更
func (s *Service) OnWPGoodsPriceChangeEvent(handler msg.XinyunEventHandler[WPGoodsPriceChangeEvent]) (service *Service) {
	var listener = &WPGoodsPriceChangeListener{handler}
	s.Register(msg.WrapListener[WPGoodsPriceChangeEvent](listener), "WP_goods", "priceChange")
	return s
}
