package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderConfirmBookGoodsOrder
// @id 1433
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=确认预约商品订单)
func (client *Client) OrderConfirmBookGoodsOrder(request *OrderConfirmBookGoodsOrderRequest) (response *OrderConfirmBookGoodsOrderResponse, err error) {
	rpcResponse := CreateOrderConfirmBookGoodsOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderConfirmBookGoodsOrderRequest struct {
	*api.BaseRequest

	NickName    string `json:"nickName,omitempty"`
	OrderNo     string `json:"orderNo,omitempty"`
	Description string `json:"description,omitempty"`
}

type OrderConfirmBookGoodsOrderResponse struct {
}

func CreateOrderConfirmBookGoodsOrderRequest() (request *OrderConfirmBookGoodsOrderRequest) {
	request = &OrderConfirmBookGoodsOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "order/confirmBookGoodsOrder", "api")
	request.Method = api.POST
	return
}

func CreateOrderConfirmBookGoodsOrderResponse() (response *api.BaseResponse[OrderConfirmBookGoodsOrderResponse]) {
	response = api.CreateResponse[OrderConfirmBookGoodsOrderResponse](&OrderConfirmBookGoodsOrderResponse{})
	return
}
