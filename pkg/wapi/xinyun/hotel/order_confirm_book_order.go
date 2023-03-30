package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderConfirmBookOrder
// @id 512
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=确认订房订单)
func (client *Client) OrderConfirmBookOrder(request *OrderConfirmBookOrderRequest) (response *OrderConfirmBookOrderResponse, err error) {
	rpcResponse := CreateOrderConfirmBookOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderConfirmBookOrderRequest struct {
	*api.BaseRequest

	RoomNumberMap []map[string]any `json:"roomNumberMap,omitempty"`
	StoreId       int64            `json:"storeId,omitempty"`
	OrderNo       string           `json:"orderNo,omitempty"`
	RoomId        int64            `json:"roomId,omitempty"`
	RoomNum       string           `json:"roomNum,omitempty"`
}

type OrderConfirmBookOrderResponse struct {
}

func CreateOrderConfirmBookOrderRequest() (request *OrderConfirmBookOrderRequest) {
	request = &OrderConfirmBookOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "order/confirmBookOrder", "api")
	request.Method = api.POST
	return
}

func CreateOrderConfirmBookOrderResponse() (response *api.BaseResponse[OrderConfirmBookOrderResponse]) {
	response = api.CreateResponse[OrderConfirmBookOrderResponse](&OrderConfirmBookOrderResponse{})
	return
}
