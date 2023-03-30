package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderClose
// @id 31
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=关闭订单)
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

	OrderNo   string `json:"order_no,omitempty"`
	ShopId    string `json:"shop_id,omitempty"`
	Reason    string `json:"reason,omitempty"`
	UpdateMan string `json:"update_man,omitempty"`
}

type OrderCloseResponse struct {
}

func CreateOrderCloseRequest() (request *OrderCloseRequest) {
	request = &OrderCloseRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "Order/Close", "api")
	request.Method = api.POST
	return
}

func CreateOrderCloseResponse() (response *api.BaseResponse[OrderCloseResponse]) {
	response = api.CreateResponse[OrderCloseResponse](&OrderCloseResponse{})
	return
}
