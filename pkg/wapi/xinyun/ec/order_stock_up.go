package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderStockUp
// @id 1120
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=自提订单备货)
func (client *Client) OrderStockUp(request *OrderStockUpRequest) (response *OrderStockUpResponse, err error) {
	rpcResponse := CreateOrderStockUpResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderStockUpRequest struct {
	*api.BaseRequest

	StoreId   int64 `json:"storeId,omitempty"`
	SiteId    int64 `json:"siteId,omitempty"`
	OrderNo   int64 `json:"orderNo,omitempty"`
	SceneType int   `json:"sceneType,omitempty"`
}

type OrderStockUpResponse struct {
}

func CreateOrderStockUpRequest() (request *OrderStockUpRequest) {
	request = &OrderStockUpRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "order/stockUp", "api")
	request.Method = api.POST
	return
}

func CreateOrderStockUpResponse() (response *api.BaseResponse[OrderStockUpResponse]) {
	response = api.CreateResponse[OrderStockUpResponse](&OrderStockUpResponse{})
	return
}
