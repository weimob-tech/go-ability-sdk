package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderUpdateOrderAmount
// @id 2805
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=订单改价)
func (client *Client) OrderUpdateOrderAmount(request *OrderUpdateOrderAmountRequest) (response *OrderUpdateOrderAmountResponse, err error) {
	rpcResponse := CreateOrderUpdateOrderAmountResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderUpdateOrderAmountRequest struct {
	*api.BaseRequest

	UpdatedItemList       []OrderUpdateOrderAmountRequestUpdatedItemList `json:"updatedItemList,omitempty"`
	OrderNo               int64                                          `json:"orderNo,omitempty"`
	UpdatedDeliveryAmount float64                                        `json:"updatedDeliveryAmount,omitempty"`
}

type OrderUpdateOrderAmountResponse struct {
}

func CreateOrderUpdateOrderAmountRequest() (request *OrderUpdateOrderAmountRequest) {
	request = &OrderUpdateOrderAmountRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "order/updateOrderAmount", "api")
	request.Method = api.POST
	return
}

func CreateOrderUpdateOrderAmountResponse() (response *api.BaseResponse[OrderUpdateOrderAmountResponse]) {
	response = api.CreateResponse[OrderUpdateOrderAmountResponse](&OrderUpdateOrderAmountResponse{})
	return
}

type OrderUpdateOrderAmountRequestUpdatedItemList struct {
	ItemId        int64   `json:"itemId,omitempty"`
	UpdatedAmount float64 `json:"updatedAmount,omitempty"`
}
