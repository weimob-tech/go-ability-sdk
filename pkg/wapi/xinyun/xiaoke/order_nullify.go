package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderNullify
// @id 2664
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=作废订单)
func (client *Client) OrderNullify(request *OrderNullifyRequest) (response *OrderNullifyResponse, err error) {
	rpcResponse := CreateOrderNullifyResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderNullifyRequest struct {
	*api.BaseRequest

	Keys               []string `json:"keys,omitempty"`
	UserWid            int64    `json:"userWid,omitempty"`
	CauseOfAbandonment string   `json:"causeOfAbandonment,omitempty"`
	Key                string   `json:"key,omitempty"`
}

type OrderNullifyResponse struct {
}

func CreateOrderNullifyRequest() (request *OrderNullifyRequest) {
	request = &OrderNullifyRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "order/nullify", "api")
	request.Method = api.POST
	return
}

func CreateOrderNullifyResponse() (response *api.BaseResponse[OrderNullifyResponse]) {
	response = api.CreateResponse[OrderNullifyResponse](&OrderNullifyResponse{})
	return
}
