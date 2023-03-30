package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderAmountUpdate
// @id 1811
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1811?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=订单改价)
func (client *Client) OrderAmountUpdate(request *OrderAmountUpdateRequest) (response *OrderAmountUpdateResponse, err error) {
	rpcResponse := CreateOrderAmountUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderAmountUpdateRequest struct {
	*api.BaseRequest

	UpdateItemAmounts    []OrderAmountUpdateRequestUpdateItemAmounts `json:"updateItemAmounts,omitempty"`
	OrderNo              int64                                       `json:"orderNo,omitempty"`
	UpdateDeliveryAmount string                                      `json:"updateDeliveryAmount,omitempty"`
}

type OrderAmountUpdateResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateOrderAmountUpdateRequest() (request *OrderAmountUpdateRequest) {
	request = &OrderAmountUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "order/amount/update", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderAmountUpdateResponse() (response *api.BaseResponse[OrderAmountUpdateResponse]) {
	response = api.CreateResponse[OrderAmountUpdateResponse](&OrderAmountUpdateResponse{})
	return
}

type OrderAmountUpdateRequestUpdateItemAmounts struct {
	ItemId       int64  `json:"itemId,omitempty"`
	UpdateAmount string `json:"updateAmount,omitempty"`
}
