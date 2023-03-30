package hotel

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type HOTELMealOrderSubmitMealOrderListener struct {
	handler msg.XinyunEventHandler[HOTELMealOrderSubmitMealOrderEvent]
}

func (s HOTELMealOrderSubmitMealOrderListener) NewMessage() msg.MsgRequest[HOTELMealOrderSubmitMealOrderEvent] {
	return &msg.XinyunOpenMessage[HOTELMealOrderSubmitMealOrderEvent]{
		MsgBody: &HOTELMealOrderSubmitMealOrderEvent{},
	}
}

func (s HOTELMealOrderSubmitMealOrderListener) OnMessage(ctx context.Context, message msg.MsgRequest[HOTELMealOrderSubmitMealOrderEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[HOTELMealOrderSubmitMealOrderEvent]))
}

type HOTELMealOrderSubmitMealOrderEvent struct {
	OrderNo       string `json:"orderNo,omitempty"`
	TableNumber   string `json:"tableNumber,omitempty"`
	MealPeopleNum string `json:"mealPeopleNum,omitempty"`
	MealRemark    string `json:"mealRemark,omitempty"`
	StoreId       string `json:"storeId,omitempty"`
}

// OnHOTELMealOrderSubmitMealOrderEvent
// @id 1441
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=堂食订单提交成功消息
func (s *Service) OnHOTELMealOrderSubmitMealOrderEvent(handler msg.XinyunEventHandler[HOTELMealOrderSubmitMealOrderEvent]) (service *Service) {
	var listener = &HOTELMealOrderSubmitMealOrderListener{handler}
	s.Register(msg.WrapListener[HOTELMealOrderSubmitMealOrderEvent](listener), "HOTEL_MealOrder", "submitMealOrder")
	return s
}
