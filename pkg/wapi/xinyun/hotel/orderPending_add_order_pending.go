package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderPendingAddOrderPending
// @id 1292
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增待付订单)
func (client *Client) OrderPendingAddOrderPending(request *OrderPendingAddOrderPendingRequest) (response *OrderPendingAddOrderPendingResponse, err error) {
	rpcResponse := CreateOrderPendingAddOrderPendingResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderPendingAddOrderPendingRequest struct {
	*api.BaseRequest

	PendingItems []OrderPendingAddOrderPendingRequestPendingItems `json:"pendingItems,omitempty"`
	OrderNo      string                                           `json:"orderNo,omitempty"`
	HandNo       string                                           `json:"handNo,omitempty"`
	GoodsName    string                                           `json:"goodsName,omitempty"`
	Nickname     string                                           `json:"nickname,omitempty"`
	OrderAmount  int                                              `json:"orderAmount,omitempty"`
	Remark       string                                           `json:"remark,omitempty"`
	StoreId      int64                                            `json:"storeId,omitempty"`
}

type OrderPendingAddOrderPendingResponse struct {
}

func CreateOrderPendingAddOrderPendingRequest() (request *OrderPendingAddOrderPendingRequest) {
	request = &OrderPendingAddOrderPendingRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "orderPending/addOrderPending", "api")
	request.Method = api.POST
	return
}

func CreateOrderPendingAddOrderPendingResponse() (response *api.BaseResponse[OrderPendingAddOrderPendingResponse]) {
	response = api.CreateResponse[OrderPendingAddOrderPendingResponse](&OrderPendingAddOrderPendingResponse{})
	return
}

type OrderPendingAddOrderPendingRequestPendingItems struct {
	GoodsNo    int     `json:"goodsNo,omitempty"`
	GoodsName  string  `json:"goodsName,omitempty"`
	GoodsPrice float64 `json:"goodsPrice,omitempty"`
	GoodsTotal int     `json:"goodsTotal,omitempty"`
}
