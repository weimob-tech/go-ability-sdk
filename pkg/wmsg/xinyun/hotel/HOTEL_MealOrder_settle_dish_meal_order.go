package hotel

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type HOTELMealOrderSettleDishMealOrderListener struct {
	handler msg.XinyunEventHandler[HOTELMealOrderSettleDishMealOrderEvent]
}

func (s HOTELMealOrderSettleDishMealOrderListener) NewMessage() msg.MsgRequest[HOTELMealOrderSettleDishMealOrderEvent] {
	return &msg.XinyunOpenMessage[HOTELMealOrderSettleDishMealOrderEvent]{
		MsgBody: &HOTELMealOrderSettleDishMealOrderEvent{},
	}
}

func (s HOTELMealOrderSettleDishMealOrderListener) OnMessage(ctx context.Context, message msg.MsgRequest[HOTELMealOrderSettleDishMealOrderEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[HOTELMealOrderSettleDishMealOrderEvent]))
}

type HOTELMealOrderSettleDishMealOrderEvent struct {
	OrderNo       string `json:"orderNo,omitempty"`
	TableNumber   string `json:"tableNumber,omitempty"`
	MealPeopleNum string `json:"mealPeopleNum,omitempty"`
	MealRemark    string `json:"mealRemark,omitempty"`
	StoreId       string `json:"storeId,omitempty"`
}

// OnHOTELMealOrderSettleDishMealOrderEvent
// @id 1706
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=堂食订单结算消息
func (s *Service) OnHOTELMealOrderSettleDishMealOrderEvent(handler msg.XinyunEventHandler[HOTELMealOrderSettleDishMealOrderEvent]) (service *Service) {
	var listener = &HOTELMealOrderSettleDishMealOrderListener{handler}
	s.Register(msg.WrapListener[HOTELMealOrderSettleDishMealOrderEvent](listener), "HOTEL_MealOrder", "settleDishMealOrder")
	return s
}
