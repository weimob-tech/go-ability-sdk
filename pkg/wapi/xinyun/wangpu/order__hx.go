package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderHx
// @id 40
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=订单核销)
func (client *Client) OrderHx(request *OrderHxRequest) (response *OrderHxResponse, err error) {
	rpcResponse := CreateOrderHxResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderHxRequest struct {
	*api.BaseRequest

	OrderNo      string `json:"order_no,omitempty"`
	WeimobOpenId string `json:"weimob_openId,omitempty"`
	HxWay        string `json:"hx_way,omitempty"`
}

type OrderHxResponse struct {
}

func CreateOrderHxRequest() (request *OrderHxRequest) {
	request = &OrderHxRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "order/Hx", "api")
	request.Method = api.POST
	return
}

func CreateOrderHxResponse() (response *api.BaseResponse[OrderHxResponse]) {
	response = api.CreateResponse[OrderHxResponse](&OrderHxResponse{})
	return
}
