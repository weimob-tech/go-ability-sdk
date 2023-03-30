package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderCancelOrder
// @id 1250
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1250?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=产品融合取消订单)
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

	Obj string `json:"obj,omitempty"`
}

type OrderCancelOrderResponse struct {
	Obj string `json:"obj,omitempty"`
}

func CreateOrderCancelOrderRequest() (request *OrderCancelOrderRequest) {
	request = &OrderCancelOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "order/cancelOrder", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderCancelOrderResponse() (response *api.BaseResponse[OrderCancelOrderResponse]) {
	response = api.CreateResponse[OrderCancelOrderResponse](&OrderCancelOrderResponse{})
	return
}
