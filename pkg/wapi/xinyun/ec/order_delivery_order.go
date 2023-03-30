package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderDeliveryOrder
// @id 569
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=订单发货（单个）)
func (client *Client) OrderDeliveryOrder(request *OrderDeliveryOrderRequest) (response *OrderDeliveryOrderResponse, err error) {
	rpcResponse := CreateOrderDeliveryOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderDeliveryOrderRequest struct {
	*api.BaseRequest

	DeliveryOrderItemList []OrderDeliveryOrderRequestDeliveryOrderItemList `json:"deliveryOrderItemList,omitempty"`
	OrderNo               int64                                            `json:"orderNo,omitempty"`
	DeliveryNo            string                                           `json:"deliveryNo,omitempty"`
	DeliveryCompanyCode   string                                           `json:"deliveryCompanyCode,omitempty"`
	DeliveryCompanyName   string                                           `json:"deliveryCompanyName,omitempty"`
	IsNeedLogistics       bool                                             `json:"isNeedLogistics,omitempty"`
	IsSplitPackage        bool                                             `json:"isSplitPackage,omitempty"`
	DeliveryRemark        string                                           `json:"deliveryRemark,omitempty"`
	DeliveryOrderId       int64                                            `json:"deliveryOrderId,omitempty"`
}

type OrderDeliveryOrderResponse struct {
}

func CreateOrderDeliveryOrderRequest() (request *OrderDeliveryOrderRequest) {
	request = &OrderDeliveryOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "order/deliveryOrder", "api")
	request.Method = api.POST
	return
}

func CreateOrderDeliveryOrderResponse() (response *api.BaseResponse[OrderDeliveryOrderResponse]) {
	response = api.CreateResponse[OrderDeliveryOrderResponse](&OrderDeliveryOrderResponse{})
	return
}

type OrderDeliveryOrderRequestDeliveryOrderItemList struct {
	ItemId         int64 `json:"itemId,omitempty"`
	DeliverySkuNum int64 `json:"deliverySkuNum,omitempty"`
}
