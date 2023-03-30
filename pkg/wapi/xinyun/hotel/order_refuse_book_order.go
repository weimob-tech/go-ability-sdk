package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderRefuseBookOrder
// @id 513
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=拒绝订房订单)
func (client *Client) OrderRefuseBookOrder(request *OrderRefuseBookOrderRequest) (response *OrderRefuseBookOrderResponse, err error) {
	rpcResponse := CreateOrderRefuseBookOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderRefuseBookOrderRequest struct {
	*api.BaseRequest

	StoreId     int64  `json:"storeId,omitempty"`
	OrderNo     string `json:"orderNo,omitempty"`
	NickName    string `json:"nickName,omitempty"`
	Description string `json:"description,omitempty"`
}

type OrderRefuseBookOrderResponse struct {
}

func CreateOrderRefuseBookOrderRequest() (request *OrderRefuseBookOrderRequest) {
	request = &OrderRefuseBookOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "order/refuseBookOrder", "api")
	request.Method = api.POST
	return
}

func CreateOrderRefuseBookOrderResponse() (response *api.BaseResponse[OrderRefuseBookOrderResponse]) {
	response = api.CreateResponse[OrderRefuseBookOrderResponse](&OrderRefuseBookOrderResponse{})
	return
}
