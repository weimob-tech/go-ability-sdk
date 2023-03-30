package hotel

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type HOTELMealOrderSubmitAddDishListener struct {
	handler msg.XinyunEventHandler[HOTELMealOrderSubmitAddDishEvent]
}

func (s HOTELMealOrderSubmitAddDishListener) NewMessage() msg.MsgRequest[HOTELMealOrderSubmitAddDishEvent] {
	return &msg.XinyunOpenMessage[HOTELMealOrderSubmitAddDishEvent]{
		MsgBody: &HOTELMealOrderSubmitAddDishEvent{},
	}
}

func (s HOTELMealOrderSubmitAddDishListener) OnMessage(ctx context.Context, message msg.MsgRequest[HOTELMealOrderSubmitAddDishEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[HOTELMealOrderSubmitAddDishEvent]))
}

type HOTELMealOrderSubmitAddDishEvent struct {
	OrderNo       string `json:"orderNo,omitempty"`
	TableNumber   string `json:"tableNumber,omitempty"`
	MealPeopleNum string `json:"mealPeopleNum,omitempty"`
	MealRemark    string `json:"mealRemark,omitempty"`
	StoreId       string `json:"storeId,omitempty"`
}

// OnHOTELMealOrderSubmitAddDishEvent
// @id 1442
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=堂食订单加餐提交成功消息
func (s *Service) OnHOTELMealOrderSubmitAddDishEvent(handler msg.XinyunEventHandler[HOTELMealOrderSubmitAddDishEvent]) (service *Service) {
	var listener = &HOTELMealOrderSubmitAddDishListener{handler}
	s.Register(msg.WrapListener[HOTELMealOrderSubmitAddDishEvent](listener), "HOTEL_MealOrder", "submitAddDish")
	return s
}
