package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderClose
// @id 1004
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=关闭线路订单)
func (client *Client) OrderClose(request *OrderCloseRequest) (response *OrderCloseResponse, err error) {
	rpcResponse := CreateOrderCloseResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderCloseRequest struct {
	*api.BaseRequest

	OrderNo string `json:"orderNo,omitempty"`
}

type OrderCloseResponse struct {
}

func CreateOrderCloseRequest() (request *OrderCloseRequest) {
	request = &OrderCloseRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "order/close", "api")
	request.Method = api.POST
	return
}

func CreateOrderCloseResponse() (response *api.BaseResponse[OrderCloseResponse]) {
	response = api.CreateResponse[OrderCloseResponse](&OrderCloseResponse{})
	return
}
