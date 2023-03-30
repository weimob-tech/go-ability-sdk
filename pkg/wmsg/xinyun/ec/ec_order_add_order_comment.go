package ec

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type EcOrderAddOrderCommentListener struct {
	handler msg.XinyunEventHandler[EcOrderAddOrderCommentEvent]
}

func (s EcOrderAddOrderCommentListener) NewMessage() msg.MsgRequest[EcOrderAddOrderCommentEvent] {
	return &msg.XinyunOpenMessage[EcOrderAddOrderCommentEvent]{
		MsgBody: &EcOrderAddOrderCommentEvent{},
	}
}

func (s EcOrderAddOrderCommentListener) OnMessage(ctx context.Context, message msg.MsgRequest[EcOrderAddOrderCommentEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[EcOrderAddOrderCommentEvent]))
}

type EcOrderAddOrderCommentEvent struct {
	GoodsComments         []EcOrderAddOrderCommentGoodsComments `json:"goodsComments,omitempty"`
	OrderNo               int64                                 `json:"orderNo,omitempty"`
	ServiceAttitudeScore  int64                                 `json:"serviceAttitudeScore,omitempty"`
	LogisticsServiceScore int64                                 `json:"logisticsServiceScore,omitempty"`
}

type EcOrderAddOrderCommentGoodsComments struct {
	UserImageUrl         string `json:"userImageUrl,omitempty"`
	UserNickname         string `json:"userNickname,omitempty"`
	GoodsTitle           string `json:"goodsTitle,omitempty"`
	GoodsIndexImage      string `json:"goodsIndexImage,omitempty"`
	CommentScore         int64  `json:"commentScore,omitempty"`
	CommentLevel         int64  `json:"commentLevel,omitempty"`
	CommentContent       string `json:"commentContent,omitempty"`
	CommentImageUrl      string `json:"commentImageUrl,omitempty"`
	CommentVideoUrl      string `json:"commentVideoUrl,omitempty"`
	CommentVideoImageUrl string `json:"commentVideoImageUrl,omitempty"`
	IsAutoComment        string `json:"isAutoComment,omitempty"`
}

// OnEcOrderAddOrderCommentEvent
// @id 1904
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=生成订单评价
func (s *Service) OnEcOrderAddOrderCommentEvent(handler msg.XinyunEventHandler[EcOrderAddOrderCommentEvent]) (service *Service) {
	var listener = &EcOrderAddOrderCommentListener{handler}
	s.Register(msg.WrapListener[EcOrderAddOrderCommentEvent](listener), "ec_order", "addOrderComment")
	return s
}
