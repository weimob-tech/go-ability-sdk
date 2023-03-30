package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderRefuseBookGoodsOrder
// @id 1434
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=拒绝预约商品订单)
func (client *Client) OrderRefuseBookGoodsOrder(request *OrderRefuseBookGoodsOrderRequest) (response *OrderRefuseBookGoodsOrderResponse, err error) {
	rpcResponse := CreateOrderRefuseBookGoodsOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderRefuseBookGoodsOrderRequest struct {
	*api.BaseRequest

	NickName    string `json:"nickName,omitempty"`
	OrderNo     string `json:"orderNo,omitempty"`
	Description string `json:"description,omitempty"`
}

type OrderRefuseBookGoodsOrderResponse struct {
}

func CreateOrderRefuseBookGoodsOrderRequest() (request *OrderRefuseBookGoodsOrderRequest) {
	request = &OrderRefuseBookGoodsOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "order/refuseBookGoodsOrder", "api")
	request.Method = api.POST
	return
}

func CreateOrderRefuseBookGoodsOrderResponse() (response *api.BaseResponse[OrderRefuseBookGoodsOrderResponse]) {
	response = api.CreateResponse[OrderRefuseBookGoodsOrderResponse](&OrderRefuseBookGoodsOrderResponse{})
	return
}
