package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderConfirm
// @id 1003
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=确认线路订单)
func (client *Client) OrderConfirm(request *OrderConfirmRequest) (response *OrderConfirmResponse, err error) {
	rpcResponse := CreateOrderConfirmResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderConfirmRequest struct {
	*api.BaseRequest

	OrderNo string `json:"orderNo,omitempty"`
}

type OrderConfirmResponse struct {
}

func CreateOrderConfirmRequest() (request *OrderConfirmRequest) {
	request = &OrderConfirmRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "order/confirm", "api")
	request.Method = api.POST
	return
}

func CreateOrderConfirmResponse() (response *api.BaseResponse[OrderConfirmResponse]) {
	response = api.CreateResponse[OrderConfirmResponse](&OrderConfirmResponse{})
	return
}
