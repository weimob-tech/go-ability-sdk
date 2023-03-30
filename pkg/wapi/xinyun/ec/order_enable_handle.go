package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderEnableHandle
// @id 1483
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=订单可发货变更)
func (client *Client) OrderEnableHandle(request *OrderEnableHandleRequest) (response *OrderEnableHandleResponse, err error) {
	rpcResponse := CreateOrderEnableHandleResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderEnableHandleRequest struct {
	*api.BaseRequest
}

type OrderEnableHandleResponse struct {
}

func CreateOrderEnableHandleRequest() (request *OrderEnableHandleRequest) {
	request = &OrderEnableHandleRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "order/enableHandle", "api")
	request.Method = api.POST
	return
}

func CreateOrderEnableHandleResponse() (response *api.BaseResponse[OrderEnableHandleResponse]) {
	response = api.CreateResponse[OrderEnableHandleResponse](&OrderEnableHandleResponse{})
	return
}
