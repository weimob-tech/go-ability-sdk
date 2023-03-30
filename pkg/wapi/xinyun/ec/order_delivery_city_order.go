package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderDeliveryCityOrder
// @id 1129
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=同城配送发单)
func (client *Client) OrderDeliveryCityOrder(request *OrderDeliveryCityOrderRequest) (response *OrderDeliveryCityOrderResponse, err error) {
	rpcResponse := CreateOrderDeliveryCityOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderDeliveryCityOrderRequest struct {
	*api.BaseRequest

	StoreId         int64  `json:"storeId,omitempty"`
	OrderNo         int64  `json:"orderNo,omitempty"`
	SelfDelivery    bool   `json:"selfDelivery,omitempty"`
	Remark          string `json:"remark,omitempty"`
	DeliveryOrderId int64  `json:"deliveryOrderId,omitempty"`
}

type OrderDeliveryCityOrderResponse struct {
}

func CreateOrderDeliveryCityOrderRequest() (request *OrderDeliveryCityOrderRequest) {
	request = &OrderDeliveryCityOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "order/deliveryCityOrder", "api")
	request.Method = api.POST
	return
}

func CreateOrderDeliveryCityOrderResponse() (response *api.BaseResponse[OrderDeliveryCityOrderResponse]) {
	response = api.CreateResponse[OrderDeliveryCityOrderResponse](&OrderDeliveryCityOrderResponse{})
	return
}
