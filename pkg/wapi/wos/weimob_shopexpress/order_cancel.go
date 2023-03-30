package weimob_shopexpress

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderCancel
// @id 1964
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1964?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=取消订单)
func (client *Client) OrderCancel(request *OrderCancelRequest) (response *OrderCancelResponse, err error) {
	rpcResponse := CreateOrderCancelResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderCancelRequest struct {
	*api.BaseRequest

	RightsId string `json:"rightsId,omitempty"`
}

type OrderCancelResponse struct {
}

func CreateOrderCancelRequest() (request *OrderCancelRequest) {
	request = &OrderCancelRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shopexpress", "v2.0", "order/cancel", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderCancelResponse() (response *api.BaseResponse[OrderCancelResponse]) {
	response = api.CreateResponse[OrderCancelResponse](&OrderCancelResponse{})
	return
}
