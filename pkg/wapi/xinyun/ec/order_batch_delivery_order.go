package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderBatchDeliveryOrder
// @id 342
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=订单发货(批量))
func (client *Client) OrderBatchDeliveryOrder(request *OrderBatchDeliveryOrderRequest) (response *OrderBatchDeliveryOrderResponse, err error) {
	rpcResponse := CreateOrderBatchDeliveryOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderBatchDeliveryOrderRequest struct {
	*api.BaseRequest

	DeliveryOrderList   []map[string]any `json:"deliveryOrderList,omitempty"`
	OrderNo             int64            `json:"orderNo,omitempty"`
	DeliveryNo          string           `json:"deliveryNo,omitempty"`
	DeliveryCompanyCode string           `json:"deliveryCompanyCode,omitempty"`
	DeliveryCompanyName string           `json:"deliveryCompanyName,omitempty"`
	IsNeedLogistics     bool             `json:"isNeedLogistics,omitempty"`
}

type OrderBatchDeliveryOrderResponse struct {
}

func CreateOrderBatchDeliveryOrderRequest() (request *OrderBatchDeliveryOrderRequest) {
	request = &OrderBatchDeliveryOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "order/batchDeliveryOrder", "api")
	request.Method = api.POST
	return
}

func CreateOrderBatchDeliveryOrderResponse() (response *api.BaseResponse[OrderBatchDeliveryOrderResponse]) {
	response = api.CreateResponse[OrderBatchDeliveryOrderResponse](&OrderBatchDeliveryOrderResponse{})
	return
}
