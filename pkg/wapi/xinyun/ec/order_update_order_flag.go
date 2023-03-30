package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderUpdateOrderFlag
// @id 341
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新订单标记)
func (client *Client) OrderUpdateOrderFlag(request *OrderUpdateOrderFlagRequest) (response *OrderUpdateOrderFlagResponse, err error) {
	rpcResponse := CreateOrderUpdateOrderFlagResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderUpdateOrderFlagRequest struct {
	*api.BaseRequest

	FlagRank    int     `json:"flagRank,omitempty"`
	FlagContent string  `json:"flagContent,omitempty"`
	OrderNoList []int64 `json:"orderNoList,omitempty"`
}

type OrderUpdateOrderFlagResponse struct {
}

func CreateOrderUpdateOrderFlagRequest() (request *OrderUpdateOrderFlagRequest) {
	request = &OrderUpdateOrderFlagRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "order/updateOrderFlag", "api")
	request.Method = api.POST
	return
}

func CreateOrderUpdateOrderFlagResponse() (response *api.BaseResponse[OrderUpdateOrderFlagResponse]) {
	response = api.CreateResponse[OrderUpdateOrderFlagResponse](&OrderUpdateOrderFlagResponse{})
	return
}
