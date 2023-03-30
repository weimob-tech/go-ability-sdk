package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderOptOrder
// @id 71
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=操作订单)
func (client *Client) OrderOptOrder(request *OrderOptOrderRequest) (response *OrderOptOrderResponse, err error) {
	rpcResponse := CreateOrderOptOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderOptOrderRequest struct {
	*api.BaseRequest

	StoreId   int64  `json:"storeId,omitempty"`
	OrderId   int64  `json:"orderId,omitempty"`
	Operation int    `json:"operation,omitempty"`
	PayMethod int    `json:"payMethod,omitempty"`
	Remarks   string `json:"remarks,omitempty"`
	OrderNo   string `json:"orderNo,omitempty"`
}

type OrderOptOrderResponse struct {
}

func CreateOrderOptOrderRequest() (request *OrderOptOrderRequest) {
	request = &OrderOptOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "order/optOrder", "api")
	request.Method = api.POST
	return
}

func CreateOrderOptOrderResponse() (response *api.BaseResponse[OrderOptOrderResponse]) {
	response = api.CreateResponse[OrderOptOrderResponse](&OrderOptOrderResponse{})
	return
}
