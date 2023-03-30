package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderCancel
// @id 1815
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1815?tab=1)
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

	CancelReason         string `json:"cancelReason,omitempty"`
	OrderNo              int64  `json:"orderNo,omitempty"`
	SpecificCancelReason string `json:"specificCancelReason,omitempty"`
	CancelReasonId       int64  `json:"cancelReasonId,omitempty"`
}

type OrderCancelResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateOrderCancelRequest() (request *OrderCancelRequest) {
	request = &OrderCancelRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "order/cancel", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderCancelResponse() (response *api.BaseResponse[OrderCancelResponse]) {
	response = api.CreateResponse[OrderCancelResponse](&OrderCancelResponse{})
	return
}
