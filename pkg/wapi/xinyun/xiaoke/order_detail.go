package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderDetail
// @id 2675
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取订单详情)
func (client *Client) OrderDetail(request *OrderDetailRequest) (response *OrderDetailResponse, err error) {
	rpcResponse := CreateOrderDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderDetailRequest struct {
	*api.BaseRequest

	UserWid     int64  `json:"userWid,omitempty"`
	OrderNumber string `json:"orderNumber,omitempty"`
}

type OrderDetailResponse struct {
}

func CreateOrderDetailRequest() (request *OrderDetailRequest) {
	request = &OrderDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "order/detail", "api")
	request.Method = api.POST
	return
}

func CreateOrderDetailResponse() (response *api.BaseResponse[OrderDetailResponse]) {
	response = api.CreateResponse[OrderDetailResponse](&OrderDetailResponse{})
	return
}
