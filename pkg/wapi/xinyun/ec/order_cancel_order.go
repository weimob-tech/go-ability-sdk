package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderCancelOrder
// @id 338
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=商家关闭订单)
func (client *Client) OrderCancelOrder(request *OrderCancelOrderRequest) (response *OrderCancelOrderResponse, err error) {
	rpcResponse := CreateOrderCancelOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderCancelOrderRequest struct {
	*api.BaseRequest

	SpecificCancelReason string `json:"specificCancelReason,omitempty"`
	OrderNo              int64  `json:"orderNo,omitempty"`
	CancelReasonId       int    `json:"cancelReasonId,omitempty"`
	CancelReason         string `json:"cancelReason,omitempty"`
}

type OrderCancelOrderResponse struct {
}

func CreateOrderCancelOrderRequest() (request *OrderCancelOrderRequest) {
	request = &OrderCancelOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "order/cancelOrder", "api")
	request.Method = api.POST
	return
}

func CreateOrderCancelOrderResponse() (response *api.BaseResponse[OrderCancelOrderResponse]) {
	response = api.CreateResponse[OrderCancelOrderResponse](&OrderCancelOrderResponse{})
	return
}
